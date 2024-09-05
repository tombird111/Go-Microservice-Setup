package service

import (
	"net/http"
	"net/url"
	"io"
	"encoding/json"
)

const (
	KEY = "" /* WRITE API KEY HERE */
	URI = "https://api.audd.io/" /* The URI for the API */
)

func Service(audiostring string) (string, int) {
	data := url.Values{ /* Create a data object that contains the values to send */
		"api_token":{KEY},
		"audio":{audiostring},
	}
	if response, err := http.PostForm(URI, data); err == nil { /* Attempt to post the form */
		defer response.Body.Close() /* Close the response body at the end */
		if body, err := io.ReadAll(response.Body); err == nil { /* Read all parts of the response */
			var responseBody map[string]interface{} /* Create a map to map the body of the response */
			json.Unmarshal(body, &responseBody) /* Unmarshal the json in the body into the map */
			if responseBody["status"].(string) == "success" { /* If the response was successful */
				result := responseBody["result"] /* Get the result from the body of the response */
				var resultMap = map[string]string{} /* Create a map of strings to map the information of the result */
				for key, result := range result.(map[string]interface{}) {
					resultMap[key] = result.(string) /* Iterate through the result, adding to the string map */
				}
				return resultMap["title"], 200 /* Return the title, and the code 200 for the array */
			} else { /* If the response was unsuccessful */
				apiError := responseBody["error"] /* Get the error information from the body */
				var errorMap = map[string]float64{} /* Create a map of floats */
				for key, apiError := range apiError.(map[string]interface{}) {
					switch variable := apiError.(type){ 
						case float64:
							errorMap[key] = variable /* Iterate through the error information, adding any floats */
					}
				}
				if errorMap["error_code"] == 300 { /* If the error code was 300 (Unable to find based on the sent .wav file) */
					return "", 404 /* Return an empty string, and a 404 */
				} else { /* For different error codes */
					return "", 400 /* Return 400 for bad requests */
				}
			}
		}
	}
	return "", 500 /* Return a blank string and 500 if the server fails to post the form*/
}