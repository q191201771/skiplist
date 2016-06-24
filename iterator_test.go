package skip_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoop(t *testing.T) {
	assert := assert.New(t)
	sl := Default()
	i := 0
	for i = 9; i >= 0; i-- {
		sl.Insert(i, i*2)
	}

	i = 0
	for iter := sl.Begin(); iter != sl.End(); iter = iter.Next() {
		assert.Equal(iter.Key().(int), i, "!")
		assert.Equal(iter.Value().(int), i*2, "!")
		i++
	}
	assert.Equal(i, 10, "!")
}
