package implementation

import (
	"context"
	usergroup "md-auth-svc/request_response/user_group"
)

func (s *service) ChangeGroup(ctx context.Context, req usergroup.UserToGroupRequest) error {
	return s.cognito.MoveUserToGroup(ctx, req)
}
