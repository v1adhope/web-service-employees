package v1

import (
	"errors"
	"net/http"
)

var _errQueryIsEmpty = errors.New("query is empty")

func (rs *Routes) writeNotFound(w http.ResponseWriter) {
	http.Error(w, "404 page not found", http.StatusNotFound)
}
