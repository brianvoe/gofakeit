package gofakeit

import (
	"errors"
	"sync"
	"time"
)

// SnowFlake will generate a unique int64 based upon timestamp
// Incremental with order: Only requires ordering, not continuity.
// Globally unique: Needs to be cross-machine and cross-time.

const (
	workerBits  uint8 = 10                      // Number of bits for worker ID
	numberBits  uint8 = 12                      // Number of bits for sequence number
	workerMax   int64 = -1 ^ (-1 << workerBits) // Maximum value for worker ID (1023)
	numberMax   int64 = -1 ^ (-1 << numberBits) // Maximum value for sequence number (4095)
	timeShift   uint8 = workerBits + numberBits // Offset for timestamp
	workerShift uint8 = numberBits              // Offset for worker ID
	epoch       int64 = 1671849600000           // Start constant timestamp (in milliseconds), chosen as 2023-11-23 00:00:00
)

type SFWorker struct {
	mu        sync.Mutex
	timeStamp int64
	workerId  int64
	number    int64
}

func NewSFWorker(workerId int64) (*SFWorker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("WorkerId exceeds the limit.！")
	}
	return &SFWorker{
		timeStamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *SFWorker) NextId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	// Current timestamp in milliseconds
	now := time.Now().UnixNano() / 1e6
	// If the timestamp is the same as the current time, increment the sequence number
	if w.timeStamp == now {
		w.number++
		// If the sequence number exceeds the maximum value, update the timestamp
		if w.number > numberMax {
			for now <= w.timeStamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else { // If the timestamp is different from the current time, update the timestamp directly
		w.number = 0
		w.timeStamp = now
	}
	// ID consists of timestamp, worker ID, and sequence number
	ID := (now-epoch)<<timeShift | (w.workerId << workerShift) | (w.number)
	return ID
}
