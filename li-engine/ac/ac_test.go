package ac_test

import (
	"context"
	"testing"
	"time"

	"github.com/BeanWei/li/li-engine/ac"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_CheckForController(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		pass, err := ac.CheckForController(context.Background(), func(ctx context.Context) (pass bool, err error) {
			t.Log("1")
			time.Sleep(time.Second * 3)
			return false, nil
		}, func(ctx context.Context) (pass bool, err error) {
			t.Log("2")
			return true, nil
		}, func(ctx context.Context) (pass bool, err error) {
			t.Log("3")
			return false, nil
		}, func(ctx context.Context) (pass bool, err error) {
			t.Log("4")
			time.Sleep(time.Second * 10)
			return false, nil
		}, func(ctx context.Context) (pass bool, err error) {
			t.Log("5")
			time.Sleep(time.Second * 10)
			return false, nil
		}, func(ctx context.Context) (pass bool, err error) {
			t.Log("6")
			return true, nil
		}, func(ctx context.Context) (pass bool, err error) {
			t.Log("7")
			time.Sleep(time.Second * 10)
			return false, nil
		})
		t.AssertEQ(pass, true)
		t.AssertEQ(err, nil)
	})
}

func Test_CheckForView(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ac.Bind("1", func(ctx context.Context) (pass bool, err error) {
			t.Log("1")
			return false, nil
		})
		ac.Bind("2", func(ctx context.Context) (pass bool, err error) {
			t.Log("2")
			return true, nil
		})
		ac.Bind("3", func(ctx context.Context) (pass bool, err error) {
			t.Log("3")
			return false, nil
		})
		ac.Bind("4", func(ctx context.Context) (pass bool, err error) {
			t.Log("4")
			return true, nil
		})
		removes, err := ac.CheckForView(context.Background(), "1", "2", "3", "4")
		t.AssertEQ(err, nil)
		t.AssertEQ(garray.NewStrArrayFrom(removes).Sort().Join(""), "13")
	})
}
