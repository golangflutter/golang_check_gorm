package response

type SizeResponse struct {
    Id        int    `json:"id"`
    Name      string `json:"name"`
    Value     string `json:"value"`
    CreatedAt string `json:"created_at"`
}
