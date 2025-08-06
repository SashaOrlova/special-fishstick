package raft

type RequestVoteArgs struct {
	Term        int
	CandidateID string
}

type RequestVoteReply struct {
	Term        int
	VoteGranted bool
}

type AppendEntriesArgs struct {
	Term     int // Leader's Терм
	LeaderID string
	Entries  []LogEntry
}

type AppendEntriesReply struct {
	Term    int // Node's Term
	Success bool
}
