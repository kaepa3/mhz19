package mhz19_test

import (
	"testing"

	"github.com/kebhr/mhz19"
	"github.com/stretchr/testify/assert"
)

func MocTest(t *testing.T) {
	val := 2
	mock := mhz19.MocMHZ19{Value: val}
	assert.Nil(t, mock.Connect())
	ret, err := mock.ReadCO2()
	assert.Equal(t, val, ret)
	assert.Nil(t, err)
}
