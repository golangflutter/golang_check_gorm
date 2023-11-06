package request

type CreateBillboardRequest struct {
    Label    string   `validate:"required,min=2,max=255" json:"label"`
    ImageUrl []string `json:"imageUrl"`
}

type UpdateBillboardRequest struct {
    Id       int      `validate:"required"`
    Label    string   `validate:"required,min=2,max=255" json:"label"`
    ImageUrl []string `json:"imageUrl"`
}

type PaginationRequestBillboard struct {
    Page     int      `json:"page"`
    PageSize int      `json:"pageSize"`
    Label    string   `json:"label"`
    ImageUrl []string `json:"imageUrl"`
}
