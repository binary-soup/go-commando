package sample_test

import (
	"testing"

	"github.com/binary-soup/go-commando/command/sample"
	"github.com/binary-soup/go-commando/test"
	"github.com/stretchr/testify/suite"
)

type SampleConfigSuite struct {
	suite.Suite
	Config sample.SampleConfig
}

func TestSampleConfigSuite(t *testing.T) {
	suite.Run(t, &SampleConfigSuite{})
}

func (s *SampleConfigSuite) SetupTest() {
	s.Config = sample.SampleConfig{
		Path:  "path/to/somewhere",
		Count: 50,
	}
}

func (s *SampleConfigSuite) TestValid() {
	test.ConfigValid(s.T(), s.Config)
}

func (s *SampleConfigSuite) TestInvalidPathEmpty() {
	s.Config.Path = ""
	test.ConfigValidateErrors(s.T(), s.Config, []string{"path", "empty"})
}

func (s *SampleConfigSuite) TestInvalidCountLowerLimit() {
	s.Config.Count = 5
	test.ConfigValid(s.T(), s.Config)

	s.Config.Count--
	test.ConfigValidateErrors(s.T(), s.Config, []string{"count", "less than min"})
}

func (s *SampleConfigSuite) TestInvalidCountUpperLimit() {
	s.Config.Count = 100
	test.ConfigValid(s.T(), s.Config)

	s.Config.Count++
	test.ConfigValidateErrors(s.T(), s.Config, []string{"count", "more than max"})
}
