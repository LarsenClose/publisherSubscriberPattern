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

// PubSub, struct for the publisher-subscriber pattern
type PubSub struct {
	topics map[string][]chan string
	mu     sync.Mutex
}

// Creating the wait group variable
var wg sync.WaitGroup

// creates and returns a new channel on a given topic, updating the PubSub struct
func (ps *PubSub) subscribe(topic string) chan string {
	ps.mu.Lock()
	ch := make(chan string)
	ps.topics[topic] = append(ps.topics[topic], ch)
	ps.mu.Unlock()
	return ch
}

// writes the given message on all the channels associated with the given topic
func (ps *PubSub) publish(topic string, msg string) {
	ps.mu.Lock()
	for _, ch := range ps.topics[topic] {
		go func(channel chan string) {
			channel <- msg
		}(ch)
	}
	ps.mu.Unlock()
}

// sends messages taken from a given array of message, one at a time and at random intervals, to all topic subscribers
func (ps *PubSub) publisher(topic string, msgs [10]string) {
	for message := range msgs {
		ps.publish(topic, msgs[message])
		time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	}
	wg.Done()
}

// reads and displays all messages received from a particular topic
func (ps *PubSub) subscriber(name string, topic string) {
	ch := ps.subscribe(topic)
	for range ps.topics[topic] {
		go printer(name, topic, ch)
	}
}

// printer for channels, takes sunscriber name, topic and the channel
func printer(name string, topic string, channel chan string) {
	for {
		var recieve = <-channel
		fmt.Printf("Subcriber: %v, Topic: %v; Message: %v\n\n", name, topic, recieve)
	}
}

func main() {

	ps := PubSub{topics: make(map[string][]chan string)}

	// TODO: create the arrays of messages to be sent on each topic
	var beesArray [10]string

	beesArray[0] = "bees are pollinators"
	beesArray[1] = "bees produce honey"
	beesArray[2] = "all worker bees are female"
	beesArray[3] = "bees have 5 eyes"
	beesArray[4] = "bees fly about 20mph"
	beesArray[5] = "Bees are insects, so they have 6 legs."
	beesArray[6] = "Number of eggs laid by queen: 2,000 per day is the high."
	beesArray[7] = "Bees have 2 pairs of wings"
	beesArray[8] = "Male bees in the hive are called drones"
	beesArray[9] = "The average forager makes about 1/12 th of a teaspoon of honey in her lifetime"

	var philosophyArray [10]string

	philosophyArray[0] = "And if you gaze for long into an abyss, the abyss gazes also into you."
	philosophyArray[1] = "For all practical purposes, the universe is a pattern generator, and the mind 'makes sense' of these patterns by encoding them according to the regularities it can find. Thus, the representation of a concept in an intelligent system is not a pointer to a 'thing in reality', but a set of hierarchical constraints over (for instance perceptual) data"
	philosophyArray[2] = "The question is not what you look at, but what you see."
	philosophyArray[3] = "Go confidently in the direction of your dreams! Live the life you’ve imagined."
	philosophyArray[4] = "Rather than love, than money, than fame, give me truth."
	philosophyArray[5] = "What lies behind us and what lies before us are tiny matters compared to what lies within us."
	philosophyArray[6] = "Is it so bad, then, to be misunderstood? Pythagoras was misunderstood, and Socrates, and Jesus, and Luther, and Copernicus, and Galileo, and Newton, and every pure and wise spirit that ever took flesh. To be great is to be misunderstood."
	philosophyArray[7] = "Be thyself! all that thou doest and thinkest and desirest, is not thyself!\n Every youthful soul hears this cry day and night, and quivers to hear it [...]"
	philosophyArray[8] = "We are what we repeatedly do. Excellence, then, is not an act, but a habit."
	philosophyArray[9] = "Some people think that a simulation can’t be conscious and only a physical system can. But they got it completely backward: a physical system cannot be conscious. Only a simulation can be conscious. Consciousness is a simulated property of the simulated self."

	// wait group 2  publishers)
	wg.Add(2)

	// publisher goroutines

	go ps.publisher("bees", beesArray)
	go ps.publisher("philosophy", philosophyArray)

	// subscriber goroutines
	go ps.subscriber("nature", "bees")
	go ps.subscriber("humanities", "philosophy")

	// wait for all publishers to be done

	wg.Wait()

}
