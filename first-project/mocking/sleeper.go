package mocking

import "time"

type Sleeper interface {
	Sleep(period time.Duration)
}

type DefaultSleeper struct{}

func (s DefaultSleeper) Sleep(period time.Duration) {
	time.Sleep(period)
}
