package skip_list

import "testing"

func TestLoop(t *testing.T) {
	sl := Default()
	i := 0
	for i = 9; i >= 0; i-- {
		sl.Insert(i, i*2)
	}

	i = 0
	for iter := sl.Begin(); iter != sl.End(); iter = iter.Next() {
		if iter.Key().(int) != i || iter.Value().(int) != i*2 {
			t.Fatal("!")
		}
		i++
	}
	if i != 10 {
		t.Fatal("!")
	}
}
