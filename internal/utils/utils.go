package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseBody(r *http.Request, x interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(x); err != nil {
		return err
	}
	return nil
}

func StringToInt(w http.ResponseWriter, r *http.Request) uint64 {
	vars := mux.Vars(r)
	bookId, err := strconv.ParseUint(vars["id"], 0, 64)

	if err != nil {
		http.Error(w, "BadRequest", http.StatusBadRequest)
	}
	return bookId
}
