package request

type CreateCategoryRequest struct {
	Name string `validate:"required,min=2,max=255" json:"name"`
}

type UpdateCategoryRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,min=2,max=255" json:"name"`
}
type PaginationRequest struct {
    Page       int    `json:"page"`
    PageSize   int    `json:"pageSize"`
    CategoryName string `json:"categoryName"`
}