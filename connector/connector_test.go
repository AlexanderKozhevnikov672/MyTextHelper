package connector

import (
	"bytes"
	"testing"
)

type testCase struct {
	description string
	input       string
	expected    string
}

var testCases = []testCase{
	{
		description: "Sample",
		input:       "+ pelun pelun pehrr\n? pe\n> h\n+ pehem pehem pehem\n? pe\n",
		expected:    "Added successfully!\npelun\npehrr\nAdded successfully!\npehem\n",
	},
	{
		description: "Test 1",
		input:       "+ abc abc abacaba cab\n? ab\n> c\n? c\n+ abbc abbc ccaba abbc cab\n? a\n> b\n> a\n? c\n> c\n? d\n",
		expected:    "Added successfully!\nabc\nabc\ncab\nAdded successfully!\nabbc\nabbc\nabacaba\ncab\nccaba\nNo candidates found!\n",
	},
}

func TestTextHelper(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			var buffer bytes.Buffer

			Run(bytes.NewReader([]byte(tc.input)), &buffer)

			if buffer.String() != tc.expected {
				t.FailNow()
			}
		})
	}
}
