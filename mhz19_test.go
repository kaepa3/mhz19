package mhz19_test

import (
	"testing"

	"github.com/kaepa3/mhz19"
	"github.com/stretchr/testify/assert"
)

func MocTest(t *testing.T) {
	val := 2
	mock := mhz19.MockMHZ19{Value: val}
	assert.Nil(t, mock.Connect())
	ret, err := mock.ReadCO2()
	assert.Equal(t, val, ret)
	assert.Nil(t, err)
}
func InterfaceTest(t *testing.T) {
	var loader mhz19.Co2Dataloader
	mock := mhz19.MockMHZ19{Value: 3}
	loader = &mock
	loader.Connect()
	loader = &mhz19.MHZ19{}
}
