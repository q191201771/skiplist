package main

import (
	"log"

	slist "github.com/q191201771/skip_list"
)

func main() {
	var (
		err     error
		v       interface{}
		exist   bool
		size    int
		isEmpty bool
	)

	sl := slist.Default()
	err = sl.Insert("a", 1)   /// err -> nil | "a": 1.
	err = sl.Insert("a", 2)   /// err -> ErrKeyAlreadyExist | "a": 1.
	sl.InsertForce("a", "a")  /// | "a": "a".
	err = sl.Erase("a")       /// err -> nil | nil
	err = sl.Erase("a")       /// err -> ErrKeyNotExist | nil
	v, exist = sl.Find("a")   /// v -> nil, exist -> false | nil
	err = sl.Insert("a", 3)   /// err -> nil | "a": 3.
	v, exist = sl.Find("a")   /// v.(int) -> 3, exist -> true | "a": 3.
	err = sl.Insert("b", "b") /// err -> nil | "a": 3, "b": "b".
	size = sl.Size()          /// size -> 2 | "a": 3, "b": "b".
	sl.Clear()                /// | nil
	isEmpty = sl.Empty()      /// isEmpty -> true | nil

	for i := 10; i > 0; i-- {
		sl.Insert(i, i*2)
	}
	for iter := sl.Begin(); iter != sl.End(); iter = iter.Next() {
		log.Println(iter.Key(), "->", iter.Value())
		/// 1 -> 2
		/// 2 -> 4
		/// ...
		/// 10 -> 20
	}
	minK, minV := sl.Min() /// minK -> 1,  minV -> 1
	maxK, maxV := sl.Max() /// maxK -> 10, maxV -> 20

	dummy(err, v, exist, size, isEmpty, minK, minV, maxK, maxV)
}

func dummy(param ...interface{}) {}
