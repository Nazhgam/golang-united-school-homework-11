package batch

import (
	"context"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var mu sync.Mutex
	errG, _ := errgroup.WithContext(context.Background())
	errG.SetLimit(pool)
	for i := n - 1; i >= 0; i-- {
		errG.Go(func(mu *sync.Mutex, id int64) {
			mu.Lock()
			res = append(res, getOne(id))
			mu.Unlock()
		}(&mu, i))
	}
	return res
}
