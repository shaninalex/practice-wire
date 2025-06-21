package config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/shaninalex/practice-wire/internal/config"
	"github.com/stretchr/testify/assert"
)

func Test_ConfigRootPath(t *testing.T) {
	conf := config.NewConfig()
	home, _ := os.UserHomeDir()
	expectedDir := fmt.Sprintf("%s/.config/notekeeper", home)
	assert.Equal(t, expectedDir, conf.RootPath())
}
