package log

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New("test"))
}

func TestNewWithZap(t *testing.T) {
	zl, _ := zap.NewProduction()
	l := NewWithZap(zl)
	assert.NotNil(t, l)
}

func buildRequest(requestID, correlationID string) *http.Request {
	req, _ := http.NewRequest("GET", "http://example.com", bytes.NewBufferString(""))
	if requestID != "" {
		req.Header.Set("X-Request-ID", requestID)
	}
	if correlationID != "" {
		req.Header.Set("X-Correlation-ID", correlationID)
	}
	return req
}

func TestNewForTest(t *testing.T) {
	logger, entries := NewForTest()
	assert.Equal(t, 0, entries.Len())
	logger.Info("msg 1")
	assert.Equal(t, 1, entries.Len())
	logger.Info("msg 2")
	logger.Info("msg 3")
	assert.Equal(t, 3, entries.Len())
	entries.TakeAll()
	assert.Equal(t, 0, entries.Len())
	logger.Info("msg 4")
	assert.Equal(t, 1, entries.Len())
}