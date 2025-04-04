package main 

import (
	"fmt"
	"time"
)

// wr: waiting room, r: receptionist, b: barber
func new_wr_r_b() {
	
}

func new_customer_loop() {

}

func main() {
	// initialze 
	go new_wr_r_b()

	fmt.Print("Hello")

	time.Sleep(3100)
}


