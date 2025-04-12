package main

import (
	"fmt"
	"sync"
	"time"
)

//

// Customer PID
var pid int = 0

// Shared variable for communication
var sharedData string

// Lock
var mu_r sync.Mutex
var mu_wr sync.Mutex
var mu_b sync.Mutex
var mu_c sync.Mutex

// Customer arriving
var customerArrives bool

// waiting room full or not
var waitingRoomFull bool

// receive signal
var receptionist_rb string
var waiting_room_rb string
var barber_rb string
var customer_rb string

// make slice to store customers in the waiting room
var sl = make([]int, 3)

func make_waiting_room() {
	fmt.Printf("Waiting room created")
	waiting_room()
}

func make_receptionist() {
	fmt.Printf("Receptionist created")
	receptionist()
}

func make_barber() {
	fmt.Printf("Barber created")
	barber()
}

func new_customer_loop(pid int) {
	fmt.Printf("Greeting Customer # %d \n", pid)
	go new_customer_loop(pid + 1)

	// set random number generator
	// time.Now().UnixNano()
	// randomNumber := rand.Intn(100)
	// time.Sleep(100 * time.Duration(randomNumber))
	// fmt.Println("Random number:", randomNumber)

}

func main() {
	// Make waiting room, receptionist, barber
	go make_waiting_room()
	go make_receptionist()
	go make_barber()

	// Customer arriving
	customerArrives = true
	waitingRoomFull = false
	go new_customer_loop(pid)
}

func receptionist() {
	for customerArrives {
		mu_r.Lock()
		if receptionist_rb == "receptionist" {
			fmt.Printf("Receptionist send customer # %d to waiting room", pid)
			// add customer to waiting room
			waiting_room_rb = "enqueue"
		}
		mu_r.Unlock()
	}
}

func waiting_room() {
	for i := 1; i > 0; i++ {
		mu_wr.Lock()
		if waiting_room_rb == "enqueue" {
			if len(sl) >= 6 {
				customer_rb = "full"
				waitingRoomFull = true
			} else {
				customer_rb = "added"
				sl[pid] = pid
			}
		} else if waiting_room_rb == "dequeue" {
			barber_rb = "customer getting haircut"
		} else if waiting_room_rb == "checkEmpty" {
			barber_rb = "no customer"

		}
		mu_wr.Unlock()
	}
}

func barber() {
	for !waitingRoomFull {
		mu_b.Lock()
		if barber_rb == "customer getting haircut" {
			fmt.Printf("Barber is cutting hair for customer # %d to waiting room", pid)
		} else if barber_rb == "no customer" {
			time.Sleep(5000)
		}
		mu_b.Unlock()
	}
}
