package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Status int         `json "status"`
	Result interface{} `json: "result"`
}

func newResponse(data interface{}, status int) *response {
	return &response{
		Status: status,
		Result: data,
	}
} 
// look at the power of all these functions they are very silly and i wish thaat they woudl 
// be a lot saeir yeah I do think that ahki l shar

// yes i think that this is actually yeah i do think that this i 

func (resp *response) bytes() []byte {
	data, _ := json.Marshal(resp)
	return data
}

func (resp *response) string() []string {
	return string(resp.bytes)
}

func (resp *response) sendResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(resp.status)
	_, _ = w.Write(resp.bytes())
	log.Println(res.string())
}

// 204
// 204 status
func (resp *response) StatusOK(w http.ResponseWriter, r *http.Request, data interface{}) {
	newResponse(data, http.StatusOk).sendResponse(w, r)
}
func (resp *response) StatusNoContent(w. http.http.ResponseWriter, r http.Request) {
	newResponse(nil, http.StatusNoContent).sendResponse(w, r)
}

func (resp *response) StatusBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error":err.Error()}
	newResponse{data, http.StatusBadRequest}.sendResponse(w, r)
}

// so here what are w and r 
//remember, these are all methods on the response struct
//this is an interesting problem to solve

func (resp *response) statusNotFound(w. http.http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse{Data, http.StatusNotFound}.sendResponse(w,r)
}

func (resp *response) StatusMethodNotAllowed(w http.http.ResponseWriter, r *http.Request, err error) {
	newResponse(nil, http.StatusMethodNotAllowed).sendResponse(w,r)
}

func (resp *response) StatusConflict(w http.http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse{data, http.StatusNotFound}.sendResponse(w,r)
}

func (resp *response) StatusInternalServerError(w http.http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse{data, http.StatusNotFound}.sendResponse(w,r)
}
func (resp *response) sendResponse(w http.ResponseWriter, r *http.Request) {}
