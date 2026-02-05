package limiter

import (
	"sync"
	"time"
)

var (
	id = 0
)

type User struct {
	Id            int
	CurWindowSize int
	MaxWindowSize int
	StartTime     time.Time
	Lock          sync.Mutex
}

func NewUser() *User {
	return &User{
		Id:            id + 1,
		CurWindowSize: 0,
		MaxWindowSize: 5,
		StartTime:     time.Now(),
	}
}
func (u *User) Check() bool {
	curTime := time.Now()
	diff := curTime.Sub(u.StartTime)

	if diff > time.Minute {
		u.StartTime = curTime
		u.CurWindowSize = 0
	}
	if u.CurWindowSize>=u.MaxWindowSize {
		return false
	}
	u.CurWindowSize++;
	return true
}
