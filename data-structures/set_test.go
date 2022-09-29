package datastructures

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testStruct struct {
	name string
	age  int
}

func TestAddValueToSet(t *testing.T) {
	s := CreateSet[testStruct]()
	v := testStruct{
		name: "Dan",
		age:  28,
	}

	s.Add(v)

	require.Equal(t, 1, s.Size)
	require.True(t, s.Exists(v))
}

func TestValueNotInSet(t *testing.T) {
	s := CreateSet[int]()

	s.Add(1)

	require.Equal(t, 1, s.Size)
	require.False(t, s.Exists(2))
}

func TestRemoveValueFromSet(t *testing.T) {
	s := CreateSet[float32]()

	s.Add(1)

	require.Equal(t, 1, s.Size)

	s.Remove(1)

	require.Equal(t, 0, s.Size)
	require.False(t, s.Exists(1))
}

func TestGetValues(t *testing.T) {
	s := CreateSet[string]()

	s.Add("Hello")
	s.Add("World")
	list := s.GetValues()

	require.Contains(t, list, "Hello")
	require.Contains(t, list, "World")
}
