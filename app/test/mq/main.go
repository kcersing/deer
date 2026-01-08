package main

import "time"

func main() {
	InitMQ()

	for {
		time.Sleep(1 * time.Second)
	}
}
