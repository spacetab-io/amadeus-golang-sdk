package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoundUp(t *testing.T) {
	t.Run("RoundUp", func(t *testing.T) {
		request := Price{
			Base:  10.6,
			Tax:   20.5,
			Total: 31.1,
		}

		request.RoundUp()
		assert.Equal(t, request.Total, float64(32))
	})
}

func TestAddUp(t *testing.T) {
	t.Run("AddUp", func(t *testing.T) {
		request := Pricing{
			Base:  10.6,
			Tax:   20.5,
			Total: 31.1,
		}
		add := Pricing{
			Base:  10.6,
			Tax:   20.5,
			Total: 31.1,
		}

		request.AddUp(&add)
		assert.Equal(t, request.Total, float64(63))
	})
}
