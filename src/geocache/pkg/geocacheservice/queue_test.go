package geocacheservice

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockMessageProcessor struct {
	messages []Message
}

func (p *MockMessageProcessor) Process(msg Message) {
	p.messages = append(p.messages, msg)
}

func TestWriter(t *testing.T) {
	mockProcessor := &MockMessageProcessor{}
	q := NewQueue(10, mockProcessor)

	expectedMessages := []Message{}

	for i := 1; i <= 10; i++ {
		msg := Message{
			ID:      i,
			Content: "Message",
		}
		q.Add(msg)
		expectedMessages = append(expectedMessages, msg)
		time.Sleep(10 * time.Millisecond)
	}

	q.Close()
	q.Reader()
	q.Wait()

	assert.Equal(t, expectedMessages, mockProcessor.messages)
}
