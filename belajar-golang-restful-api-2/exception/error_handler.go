package exception

import (
	"net/http"

	"github.com/RyouRio/belajar-golang-restful-api-2/helper"
	"github.com/RyouRio/belajar-golang-restful-api-2/model/web"
	"github.com/go-playground/validator"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err any) {
	if notFoundError(w, r , err){
		return
	}

	if validationErrors(w, r, err) {
		return
	}
	internalServerError(w, r, err)
}

func validationErrors(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
		Code: http.StatusBadRequest,
		Status: "BAD REQUEST",
		Data: exception.Error(),
	}
	helper.WriteToResponseBody(w, webResponse)
		return true
	}else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
		Code: http.StatusNotFound,
		Status: "NOT FOUND",
		Data: exception.Error,
	}
	helper.WriteToResponseBody(w, webResponse)
		return true
	}else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code: http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data: err,
	}
	helper.WriteToResponseBody(w, webResponse)
}