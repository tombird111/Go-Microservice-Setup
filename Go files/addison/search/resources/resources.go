package resources

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"search/service"
)

type Title struct {
	Id string
}

func Search(w http.ResponseWriter, r *http.Request) {
	var a service.MusicInfo /* Create an object to store the audio of a request */
	if err := json.NewDecoder(r.Body).Decode(&a); err == nil { /* Decode the request */
		if id, rnum := service.Service(a.Audio); rnum == 200 { /* Get the ID and response number from the request */
			u := Title{id}
			w.WriteHeader(rnum)
			json.NewEncoder(w).Encode(u)
		} else {
			w.WriteHeader(rnum)
		}
	} else { /* If there was an error decoding a request, return a 400 error */
		w.WriteHeader(400)
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	/* Use the search function if requested */
	r.HandleFunc("/search", Search).Methods("POST")
	return r
}
