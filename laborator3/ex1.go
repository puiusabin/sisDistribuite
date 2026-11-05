package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	delta = t4.Sub(t1) - t3.Sub(t2)
	theta = t3.Sub(t4) + delta/2
	return theta, delta
}

func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t2 := s_clock.Now()

	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)
	t3 := s_clock.Now()

	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()
	return estimate(t1, t2, t3, t4)
}

func main() {
	server := &Clock{offset: 0}

	client := &Clock{offset: -500 * time.Millisecond}
	theta, delta := measure(client, server)
	client.offset += theta
	fmt.Println("a) theta:", theta, "delta:", delta, "offset:", client.offset)

	client = &Clock{offset: -500 * time.Millisecond}
	minTheta, minDelta := measure(client, server)
	thetaMinDelta := minTheta

	for i := 2; i <= 10; i++ {
		theta, delta := measure(client, server)

		if theta < minTheta {
			minTheta = theta
		}
		if delta < minDelta {
			minDelta = delta
			thetaMinDelta = theta
		}
	}

	client.offset += thetaMinDelta
	fmt.Println("b) theta:", minTheta, "delta:", minDelta, "offset:", client.offset)
}
