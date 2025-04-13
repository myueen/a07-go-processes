package main

import (
	"fmt"
	"time"
)

// Customer PID
var pid int = 0

// channels
var customer_arrive_ch = make(chan int)
var waiting_room_ch = make(chan int, 6) // buffered channel for 6 customers
var barber_cut_hair_ch = make(chan int)

func main() {
	// Make waiting room, receptionist, barber
	go waiting_room()
	go receptionist()
	go barber()

	// Customer arriving
	for i := 0; i >= 0; i++ {
		customer_pid := pid
		pid++
		fmt.Printf("Greeting Customer %d \n", customer_pid)
		go customer(customer_pid)
		time.Sleep(1 * time.Second)
	}
}

func customer(customer_pid int) {
	// Messages (wait, full, done)
	// When arrive, sent to the receptionist
	customer_arrive_ch <- customer_pid

}

func receptionist() {
	fmt.Printf("Receptionist created")
	for {
		var customer_pid = <-customer_arrive_ch
		fmt.Printf("Greeting %d \n", customer_pid)

		select {
		case waiting_room_ch <- customer_pid:
			fmt.Printf("Customer %d sent to waiting room \n", customer_pid)
		default:
			fmt.Printf("Waiting room full, customer %d sent away \n", customer_pid)
		}
	}
	// Messages (receive_customer, send_customer_to_waiting_room, tell_barber_added, tell_customer_to_wait)

}

func waiting_room() {
	fmt.Printf("Waiting room created")
	// Messages (receive_customer, enqueue, dequeue, check_empty)

}

func barber() {
	fmt.Printf("Barber created")
	// Messages (receive_customer, cut_hair_random_time, tell customer_done)

}
