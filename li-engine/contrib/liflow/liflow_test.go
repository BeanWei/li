package liflow

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gutil"
)

func Test_LiFlow(t *testing.T) {
	a := map[string]interface{}{"a": 1, "c": 3}
	b := map[string]interface{}{"a": 2, "d": 4}
	gutil.MapMerge(a, b)
	g.Dump(
		a,
	)
}
