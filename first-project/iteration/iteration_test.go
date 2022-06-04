package iteration

import "testing"

func TestRepeat(t *testing.T) {
	type args struct {
		c           string
		repeatCount int
	}
	type testCase struct {
		name string
		args args
		want string
	}

	tests := []testCase{
		{
			name: "Should repeat the character 5 times",
			args: args{
				c:           "a",
				repeatCount: 5,
			},
			want: "aaaaa",
		},
		{
			name: "Should repeat the substring 5 times",
			args: args{
				c:           "ab",
				repeatCount: 5,
			},
			want: "ababababab",
		},
		{
			name: "Should return an empty string if it's repeated 0 times",
			args: args{
				c:           "ab",
				repeatCount: 0,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.c, tt.args.repeatCount); got != tt.want {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
