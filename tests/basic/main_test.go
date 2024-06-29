package basic

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAddOne1(t *testing.T) {
	assert.Equal(t, AddOne(1), 2, "Incorrect addition")
}

func TestAddOne2(t *testing.T) {
	assert.Equal(t, AddOne(2), 3, "Incorrect addition")
}

func TestAddOne3(t *testing.T) {
	assert.Equal(t, AddOne(3), 4, "Incorrect addition")
}

func TestAddOne5(t *testing.T) {
	assert.Equal(t, AddOne(5), 6, "Incorrect addition")
}
