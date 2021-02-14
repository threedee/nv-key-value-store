package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

// KeyValueStore stores score information about players.
type KeyValueStore interface {
	SetValue(key, value string)
	FetchValue(key string) string
	DeleteKey(key string)
}

// KeyValueServer is a HTTP interface for player information.
type KeyValueServer struct {
	store KeyValueStore
	http.Handler
}

const jsonContentType = "application/json"

// NewKeyValueServer creates a KeyValueServer with routing configured.
func NewKeyValueServer(store KeyValueStore) *KeyValueServer {
	p := new(KeyValueServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(p.keyValueHandler))

	p.Handler = router

	return p
}

func (k *KeyValueServer) keyValueHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/"):]

	switch r.Method {
	case http.MethodPut:
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		k.processSet(w, key, string(body))
	case http.MethodGet:
		k.processFetch(w, key)
	case http.MethodDelete:
		k.processDelete(w, key)
	}
}

func (k *KeyValueServer) processSet(w http.ResponseWriter, key, value string) {
	k.store.SetValue(key, value)
	w.WriteHeader(http.StatusNoContent)

}

func (k *KeyValueServer) processFetch(w http.ResponseWriter, key string) {
	if value := k.store.FetchValue(key); value != "" {

		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", strconv.Itoa(len(value)))
		w.Write([]byte(value))
		// fmt.Fprint(w, value)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (k *KeyValueServer) processDelete(w http.ResponseWriter, key string) {
	k.store.DeleteKey(key)
	w.WriteHeader(http.StatusNoContent)
}
