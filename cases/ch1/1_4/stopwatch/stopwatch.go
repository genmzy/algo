package stopwatch

import "time"

type StopWatch struct {
	start time.Time
}

func (sw StopWatch) ElapsedTime() time.Duration {
	return time.Now().Sub(sw.start)
}

func NewStopWatch() StopWatch {
	return StopWatch{
		start: time.Now(),
	}
}
