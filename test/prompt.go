package test

import "testing"

// Assert the prompt is repeated the expected amount of times.
func PromptCount(t *testing.T, test, prompt string, count int) {
	ContainsSubstringCount(t, test, prompt, count, "wrong number of prompts")
}

// Assert the prompt is an overwrite prompt with the expected number of repetitions.
func PromptOverwrite(t *testing.T, test string, count int) {
	PromptCount(t, test, "Overwrite [Y/n]?", count)
}
