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

package queue

import (
	"errors"
	"fmt"
)

var (
	// ErrQueue is returned when an operation cannot be performed
	// e.g. a pop from an empty queue.
	ErrQueue = errors.New("opperation cannot be performed")
)

// Queue implements a generic FIFO queue.
type Queue[T any] struct {
	items []T
}

// New creates a new FIFO queue.
func New[T any](t ...T) *Queue[T] {
	q := &Queue[T]{}

	for _, i := range t {
		q.Push(i)
	}

	return q
}

// Empty is true of the queue contains no items.
func (q *Queue[T]) Empty() bool {
	return len(q.items) == 0
}

// Len returns the number of items in the queue.
func (q *Queue[T]) Len() int {
	return len(q.items)
}

// Push adds an item to the tail of the queue.
func (q *Queue[T]) Push(value T) {
	q.items = append(q.items, value)
}

// Peek fetches the head item from the queue.
func (q *Queue[T]) Peek() (T, error) {
	if q.Empty() {
		var t T

		return t, fmt.Errorf("%w: attempted to read from an empty list", ErrQueue)
	}

	return q.items[0], nil
}

// Pop removes the head item from the queue and returns it.
func (q *Queue[T]) Pop() (T, error) {
	head, err := q.Peek()
	if err != nil {
		return head, err
	}

	q.items = q.items[1:]

	return head, nil
}

// Clear removes all items from the queue.
func (q *Queue[T]) Clear() {
	q.items = nil
}
