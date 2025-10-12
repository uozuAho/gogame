package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBullet_speed_works(t *testing.T) {
	bullet := *NewBullet(0, 0, 1, 0)
	bullet.Update(nil, nil)

	assert.Equal(t, 20.0, bullet.Pos.X)
	assert.Equal(t, 0.0, bullet.Pos.Y)
}

func TestBullet_speed_works2(t *testing.T) {
	bullet := *NewBullet(0, 0, 1, 1)
	assert.Equal(t, 0.0, bullet.Pos.Len())
	bullet.Update(nil, nil)
	assert.Equal(t, 20.0, bullet.Pos.Len())
}
