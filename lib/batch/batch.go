package batch

import (
	"context"
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
	res = make([]user, n)
	errG, _ := errgroup.WithContext(context.Background())
	errG.SetLimit(int(pool))
	for i := range res {
		i := i
		errG.Go(func() error {
			res[i] = getOne(int64(i))
			return nil
		})
	}
	errG.Wait()
	return res

}
