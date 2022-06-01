package numbers

import (
	"math/rand"
	"testing"
)

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

// This function compares the stdout with the value provided in the output comment.
// An elegant solution to allow this kind of testing.
func ExampleShowNumber() {
	ShowNumber()
	// Output: The number is 10
}

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}
