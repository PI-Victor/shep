package services

// VCSService is an interface that is used to inject specific application
// components related to CI/CD servers.
type VCSService interface {
	Start() error
	Stop() error
}

// Service is a generic service that is used to inject different service in the
// mechanism that starts the application.
type Service struct {
	VCSServices []VCSService
}
