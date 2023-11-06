package response

type BillboardResponse struct {
    Id        int    `json:"id"`
    Label     string `json:"label"`
    ImageUrl  []string `json:"imageUrl"`
    CreatedAt string `json:"created_at"`
}
