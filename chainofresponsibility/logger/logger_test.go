package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewLogger()

	logger.Debug("debug message")
	logger.Debugf("debug message: %v", 301)

	logger.Info("info message")
	logger.Infof("info message: %v", 200)

	logger.Error("error message")
	logger.Errorf("error message: %v", 400)
}
