package govals

import "testing"

type structValid struct {
	Name           string `govals:"len, max=22, min=10"`
	Age            int8   `govals:"len, min=2, max=20"`
	FavoriteNumber uint   `govals:"len, min=2, max=20"`
	Email          string `govals:"email"`
	Date           string `govals:"date"`
	Times          string `govals:"time"`
	Hoby           string `govals:"alpha"`
	Max            string `govals:"numeric"`
	Text           string `govals:"alphaNumeric"`
	Phone          string `govals:"phone"`
	Intrest        string `govals:"-"`
	Address        string `govals:"custom, regex=^[a-zA-Z]*$"`
}

func TestStruct(t *testing.T) {
	validStruct := &structValid{
		Name:           "Ludin Nento",
		Age:            18,
		Email:          "ludyyn@gmail.com",
		FavoriteNumber: 10,
		Times:          "12:00:01",
		Date:           "12-31-1996",
		Hoby:           "Music",
		Max:            "233330",
		Phone:          "+6282290102726",
		Text:           "Ludin123",
		Intrest:        "Football",
		Address:        "Jlirigasi",
	}

	fileds := []string{"Name", "Email", "Age", "Address", "FavoriteNumber", "Times", "Date", "Hoby", "Max", "Phone", "Text"}

	tc := []struct {
		name  string
		value interface{}
		field []string
		err   string
	}{
		{name: "empty struct", value: &structValid{}, field: fileds, err: "invalid struct validation format"},
		{name: "valid struct", value: validStruct, field: fileds},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			str := Struct(tt.value)
			for _, v := range tt.field {
				ok, err := str.GetStatus(v)
				if tt.err != "" {
					if err == nil && ok {
						t.Errorf("%s error expects %s got %v\n", v, tt.err, err)
					}
					return
				}
				if !ok {
					t.Errorf("%s invalid result", v)
				}
			}

		})
	}

}
