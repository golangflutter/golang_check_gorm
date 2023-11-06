package request

type CreateColorRequest struct {
    Name  string `validate:"required,min=2,max=255" json:"name"`
    Value string `json:"value"`
}

type UpdateColorRequest struct {
    Id    int    `validate:"required"`
    Name  string `validate:"required,min=2,max=255" json:"name"`
    Value string `json:"value"`
}

type PaginationRequestColor struct {
    Page     int    `json:"page"`
    PageSize int    `json:"pageSize"`
    ColorName string `json:"colorName"`
    Value    string `json:"value"`
}
