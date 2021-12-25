package main

import "testing"

func Test_validateCountry(t *testing.T) {
	type args struct {
		s           string
		countryCode string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Ländercode 01",
			args: args{"756.3903.6825.80", "756"},
			want: true,
		},
		{
			name: "Ländercode 02",
			args: args{"aaa.3903.6825.80", "756"},
			want: false,
		},
		{
			name: "Ländercode 03",
			args: args{"666.3903.6825.80", "756"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateCountry(tt.args.s, tt.args.countryCode); got != tt.want {
				t.Errorf("validateCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}
