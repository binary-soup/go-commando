package sample_test

import (
	"testing"

	"github.com/binary-soup/go-command/command"
	"github.com/binary-soup/go-command/command/sample"
	"github.com/binary-soup/go-command/config"
	"github.com/binary-soup/go-command/d"
	"github.com/binary-soup/go-command/d/fs"
	"github.com/binary-soup/go-command/data"
	"github.com/binary-soup/go-command/test"
	"github.com/stretchr/testify/suite"
)

const CONFIG_FILE = "config.json"

type ConfigTestSuite struct {
	test.CommandSuite
	Data fs.StaticMap
}

func TestConfigCommandSuite(t *testing.T) {
	s := ConfigTestSuite{
		CommandSuite: test.NewCommandSuite(command.NewConfigCommand[sample.SampleConfig]()),
	}

	d.FileSystem.Override(&s.Data)
	defer d.FileSystem.Restore()

	suite.Run(t, &s)
}

func (s *ConfigTestSuite) SetupTest() {
	s.Data = fs.StaticMap{}
}

func (s *ConfigTestSuite) StoreConfig(path string, cfg config.Config) {
	err := data.SaveJSON("testing config", cfg, path)
	s.Require().NoError(err)
}

func (s *ConfigTestSuite) TestMissing() {
	s.RequireCommandFail([]string{"-cfg", CONFIG_FILE}, []string{"error opening", "config"})
}

func (s *ConfigTestSuite) TestValid() {
	s.StoreConfig(CONFIG_FILE, sample.SampleConfig{
		Path:  CONFIG_FILE,
		Count: 50,
	})

	p := test.OpenStdoutPipe()
	defer p.Close()

	s.RequireCommandPass([]string{"-cfg", CONFIG_FILE})
	p.CloseInput()

	p.AssertLineContainsSubstrings(s.T(), []string{"VALID"})
}

func (s *ConfigTestSuite) TestInvalidValid() {
	s.StoreConfig(CONFIG_FILE, sample.SampleConfig{
		Path:  "does/not/exist",
		Count: 105,
	})

	s.RequireCommandFail([]string{"-cfg", CONFIG_FILE}, []string{"invalid config"})
}
