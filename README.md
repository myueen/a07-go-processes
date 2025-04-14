# a07-go-processes
team("edlyn jeyaraj, yueen ma, thomas kung")

edlyn jeyaraj: 730569503 yueen ma: 730572152 thomas kung: 730620459


![whiteboard](<Screenshot 2025-04-13 at 9.33.40 PM.png>)

Description:

In this Sleeping Barber assignment, we implement four separate goroutines: customers, a waiting room, a receptionist, and a barber. We use a waitingRoom channel structure to manage customers waiting for haircuts.

We created four more channels to help communications between goroutines, including customerArrivesCh, barberWorkCh, barberSleepCh, and haircutCompleteCh. The customerArrivesCh helps send the customer pid between customer and receptionist. The barberWorkCh and barberSleepCh help communicate between waiting room andd barber goroutines. The haircutCompleteCh help send the customer pid whose haircut is done to customer. 

The goroutines works as follows:

New customers are created at random time intervals.

The receptionist goroutines handles each new customer by directing them to the waiting room if the waitingRoom channel contains fewer than 6 customers. If the waitingRoom channel is full (6 customers), the receptionist sends the customer away.

The waiting room goroutines manages the waitingRoom channel of waiting customers. The waiting room adds the new customer to the waitingRoom channel.


The barber goroutines handles haircuts:

When customers are waiting, the barber a) Cuts the customer's hair (taking a random amount of time) c) Send the customer pid to haircutComplete channel when their haircut is complete d) Checks with the waiting room if more customers are waiting

If the waitingRoom channel is empty, the barber goes to sleep.

If the waitingRoom channel is not empty, the barber continues serving the next customer.