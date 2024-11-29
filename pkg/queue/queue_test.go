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

package queue_test

import (
	"testing"

	"github.com/spjmurray/go-util/pkg/queue"
	"github.com/stretchr/testify/require"
)

// TestQueueBasic checks basic queue initialization, addition and deletion.
func TestSetBasic(t *testing.T) {
	t.Parallel()

	q := queue.New[int](1, 2, 3)
	require.Equal(t, 3, q.Len())
	require.False(t, q.Empty())

	q.Push(4)
	require.Equal(t, 4, q.Len())

	head, err := q.Peek()
	require.NoError(t, err)
	require.Equal(t, 4, q.Len())
	require.Equal(t, 1, head)

	head, err = q.Pop()
	require.NoError(t, err)
	require.Equal(t, 3, q.Len())
	require.Equal(t, 1, head)

	q.Clear()
	require.Equal(t, 0, q.Len())
	require.True(t, q.Empty())

	_, err = q.Pop()
	require.Error(t, err)
}
