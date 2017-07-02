package services

// CIServer is an interface that is used for integrating CI/CD servers such as
// Jenkins and Concourse.
type CIServer interface {
	Start() error
	Stop() error
}

type CIJobs interface {
	Start() error
	Restart() error
	Cancel() error
}
