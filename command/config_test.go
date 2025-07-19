package command_test

import (
	"path/filepath"
	"testing"

	"github.com/binary-soup/go-commando/command"
	"github.com/binary-soup/go-commando/command/sample"
	"github.com/binary-soup/go-commando/test"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	test.CommandSuite[command.ConfigCommand[sample.SampleConfig]]
	test.Rand
}

func TestConfigCommandSuite(t *testing.T) {
	suite.Run(t, &ConfigTestSuite{
		CommandSuite: test.NewCommandSuite(command.NewConfigCommand[sample.SampleConfig]()),
		Rand:         test.NewRandFromTime(),
	})
}

func (s *ConfigTestSuite) TestGetConfigPath() {
	var ENV = s.Rand.ASCII(10)
	var CONFIG = s.Rand.ASCII(15)

	// default env
	s.Cmd.SubmitArgs([]string{})
	s.Equal("main.json", filepath.Base(s.Cmd.GetConfigPath()))

	// env override
	s.Cmd.SubmitArgs([]string{"-env", ENV})
	s.Equal(ENV+".json", filepath.Base(s.Cmd.GetConfigPath()))

	// config override
	s.Cmd.SubmitArgs([]string{"-env", ENV, "-cfg", CONFIG})
	s.Equal(CONFIG, s.Cmd.GetConfigPath())
}

func (s *ConfigTestSuite) TestMissing() {
	var CONFIG_FILE = test.TempFile(s.T(), "missing.json")
	s.RequireCommandFail([]string{"-cfg", CONFIG_FILE}, []string{"error opening", "config"})
}

func (s *ConfigTestSuite) TestValid() {
	var CONFIG_FILE = test.CreateJSONTempFile(s.T(), "config.json", sample.SampleConfig{
		Path:  "path/to/somewhere",
		Count: 50,
	})

	pipe := test.OpenStdoutPipe()
	defer pipe.Close()

	s.RequireCommandPass([]string{"-cfg", CONFIG_FILE})
	pipe.CloseInput()

	test.ContainsSubstrings(s.T(), pipe.NextLine(s.T()), []string{"VALID"})
}

func (s *ConfigTestSuite) TestInvalid() {
	var CONFIG_FILE = test.CreateJSONTempFile(s.T(), "config.json", sample.SampleConfig{
		Path:  "",
		Count: 105,
	})

	s.RequireCommandFail([]string{"-cfg", CONFIG_FILE}, []string{"invalid config"})
}

func (s *ConfigTestSuite) TestNewNotExists() {
	var PATH = test.TempFile(s.T(), "config.json")

	s.RequireCommandPass([]string{"-new", "-cfg", PATH})
	test.CompareJSONFile(s.T(), PATH, sample.SampleConfig{})
}

func (s *ConfigTestSuite) TestNewExistsNoOverwrite() {
	var CONFIG = sample.SampleConfig{
		Path:  s.Rand.ASCII(15),
		Count: s.IntRange(-1000, 1000),
	}
	var PATH = test.CreateJSONTempFile(s.T(), "config.json", CONFIG)

	in := test.OpenStdinPipe([]any{"n"})
	defer in.Close()

	pipe := test.OpenStdoutPipe()
	defer pipe.Close()

	s.RequireCommandPass([]string{"-new", "-cfg", PATH})
	pipe.CloseInput()

	test.PromptOverwrite(s.T(), pipe.NextLine(s.T()), 1)
	test.CompareJSONFile(s.T(), PATH, CONFIG)
}

func (s *ConfigTestSuite) TestNewExistsYesOverwrite() {
	var CONFIG = sample.SampleConfig{
		Path:  s.Rand.ASCII(15),
		Count: s.IntRange(-1000, 1000),
	}
	var PATH = test.CreateJSONTempFile(s.T(), "config.json", CONFIG)

	in := test.OpenStdinPipe([]any{"Y"})
	defer in.Close()

	pipe := test.OpenStdoutPipe()
	defer pipe.Close()

	s.RequireCommandPass([]string{"-new", "-cfg", PATH})
	pipe.CloseInput()

	test.PromptOverwrite(s.T(), pipe.NextLine(s.T()), 1)
	test.CompareJSONFile(s.T(), PATH, sample.SampleConfig{})
}
