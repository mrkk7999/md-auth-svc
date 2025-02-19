package validatetoken

type ValidateTokenResponse struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}
