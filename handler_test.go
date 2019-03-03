package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	err := handleRequest(context.Background())
	assert.NoError(t, err)
}
