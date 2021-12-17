package days

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trickShot(t *testing.T) {
	// tests
	assert.Equal(t, 45, trickShotP1("target area: x=20..30, y=-10..-5"))

	// solve
	assert.Equal(t, 4278, trickShotP1("target area: x=195..238, y=-93..-67"))
}

func Test_trickShotP2(t *testing.T) {
	// tests
	assert.Equal(t, 112, trickShotP2("target area: x=20..30, y=-10..-5"))
	// solve
	assert.Equal(t, 1994, trickShotP2("target area: x=195..238, y=-93..-67"))
}
