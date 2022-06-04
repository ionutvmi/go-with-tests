package integers

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}

	type testCase struct {
		name string
		args args
		want int
	}

	tests := []testCase{
		{
			name: "should add two positive numbers",
			args: args{a: 4, b: 5},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
