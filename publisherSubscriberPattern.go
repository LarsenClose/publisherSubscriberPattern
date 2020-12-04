/*
 * CS3210 - Principles of Programming Languages - Fall 2020
 * Instructor: Thyago Mota
 * Description: Prg04 - Publish Subscribe Simulation
 * Student(s) Name(s):  Larsen Close and Matt Hurt
 */

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// PubSub is our struct for the publisher-subscriber pattern
type PubSub struct {
	topics map[string][]chan string
	mu sync.RWMutex
}

func NewPubSub() *PubSub{
	ps := &PubSub{}
	ps.topics = make(map[string][]chan string)
	return ps
}


var wg sync.WaitGroup

// TODO: creates and returns a new channel on a given topic, updating the PubSub struct
func (ps *PubSub) subscribe(topic string, ch chan string) chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.topics[topic] = append(ps.topics[topic], ch)
	return ch
}

// TODO: writes the given message on all the channels associated with the given topic
func (ps PubSub) publish(topic string, msg string) {
	for _, ch := range ps.topics[topic] {
		ch <- msg
	}
}

// TODO: sends messages taken from a given array of message, one at a time and at random intervals, to all topic subscribers
func (ps *PubSub) publisher( topic string, msgs [5]string) {
	ps.mu.RLock()
	defer ps.mu.Unlock()

	for message := range msgs {
		ps.publish(topic, msgs[message])
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
}


// TODO: reads and displays all messages received from a particular topic
func (ps *PubSub) subscriber(name string, topic string) {

	ps.mu.Lock()
	defer ps.mu.Unlock()

	for topic, slice := range ps.topics {
		go printer(topic, slice)

	}
}

func printer(topic string, channel []chan string,) {
	for item := range channel {
		var receive = <-channel[item]
		fmt.Printf("Topic: %v; Message: %v\n", topic, receive)
	}
}

func main() {

	// TODO: create the ps struct

	var ps PubSub

	// TODO: create the arrays of messages to be sent on each topic
	var beesArray [5]string

	beesArray[0] = "bees are pollinators"
	beesArray[1] = "bees produce honey"
	beesArray[2] = "all worker bees are female"
	beesArray[3] = "bees have 5 eyes"
	beesArray[4] = "bees fly about 20mph"

	var philosophyArray [5]string

	philosophyArray[0] = "And if you gaze for long into an abyss, the abyss gazes also into you."
	philosophyArray[1] = "Dreams are the touchstones of our characters."
	philosophyArray[2] = "The question is not what you look at, but what you see."
	philosophyArray[3] = "Go confidently in the direction of your dreams! Live the life you’ve imagined."
	philosophyArray[4] = "Rather than love, than money, than fame, give me truth."

	// TODO: set wait group to 2 (# of publishers)
	wg.Add(2)

	// TODO: create the publisher goroutines

	go ps.publisher("bees", beesArray)
	go ps.publisher("philosophy", philosophyArray)


	// TODO: create the subscriber goroutines
	go ps.subscriber("nature","bees")
	go ps.subscriber("humanities", "philosophy")

	// TODO: wait for all publishers to be done
*
	wg.Wait()



}
