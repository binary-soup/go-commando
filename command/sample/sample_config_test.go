package sample_test

import (
	"math/rand"

	"github.com/binary-soup/go-command/command/sample"
	"github.com/stretchr/testify/suite"
)

type SampleConfigSuite struct {
	suite.Suite
	Rand rand.Source

	Config sample.SampleConfig
}

// func TestSampleConfigSuite(t *testing.T) {
// 	fs := fs.StaticMap{}
// 	fs.Create(CONFIG_FILE)

// 	d.FileSystem.Override(&fs)
// 	defer d.FileSystem.Restore()

// 	suite.Run(t, &SampleConfigSuite{
// 		Rand: rand.NewSource(time.Now().UnixNano()),
// 	})
// }

// func (s *SampleConfigSuite) SetupTest() {
// 	s.Config = sample.SampleConfig{
// 		Path:  CONFIG_FILE,
// 		Count: 50,
// 	}
// }

// func (s *SampleConfigSuite) TestValid() {
// 	verrs, err := s.Config.Validate()
// 	s.Require().NoError(err)

// 	s.Empty(verrs)
// }
