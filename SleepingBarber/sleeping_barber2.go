package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Customer PID
var pid int = 0

// channels
var customerArrivesCh = make(chan int)
var waitingRoomCh = make(chan int, 6) // buffered channel for 6 customers
var barberWorkingCh = make(chan int)
var barberSleepCh = make(chan bool)
var haircutComplete = make(chan int)

func main() {
	// Make waiting room, receptionist, barber
	go waiting_room()
	go receptionist()
	go barber()

	// Customer arriving
	for i := 0; i >= 0; i++ {
		customer_pid := pid
		pid++
		go customer(customer_pid)
		r := rand.Intn(10)
		time.Sleep(time.Duration(r) * time.Second)
	}
}

func customer(customer_pid int) {
	// Messages (wait, full, done)
	// When arrive, sent to the receptionist
	customerArrivesCh <- customer_pid
	fmt.Printf("Customer %d arrives \n", customer_pid)

	// Check if customer is done
	for {
		var pid_done = <-haircutComplete
		if customer_pid == pid_done {
			fmt.Printf("Customer %d haircut complete \n", customer_pid)
		} else {
			haircutComplete <- pid_done
		}
	}
}

func receptionist() {
	fmt.Printf("Receptionist created")
	for {
		var customer_pid = <-customerArrivesCh
		fmt.Printf("Greeting %d \n", customer_pid)

		if len(waitingRoomCh) < 6 {
			waitingRoomCh <- customer_pid
			fmt.Printf("Customer %d sent to waiting room \n", customer_pid)
		} else {
			fmt.Printf("Waiting room full, customer %d sent away \n", customer_pid)
		}
	}
	// Messages (receive_customer, send_customer_to_waiting_room, tell_barber_added, tell_customer_to_wait)
}

func waiting_room() {
	fmt.Printf("Waiting room created")
	// Messages (receive_customer, enqueue, dequeue, check_empty)
	for {
		var customer_pid = <-waitingRoomCh
		barberWorkingCh <- customer_pid
	}

}

func barber() {
	fmt.Printf("Barber created")
	// Messages (receive_customer, cut_hair_random_time, tell customer_done)
	for {
		var customer_pid = <-barberWorkingCh
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("Customer %d sent to barber \n", customer_pid)
		r := rand.Intn(20)
		time.Sleep(time.Duration(r) * time.Second)
		haircutComplete <- customer_pid
		time.Sleep(10 * time.Millisecond)
		if len(waitingRoomCh) == 0 {
			fmt.Printf("Barber sleeping\n")
		}
	}
}
