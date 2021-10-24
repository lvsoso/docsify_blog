package main


import "time"

var LimitQueue map[string][]int64
var ok bool

func LimitFreqSingle(queueName string, count uint, timeWindow int64) bool {
	currTime := time.Now().Unix()
	if LimitQueue == nil {
		LimitQueue = make(map[string][]int64)
	}

	if _, ok = LimitQueue[queueName];!ok{
		LimitQueue[queueName] = make([]int64, 0)
	}

	if uint(len(LimitQueue[queueName])) < count {
		LimitQueue[queueName] append(LimitQueue[queueName], currTime)
		return true
	}

	earlyTime := LimitQueue[queueName][0]

	if currTime - earlyTime <= timeWindow {
		return false
	}else{
		LimitQueue[queueName] = LimitQueue[queueName][1:]
		LimitQueue[queueName] = append(LimitQueue[queueName], currTime)
	}

	return true
}


