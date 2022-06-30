package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	var i int64
	userChan := make(chan user, pool)

	for i = 0; i < n; i++ {
		wg.Add(1)
		go func(id int64, wg *sync.WaitGroup, uc *chan user, mtx *sync.Mutex) {
			defer wg.Done()
			userChan <- getOne(id)
			mtx.Lock()
			res = append(res, <-*uc)
			mtx.Unlock()
		}(i, &wg, &userChan, &mtx)
	}

	wg.Wait()

	return res
}

func GetBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	var i int64
	userChan := make(chan user, pool)

	for i = 0; i < n; i++ {
		wg.Add(1)
		go func(id int64, wg *sync.WaitGroup, uc chan user, mtx *sync.Mutex) {
			defer wg.Done()
			userChan <- getOne(id)
			mtx.Lock()
			res = append(res, <-uc)
			mtx.Unlock()
		}(i, &wg, userChan, &mtx)
	}

	wg.Wait()

	return res
}
