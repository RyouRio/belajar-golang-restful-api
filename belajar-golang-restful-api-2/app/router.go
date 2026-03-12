package app

import (
	"github.com/RyouRio/belajar-golang-restful-api-2/controller"
	"github.com/RyouRio/belajar-golang-restful-api-2/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll) // GET ALL DATAS
	router.GET("/api/categories/:categoryId", categoryController.FindById) // GET DATA BY ID
	router.POST("/api/categories", categoryController.Create) // CREATE DATA  
	router.PUT("/api/categories/:categoryId", categoryController.Update) // UPDATE DATA
	router.DELETE("/api/categories/:categoryId", categoryController.Delete) // DELETE DATA

	router.PanicHandler = exception.ErrorHandler 
	return  router
}