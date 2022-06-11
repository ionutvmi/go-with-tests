package mocking

import (
	"bytes"
	"testing"
	"time"
)

type SleeperSpy struct {
	Calls int
}

func (s *SleeperSpy) Sleep(period time.Duration) {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	type args struct {
		sleeper *SleeperSpy
	}
	tests := []struct {
		name        string
		args        args
		wantWriter  string
		wantSleeper int
	}{
		{
			name: "should show the countdown messages",
			args: args{
				sleeper: &SleeperSpy{},
			},
			wantWriter:  "3\n2\n1\nGO!\n",
			wantSleeper: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			Countdown(writer, tt.args.sleeper)
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("Countdown() = %v, want %v", gotWriter, tt.wantWriter)
			}
			if tt.args.sleeper.Calls != tt.wantSleeper {
				t.Errorf("Countdown > Sleep() = %v, want %v", tt.args.sleeper.Calls, tt.wantSleeper)
			}
		})
	}
}
