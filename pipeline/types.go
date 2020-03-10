package pipeline

import (
	"sync"

	"github.com/machester4/jane/chain"
)

type Word = chain.Word

type Stage struct {
	Name   string
	Steps []func()
	Lifo bool // TODO: Implement
}

type Pipeline struct {
	wg     sync.WaitGroup
	Stages []*Stage
}