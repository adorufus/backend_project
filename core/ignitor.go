package core

import (
	"net/http"
)

func Ignite(handler http.Handler) {
	http.ListenAndServe(":8000", handler)
}
