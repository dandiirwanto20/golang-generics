package gogeneric

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// implement Generic ([T] is Type params and interface{} is constraint)
// alternative constraint [any] = interface{}
// type data dinamis and strict
func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

// example without type generic and use interface{}
func LengthGeneral(param interface{}) interface{} {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	resultString := Length[string]("Dandi") // initialize for test the generic func
	assert.Equal(t, "Dandi", resultString)

	resultNumber := Length[int](100) // initialize for test the generic func
	assert.Equal(t, 100, resultNumber)

	// initialize for test LengthGeneral func
	// LengthGeneral("dandi")
}
