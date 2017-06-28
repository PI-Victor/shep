package services

// CIServer is an interface that is used for integrating CI/CD servers such as
// Jenkins and Concourse.
type CIServer interface {
	StartJob() error
	RestartJob() error
	CancelJob() error
}