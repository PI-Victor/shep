package services

type VCSService interface {
	Start() error
	Stop() error
}
