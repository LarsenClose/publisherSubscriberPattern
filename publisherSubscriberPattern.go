/*
 * CS3210 - Principles of Programming Languages - Fall 2020
 * Instructor: Thyago Mota
 * Description: Prg04 - Publish Subscribe Simulation
 * Student(s) Name(s):  Larsen Close and Matt Hurt
 */

package main
import (
	"sync"
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

	if channel, matches := ps.topics[topic]; matches {
		for _, msg := range messages{
			channel <- msg
		}
	}	

	ps.mutex.Unlock()
}

// TODO: sends messages taken from a given array of message, one at a time and at random intervals, to all topic subscribers
func publisher(ps PubSub, topic string, msgs []string) {
	ps.mutex.Lock()

	for message := range msgs {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		if channel, matches := ps.topics[topic]; matches {
			ps.topics
			

				}
		}



	ps.mutex.Unlock()
}

// 	// func (ps *PubSub) StreetSubscribe() chan interface{} {
	// 	msgCh := make(chan interface{}, 5) // Create slice
	// 	ps.subCh <- msgCh
	// 	return msgCh
	// }

	// func NewPubSub() *PubSub {
	// 	return &PubSub{
	// 		stopCh:    make(chan struct{}),
	// 		publishCh: make(chan interface{}, 1),
	// 		subCh:     make(chan chan interface{}, 1),
	// 		unsubCh:   make(chan chan interface{}, 1),
	// 		topics:    make(map[string][]chan string),
	// 	}
	// }

	// // TODO: writes the given message on all the channels associated with the given topic.
	// func (ps *PubSub) Start() {
	// 	defer wg.Done()
	// 	subs := map[chan interface{}]struct{}{}
	// 	for {
	// 		select {
	// 		case <-ps.stopCh:
	// 			return
	// 		case msgCh := <-ps.subCh:
	// 			subs[msgCh] = struct{}{}
	// 		case msgCh := <-ps.unsubCh:
	// 			delete(subs, msgCh)
	// 		case msg := <-ps.publishCh:
	// 			for msgCh := range subs {
	// 				// msgCh is buffered, use non-blocking send to protect the PubSub:
	// 				select {
	// 				case msgCh <- msg:
	// 				default:
	// 				}
	// 			} // end case
	// 		} // end select
	// 	} // end for
	// } // end Start

	// // TODO: creates and returns a new channel on a given topic, updating the PubSub struct
	// func (ps *PubSub) StreetSubscribe() chan interface{} {
	// 	msgCh := make(chan interface{}, 5) // Create slice
	// 	ps.subCh <- msgCh
	// 	return msgCh
	// }

	// func (ps *PubSub) BeeSubscribe() chan interface{} {
	// 	msgCh := make(chan interface{}, 5) // Create slice
	// 	ps.subCh <- msgCh
	// 	return msgCh
	// }

	// func (ps *PubSub) Unsubscribe(msgCh chan interface{}) {
	// 	ps.unsubCh <- msgCh
	// }

	// func (ps *PubSub) StreetPublish(msg interface{}) {
	// 	ps.publishCh <- msg
	// }

	// func (ps *PubSub) BeePublish(msg interface{}) {
	// 	ps.publishCh <- msg
	// }

	// func main() {
	// 	// TODO: create the ps struct
	// 	// Create and start a PubSub:
	// 	ps := NewPubSub()
	// 	//pub1 := make(chan string)
	// 	//pub2 := make(chan string)
	// 	//
	// 	//sub1 := make(chan string)
	// 	//sub2 := make(chan string)
	// 	//sub3 := make(chan string)

	// 	// TODO: create the subscriber goroutines
	// 	go ps.Start()

	// 	// TODO: create the arrays of messages to be sent on each topic
	// 	// Create and subscribe 3 clients:

	// 	subscriber := func(name string, topic string) {
	// 		msgCh := ps.BeeSubscribe()

	// 		if topic == "streets" {
	// 			for {
	// 				fmt.Printf("* %s got message: %v\n", name, <-msgCh)
	// 				//time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	// 			}
	// 		} else if topic == "bees" {
	// 			for {
	// 				fmt.Printf("* %s got message: %v\n", name, <-msgCh)
	// 				//time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	// 			}
	// 		}

	// 	}

	// 	go subscriber("Marry", "bees")
	// 	go subscriber("Tom", "streets")

	// 	//go publish(beefacts)

	// 	// TODO: set wait group to 2 (# of publishers)
	// 	// TODO: create the publisher goroutines
	// 	// Start publishing messages:
	// 	go func() {
	// 		for publisher := 0; ; publisher++ {
	// 			//time.Sleep(200 * time.Millisecond)
	// 			ps.mutex.Lock()

	// 			ps.BeePublish(fmt.Sprintf("msg#%d", publisher))
	// 			//time.Sleep(200 * time.Millisecond)
	// 			ps.BeePublish("bees are pollinators.")
	// 			//time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	// 			//time.Sleep(200 * time.Millisecond)
	// 			ps.BeePublish("bees produce honey")
	// 			//time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	// 			time.Sleep(200 * time.Millisecond)
	// 			ps.BeePublish("all worker bees are female")
	// 			//time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	// 			time.Sleep(200 * time.Millisecond)
	// 			ps.BeePublish("bees have 5 eyes,")
	// 			//time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	// 			time.Sleep(200 * time.Millisecond)
	// 			ps.BeePublish("bees fly about 20mph")
	// 			//time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	// 			time.Sleep(200 * time.Millisecond)
	// 			ps.StreetPublish("Streets are cool")
	// 			//time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	// 			time.Sleep(200 * time.Millisecond)
	// 			ps.mutex.Unlock()
	// 		}
	// 	}()
	// 	// TODO: wait for all publishers to be done
	// 	time.Sleep(time.Second)
}
