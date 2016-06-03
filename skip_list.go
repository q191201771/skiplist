package skip_list

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	ErrKeyAlreadyExist = errors.New("key already exist")
	ErrKeyNotExist     = errors.New("key not exist")
)

var kDefaultMaxHeight = 32
var kChance = 0.5

type Compare func(l, r interface{}) int

type SkipList interface {
	Insert(k, v interface{}) error /// maybe `ErrKeyAlreadyExist`.
	InsertForce(k, v interface{})  /// if exist,update it's `v`.
	Erase(k interface{}) error     /// maybe `ErrKeyNotExist`,most of time you can ignore it.
	Find(k interface{}) (v interface{}, exist bool)
	Size() int
	print()
}

type skiplist struct {
	maxHeight  int
	curHeight  int
	numOfNodes int
	head       *node
	cache      []*node
	compare    Compare
}

type level struct {
	prev *node
	next *node
}

type node struct {
	levels []*level
	height int
	k, v   interface{}
}

func Default(c Compare) SkipList {
	return New(c, kDefaultMaxHeight)
}

func New(c Compare, maxHeight int) SkipList {
	return &skiplist{
		maxHeight:  maxHeight,
		curHeight:  1,
		numOfNodes: 0,
		head:       createNode(maxHeight, nil, nil),
		cache:      make([]*node, maxHeight),
		compare:    c,
	}
}

func (sl *skiplist) InsertForce(k, v interface{}) {
	sl.insert(k, v, true)
}

func (sl *skiplist) Insert(k, v interface{}) error {
	return sl.insert(k, v, false)
}

func (sl *skiplist) Erase(k interface{}) error {
	node := sl.find(k)
	if node == nil {
		return ErrKeyNotExist
	}

	for i := node.height - 1; i >= 0; i-- {
		node.levels[i].prev.levels[i].next = node.levels[i].next
		if node.levels[i].next != nil {
			node.levels[i].next.levels[i].prev = node.levels[i].prev
		}
	}

	if node.height == sl.curHeight && node.height != 1 {
		for i := node.height - 1; i > 0; i-- {
			if sl.head.levels[i].next == nil {
				sl.curHeight--
			} else {
				break
			}
		}
	}

	sl.numOfNodes--
	return nil
}

func (sl *skiplist) Find(k interface{}) (v interface{}, exist bool) {
	node := sl.find(k)
	if node == nil {
		return nil, false
	}
	return node.v, true
}

func (sl *skiplist) Size() int {
	return sl.numOfNodes
}

func createNode(height int, k, v interface{}) *node {
	n := &node{
		k:      k,
		v:      v,
		height: height,
		levels: make([]*level, height),
	}
	for i := 0; i < height; i++ {
		n.levels[i] = &level{}
	}
	return n
}

func (sl *skiplist) find(k interface{}) *node {
	node := sl.head
	for i := sl.curHeight - 1; i >= 0; i-- {
		for ; node.levels[i].next != nil; node = node.levels[i].next {
			cmpRes := sl.compare(k, node.levels[i].next.k)
			if cmpRes < 0 {
				break
			} else if cmpRes == 0 {
				return node.levels[i].next
			}
		}
	}
	return nil
}

func (sl *skiplist) insert(k, v interface{}, force bool) error {
	node := sl.head
	for i := sl.curHeight - 1; i >= 0; i-- {
		for ; node.levels[i].next != nil; node = node.levels[i].next {
			cmpRes := sl.compare(k, node.levels[i].next.k)
			if cmpRes < 0 {
				break
			} else if cmpRes == 0 {
				if force {
					node.levels[i].next.v = v
				}
				return ErrKeyAlreadyExist
			}
		}
		sl.cache[i] = node
	}

	height := sl.randomHeight()
	if height > sl.curHeight {
		for i := height - 1; i >= sl.curHeight; i-- {
			sl.cache[i] = sl.head
		}
		sl.curHeight = height
	}

	newNode := createNode(height, k, v)
	for i := height - 1; i >= 0; i-- {
		cacheNode := sl.cache[i]
		newNode.levels[i].next = cacheNode.levels[i].next
		newNode.levels[i].prev = sl.cache[i]
		if cacheNode.levels[i].next != nil {
			cacheNode.levels[i].next.levels[i].prev = newNode
		}
		cacheNode.levels[i].next = newNode
	}

	sl.numOfNodes++
	return nil
}

func (sl *skiplist) print() {
	fmt.Println("-------------------")
	for i := sl.curHeight - 1; i >= 0; i-- {
		node := sl.head.levels[i].next
		for ; node != nil; node = node.levels[i].next {
			fmt.Printf("%v ", node.k)
		}
		fmt.Printf("\n")
	}
	fmt.Println("-------------------")
}

func (sl *skiplist) randomHeight() (n int) {
	//return 2
	for n = 1; n < sl.maxHeight && rand.Float64() < kChance; n++ {
	}
	return
}

func init() {
	rand.Seed(time.Now().Unix())
}
