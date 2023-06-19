package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlValidate(t *testing.T) {
	assert := assert.New(t)
	var err error

	// Test empty input
	err = urlValidate("")
	assert.Error(err, "Expected error for empty input, but got nil")

	// Test invalid input format
	err = urlValidate("invalid")
	assert.Error(err, "Expected error for invalid input format, but got nil")

	// Test empty host
	err = urlValidate(":8080")
	assert.Error(err, "Expected error for empty host, but got nil")

	// Test invalid port
	err = urlValidate("localhost:abc")
	assert.Error(err, "Expected error for invalid port, but got nil")

	// Test port out of range
	err = urlValidate("localhost:70000")
	assert.Error(err, "Expected error for port out of range, but got nil")

	// Test valid input
	err = urlValidate("localhost:8080")
	assert.Nil(err, "Expected nil error, but got error")
}

func TestAliasValidate(t *testing.T) {
	assert := assert.New(t)
	var err error

	// Test empty input
	err = aliasValidate("")
	assert.Error(err, "Expected error for empty input, but got nil")

	// Test non-empty input
	err = aliasValidate("test")
	assert.Nil(err, "Expected nil error, but got error")
}
