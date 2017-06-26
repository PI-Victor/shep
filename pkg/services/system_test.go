package services

import (
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewConfig(t *testing.T) {
	testCfg := &Config{}
	newCfg := NewCfg()
	if !reflect.DeepEqual(testCfg, newCfg) {
		t.Errorf("Expected %#v to match %#v", testCfg, newCfg)
	}
}

func TestNewDefaultConfig(t *testing.T) {
	newDefaultCfg := newDefaultCfg()
	if newDefaultCfg.DebugLevel != logrus.InfoLevel {
		t.Errorf("Expected DebugLevel to be %d", logrus.InfoLevel)
	}
}
