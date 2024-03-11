package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
}

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, store.Fetch())
	}
}
