package service

import (
	"context"

	"github.com/RyouRio/belajar-golang-restful-api-2/model/web"
)

type CategoryService interface { // bussiness logic
	Create(ctx context.Context, request *web.CategoryCreateRequest) *web.CategoryResponse
	Update(ctx context.Context, request *web.CategoryUpdateRequest) *web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) *web.CategoryResponse
	FindAll(ctx context.Context) [] *web.CategoryResponse
}