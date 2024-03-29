package trace

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestLogrusHook(t *testing.T) {
	resetTraceKeys()

	key := "key"
	value := "value"

	SetTraceIDKey(key)
	SetTraceIDGenerator(func() string {
		return value
	})

	var buffer bytes.Buffer
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(&buffer)
	log.AddHook(NewLogrusHook())

	ctx := Trace(context.Background())
	log.WithContext(ctx).Error("message")
	var fields logrus.Fields
	err := json.Unmarshal(buffer.Bytes(), &fields)
	require.Nil(t, err)
	require.Equal(t, value, fields[key])
}
