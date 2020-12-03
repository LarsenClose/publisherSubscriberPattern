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
	mu     sync.Mutex
	topics map[string][]chan string
}

var wg sync.WaitGroup

// TODO: creates and returns a new channel on a given topic, updating the PubSub struct
func (ps *PubSub) subscribe(topic string) chan string {
	ps.mutex.Lock()

	if previous, add := ps.topics[topic]; add{
		ps.topics[topic] = append(previous, chan topic)
	} else {
		ps.topics[topic] = append([] chan topic)
	}
	return ps.topics[topic]
	
	ps.mutex.Unlock()
}

// TODO: writes the given message on all the channels associated with the given topic
func (ps PubSub) publish(topic string, msg string) {
	ps.mutex.Lock()

	if channel, associated := ps.topics[topic]; associated {
		for _, msg := range messages{
			channel <- msg
		}
	}	

	ps.mutex.Unlock()
}

// TODO: sends messages taken from a given array of message, one at a time and at random intervals, to all topic subscribers
func publisher(ps PubSub, topic string, msgs []string) {
	ps.mutex.Lock()

	for message  {
		ps.publish(topic, msg)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
	
	ps.mutex.Unlock()
}


// TODO: reads and displays all messages received from a particular topic
func subscriber(ps PubSub, name string, topic string) {

	for channel, topic := range ps.topics {
			if channel, matches :=  ps.topics[topic]; matches {
					printer(channel, topic)
			}
		}
}

func printer(ch string, topic string) {
	fmt.Printf("Channel: %s; Topic: %s\n", ch, topic)
}



func main() {

	// TODO: create the ps struct
	// Create and start a PubSub:
	ps := NewPubSub()
	

	// TODO: create the arrays of messages to be sent on each topic

	"bees are pollinators" "bees produce honey" 
	"all worker bees are female" "all worker bees are female"
	"bees have 5 eyes" "bees fly about 20mph"

	// TODO: set wait group to 2 (# of publishers)

	go publish("topic1", "Hi topic 1")
	go publish("topic2", "Welcome to topic 2")

	// TODO: create the publisher goroutines

	// TODO: create the subscriber goroutines

	// TODO: wait for all publishers to be done
	ch1 := make(chan Data)
	ch2 := make(chan Data)
	ch3 := make(chan Data)

	ps.subscribe("topic1", ch1)
	ps.subscribe("topic2", ch2)
	ps.subscribe("topic2", ch3)


	for {
		select {
		case d := <-ch1:
			go printer("ch1", d)
		case d := <-ch2:
			go printer("ch2", d)
		case d := <-ch3:
			go printer("ch3", d)
		}
	}
}
