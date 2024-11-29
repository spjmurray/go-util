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

package graph_test

import (
	"fmt"

	"github.com/spjmurray/go-util/pkg/graph"
)

type Visitor struct {
	peers map[string][]string
}

func (v *Visitor) Visit(node string, enqueue func(string)) error {
	fmt.Println(node)

	for _, peer := range v.peers[node] {
		enqueue(peer)
	}

	return nil
}

// ExampleWalker does a very simple walk of a directed graph with four nodes.
func ExampleWalker() {
	v := &Visitor{
		peers: map[string][]string{
			"a": {"b", "c"},
			"b": {"d"},
			"c": {"d"},
		},
	}

	w := graph.NewWalker("a")

	if err := w.Visit(v); err != nil {
		panic(err)
	}

	// Output:
	// a
	// b
	// c
	// d
}
