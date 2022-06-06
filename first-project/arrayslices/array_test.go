package arrayslices

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		numbers []int
	}

	type testCase struct {
		name string
		args args
		want int
	}

	tests := []testCase{
		{
			name: "should return the sum of the provided numbers",
			args: args{
				numbers: []int{2, 3},
			},
			want: 5,
		},
		{
			name: "should return 0 if no numbers are provided",
			args: args{
				numbers: []int{},
			},
			want: 0,
		},
		{
			name: "should return the same value if we have only one element in the list",
			args: args{
				numbers: []int{99},
			},
			want: 99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.numbers); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
