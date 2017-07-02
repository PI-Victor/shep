package services

type VCSService interface {
	Start() error
	Stop() error
}

type Service struct {
	VCSServices []VCSService
}
