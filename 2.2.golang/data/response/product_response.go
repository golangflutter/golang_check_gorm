package response

type ProductResponse struct {
    Id        int    `json:"id"`
    Name      string    `json:"name"`
    Price     float64   `json:"price"`
    IsFeatured bool     `json:"isFeatured"`
    IsArchived bool     `json:"isArchived"`
    ImageURLs []string  `json:"imageUrls"`
    CreatedAt string    `json:"createdAt"`
    UpdatedAt string    `json:"updatedAt"`
    Category   CategoryResponse `json:"category"`
    Size       SizeResponse     `json:"size"`
    Color      ColorResponse    `json:"color"`
}

