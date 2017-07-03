package services

// CIServer is an interface that is used for integrating CI/CD servers such as
// Jenkins and Concourse.
type CIServer interface {
	Start() error
	Stop() error
}

// CIJobs is an interface used for abstracting Jobs in a CI/CD server
// environment.
type CIJobs interface {
	Start() error
	Cancel() error
}
