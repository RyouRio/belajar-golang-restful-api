package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result any)error{
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	return err
}
func WriteToResponseBody(w http.ResponseWriter, response any)  {
	w.Header().Add("Content-Type", "application/json") // memberitahu bentuknya itu akan json
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err)
}