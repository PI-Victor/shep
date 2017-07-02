package services

import (
	"io/ioutil"

	"github.com/bndr/gojenkins"
	"github.com/sirupsen/logrus"
)

// JenkinsSettings holds the information about one or more Jenkins servers that
// the bot should send or retrieve information from.
type Jenkins struct {
	URL            string             `json:"jenkinsURL"`
	User           string             `json:"jenkinsUser"`
	Token          string             `json:"jenkinsToken"`
	CACertFilePath string             `json:"caCertFilePath"`
	Client         *gojenkins.Jenkins `json:"-"`
	Jobs           []Job              `json:"-"`
}

type Job struct{}

func newJenkinsInstance() (*Jenkins, error) {
	return nil, nil
}

func (j *Jenkins) Start() error {
	client := gojenkins.CreateJenkins(j.URL, j.User, j.Token)

	if j.CACertFilePath != "" {
		caCert, err := ioutil.ReadFile(j.CACertFilePath)
		if err != nil {
			return err
		}
		if len(caCert) == 0 {
			logrus.Warnf("Specified CA Certificate file (%s) is empty. Using unencrypted connection", j.CACertFilePath)
		} else {
			client.Requester.CACert = caCert
		}
	}
	instance, err := client.Init()

	if err != nil {
		return err
	}

	logrus.Debugf("This is the jenkins instance: %#v", instance)
	return nil
}

func (j *Jenkins) Stop() error { return nil }

func (j *Job) StartJob() error { return nil }

func (j *Job) CancelJob() error { return nil }
