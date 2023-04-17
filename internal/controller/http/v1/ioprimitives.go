// INFO: Out plain/text if problem with marshal
package v1

import (
	"encoding/json"
	"errors"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, data any) {
	buf, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json")

	status := http.StatusOK

	if code != -1 {
		status = code
	}

	w.WriteHeader(status)
	w.Write(buf)
}

func BindJSON(w http.ResponseWriter, r *http.Request, placeholder any) error {
	err := json.NewDecoder(r.Body).Decode(&placeholder)
	if err != nil {
		buf, err := json.Marshal(map[string]any{
			"msg": err.Error(),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(buf)

		return errors.New("-1") // Bad state
	}

	return nil
}
