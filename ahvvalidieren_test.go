package ahvvalidieren

import (
	"testing"
)

func TestValidate(t *testing.T) {
	type args struct {
		ahvnr string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		// TODO: Add test cases.
		{
			name:    "Ländercode 01",
			args:    args{"756.3903.6825.80"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Ländercode 02",
			args:    args{"666.3903.6825.80"},
			want:    false,
			wantErr: true,
		}, {
			name:    "Ländercode 03",
			args:    args{"aaa.3903.6825.aa"},
			want:    false,
			wantErr: true,
		}, {
			name:    "Ländercode 04",
			args:    args{"666.3903.6444825.80"},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Validate(tt.args.ahvnr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
