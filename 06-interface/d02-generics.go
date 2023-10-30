package main

import "fmt"

// 泛型函数，有两个类型参数：
//          一个 K (comparable 类型，可比较的类型), 
//          一个 V (任意类型，any 是 interface{} 的别称)
// 所以，MapKeys 可以接收任意类型的 map，返回 此 map 的 key 的切片
func MapKeys[K comparable, V any](m map[K]V) []K {
    r := make([]K, 0, len(m))
    for k := range m {
        r = append(r, k)
    }
    return r
}

// 泛型类型，有一个类型参数：T (任意类型)
type List[T any] struct {
    head, tail *element[T]
}

type element[T any] struct {
    next *element[T]
    val  T
}
// 泛型类型的方法
func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) GetAll() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func main() {
    var m = map[int]string{1: "2", 2: "4", 4: "8"}

    fmt.Println("keys m:", MapKeys(m))

	// 调用泛型函数时，可以不传入 类型参数，因为有类型推断
    _ = MapKeys[int, string](m)

    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list:", lst.GetAll())
}