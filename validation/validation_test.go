package validation

import "testing"

func TestIsIntervalStartOverIntervalStop(t *testing.T) {

	type testCase struct {
		startInterval string
		stopInterval  string
		expect        bool
	}
	cases := []testCase{
		{startInterval: "00:00:00", stopInterval: "00:00:00", expect: false},
		{startInterval: "00:00:00", stopInterval: "00:00:00", expect: false},
		{startInterval: "00:01:00", stopInterval: "00:00:00", expect: true},
		{startInterval: "00::00", stopInterval: "00:00:00", expect: true},
		{startInterval: "00:01:00", stopInterval: "00::00", expect: true},
	}

	for _, element := range cases {
		got, _ := IsIntervalStartOverIntervalStop(element.startInterval, element.stopInterval)

		if got == element.expect {
			t.Logf("%s - %s = %v; want %v", element.startInterval, element.stopInterval, got, element.expect)
		}else{
			t.Errorf("%s - %s = %v; want %v", element.startInterval, element.stopInterval, got, element.expect)
		}
	}

}
