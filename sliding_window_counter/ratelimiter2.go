package slidingwindowlimiter

import (
	"sync"
	"time"
)

var (
	id = 0
)

type User struct {
	Id            int
	WindowSize    int64
	PrevCount     float64
	CurCount      float64
	MaxWindowSize float64
	StartTime     int64
	Lock          sync.Mutex
}

func NewSlider() *User {
	return &User{
		Id:            id + 1,
		CurCount:      0,
		PrevCount:     0,
		MaxWindowSize: 5,
		StartTime:     time.Now().Unix(),
	}
}

func (u *User) Check() bool {
	u.Lock.Lock()
	defer u.Lock.Unlock()
	curTime := time.Now().Unix()

	if  curTime> u.WindowSize+u.StartTime{
		u.StartTime = curTime
		u.PrevCount = u.CurCount
		u.CurCount = 0
	}
	weight := 1.1
	if u.CurCount+u.PrevCount*(weight) >= u.MaxWindowSize {
		return false
	}
	u.CurCount += 1
	return true
}
