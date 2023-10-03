package counter

import (
	"io"
	"fmt"
	"time"
)

const finalWord = "Go!"
const CountdownStart = 3

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct {}

type DelayedWriter struct {
	Writer io.Writer
	Sleeper Sleeper
}

func (dw *DelayedWriter) Write(value any) {
	fmt.Fprintln(dw.Writer, value)
	dw.Sleeper.Sleep()
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, s Sleeper) {
	delayedWriter := DelayedWriter{out, s}

	for i := CountdownStart; i > 0; i-- {
		delayedWriter.Write(i)
	}

	fmt.Fprint(out, finalWord)
}
