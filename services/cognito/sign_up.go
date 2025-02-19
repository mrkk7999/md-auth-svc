package cognito

import (
	"context"
	"fmt"
	"log"
	signup "md-auth-svc/request_response/sign_up"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/aws-sdk-go/aws"
)

// SignUp - Registers a user in Cognito and assigns them to a group
func (s *AuthService) SignUp(ctx context.Context, req *signup.SignUpRequest) (*signup.SignUpResponse, error) {
	if req.Password != req.ConfirmPassword {
		return nil, fmt.Errorf("passwords do not match")
	}

	secretHash := s.cognito.GenerateSecretHash(req.Username)

	// Step 1:
	// Sign up the user
	input := &cognitoidentityprovider.SignUpInput{
		ClientId:   &s.cognito.ClientID,
		Username:   &req.Username,
		Password:   &req.Password,
		SecretHash: &secretHash,
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: &req.Email},
			{Name: aws.String("given_name"), Value: &req.GivenName},
			{Name: aws.String("family_name"), Value: &req.FamilyName},
		},
	}

	result, err := s.cognito.Client.SignUp(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("sign-up failed: %v", err)
	}

	// Step 2:
	// Assign user to Cognito User Pool Group (only if group is provided)
	log.Println("user pool group", req.UserPoolGroup)
	if req.UserPoolGroup != "" {
		groupInput := &cognitoidentityprovider.AdminAddUserToGroupInput{
			UserPoolId: &s.cognito.UserPoolID,
			Username:   &req.Username,
			GroupName:  &req.UserPoolGroup,
		}

		log.Println(groupInput)
		_, err := s.cognito.Client.AdminAddUserToGroup(ctx, groupInput)
		if err != nil {
			return nil, fmt.Errorf("user signed up but failed to assign to group: %v", err)
		}
	}

	return &signup.SignUpResponse{
		UserSub: *result.UserSub,
		Message: "Sign-up successful and user assigned to group",
	}, nil
}
