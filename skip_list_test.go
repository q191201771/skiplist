package skip_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	err   error
	exist bool
	k     interface{}
	v     interface{}
)

func fill(sl SkipList, t *testing.T) {
	for i := 101000; i > 100000; i-- {
		err := sl.Insert(i, i)
		assert.Nil(t, err, "!")
	}
}

func forceFill(sl SkipList, t *testing.T) {
	for i := 101000; i > 100000; i-- {
		sl.InsertForce(i, i)
	}
}

func TestMaxHeight(t *testing.T) {
	for i := 1; i <= 32; i++ {
		sl := New(CommonCompare, i)
		for j := 0; j < 1000; j++ {
			err := sl.Insert(j, i)
			assert.Nil(t, err, "!")
			v, exist := sl.Find(j)
			assert.True(t, exist, "!")
			assert.Zero(t, CommonCompare(v, i), "!")
		}
		assert.Equal(t, sl.Size(), 1000)
	}
}

func TestInsert(t *testing.T) {
	sl := Default()
	fill(sl, t)
	for i := 5; i >= 1; i-- {
		for j := 1; j <= 100; j *= 10 {
			err := sl.Insert(i*j, i*j*2)
			assert.Nil(t, err, "!")
			v, exist := sl.Find(i * j)
			assert.True(t, exist, "!")
			assert.Zero(t, CommonCompare(v, i*j*2), "!")
		}
	}
	forceFill(sl, t)
	assert.Equal(t, sl.Size(), 1015)
	//sl.debugPrint()
}

func TestInsertError(t *testing.T) {
	assert := assert.New(t)
	sl := Default()
	fill(sl, t)

	err = sl.Insert(1, 1)
	assert.Nil(err, "!")

	forceFill(sl, t)
	v, exist = sl.Find(1)
	assert.True(exist, "!")
	assert.Zero(CommonCompare(v, 1), "!")

	err = sl.Insert(1, 1)
	assert.Equal(err, ErrKeyAlreadyExist, "!")

	v, exist = sl.Find(1)
	assert.True(exist, "!")
	assert.Zero(CommonCompare(v, 1), "!")

	err = sl.Insert(2, 2)
	assert.Nil(err, "!")

	v, exist = sl.Find(2)
	assert.True(exist, "!")
	assert.Zero(CommonCompare(v, 2), "!")

	err = sl.Insert(2, 3)
	assert.Equal(err, ErrKeyAlreadyExist, "!")

	v, exist = sl.Find(2)
	assert.True(exist, "!")
	assert.Zero(CommonCompare(v, 2), "!")

	err = sl.Insert(1, 3)
	assert.Equal(err, ErrKeyAlreadyExist, "!")

	v, exist = sl.Find(1)
	assert.True(exist, "!")
	assert.Zero(CommonCompare(v, 1), "!")

	assert.Equal(sl.Size(), 1002)
}

func TestInsertForce(t *testing.T) {
	assert := assert.New(t)
	sl := Default()
	fill(sl, t)

	err = sl.Insert(1, 1)
	assert.Nil(err, "!")

	v, exist = sl.Find(1)
	assert.True(exist, "!")
	assert.Zero(CommonCompare(v, 1), "!")

	sl.InsertForce(1, 2)
	v, exist = sl.Find(1)
	assert.True(exist, "!")
	assert.Zero(CommonCompare(v, 2), "!")

	forceFill(sl, t)
	sl.InsertForce(2, 2)
	v, exist = sl.Find(2)
	assert.True(exist, "!")
	assert.Zero(CommonCompare(v, 2), "!")

	sl.InsertForce(2, 2)
	v, exist = sl.Find(2)
	assert.True(exist, "!")
	assert.Zero(CommonCompare(v, 2), "!")

	sl.InsertForce(2, "2")
	v, exist = sl.Find(2)
	assert.True(exist, "!")
	assert.Equal(v.(string), "2", "!")

	assert.Equal(sl.Size(), 1002, "!")
	//sl.debugPrint()
}

func TestFind(t *testing.T) {
	assert := assert.New(t)
	sl := Default()

	fill(sl, t)
	v, exist = sl.Find(1)
	assert.False(exist, "!")
	assert.Nil(v, "!")

	err = sl.Insert(1, 1)
	assert.Nil(err, "!")

	forceFill(sl, t)
	v, exist = sl.Find(1)
	assert.True(exist, "!")
	assert.Zero(CommonCompare(v, 1), "!")

	assert.Equal(sl.Size(), 1001)
}

func TestEraseError(t *testing.T) {
	assert := assert.New(t)

	sl := Default()
	fill(sl, t)
	err = sl.Erase(1)
	assert.Equal(err, ErrKeyNotExist, "!")

	err = sl.Insert(1, 1)
	assert.Nil(err, "!")

	forceFill(sl, t)
	err = sl.Erase(1)
	assert.Nil(err, "!")

	v, exist = sl.Find(1)
	assert.False(exist, "!")
	assert.Nil(v)

	assert.Equal(sl.Size(), 1000)
}

func TestEmpty(t *testing.T) {
	assert := assert.New(t)
	sl := Default()

	assert.True(sl.Empty(), "!")

	sl.Insert(1, 1)
	assert.False(sl.Empty(), "!")

	sl.Erase(2)
	assert.False(sl.Empty(), "!")

	sl.Erase(1)
	assert.True(sl.Empty(), "!")
}

func TestClear(t *testing.T) {
	assert := assert.New(t)
	sl := Default()

	sl.Clear()
	sl.Insert("aaa", 1)
	sl.Insert("baa", 1)
	sl.InsertForce("aaa", "bbb")
	assert.Equal(sl.Size(), 2, "!")

	sl.Clear()
	assert.True(sl.Empty(), "!")
}

func TestMinMax(t *testing.T) {
	assert := assert.New(t)
	sl := Default()

	k, v = sl.Min()
	assert.Nil(k, "!")
	assert.Nil(v, "!")

	k, v = sl.Max()
	assert.Nil(k, "!")
	assert.Nil(v, "!")

	sl.Insert(2, 200)
	sl.Insert(1, 100)
	sl.Insert(3, 300)
	sl.Insert(6, 600)
	sl.Insert(4, 400)
	sl.Insert(5, 500)

	k, v = sl.Min()
	assert.Equal(k.(int), 1)
	assert.Equal(v.(int), 100)

	k, v = sl.Max()
	assert.Equal(k.(int), 6)
	assert.Equal(v.(int), 600)
}
