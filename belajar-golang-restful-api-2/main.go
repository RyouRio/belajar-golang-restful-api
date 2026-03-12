package main

import (
	"net/http"

	"github.com/RyouRio/belajar-golang-restful-api-2/app"
	"github.com/RyouRio/belajar-golang-restful-api-2/controller"
	"github.com/RyouRio/belajar-golang-restful-api-2/helper"
	"github.com/RyouRio/belajar-golang-restful-api-2/middleware"
	"github.com/RyouRio/belajar-golang-restful-api-2/repository"
	"github.com/RyouRio/belajar-golang-restful-api-2/service"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	defer db.Close()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: ":3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}