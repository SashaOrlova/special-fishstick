package raft_test

import (
	"testing"

	"github.com/SashaOrlova/special-fishstick/src/raft"
	"github.com/stretchr/testify/assert"
)

// Every Node start life from "Follower" state
func TestNodeStartsAsFollower(t *testing.T) {
	node := raft.NewNode("node-A")
	assert.Equal(t, raft.Follower, node.State)
	assert.Equal(t, 0, node.CurrentTerm)
}

// Test correct enum to string formating
func TestNodeStatesCorrect(t *testing.T) {
	assert.Equal(t, raft.Follower.String(), "Follower")
	assert.Equal(t, raft.Leader.String(), "Leader")
	assert.Equal(t, raft.Candidate.String(), "Candidate")
}
