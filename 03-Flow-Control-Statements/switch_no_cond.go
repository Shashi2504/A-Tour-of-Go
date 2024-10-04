package main

import (
	"fmt"
	"time"
)

func main() {
	timeuhhhh := time.Now()
	switch {
	case timeuhhhh.Hour() < 12:
		fmt.Println("Your completion of task has been completed")

	case timeuhhhh.Hour() < 16:
		fmt.Println("Your next task will be started in the afternoon session. Be ready")

	default:
		fmt.Println("The tasks day has been completed. Thank you for attending the task")
	}
}
