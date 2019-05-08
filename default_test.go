package govals

import "testing"

type testCase struct {
	name string
	val  string
	err  error
}

func TestEmail(t *testing.T) {
	tc := []testCase{
		{name: "valid email", val: "ludyyn@gmail.com"},
		{name: "invalid email", val: "ludyyn$$##gmail.com", err: errInvalidFormat},
		{name: "empty email", val: "", err: errEmptyValue},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := Email(tt.val)
			if tt.err != nil {
				if tt.err != err && !ok {
					t.Errorf("error email expects %v, got %v\n", tt.err, err)
				}
			}
		})
	}

}

func TestNumeric(t *testing.T) {
	tc := []testCase{
		{name: "valid numeric", val: "122323"},
		{name: "invalid numeric", val: "23dsdfdfsdf", err: errInvalidFormat},
		{name: "empty numeric", val: "", err: errEmptyValue},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := Numeric(tt.val)
			if tt.err != nil {
				if tt.err != err && !ok {
					t.Errorf("error numeric expects %v, got %v\n", tt.err, err)
				}
			}
		})
	}

}

func TestAlpha(t *testing.T) {
	tc := []testCase{
		{name: "valid alpha", val: "Ludin Nento"},
		{name: "invalid alpha", val: "sdfnj23434", err: errInvalidFormat},
		{name: "empty alpha", val: "", err: errEmptyValue},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := Alpha(tt.val)
			if tt.err != nil {
				if tt.err != err && !ok {
					t.Errorf("error alpha expects %v, got %v\n", tt.err, err)
				}
			}
		})
	}

}
func TestAlphaNumeric(t *testing.T) {
	tc := []testCase{
		{name: "valid alpha numeric", val: "Ludin Nento2323"},
		{name: "invalid alpha numeric", val: "sdfnj23434^^&9", err: errInvalidFormat},
		{name: "empty alpha numeric", val: "", err: errEmptyValue},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := AlphaNumeric(tt.val)
			if tt.err != nil {
				if tt.err != err && !ok {
					t.Errorf("error alpha numeric expects %v, got %v\n", tt.err, err)
				}
			}
		})
	}

}

func TestDate(t *testing.T) {
	tc := []testCase{
		{name: "valid date", val: "12-12-1996"},
		{name: "invalid date", val: "12-", err: errInvalidFormat},
		{name: "empty date", val: "", err: errEmptyValue},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := Date(tt.val)
			if tt.err != nil {
				if tt.err != err && !ok {
					t.Errorf("error date expects %v, got %v\n", tt.err, err)
				}
			}
		})
	}

}

func TestTimes(t *testing.T) {
	tc := []testCase{
		{name: "valid time", val: "12:12:01"},
		{name: "invalid time", val: "12:", err: errInvalidFormat},
		{name: "empty time", val: "", err: errEmptyValue},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := Times(tt.val)
			if tt.err != nil {
				if tt.err != err && !ok {
					t.Errorf("error time expects %v, got %v\n", tt.err, err)
				}
			}
		})
	}

}
func TestPhoneNumber(t *testing.T) {
	tc := []testCase{
		{name: "valid phone number", val: "+6282290202728"},
		{name: "invalid phone number", val: "12--&&809098^^%$:", err: errInvalidFormat},
		{name: "empty phone number", val: "", err: errEmptyValue},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := PhoneNumber(tt.val)
			if tt.err != nil {
				if tt.err != err && !ok {
					t.Errorf("error phone number expects %v, got %v\n", tt.err, err)
				}
			}
		})
	}

}
