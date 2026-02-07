package slidingwindowlimiter

import (
	"sync"
	"time"
)

var (
	id     = 0
	weight = 0.35
)

type User struct {
	Id            int
	PrevCount     float64
	CurCount      float64
	MaxWindowSize float64
	StartTime     time.Time
	Lock          sync.Mutex
}

func NewUser() *User {
	return &User{
		Id:            id + 1,
		CurCount:      0,
		PrevCount:     0,
		MaxWindowSize: 5,
		StartTime:     time.Now(),
	}
}

func (u *User) Check() bool {
	u.Lock.Lock()
	defer u.Lock.Unlock()
	curTime := time.Now()
	diff := curTime.Sub(u.StartTime)

	if diff > time.Minute {
		u.StartTime = curTime
		u.PrevCount = u.CurCount
		u.CurCount = 0
	}
	if u.CurCount+u.PrevCount*(weight) >= u.MaxWindowSize {
		return false
	}
	u.CurCount+=1
	return true
}
