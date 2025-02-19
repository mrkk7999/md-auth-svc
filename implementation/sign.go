package implementation

import (
	"context"
	"fmt"
	signin "md-auth-svc/request_response/sign_in"
	signup "md-auth-svc/request_response/sign_up"
	"net/http"
)

// UserSignUp
func (s *service) UserSignUp(ctx context.Context, req signup.UserSignUpRequest) (*signup.SignUpResponse, error) {
	
	// Tenant verification
	if err := s.verifyTenant(ctx, req.TenantID); err != nil {
		return nil, fmt.Errorf("tenant verification failed: %v", err)
	}

	signUpReq := signup.SignUpRequest{
		Username:        req.Username,
		Email:           req.Email,
		GivenName:       req.GivenName,
		FamilyName:      req.FamilyName,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
		UserPoolGroup:   req.UserPoolGroup,
	}

	cogRes, err := s.cognito.SignUp(ctx, &signUpReq)
	if err != nil {
		return nil, err
	}
	res, err := s.repository.UserSignUp(ctx, &req, cogRes.UserSub)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SysAdminSignUp
func (s *service) SysAdminSignUp(ctx context.Context, req signup.SysAdminSignUpRequest) (*signup.SignUpResponse, error) {
	signUpReq := signup.SignUpRequest{
		Username:        req.Username,
		Email:           req.Email,
		GivenName:       req.GivenName,
		FamilyName:      req.FamilyName,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
		UserPoolGroup:   req.UserPoolGroup,
	}
	cogRes, err := s.cognito.SignUp(ctx, &signUpReq)
	if err != nil {
		return nil, err
	}
	res, err := s.repository.SysAdminSignUp(ctx, &req, cogRes.UserSub)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *service) SignIn(ctx context.Context, req *signin.SignInRequest) (*signin.SignInResponse, error) {
	res, err := s.cognito.SignIn(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *service) verifyTenant(ctx context.Context, tenantID string) error {
	apiURL := fmt.Sprintf("http://localhost:9002/api/v1/tenants/%v", tenantID)

	// Create request with Authorization header
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	// req.Header.Set("Authorization", authHeader)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to call tenant verification API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("tenant does not exist or service error (status: %d)", resp.StatusCode)
	}

	return nil
}
