package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getReady(user string) {
	t := random(60, 90)
	fmt.Printf("%s started getting ready.\n", string(user))
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Printf("%s spent %d seconds getting ready.\n", string(user), t)
}
func putOnShoe(user string) {

	t := random(35, 45)
	fmt.Printf("%s started putting on shoes.\n", string(user))
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Printf("%s spent %d seconds putting on shoes.\n", string(user), t)
}

func armAlarm(alarmStarted, alarmFinished chan struct{}) {
	fmt.Println("Arming alarm.")
	close(alarmStarted)
	time.AfterFunc(60*time.Second, func() {
		fmt.Println("Alarm is armed.")
		close(alarmFinished)
	})
	fmt.Println("Alarm is couting down.")
}

var users = [2]string{"Alice", "Bob"}

func main() {
	var waitForReady, waitForPutOnShoe sync.WaitGroup
	alarmStarted, alarmFinished := make(chan struct{}), make(chan struct{})
	fmt.Println("Let's go for a walk.")
	waitForReady.Add(2)
	waitForPutOnShoe.Add(2)
	for _, user := range users {
		go func(user string) {
			getReady(user)
			waitForReady.Done()
			<-alarmStarted
			putOnShoe(user)
			waitForPutOnShoe.Done()
		}(user)
	}
	waitForReady.Wait()
	armAlarm(alarmStarted, alarmFinished)
	waitForPutOnShoe.Wait()
	fmt.Println("Exiting and locking the door.")
	<-alarmFinished
}
func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
