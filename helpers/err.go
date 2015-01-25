// Steve Phillips / elimisteve
// 2015.01.25

package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/222Labs/help"
)

func WriteJSON(w http.ResponseWriter, structure interface{}) {
	jsonData, err := json.Marshal(structure)
	if err != nil {
		help.WriteError(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonData)
}
