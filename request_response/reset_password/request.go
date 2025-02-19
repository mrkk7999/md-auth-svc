package resetpassword

type ResetPasswordRequest struct {
	Username         string `json:"username"`
	ConfirmationCode string `json:"confirmation_code"`
	NewPassword      string `json:"new_password"`
}
