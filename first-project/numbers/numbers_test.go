package numbers

import "testing"

func Test_getNumber(t *testing.T) {
	type testCase struct {
		name    string
		want    int
		wantErr bool
	}

	tests := []testCase{
		{
			name:    "get a number",
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getNumber()
			if (err != nil) != tt.wantErr {
				t.Errorf("getNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
