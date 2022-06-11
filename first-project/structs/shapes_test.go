package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	type args struct {
		r Rectangle
	}

	type testCase struct {
		name string
		args args
		want float64
	}

	tests := []testCase{
		{
			name: "should calculate the perimeter of a square",
			args: args{
				r: Rectangle{
					Width:  10.0,
					Height: 10.0,
				},
			},
			want: 40.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Perimeter(tt.args.r); got != tt.want {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	type args struct {
		r Rectangle
	}

	type testCase struct {
		name string
		args args
		want float64
	}

	tests := []testCase{
		{
			name: "Should calculate the area of a square",
			args: args{
				r: Rectangle{
					Width:  10.0,
					Height: 10.0,
				},
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Area(tt.args.r); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}
