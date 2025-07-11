package sample_test

import (
	"path/filepath"
	"testing"

	"github.com/binary-soup/go-command/command"
	"github.com/binary-soup/go-command/command/sample"
	"github.com/binary-soup/go-command/config"
	"github.com/binary-soup/go-command/data"
	"github.com/binary-soup/go-command/test"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	test.CommandSuite
}

func TestConfigCommandSuite(t *testing.T) {
	suite.Run(t, &ConfigTestSuite{
		CommandSuite: test.NewCommandSuite(command.NewConfigCommand[sample.SampleConfig]()),
	})
}

func (s *ConfigTestSuite) StoreConfig(path string, cfg config.Config) {
	err := data.SaveJSON("testing config", cfg, path)
	s.Require().NoError(err)
}

func (s *ConfigTestSuite) TestMissing() {
	var CONFIG_FILE = filepath.Join(s.T().TempDir(), "missing.json")
	s.RequireCommandFail([]string{"-cfg", CONFIG_FILE}, []string{"error opening", "config"})
}

func (s *ConfigTestSuite) TestValid() {
	var CONFIG_FILE = filepath.Join(s.T().TempDir(), "config.json")

	s.StoreConfig(CONFIG_FILE, sample.SampleConfig{
		Path:  "path/to/somewhere",
		Count: 50,
	})

	p := test.OpenStdoutPipe()
	defer p.Close()

	s.RequireCommandPass([]string{"-cfg", CONFIG_FILE})
	p.CloseInput()

	p.AssertLineContainsSubstrings(s.T(), []string{"VALID"})
}

func (s *ConfigTestSuite) TestInvalidValid() {
	var CONFIG_FILE = filepath.Join(s.T().TempDir(), "config.json")

	s.StoreConfig(CONFIG_FILE, sample.SampleConfig{
		Path:  "does/not/exist",
		Count: 105,
	})

	s.RequireCommandFail([]string{"-cfg", CONFIG_FILE}, []string{"invalid config"})
}
