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

package set_test

import (
	"testing"

	"github.com/spjmurray/go-util/pkg/set"
	"github.com/stretchr/testify/require"
)

// TestSetBasic checks basic set initialization, addition and deletion.
func TestSetBasic(t *testing.T) {
	t.Parallel()

	s := set.New[int](1, 2, 3)
	require.Equal(t, 3, s.Len())
	require.True(t, s.Contains(1))
	require.True(t, s.Contains(2))
	require.True(t, s.Contains(3))

	s.Add(3)
	require.Equal(t, 3, s.Len())

	s.Add(4)
	require.Equal(t, 4, s.Len())
	require.True(t, s.Contains(4))

	s.Delete(1)
	require.Equal(t, 3, s.Len())
	require.False(t, s.Contains(1))

	s.Clear()
	require.Equal(t, 0, s.Len())
}

// TestSetArithmetic checks the set aritmetic operations.
func TestSetArithmetic(t *testing.T) {
	t.Parallel()

	a := set.New[int](1, 2, 3)
	b := set.New[int](3, 4, 5)

	u := a.Union(b)
	require.Equal(t, 5, u.Len())

	i := a.Intersection(b)
	require.Equal(t, 1, i.Len())
	require.True(t, i.Contains(3))

	d := a.Difference(b)
	require.Equal(t, 2, d.Len())
	require.True(t, d.Contains(1))
	require.True(t, d.Contains(2))

	s := a.SymmetricDifference(b)
	require.Equal(t, 4, s.Len())
	require.True(t, s.Contains(1))
	require.True(t, s.Contains(2))
	require.False(t, s.Contains(3))
	require.True(t, s.Contains(4))
	require.True(t, s.Contains(5))
}
