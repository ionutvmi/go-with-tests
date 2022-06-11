package structs

import (
	"math"
	"testing"
)

func TestRectangle_Perimeter(t *testing.T) {
	type testCase struct {
		name string
		r    Rectangle
		want float64
	}

	tests := []testCase{
		{
			name: "should calculate the perimeter of a square",
			r: Rectangle{
				Width:  10.0,
				Height: 10.0,
			},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Perimeter(); got != tt.want {
				t.Errorf("Rectangle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Area(t *testing.T) {
	type testCase struct {
		name string
		r    Rectangle
		want float64
	}

	tests := []testCase{
		{
			name: "should calculate the area of a square",
			r:    Rectangle{10, 10},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Area(); got != tt.want {
				t.Errorf("Rectangle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Area(t *testing.T) {
	type testCase struct {
		name string
		c    Circle
		want float64
	}

	tests := []testCase{
		{
			name: "should calculate the area of a square",
			c:    Circle{10},
			want: 314.1593,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roundFloat(tt.c.Area(), 4); got != tt.want {
				t.Errorf("Circle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
