package cognito

import mdgeotrack "md-auth-svc/iface"

// AuthService
type AuthService struct {
	cognito *CognitoClient
}

// NewAuthService
func NewAuthService(cognito *CognitoClient) mdgeotrack.CognitoServiceInterface {
	return &AuthService{cognito: cognito}
}

// Example: Implement SignUp, SignIn, ValidateToken here
