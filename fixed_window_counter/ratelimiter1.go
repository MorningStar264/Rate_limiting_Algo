package fixed_window_limiter

import (
	"sync"
	"time"
)

var (
	id = 0
)

type User struct {
	Id            int
	CurCount      int
	MaxWindowSize int
	StartTime     time.Time
	Lock          sync.Mutex
}

func NewUser() *User {
	return &User{
		Id:            id + 1,
		CurCount:      0,
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
		u.CurCount = 0
	}
	if u.CurCount >= u.MaxWindowSize {
		return false
	}
	u.CurCount++
	return true
}
