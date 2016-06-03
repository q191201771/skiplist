package skip_list

import "testing"

func cmp(l, r interface{}) int {
	il := l.(int)
	ir := r.(int)
	if il < ir {
		return -1
	} else if il == ir {
		return 0
	}
	return 1
}

func fill(sl SkipList, t *testing.T) {
	for i := 101000; i > 100000; i-- {
		if err := sl.Insert(i, i); err != nil {
			t.Fatal("!")
		}
	}
}

func TestMaxHeight(t *testing.T) {
	for i := 1; i <= 32; i++ {
		sl := New(cmp, i)
		for j := 0; j < 1000; j++ {
			if err := sl.Insert(j, i); err != nil {
				t.Fatal("!")
			}
			if v, exist := sl.Find(j); !exist || cmp(v, i) != 0 {
				t.Fatal("!")
			}
		}
		if sl.Size() != 1000 {
			t.Fatal("!")
		}
	}
}

func TestInsert(t *testing.T) {
	sl := Default(cmp)
	fill(sl, t)
	for i := 5; i >= 1; i-- {
		for j := 1; j <= 100; j *= 10 {
			if err := sl.Insert(i*j, i*j*2); err != nil {
				t.Fatal("!")
			}
			if v, exist := sl.Find(i * j); !exist || cmp(v, i*j*2) != 0 {
				t.Fatal("!")
			}
		}
	}
	if sl.Size() != 1015 {
		t.Fatal("!")
	}
	//sl.print()
}

func TestInsertError(t *testing.T) {
	sl := Default(cmp)
	fill(sl, t)
	if err := sl.Insert(1, 1); err != nil {
		t.Fatal("!")
	}
	if v, exist := sl.Find(1); !exist || cmp(v, 1) != 0 {
		t.Fatal("!")
	}
	if err := sl.Insert(1, 1); err != ErrKeyAlreadyExist {
		t.Fatal("!")
	}
	if v, exist := sl.Find(1); !exist || cmp(v, 1) != 0 {
		t.Fatal("!")
	}
	if err := sl.Insert(2, 2); err != nil {
		t.Fatal("!")
	}
	if v, exist := sl.Find(2); !exist || cmp(v, 2) != 0 {
		t.Fatal("!")
	}
	if err := sl.Insert(2, 3); err != ErrKeyAlreadyExist {
		t.Fatal("!")
	}
	if v, exist := sl.Find(2); !exist || cmp(v, 2) != 0 {
		t.Fatal("!")
	}
	if err := sl.Insert(1, 3); err != ErrKeyAlreadyExist {
		t.Fatal("!")
	}
	if v, exist := sl.Find(1); !exist || cmp(v, 1) != 0 {
		t.Fatal("!")
	}
	if sl.Size() != 1002 {
		t.Fatal("!")
	}
}

func TestInsertForce(t *testing.T) {
	sl := Default(cmp)
	fill(sl, t)
	if err := sl.Insert(1, 1); err != nil {
		t.Fatal("!")
	}
	if v, exist := sl.Find(1); !exist || cmp(v, 1) != 0 {
		t.Fatal("!")
	}
	sl.InsertForce(1, 2)
	if v, exist := sl.Find(1); !exist || cmp(v, 2) != 0 {
		t.Fatal("!")
	}
	sl.InsertForce(2, 2)
	if v, exist := sl.Find(2); !exist || cmp(v, 2) != 0 {
		t.Fatal("!")
	}
	sl.InsertForce(2, 2)
	if v, exist := sl.Find(2); !exist || cmp(v, 2) != 0 {
		t.Fatal("!")
	}
	sl.InsertForce(2, "2")
	if v, exist := sl.Find(2); !exist || v.(string) != "2" {
		t.Fatal("!")
	}
	if sl.Size() != 1002 {
		t.Fatal("!")
	}
	//sl.print()
}

func TestFind(t *testing.T) {
	sl := Default(cmp)
	fill(sl, t)
	if v, exist := sl.Find(1); exist || v != nil {
		t.Fatal("!")
	}
	if err := sl.Insert(1, 1); err != nil {
		t.Fatal("!")
	}
	if v, exist := sl.Find(1); !exist || cmp(v, 1) != 0 {
		t.Fatal("!")
	}
	if sl.Size() != 1001 {
		t.Fatal("!")
	}
}

func TestEraseError(t *testing.T) {
	sl := Default(cmp)
	fill(sl, t)
	if err := sl.Erase(1); err != ErrKeyNotExist {
		t.Fatal("!")
	}
	if err := sl.Insert(1, 1); err != nil {
		t.Fatal(err)
	}
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
