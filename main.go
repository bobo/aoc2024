package main

import (
	"adventOfCode/day1"
	"adventOfCode/day2"
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func main() {
	RunTimed(day1.Run)
	RunTimed(day2.Run)
}

func RunTimed(f func()) {
	startTime := time.Now()
	f()
	duration := time.Since(startTime)

	funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	fmt.Printf("Function %s duration: %s\n", funcName, duration)

}
