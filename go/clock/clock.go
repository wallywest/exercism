package clock

import "fmt"

const testVersion = 4

type Clock int

func New(hour, minute int) Clock {

	//https://github.com/golang/go/issues/448
	//coming from ruby this returns a negative time value which is odd
	time := (hour*60 + minute) % 1440

	if time <= 0 {
		time += 1440
	}

	return Clock(time)
}
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", (c/60)%24, c%60)
}

func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}
