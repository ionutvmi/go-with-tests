package helloworld

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		name string
	}

	type testCase struct {
		name string
		args args
		want string
	}

	tests := []testCase{
		{
			name: "should show use the name if it was provided",
			args: args{name: "Mihai"},
			want: "Hello, Mihai",
		},
		{
			name: "should say hello world if no name was provided",
			args: args{name: ""},
			want: "Hello, World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.name); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
