package liflow

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"
)

func Test_LiFlow(t *testing.T) {
	type Bar struct {
		Item int
	}
	type Foo struct {
		List []Bar
	}
	foos := &Foo{
		List: make([]Bar, 0),
	}
	bar1 := &Bar{Item: 10}
	foos.List = append(foos.List, *bar1)
	bar1.Item = 11
	g.Dump(foos.List)
	bar2 := &Bar{Item: 20}
	foos.List = append(foos.List, *bar2)
	bar2.Item = 22
	g.Dump(foos.List)
}
