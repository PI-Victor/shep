package service

import (
	"github.com/PI-Victor/shep/pkg/fs"
	"github.com/Sirupsen/logrus"
)

// Service is the general service
type Service struct {
	//tasks []tasks
}

type task struct {
}

// NewService returns a new instance of the service structure along with
// defaults.
func NewService() *Service {
	return &Service{}
}

// Start starts the application bot service.
func (s *Service) Start(cfg *fs.Config) {
	logrus.Info("Starting Shep...")
}
