package arrayslices

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	type args struct {
		args [][]int
	}
	type testCase struct {
		name string
		args args
		want []int
	}

	tests := []testCase{
		{
			name: "should sum each list and return the result",
			args: args{
				args: [][]int{
					{2, 3},
					{4, 7},
				},
			},
			want: []int{5, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumAll(tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumTail(t *testing.T) {
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
			name: "Should gives us the sum of all elements except the first one",
			args: args{
				numbers: []int{99, 2, 3},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumTail(tt.args.numbers); got != tt.want {
				t.Errorf("SumTail() = %v, want %v", got, tt.want)
			}
		})
	}
}
