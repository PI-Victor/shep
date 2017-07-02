package services

import (
	"github.com/bndr/gojenkins"
)

// JenkinsSettings holds the information about one or more Jenkins servers that
// the bot should send or retrieve information from.
type Jenkins struct {
	URL    string             `json:"jenkinsURL"`
	User   string             `json:"jenkinsUser"`
	Token  string             `json:"jenkinsToken"`
	Client *gojenkins.Jenkins `json:"-"`
}

func newJenkinsInstance() (*Jenkins, error) {
	return nil, nil
}

func (j *Jenkins) Start() error {
	var err error

	j.Client, err = gojenkins.CreateJenkins(j.URL, j.User, j.Token).Init()
	if err != nil {
		return err
	}
	return nil
}

func (j *Jenkins) Stop() error { return nil }

func (j *Jenkins) StartJob() error { return nil }

func (j *Jenkins) StopJob() error { return nil }

func (j *Jenkins) CancelJob() error { return nil }
