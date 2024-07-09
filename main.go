package main

import "fmt"

func main() {

	//Deadlock conditon : This is because the program is stuck on sending a value to the channel: mychannel <- 10.
	//The sending operation is a blocking operation and requires the receive channel to be ready before sending data to the channel.
	mychannel := make(chan int)
	mychannel <- 10
	fmt.Println(<-mychannel)
}
