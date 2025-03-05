package messages

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

var defaultMessages = []string{
	"Update code",
	"Fix bugs",
	"Add feature",
	"Refactor code",
	"Improve performance",
}

type MessageGenerator struct {
	messages []string
	index    int
	repeat   bool
}

func NewMessageGenerator(filename string, repeat bool) (*MessageGenerator, error) {
	if filename == "" {
		return &MessageGenerator{
			messages: defaultMessages,
			repeat:   true,
		}, nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var messages []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		msg := strings.TrimSpace(scanner.Text())
		if msg != "" {
			messages = append(messages, msg)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &MessageGenerator{
		messages: messages,
		repeat:   repeat,
	}, nil
}

func (mg *MessageGenerator) Next() string {
	if len(mg.messages) == 0 {
		return "Random commit"
	}

	if mg.index >= len(mg.messages) {
		if mg.repeat {
			mg.index = 0
		} else {
			return "Random commit " + string(rune(rand.Intn(1000)))
		}
	}

	msg := mg.messages[mg.index]
	mg.index++
	return msg
}
