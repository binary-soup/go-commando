package test

import "testing"

// Assert the prompt is repeated the expected amount of times.
func PromptCount(t *testing.T, test, prompt string, count int) {
	ContainsSubstringCount(t, test, prompt, count, "wrong number of prompts")
}
