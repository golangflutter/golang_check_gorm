package request

type CreateProductRequest struct {
    Name       string `validate:"required,min=2,max=255" json:"name"`
    Price      float64 `json:"price"`
    IsFeatured bool `json:"isFeatured"`
    IsArchived bool `json:"isArchived"`
    ImageUrls  []string `json:"imageUrls"`
    CategoryID int `json:"categoryID"`
    SizeID     int `json:"sizeID"`
    ColorID    int `json:"colorID"`
}

type UpdateProductRequest struct {
    Id        int `validate:"required"`
    Name       string `validate:"min=2,max=255" json:"name"`
    Price      float64 `json:"price"`
    IsFeatured bool `json:"isFeatured"`
    IsArchived bool `json:"isArchived"`
    ImageUrls  []string `json:"imageUrls"`
    CategoryId int `json:"categoryID"`
    SizeId     int `json:"sizeID"`
    ColorId   int `json:"colorID"`
}

