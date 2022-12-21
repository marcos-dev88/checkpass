package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/marcos-dev88/checkpass/application"
	app "github.com/marcos-dev88/checkpass/application"
)

type Handler interface {
	HandlePassword(rw http.ResponseWriter, req *http.Request)
}

type handler struct {
	application application.Application
}

type response struct {
	Verify  bool     `json:"verify"`
	NoMatch []string `json:"noMatch"`
}

func NewHandler(app app.Application) Handler {
	return handler{application: app}
}

func (h handler) HandlePassword(rw http.ResponseWriter, req *http.Request) {

	var reqPassword app.Password

	data, err := io.ReadAll(req.Body)

	if err != nil {
		DefaultErrorResponse(rw, err, http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(data, &reqPassword); err != nil {
		DefaultErrorResponse(rw, err, http.StatusInternalServerError)
		return
	}

	if err = h.application.SetPassword(&reqPassword); err != nil {
		DefaultErrorResponse(rw, err, http.StatusInternalServerError)
		return
	}

	h.application.SetVerify(reqPassword) // updating verify rules with new request data

	failValidations, err := h.application.RunCheck()

	if err != nil {
		DefaultErrorResponse(rw, err, http.StatusInternalServerError)
		return
	}

	if len(failValidations) > 0 {
		r := response{Verify: false, NoMatch: failValidations}
		DefaultSuccessResponse(rw, r, http.StatusOK)
		return
	}

	r := response{Verify: true, NoMatch: []string{}}
	DefaultSuccessResponse(rw, r, http.StatusOK)

}

func DefaultSuccessResponse(rw http.ResponseWriter, resp response, httpStatus int) {
	b, err := json.Marshal(resp)

	if err != nil {
		DefaultErrorResponse(rw, err, http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(b)
}

func DefaultErrorResponse(rw http.ResponseWriter, err error, httpStatus int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(httpStatus)
	rw.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
}
