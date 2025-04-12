package main

import "fmt"

// channels
var ch_c = make(chan string)
var ch_r = make(chan string)
var ch_wr = make(chan string)
var ch_b = make(chan string)

func new_customer_loop() {
	fmt.Printf("Greeting Customer: new customer arrives\n")
	go customer()
	go new_customer_loop()
}

func main() {
	// Make waiting room, receptionist, barber
	go waiting_room()
	go receptionist()
	go barber()

	// Customer arriving
	go new_customer_loop()
}

func customer() {
	// Messages (wait, full, done)
	// When arrive, sent to the receptionist
	ch_r <- ""
	for {
		ch_c
	}

}

func receptionist() {
	fmt.Printf("Receptionist created")

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
