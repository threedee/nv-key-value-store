package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `{}`)
	defer cleanDatabase()
	store, err := NewFileSystemKeyValueStore(database)

	assertNoError(t, err)

	server := NewKeyValueServer(store)

	server.ServeHTTP(httptest.NewRecorder(), newSetValueRequest("ABC", "hurz"))
	server.ServeHTTP(httptest.NewRecorder(), newSetValueRequest("DEF", "test"))

	t.Run("get value", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newValueRequest("ABC"))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, []byte(response.Body.String()), "hurz")
	})
	t.Run("Key not found", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newValueRequest("XYZ"))
		assertStatus(t, response.Code, http.StatusNotFound)
	})
	t.Run("Set key", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newSetValueRequest("XYZ", "hat geklappt"))
		assertStatus(t, response.Code, http.StatusNoContent)
	})

	t.Run("Delete key", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newDeleteKeyRequest("ABC"))
		assertStatus(t, response.Code, http.StatusNoContent)
	})
	t.Run("Key not found", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newValueRequest("ABC"))
		assertStatus(t, response.Code, http.StatusNotFound)
	})

}
