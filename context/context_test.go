package context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return http.Header{}
}
func (s *SpyResponseWriter) Write(b []byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}
func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	go func() {
		result := ""
		for _, char := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store cancelled")
				log.Println(ctx.Err())
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(char)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func TestServer(t *testing.T) {
	t.Run("should return correct response", func(t *testing.T) {
		data := "hello, world"
		server := Server(&SpyStore{response: data, t: t})
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertEqual(t, data, response.Body.String())
	})
	t.Run("should cancel context", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)
		response := &SpyResponseWriter{}
		server.ServeHTTP(response, request)
		assertEqual(t, false, response.written)
		/*
			if response.written {
				t.Error("response should not be written")
			}
		*/
	})
	/*
		t.Run("should return data from store", func(t *testing.T) {
			data := "hello, world"
			store := &SpyStore{response: data, t: t}
			server := Server(store)

			request := httptest.NewRequest(http.MethodGet, "/", nil)
			response := httptest.NewRecorder()
			server.ServeHTTP(response, request)
			assertEqual(t, data, response.Body.String())
			store.assertWasNotCancelled()
		})
	*/
}

func assertEqual(t *testing.T, expected, got interface{}) {
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}
