package implementation

import (
	"context"
	emailverification "md-auth-svc/request_response/email_verification"
)

func (s *service) EmailVerification(ctx context.Context, req emailverification.VerifyEmailRequest) (*emailverification.VerifyEmailResponse, error) {
	return s.cognito.VerifyEmail(ctx, &req)
}
