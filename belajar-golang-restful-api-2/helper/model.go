package helper

import (
	"github.com/RyouRio/belajar-golang-restful-api-2/model/domain"
	"github.com/RyouRio/belajar-golang-restful-api-2/model/web"
)

func ToCategoryResponse(category *domain.Category) *web.CategoryResponse {
	return &web.CategoryResponse{
		Id: category.Id,
		Name: category.Name,
	}
}
func ToCategoryResponses(categories []*domain.Category) []*web.CategoryResponse {
	var categoryResponses []*web.CategoryResponse

	for _, category := range categories {
		categoryResponses = append(categoryResponses,  ToCategoryResponse(category))
	}
	return categoryResponses
}