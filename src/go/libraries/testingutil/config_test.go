package testingutil_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ParseFile(t *testing.T) {
	// GIVEN
	path := "resources/test_config.yaml"

	// WHEN
	config, err := ParseFile[Config](path)

	// THEN
	assert.NoError(t, err)

	expectedConfig := &Config{
		BaseURL: "http://localhost:8080",
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     9432,
			Database: "dev_db",
			Username: "test",
			Password: "test",
		},
	}
	assert.Equal(t, expectedConfig, config)
}

func Test_ParseFile_FileNotFound(t *testing.T) {
	// GIVEN
	path := "file_not_found.yaml"

	// WHEN
	_, err := ParseFile[Config](path)

	// THEN
	assert.Error(t, err)
}

func Test_ParseFile_UnmarshalFailed(t *testing.T) {
	// GIVEN
	path := "resources/unknown_format.text"

	// WHEN
	_, err := ParseFile[Config](path)

	// THEN
	assert.Error(t, err)
}
