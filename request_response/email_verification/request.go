package emailverification

type VerifyEmailRequest struct {
	Username         string `json:"username"`
	VerificationCode string `json:"verification_code"`
}
