package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddSuccess(t *testing.T) {
	c := require.New(t)

	result := Add(20, 3)

	expect := 23

	//value expected, values result
	c.Equal(expect, result)
}
