package challenge02

import (
	"reflect"
	"testing"
)

type TestCase struct {
	name       string
	input1     []Person
	input2     string
	expected01 []string
	expected02 []string
	expected03 []string
}

func TestGetSocialBubbles(t *testing.T) {

	testCases := []TestCase{
		{
			name: "Testing for normal ladder",
			input1: []Person{
				Person{ID: "0", Friends: []string{"4", "1"}},
				Person{ID: "1", Friends: []string{"2", "3"}},
				Person{ID: "2", Friends: []string{"1"}},
				Person{ID: "3", Friends: []string{"2", "4"}},
				Person{ID: "4", Friends: []string{"3"}},
			},
			input2:     "2",
			expected01: []string{"1"},
			expected02: []string{"2", "3"},
			expected03: []string{"1", "2", "4"},
		},
		{
			name: "Testing for normal ladder with more elements",
			input1: []Person{
				Person{ID: "0", Friends: []string{"4", "1"}},
				Person{ID: "1", Friends: []string{"2", "3"}},
				Person{ID: "2", Friends: []string{"1"}},
				Person{ID: "3", Friends: []string{"2", "4"}},
				Person{ID: "4", Friends: []string{"3"}},
			},
			input2:     "3",
			expected01: []string{"2", "4"},
			expected02: []string{"1", "3"},
			expected03: []string{"2", "3", "4"},
		},
		{
			name: "Testint non-existing ID",
			input1: []Person{
				Person{ID: "0", Friends: []string{"4", "1"}},
				Person{ID: "1", Friends: []string{"2", "3"}},
				Person{ID: "2", Friends: []string{"1"}},
				Person{ID: "3", Friends: []string{"2", "4"}},
				Person{ID: "4", Friends: []string{"3"}},
			},
			input2:     "5",
			expected01: []string{},
			expected02: []string{},
			expected03: []string{},
		},
	}

	for _, tc := range testCases {
		e1, e2, e3 := GetSocialBubbles(tc.input1, tc.input2)

		if (len(e1) > 0 && len(tc.expected01) > 0) && !reflect.DeepEqual(e1, tc.expected01) {
			t.Errorf("%s for output1 got %v, wanted %v, of type %s", tc.name, e1, tc.expected01, "[]string")
		}
		if (len(e2) > 0 && len(tc.expected02) > 0) && !reflect.DeepEqual(e2, tc.expected02) {
			t.Errorf("%s for output2 got %v, wanted %v, of type %s", tc.name, e2, tc.expected02, "[]string")
		}
		if (len(e3) > 0 && len(tc.expected03) > 0) && !reflect.DeepEqual(e3, tc.expected03) {
			t.Errorf("%s for output3 got %v, wanted %v, of type %s", tc.name, e3, tc.expected03, "[]string")
		}
	}
}
