package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnectionDBWithoutConfig(t *testing.T)  {
	_, err := ConnectionDB()

	if assert.NotNil(t, err) {
		assert.Equal(t,"dial tcp: lookup yourhost: no such host", err.Error())
	} else {
		assert.Nil(t, err)
	}
}

func TestCreateRouter(t *testing.T) {
	router := CreateRouter()
	assert.NotNil(t, router)
}
