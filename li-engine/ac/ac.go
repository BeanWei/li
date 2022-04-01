package ac

import (
	"context"
	"runtime"

	"github.com/BeanWei/li/li-engine/pkg/errgroup"
)

type AC func(ctx context.Context) (pass bool, err error)

var acl = make(map[string]AC)

func Bind(path string, f AC) {
	if path != "" && f != nil {
		acl[path] = f
	}
}

func Get(path string) AC {
	return acl[path]
}

func GetAll() map[string]AC {
	return acl
}

// CheckForController 接口访问的鉴权，有一个通过即可
func CheckForController(ctx context.Context, acs ...AC) (pass bool, err error) {
	if len(acs) == 0 {
		return true, nil
	}
	g := errgroup.WithCancel(context.Background())
	g.GOMAXPROCS(runtime.NumCPU())
	for _, f := range acs {
		if f == nil {
			continue
		}
		f := f
		g.Go(func(_ context.Context) error {
			if pass {
				return nil
			}
			p, e := f(ctx)
			if p {
				pass = true
			}
			return e
		})
	}
	if pass {
		return true, nil
	}
	err = g.Wait()
	return
}

// CheckForView 视图访问的鉴权，执行全部并返回需要移除的无权限的路径
func CheckForView(ctx context.Context, paths ...string) (removes []string, err error) {
	if len(paths) == 0 {
		return
	}
	g := errgroup.WithCancel(context.Background())
	g.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan string, len(paths))
	for _, path := range paths {
		f := acl[path]
		if f == nil {
			continue
		}
		path := path
		g.Go(func(_ context.Context) error {
			p, e := f(ctx)
			if !p {
				ch <- path
			}
			return e
		})
	}
	err = g.Wait()
	close(ch)
	if err != nil {
		return
	}
	for path := range ch {
		removes = append(removes, path)
	}
	return
}
