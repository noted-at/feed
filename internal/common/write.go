package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, response any) {
	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "failed to serialize response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, "%s\n", string(b))
}
