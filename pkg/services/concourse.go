package services

import (
	"net/http"

	"github.com/concourse/go-concourse/concourse"
	"github.com/sirupsen/logrus"
)

// ConcourseInst holds information about a Concourse CI instance.
type ConcourseInst struct {
	URL   string         `json:"url,omitempty"`
	User  string         `json:"user,omitempty"`
	Token string         `json:"token,omitempty"`
	Jobs  []ConcourseJob `json:"-"`
}

// ConcourseJob holds information about a job in Concourse.
type ConcourseJob struct{}

// Start will create a new instance of Concourse CI.
func (c *ConcourseInst) Start() error {
	httpClient := http.Client{}
	newClient := concourse.NewClient(c.URL, &httpClient)
	logrus.Debug(newClient.GetInfo())
	return nil
}

// Stop will gracefully stop a Concourse instance.
func (c *ConcourseInst) Stop() error { return nil }

// StartJob will start a new Concourse job.
func (j *ConcourseJob) StartJob() error { return nil }

// CancelJob will cancel a currently running Concourse job.
func (j *ConcourseJob) CancelJob() error { return nil }
