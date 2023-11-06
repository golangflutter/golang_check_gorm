package request

type CreateSizeRequest struct {
    Name  string `validate:"required,min=2,max=255" json:"name"`
    Value string `json:"value"`
}

type UpdateSizeRequest struct {
    Id    int    `validate:"required"`
    Name  string `validate:"required,min=2,max=255" json:"name"`
    Value string `json:"value"`
}


type PaginationRequestSize struct {
    Page     int    `json:"page"`
    PageSize int    `json:"pageSize"`
    SizeName string `json:"sizeName"`
    Value    string `json:"value"`
}
