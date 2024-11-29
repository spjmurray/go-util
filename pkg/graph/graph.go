/*
Copyright 2024 Simon Murray

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package graph

import (
	"github.com/spjmurray/go-util/pkg/queue"
	"github.com/spjmurray/go-util/pkg/set"
)

// Visitor is used to visit a node in the graph.
type Visitor[T comparable] interface {
	// Visit is called when a new node is encountered, it accepts
	// the node itself and an enqueue function used to add any peers.
	Visit(node T, enqueue func(T)) error
}

// Walker implements a generic graph walker that guarantees a
// given node is only visited once.
type Walker[T comparable] struct {
	queue *queue.Queue[T]
	seen  set.Set[T]
}

// NewWalker creates a new graph walker.
func NewWalker[T comparable](t ...T) *Walker[T] {
	return &Walker[T]{
		queue: queue.New[T](t...),
		seen:  set.New[T](),
	}
}

// Push adds a node to the FIFO queue.  This should only be used for initialization
// before calling Visit.
func (g *Walker[T]) Push(t T) {
	g.queue.Push(t)
}

// Visit visits all nodes in the graph exactly once.  Peer nodes are added as nodes
// are visited by a user defined Visitor.
func (g *Walker[T]) Visit(visitor Visitor[T]) error {
	for !g.queue.Empty() {
		t, err := g.queue.Pop()
		if err != nil {
			return err
		}

		if g.seen.Contains(t) {
			continue
		}

		g.seen.Add(t)

		if err := visitor.Visit(t, g.Push); err != nil {
			return err
		}
	}

	return nil
}
