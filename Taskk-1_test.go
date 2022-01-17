package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_storyStats(t *testing.T) {
	testCases := []struct {
		desc        string
		input       string
		expShortest string
		expLargest  string
		avgWord     []string
	}{
		{
			desc:        "should return all value",
			input:       "ab-abc-abcd-abcde-abcdef",
			expShortest: "ab",
			expLargest:  "abcdef",
			avgWord:     []string{"abcd"},
		},
		{
			desc:        "should return empty average word list",
			input:       "ab-abc-xyz-abcde-abcdef",
			expShortest: "ab",
			expLargest:  "abcdef",
			avgWord:     []string{},
		},
		{
			desc:        "should return average word list",
			input:       "ab-abc-axyz-abcd-abcdefg",
			expShortest: "ab",
			expLargest:  "abcdefg",
			avgWord:     []string{"axyz", "abcd"},
		},
		{
			desc:        "no word",
			input:       "12-23-23",
			expShortest: "",
			expLargest:  "",
			avgWord:     []string{},
		},
		{
			desc:        "one word",
			input:       "123-abcd-123-21",
			expShortest: "abcd",
			expLargest:  "abcd",
			avgWord:     []string{"abcd"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			shortest, largest, avgList := storyStats(tc.input)
			assert.Equal(t, tc.expShortest, shortest)
			assert.Equal(t, tc.expLargest, largest)
			assert.EqualValues(t, tc.avgWord, avgList)
		})
	}
}

func Test_averageNumber(t *testing.T) {
	testCases := []struct {
		desc    string
		input   string
		average int
	}{
		{
			desc:    "should success",
			input:   "23-ab-48-caba-56-haha",
			average: 42,
		},
		{
			desc:    "should return only valid number average",
			input:   "23-ab-48-4a5ba-56-haha",
			average: 42,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := averageNumber(tc.input)
			assert.Equal(t, tc.average, result)
		})
	}
}

func Test_testValidity(t *testing.T) {
	testCases := []struct {
		desc      string
		input     string
		expResult bool
	}{
		{
			desc:      "should success",
			input:     "23-ab-48-caba-56-haha",
			expResult: true,
		},
		{
			desc:      "invalid input",
			input:     "23-ab--48-caba-56-haha",
			expResult: false,
		},
		{
			desc:      "invalid number",
			input:     "23-ab-48a-caba-56-haha",
			expResult: false,
		},
		{
			desc:      "invalid world",
			input:     "23-ab-48-c5aba-56-haha",
			expResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := testValidity(tc.input)
			assert.Equal(t, tc.expResult, result)
		})
	}
}

func Test_wholeStory(t *testing.T) {
	testCases := []struct {
		desc      string
		input     string
		expResult string
	}{
		{
			desc:      "should join all words",
			input:     "1-hello-2-world",
			expResult: "hello world",
		},
		{
			desc:      "should join only valid words",
			input:     "1-hel22lo-2-world",
			expResult: "world",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := wholeStory(tc.input)
			assert.Equal(t, tc.expResult, result)
		})
	}
}

func Test_isValidNumber(t *testing.T) {
	t.Run("should success", func(t *testing.T) {
		result := isValidNumber("123")
		assert.True(t, result)
	})

	t.Run("start with number but contain alphabits", func(t *testing.T) {
		result := isValidNumber("1sdh23")
		assert.False(t, result)
	})

	t.Run("start with  alphabits", func(t *testing.T) {
		result := isValidNumber("aaa")
		assert.False(t, result)
	})
}

func Test_isValidWord(t *testing.T) {
	t.Run("should success", func(t *testing.T) {
		result := isValidWord("abcd")
		assert.True(t, result)
	})

	t.Run("start with alphabits but contain number", func(t *testing.T) {
		result := isValidNumber("a2df")
		assert.False(t, result)
	})

	t.Run("start with number", func(t *testing.T) {
		result := isValidWord("123")
		assert.False(t, result)
	})

}
