## skiplist(跳表)

### 接口

```
func Default() SkipList
func New(c Compare, maxHeight int) SkipList

/// skip_list.go
type SkipList interface {
    ...
	SkipListIterator
}

/// iterator.go
type Iterator interface {
    ...
}
```

### 注意

* 迭代器按C++ STL设计，插入/删除后迭代器失效。

### TODO

* vargrind
