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

package slices_test

import (
	"testing"

	"github.com/spjmurray/go-util/pkg/slices"
	"github.com/stretchr/testify/require"
)

type Pair struct {
	x, y int
}

// TestPermute esnures all possible combinations of slice members are
// returned by the Permute function.
func TestPermute(t *testing.T) {
	t.Parallel()

	in := []int{1, 2, 3}

	//nolint:prealloc
	var out []Pair

	for x, y := range slices.Permute(in) {
		out = append(out, Pair{x: x, y: y})
	}

	expected := []Pair{
		{x: 1, y: 2},
		{x: 1, y: 3},
		{x: 2, y: 3},
	}

	require.Equal(t, expected, out)
}
