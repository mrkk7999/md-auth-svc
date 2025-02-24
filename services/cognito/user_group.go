package cognito

import (
	"context"
	"fmt"
	usergroup "md-auth-svc/request_response/user_group"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// AddUserToGroup - Adds a user to a specific Cognito group
func (s *AuthService) AddUserToGroup(ctx context.Context, username, groupName string) error {
	input := &cognitoidentityprovider.AdminAddUserToGroupInput{
		UserPoolId: &s.cognito.UserPoolID,
		Username:   &username,
		GroupName:  &groupName,
	}

	_, err := s.cognito.Client.AdminAddUserToGroup(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to add user to group: %v", err)
	}

	return nil
}

// MoveUserToGroup - Moves a user from one group to another
func (s *AuthService) MoveUserToGroup(ctx context.Context, req usergroup.UserToGroupRequest) error {
	// Remove user from the old group
	err := s.RemoveUserFromGroup(ctx, req.Username, req.OldGroup)
	if err != nil {
		return fmt.Errorf("failed to remove user from old group: %v", err)
	}

	// Add user to the new group
	err = s.AddUserToGroup(ctx, req.Username, req.NewGroup)
	if err != nil {
		return fmt.Errorf("failed to add user to new group: %v", err)
	}

	return nil
}

// RemoveUserFromGroup - Removes a user from a specific Cognito group
func (s *AuthService) RemoveUserFromGroup(ctx context.Context, username, groupName string) error {
	input := &cognitoidentityprovider.AdminRemoveUserFromGroupInput{
		UserPoolId: &s.cognito.UserPoolID,
		Username:   &username,
		GroupName:  &groupName,
	}

	_, err := s.cognito.Client.AdminRemoveUserFromGroup(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to remove user from group: %v", err)
	}

	return nil
}
