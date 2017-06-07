package connector

import (
	"os"
	"time"

	"github.com/apex/log"
)

const (
	defaultBufferSize = 500
	defaultRedisAddr  = "127.0.0.1:6379"
)

// Config vars for the application
type Config struct {
	// AppName is the application name and checkpoint namespace.
	AppName string

	// StreamName is the Kinesis stream.
	StreamName string

	// FlushInterval is a regular interval for flushing the buffer. Defaults to 1s.
	FlushInterval time.Duration

	// BufferSize determines the batch request size. Must not exceed 500. Defaults to 500.
	BufferSize int

	// Logger is the logger used. Defaults to log.Log.
	Logger log.Interface

	// Checkpoint for tracking progress of consumer.
	Checkpoint Checkpoint
}

// defaults for configuration.
func (c *Config) setDefaults() {
	if c.Logger == nil {
		c.Logger = log.Log
	}

	c.Logger = c.Logger.WithFields(log.Fields{
		"package": "kinesis-connectors",
	})

	if c.AppName == "" {
		c.Logger.WithField("type", "config").Error("AppName required")
		os.Exit(1)
	}

	if c.StreamName == "" {
		c.Logger.WithField("type", "config").Error("StreamName required")
		os.Exit(1)
	}

	c.Logger = c.Logger.WithFields(log.Fields{
		"app":    c.AppName,
		"stream": c.StreamName,
	})

	if c.BufferSize == 0 {
		c.BufferSize = defaultBufferSize
	}

	if c.FlushInterval == 0 {
		c.FlushInterval = time.Second
	}
}
