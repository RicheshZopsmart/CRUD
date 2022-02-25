package Pallindrome

import (
	"reflect"
	"testing"
)

func TestPall(t *testing.T) {
	testcases := []struct {
		desc   string
		text   string
		output bool
	}{
		{
			desc:   "Even no. and same chars",
			text:   "AAAAAA",
			output: true,
		},
		{
			desc:   "Odd nums of char with middle one different",
			text:   "AAAHAAA",
			output: true,
		},
		{
			desc:   "Normal scribbled text",
			text:   "KAKAAKAK",
			output: true,
		},
		{
			desc:   "Not a pallindrome text",
			text:   "Hello",
			output: false,
		},
		{
			desc:   "Empty String",
			text:   "",
			output: true,
		},
	}
	for _, tt := range testcases {

		t.Run(tt.desc, func(t *testing.T) {
			got := isPall(tt.text)
			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("got %v, want: %v", got, tt.output)
			}
		})
	}
}
