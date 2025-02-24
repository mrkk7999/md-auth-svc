package cognito

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func (s *AuthService) SignOut(ctx context.Context, accessToken string) error {
	input := &cognitoidentityprovider.GlobalSignOutInput{
		AccessToken: &accessToken,
	}
	_, err := s.cognito.Client.GlobalSignOut(ctx, input)
	if err != nil {
		return fmt.Errorf("error sign-out : %v", err)
	}
	return nil
}
