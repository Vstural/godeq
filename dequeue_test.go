package godeque

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQueue(t *testing.T) {
	Convey("TestQueue", t, func() {
		que := NewDequeue()
		So(que, ShouldNotBeNil)

		b := que.Back()
		So(b, ShouldBeNil)
		f := que.Front()
		So(f, ShouldBeNil)

		printFunc := func(i interface{}) bool {
			if val, ok := i.(int); ok {
				fmt.Print(val, " ")
			} else {
				fmt.Println("Refelect fail")
			}
			return true
		}

		que.PushBack(1)
		que.Foreach(printFunc)
		fmt.Println("")
		que.PushBack(2)
		que.Foreach(printFunc)
		fmt.Println("")
		que.PushBack(3)
		que.Foreach(printFunc)
		fmt.Println("")
		que.PushBack(4)
		que.Foreach(printFunc)
		fmt.Println("")
		que.PushFront(5)
		que.Foreach(printFunc)
		fmt.Println("")
		que.PushFront(6)
		que.Foreach(printFunc)
		fmt.Println("")
		que.PushFront(7)
		que.Foreach(printFunc)
		fmt.Println("")
		// 4 3 2 1 5 6 7

		v := que.Back().(int)
		So(v, ShouldEqual, 4)
		v = que.Front().(int)
		So(v, ShouldEqual, 7)

		que.PopBack()
		v = que.Back().(int)
		que.Foreach(printFunc)
		fmt.Println("")
		So(v, ShouldEqual, 3)

		que.PopFront()
		v = que.Front().(int)
		que.Foreach(printFunc)
		fmt.Println("")
		So(v, ShouldEqual, 6)

		que.PopFront()
		que.PopFront()
		que.PopFront()
		que.PopFront()
		que.PopFront()
		que.PopFront()
		que.PopFront()
		que.PopFront()
		que.PopFront()

	})
}
