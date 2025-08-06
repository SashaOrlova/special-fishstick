package raft

type LogEntry struct {
	Term    int
	Command interface{}
}

type NodeState int

const (
	Follower NodeState = iota
	Candidate
	Leader
)

func (s NodeState) String() string {
	return [...]string{"Follower", "Candidate", "Leader"}[s]
}

type Node struct {
	ID          string    // Unique indentifier
	State       NodeState // "Follower", "Candidate", "Leader"
	CurrentTerm int       // Node's current term
	VotedFor    string    // Voted candidate ID in current term
	Log         []LogEntry
	CommitIndex int
	LastApplied int
}

func NewNode(id string) *Node {
	return &Node{
		ID:          id,
		State:       Follower, // All nodes start as followers
		CurrentTerm: 0,
		VotedFor:    "",
		Log:         []LogEntry{},
		CommitIndex: -1, // -1 means empty log
		LastApplied: -1,
	}
}
