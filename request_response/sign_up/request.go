package signup

type SignUpRequest struct {
	Username        string `json:"username"  binding:"required"`
	Email           string `json:"email"  binding:"required"`
	GivenName       string `json:"given_name"  binding:"required"`
	FamilyName      string `json:"family_name"  binding:"required"`
	Password        string `json:"password"  binding:"required"`
	ConfirmPassword string `json:"confirm_password"  binding:"required"`
	UserPoolGroup   string `json:"user_pool_group"`
}

type UserSignUpRequest struct {
	SignUpRequest
	// Tenant ID to validate the tenant
	TenantID string `json:"tenant_id" binding:"required,uuid"`
}

type SysAdminSignUpRequest struct {
	SignUpRequest
}
