package main

import (
	"testing"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	corev2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
}

func TestCheckArgs(t *testing.T) {
	assert := assert.New(t)
	event := corev2.FixtureEvent("entity1", "check1")
	i, e := checkArgs(event)
	assert.Equal(sensu.CheckStateWarning, i)
	assert.Error(e)
	plugin.Critical = float64(20)
	i, e = checkArgs(event)
	assert.Equal(sensu.CheckStateWarning, i)
	assert.Error(e)
	plugin.Warning = float64(10)
	i, e = checkArgs(event)
	assert.Equal(sensu.CheckStateOK, i)
	assert.NoError(e)
	plugin.Critical = float64(5)
	i, e = checkArgs(event)
	assert.Equal(sensu.CheckStateWarning, i)
	assert.Error(e)
}
