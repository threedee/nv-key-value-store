package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type StubPlayerStore struct {
	keyvalue map[string]string
}

// GetLeague returns the scores of all the players.
func (s *StubPlayerStore) FetchValue(key string) string {
	return s.keyvalue[key]
}

func (s *StubPlayerStore) SetValue(key string, value string) {
	s.keyvalue[key] = value
}

func (s *StubPlayerStore) DeleteKey(key string) {
	delete(s.keyvalue, key)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]string{
			"ABC":  "hurz",
			"TEST": "Noch ein test",
		},
	}
	server := NewKeyValueServer(&store)
	t.Run("returns value of for 'ABC' key", func(t *testing.T) {

		request := newValueRequest("ABC")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.Bytes(), "hurz")
	})
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Header().Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.HeaderMap)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func newValueRequest(key string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/%s", key), nil)
	return req
}

func newSetValueRequest(key, value string) *http.Request {
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/%s", key), strings.NewReader(value))
	return req
}

func newDeleteKeyRequest(key string) *http.Request {
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/%s", key), nil)
	return req
}
func assertResponseBody(t testing.TB, got []byte, want string) {
	t.Helper()
	if !cmp.Equal(got, []byte(want)) {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
