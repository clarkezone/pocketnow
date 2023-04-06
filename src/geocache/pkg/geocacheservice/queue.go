package geocacheservice

import (
	"sync"

	clarkezoneLog "github.com/clarkezone/boosted-go/log"
)

type Message struct {
	ID        int
	Content   string
	locations *Locations
}

type MessageProcessor interface {
	Process(msg Message)
}

type Queue struct {
	channel          chan Message
	wg               sync.WaitGroup
	messageProcessor MessageProcessor
}

func NewQueue(size int, messageProcessor MessageProcessor) *Queue {
	return &Queue{
		channel:          make(chan Message, size),
		messageProcessor: messageProcessor,
	}
}

func (q *Queue) Add(item Message) {
	clarkezoneLog.Debugf("Queue: Added item")
	q.channel <- item
}

func (q *Queue) Remove() (Message, bool) {
	msg, ok := <-q.channel
	return msg, ok
}

func (q *Queue) Close() {
	close(q.channel)
}

func (q *Queue) Reader() {
	clarkezoneLog.Debugf("Queue: started reader")
	q.wg.Add(1)
	go func() {
		defer q.wg.Done()
		for {
			msg, ok := q.Remove()
			if !ok {
				break
			}
			q.messageProcessor.Process(msg)
		}
	}()
}

func (q *Queue) Wait() {
	q.wg.Wait()
}
