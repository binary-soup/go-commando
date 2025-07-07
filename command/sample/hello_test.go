package sample_test

import (
	"bufio"
	"testing"

	"github.com/binary-soup/go-command/command/sample"
	"github.com/binary-soup/go-command/test"
	"github.com/stretchr/testify/require"
)

func TestNameNotEmpty(t *testing.T) {
	err := sample.NewHelloCommand().Run([]string{})
	require.Error(t, err)
	test.AssertErrorContainsAll(t, err, []string{"name", "cannot", "empty"})
}

func TestPrintName(t *testing.T) {
	s := test.CaptureStdout()
	defer s.Close()

	const NAME = "Bob"

	err := sample.NewHelloCommand().Run([]string{"-name", NAME})
	require.NoError(t, err)

	scanner := bufio.NewScanner(s)

	scanner.Scan()
	test.AssertContainsAll(t, scanner.Text(), []string{"Hello", NAME})
}
