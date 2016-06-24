## skiplist(跳表)

### 接口

```golang
func Default() SkipList
func New(c Compare, maxHeight int) SkipList

/// skip_list.go
type SkipList interface {
	Insert(k, v interface{}) error /// nil | `ErrKeyAlreadyExist`.
	InsertForce(k, v interface{})  /// if exist already,update its `v`.
	Erase(k interface{}) error     /// nil | `ErrKeyNotExist`,most of scenario you can ignore it.
	Find(k interface{}) (v interface{}, exist bool)
	Clear()
	Size() int
	Empty() bool
	Min() (k, v interface{}) /// if Empty(),return nil,nil
	Max() (k, v interface{}) /// if Empty(),return nil,nil

	SkipListIterator
}

type SkipListIterator interface {
	Begin() Iterator /// minimum,or nil if Empty()
	End() Iterator   /// same as nil
}

/// iterator.go
type Iterator interface {
	Next() Iterator /// return's key is bigger than caller's,or nil if not exist
	Key() interface{}
	Value() interface{}
}
```

简单示例看eg/basic.go，详细请看源码以及对应test。

### 注意

* 迭代器按C++ STL设计，插入/删除后迭代器失效。

### TODO

* vargrind
* benchmark
