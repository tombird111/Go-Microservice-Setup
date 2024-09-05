package resources

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"tracks/repository"
)

func updateTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var c repository.Cell
	if err := json.NewDecoder(r.Body).Decode(&c); err == nil {
		if id == c.Id {
			if n := repository.Update(c); n > 0 {
				w.WriteHeader(204) /* No Content */
			} else if n := repository.Insert(c); n > 0 {
				w.WriteHeader(201) /* Created */
			} else {
				w.WriteHeader(500) /* Internal Server Error */
			}
		} else {
			w.WriteHeader(400) /* Bad Request */
		}
	} else {
		w.WriteHeader(400) /* Bad Request */
	}
}

func getTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if c, n := repository.Read(id); n > 0 {
		d := repository.Cell{Id: c.Id, Audio: c.Audio}
		w.WriteHeader(200) /* OK */
		json.NewEncoder(w).Encode(d)
	} else if n == 0 {
		w.WriteHeader(404) /* Not Found */
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func listTracks(w http.ResponseWriter, r *http.Request) {
	if c, n := repository.GetIDs(); n > 0 {
		w.WriteHeader(200) /* OK */
		json.NewEncoder(w).Encode(c)
	} else {
		w.WriteHeader(500)
	}
}

func delTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if c, n := repository.Read(id); n > 0 {
		if n := repository.Delete(c); n > 0 {
			w.WriteHeader(204) /* OK */
		} else {
			w.WriteHeader(500) /* Internal server error */
		}
	} else if n == 0 {
		w.WriteHeader(404) /* Not Found */
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	/* Store */
	r.HandleFunc("/tracks/{id}", updateTrack).Methods("PUT")
	/* Get a specific track */
	r.HandleFunc("/tracks/{id}", getTrack).Methods("GET")
	/* Get a list of all tracks */
	r.HandleFunc("/tracks", listTracks).Methods("GET")
	/* Delete a track */
	r.HandleFunc("/tracks/{id}", delTrack).Methods("DELETE")
	return r
}