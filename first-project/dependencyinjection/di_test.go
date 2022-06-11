package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreetWithDI(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
	}{
		{
			name: "should greet with the person name",
			args: args{
				name: "Mihai",
			},
			wantWriter: "Hello, Mihai!\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			GreetWithDI(writer, tt.args.name)
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("GreetWithDI() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
