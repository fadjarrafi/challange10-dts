package dto

type NewProductRequest struct {
	Title       string `json:"title" valid:"required~title cannot be empty" example:"Rokok Surya 16"`
	Description string `json:"description" valid:"required~description cannot be empty" example:"Deskripsi dari Rokok Surya 16"`
}

type NewProductResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
