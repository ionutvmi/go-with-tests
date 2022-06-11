package mocking

import (
	"fmt"
	"io"
	"time"
)

const startingPoint = 3
const finalWord = "GO!"

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := startingPoint; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep(1 * time.Second)
	}
	fmt.Fprintln(writer, finalWord)
}
