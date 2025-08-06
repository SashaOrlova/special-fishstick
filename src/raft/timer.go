package raft

import (
	"math/rand"
	"time"
)

// Here is some consts for elections
// it may be replace to config
const (
	ElectionTimeoutMin = 150 * time.Millisecond
	ElectionTimeoutMax = 300 * time.Millisecond
)

func (n *Node) resetElectionTimer() time.Duration {
	return time.Duration(rand.Int63n(
		int64(ElectionTimeoutMax-ElectionTimeoutMin),
	) + int64(ElectionTimeoutMin))
}
