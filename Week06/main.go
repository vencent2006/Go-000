package main

import (
	"flag"
)

var windowPeriod = flag.Int64("window_period", 10, "Input Slide Window Period")
var statPeriod = flag.Int("stat_period", 1, "Input Stat Display Period")
var testTimes = flag.Int("test_times", 20, "Input Test Times")

func main() {
	flag.Parse()

	counter := NewSlideWindow(*windowPeriod)

	go counter.Stat(*statPeriod)

	counter.TestIncr(*testTimes)
}
