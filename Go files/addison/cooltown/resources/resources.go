package resources

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"bytes"
	"io"
	"strings"
)

type AudioClip struct {
	Audio string
}

type AudioId struct {
	Id string
}

func SearchFromClip(w http.ResponseWriter, r *http.Request) {
	var a AudioClip
	if err := json.NewDecoder(r.Body).Decode(&a); err == nil { /* Decode the request */
		postBody, _ := json.Marshal(a)
		responseBody := bytes.NewBuffer(postBody)
		if response, err := http.Post("http://127.0.0.1:3001/search", "application/json", responseBody); err == nil {
			defer response.Body.Close()
			if response.StatusCode == 200 {
				if body, err := io.ReadAll(response.Body); err == nil {
					var idBody AudioId /* Create an AudioId to store the id from the responses body */
					json.Unmarshal(body, &idBody) /* Unmarshal the JSON object into the AudioId struct */
					targetId := NormaliseInput(idBody.Id) /* Normalises the input. */
					if track, code := RetrieveWithId(targetId); code == 200 {
						TrackBody := AudioClip{track}
						w.WriteHeader(code)
						json.NewEncoder(w).Encode(TrackBody)
					} else {
						w.WriteHeader(code)
					}
				}
			} else {
				w.WriteHeader(response.StatusCode) /* Return an error based on the response from the search microservice */
			}
		} else {
			w.WriteHeader(500) /* If the server fails to contact the other microservice, return a 500 error */
		}
	} else { /* If there was an error decoding a request, return a 400 error */
		w.WriteHeader(400)
	}
}

func NormaliseInput(input string) string {
	return strings.ReplaceAll(input, " ", "+")
}

func RetrieveWithId(input string) (string, int) {
	if response, err := http.Get(("http://127.0.0.1:3000/tracks/" + input)); err == nil { /* Make a get request to the tracks microservice */
		defer response.Body.Close() /* Close the body when it is no longer needed */
		if body, err := io.ReadAll(response.Body); err == nil {
			if response.StatusCode == 200 {
				var audioBody AudioClip /* Create an AudioId to store the audio from the responses body */
				json.Unmarshal(body, &audioBody) /* Unmarshal the JSON object into the AudioId struct */
				return audioBody.Audio, 200
			} else {
				return "", response.StatusCode /* Return an empty string and a code based on the tracks microservices response */
			}
		} else {
			return "", 500 /* Return an empty string and a 500 code if there is a problem reading the track microservices response*/
		}
	} else {
		return "", 500 /* Return an empty string with a 500 code to say there was a problem contacting the tracks microservice */
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	/* Use the search function based on a recieved clip*/
	r.HandleFunc("/cooltown", SearchFromClip).Methods("POST")
	return r
}