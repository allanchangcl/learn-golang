package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	myconfig := GetConfig()
	fmt.Printf("%+v\n", myconfig.DB.Dialect)
	result := "mysql"
	// assert.Equal(t, myconfig, result, "The values should be the same")
	assert.Equal(t, myconfig.DB.Dialect, result, "The values should be the same")
}
