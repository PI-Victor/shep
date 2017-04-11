package fs

import (
	"reflect"
	"testing"

	"github.com/Sirupsen/logrus"
)

func TestNewConfig(t *testing.T) {
	testCfg := &Config{}
	newCfg := NewConfig()
	if !reflect.DeepEqual(testCfg, newCfg) {
		t.Errorf("Expected %#v to match %#v", testCfg, newCfg)
	}
}

func TestNewDefaultConfig(t *testing.T) {
	newDefaultCfg := NewDefaultConfig()
	if newDefaultCfg.DebugLevel != logrus.InfoLevel {
		t.Errorf("Expected DebugLevel to be %d", logrus.InfoLevel)
	}
}
