package cognito

import (
	"context"
	validatetoken "md-auth-svc/request_response/validate_token"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// ValidateToken checks the validity of a given token.
func (s *AuthService) ValidateToken(ctx context.Context, req *validatetoken.ValidateTokenRequest) (*validatetoken.ValidateTokenResponse, error) {
	input := &cognitoidentityprovider.GetUserInput{
		AccessToken: &req.Token,
	}

	_, err := s.cognito.Client.GetUser(ctx, input)
	if err != nil {
		return &validatetoken.ValidateTokenResponse{Valid: false, Message: "Invalid token"}, nil
	}

	return &validatetoken.ValidateTokenResponse{Valid: true, Message: "Valid token"}, nil
}
