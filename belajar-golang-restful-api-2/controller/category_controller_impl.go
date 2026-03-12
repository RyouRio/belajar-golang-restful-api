package controller

import (
	"net/http"
	"strconv"

	"github.com/RyouRio/belajar-golang-restful-api-2/helper"
	"github.com/RyouRio/belajar-golang-restful-api-2/model/web"
	"github.com/RyouRio/belajar-golang-restful-api-2/service"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}
func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateRequest := &web.CategoryCreateRequest{} // membuat struct kosong untuk menampung json

	helper.ReadFromRequestBody(r, categoryCreateRequest) // memanggil decoder (json to struct go)
	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest) // Data yang dibaca dari json akan di proses oleh service
	webResponse := web.WebResponse{  // selesai dimasak ini struk nya dan akan return beberapa data
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse) // Encode struct go to json
	
}
func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	categoryUpdateRequest := &web.CategoryUpdateRequest{} // Membuat struct kosong CategoryUpdateRequest
	if err := helper.ReadFromRequestBody(r, categoryUpdateRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	} // membaca JSON dari body request, lalu decode ke dalam struct

	categoryId := params.ByName("categoryId") // ambil nilai categoryId dari URL, contohnya /category/5 → dapat "5"

	id, err := strconv.Atoi(categoryId) 
	if err != nil { //  cek err jika ada id yang tidak valid
		http.Error(w, "Invalid category id", http.StatusBadRequest)
		return
	}

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}
func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
	http.Error(w, "Invalid category id", http.StatusBadRequest)
	return
	}

	controller.CategoryService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		
	}

	helper.WriteToResponseBody(w, webResponse)

}
func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		http.Error(w, "Invalid category id", http.StatusBadRequest)
		return
	}

	categoryResponse := controller.CategoryService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",  
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponses,
	}

	helper.WriteToResponseBody(w, webResponse)

}