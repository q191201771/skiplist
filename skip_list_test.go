package skip_list

import "testing"

func fill(sl SkipList, t *testing.T) {
	for i := 101000; i > 100000; i-- {
		if err := sl.Insert(i, i); err != nil {
			t.Fatal("!")
		}
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
			if err := sl.Insert(j, i); err != nil {
				t.Fatal("!")
			}
			if v, exist := sl.Find(j); !exist || CommonCompare(v, i) != 0 {
				t.Fatal("!")
			}
		}
		if sl.Size() != 1000 {
			t.Fatal("!")
		}
	}
}

func TestInsert(t *testing.T) {
	sl := Default()
	fill(sl, t)
	for i := 5; i >= 1; i-- {
		for j := 1; j <= 100; j *= 10 {
			if err := sl.Insert(i*j, i*j*2); err != nil {
				t.Fatal("!")
			}
			if v, exist := sl.Find(i * j); !exist || CommonCompare(v, i*j*2) != 0 {
				t.Fatal("!")
			}
		}
	}
	forceFill(sl, t)
	if sl.Size() != 1015 {
		t.Fatal("!")
	}
	//sl.debugPrint()
}

func TestInsertError(t *testing.T) {
	sl := Default()
	fill(sl, t)
	if err := sl.Insert(1, 1); err != nil {
		t.Fatal("!")
	}
	forceFill(sl, t)
	if v, exist := sl.Find(1); !exist || CommonCompare(v, 1) != 0 {
		t.Fatal("!")
	}
	if err := sl.Insert(1, 1); err != ErrKeyAlreadyExist {
		t.Fatal("!")
	}
	if v, exist := sl.Find(1); !exist || CommonCompare(v, 1) != 0 {
		t.Fatal("!")
	}
	if err := sl.Insert(2, 2); err != nil {
		t.Fatal("!")
	}
	if v, exist := sl.Find(2); !exist || CommonCompare(v, 2) != 0 {
		t.Fatal("!")
	}
	if err := sl.Insert(2, 3); err != ErrKeyAlreadyExist {
		t.Fatal("!")
	}
	if v, exist := sl.Find(2); !exist || CommonCompare(v, 2) != 0 {
		t.Fatal("!")
	}
	if err := sl.Insert(1, 3); err != ErrKeyAlreadyExist {
		t.Fatal("!")
	}
	if v, exist := sl.Find(1); !exist || CommonCompare(v, 1) != 0 {
		t.Fatal("!")
	}
	if sl.Size() != 1002 {
		t.Fatal("!")
	}
}

func TestInsertForce(t *testing.T) {
	sl := Default()
	fill(sl, t)
	if err := sl.Insert(1, 1); err != nil {
		t.Fatal("!")
	}
	if v, exist := sl.Find(1); !exist || CommonCompare(v, 1) != 0 {
		t.Fatal("!")
	}
	sl.InsertForce(1, 2)
	if v, exist := sl.Find(1); !exist || CommonCompare(v, 2) != 0 {
		t.Fatal("!")
	}
	forceFill(sl, t)
	sl.InsertForce(2, 2)
	if v, exist := sl.Find(2); !exist || CommonCompare(v, 2) != 0 {
		t.Fatal("!")
	}
	sl.InsertForce(2, 2)
	if v, exist := sl.Find(2); !exist || CommonCompare(v, 2) != 0 {
		t.Fatal("!")
	}
	sl.InsertForce(2, "2")
	if v, exist := sl.Find(2); !exist || v.(string) != "2" {
		t.Fatal("!")
	}
	if sl.Size() != 1002 {
		t.Fatal("!")
	}
	//sl.debugPrint()
}

func TestFind(t *testing.T) {
	sl := Default()
	fill(sl, t)
	if v, exist := sl.Find(1); exist || v != nil {
		t.Fatal("!")
	}
	if err := sl.Insert(1, 1); err != nil {
		t.Fatal("!")
	}
	forceFill(sl, t)
	if v, exist := sl.Find(1); !exist || CommonCompare(v, 1) != 0 {
		t.Fatal("!")
	}
	if sl.Size() != 1001 {
		t.Fatal("!")
	}
}

func TestEraseError(t *testing.T) {
	sl := Default()
	fill(sl, t)
	if err := sl.Erase(1); err != ErrKeyNotExist {
		t.Fatal("!")
	}
	if err := sl.Insert(1, 1); err != nil {
		t.Fatal(err)
	}
	forceFill(sl, t)
	if err := sl.Erase(1); err != nil {
		t.Fatal("!")
	}
	if v, exist := sl.Find(1); exist || v != nil {
		t.Fatal("!")
	}
	if sl.Size() != 1000 {
		t.Fatal(sl.Size())
	}
}

func TestEmpty(t *testing.T) {
	sl := Default()
	if !sl.Empty() {
		t.Fatal("!")
	}
	sl.Insert(1, 1)
	if sl.Empty() {
		t.Fatal("!")
	}
	sl.Erase(2)
	if sl.Empty() {
		t.Fatal("!")
	}
	sl.Erase(1)
	if !sl.Empty() {
		t.Fatal("!")
	}
}

func TestClear(t *testing.T) {
	sl := Default()
	sl.Clear()
	sl.Insert("aaa", 1)
	sl.Insert("baa", 1)
	sl.InsertForce("aaa", "bbb")
	if sl.Size() != 2 {
		t.Fatal("!")
	}
	sl.Clear()
	if !sl.Empty() {
		t.Fatal("!")
	}
}

func TestMinMax(t *testing.T) {
	sl := Default()
	if k, v := sl.Min(); k != nil || v != nil {
		t.Fatal("!")
	}
	if k, v := sl.Max(); k != nil || v != nil {
		t.Fatal("!")
	}
	sl.Insert(2, 200)
	sl.Insert(1, 100)
	sl.Insert(3, 300)
	sl.Insert(6, 600)
	sl.Insert(4, 400)
	sl.Insert(5, 500)
	if k, v := sl.Min(); k.(int) != 1 || v.(int) != 100 {
		t.Fatal("!")
	}
	if k, v := sl.Max(); k.(int) != 6 || v.(int) != 600 {
		t.Fatal("!")
	}
}
