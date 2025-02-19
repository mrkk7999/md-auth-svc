package implementation

import (
	mdgeotrack "md-auth-svc/iface"
)

type service struct {
	repository mdgeotrack.Repository
	cognito    mdgeotrack.CognitoServiceInterface
}

func New(repository mdgeotrack.Repository, cognito mdgeotrack.CognitoServiceInterface) mdgeotrack.Service {
	return &service{
		repository: repository,
		cognito:    cognito,
	}
}
