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

package set

import (
	"iter"
	"maps"
	"slices"
)

// Set allows O(log N) insertion and deletion.
type Set[T comparable] map[T]any

// New creates a new set.
func New[T comparable](t ...T) Set[T] {
	s := Set[T]{}

	for _, i := range t {
		s.Add(i)
	}

	return s
}

// Len is the size of the set.
func (s Set[T]) Len() int {
	return len(s)
}

// Add adds a new member.
func (s Set[T]) Add(t T) {
	s[t] = nil
}

// Delete removes and existing member.
func (s Set[T]) Delete(t T) {
	delete(s, t)
}

// Contains checks whether a value is in the set.
func (s Set[T]) Contains(t T) bool {
	_, ok := s[t]

	return ok
}

// Clear removes all set elements.
func (s *Set[T]) Clear() {
	*s = map[T]any{}
}

// All provides non-deterministic iteration.
func (s Set[T]) All() iter.Seq[T] {
	return func(yield func(t T) bool) {
		for k := range s {
			if !yield(k) {
				return
			}
		}
	}
}

// AllSortedFunc provides deterministic iteration.  This is a lot slower
// than non-deterministic, but useful for debugging.
func (s Set[T]) AllSortedFunc(cmp func(T, T) int) iter.Seq[T] {
	return func(yield func(t T) bool) {
		for _, k := range slices.SortedFunc(maps.Keys(s), cmp) {
			if !yield(k) {
				return
			}
		}
	}
}

// Union returns a set that contains all elements in both sets.
func (s Set[T]) Union(o Set[T]) Set[T] {
	out := Set[T]{}

	for i := range s.All() {
		out.Add(i)
	}

	for i := range o.All() {
		out.Add(i)
	}

	return out
}

// Intersection returns a set that contains all elements that exist in both sets.
func (s Set[T]) Intersection(o Set[T]) Set[T] {
	out := Set[T]{}

	for i := range s.All() {
		if o.Contains(i) {
			out.Add(i)
		}
	}

	return out
}

// Difference returns a set that contains only elements in the first set that do not
// occur in the second.
func (s Set[T]) Difference(o Set[T]) Set[T] {
	out := Set[T]{}

	for i := range s.All() {
		if !o.Contains(i) {
			out.Add(i)
		}
	}

	return out
}

// SymmetricDifference returns a set that contains only elements that appear in one
// set.
func (s Set[T]) SymmetricDifference(o Set[T]) Set[T] {
	return s.Difference(o).Union(o.Difference(s))
}
