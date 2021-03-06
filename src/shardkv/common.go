package shardkv

//
// Sharded key/value server.
// Lots of replica groups, each running op-at-a-time paxos.
// Shardmaster decides which group serves each shard.
// Shardmaster may change shard assignment from time to time.
//
// You will have to modify these definitions.
//

const (
	OK            = "OK"
	ErrNoKey      = "ErrNoKey"
	ErrWrongGroup = "ErrWrongGroup"
)

type Err string

// Put or Append
type PutAppendArgs struct {
	// You'll have to add definitions here.
	Key   string
	Value string
	Op    string // "Put" or "Append"
	// You'll have to add definitions here.
	// Field names must start with capital letters,
	// otherwise RPC will break.
	Identifier int64
	Counter    int
}

type PutAppendReply struct {
	WrongLeader bool
	Err         Err
}

type GetArgs struct {
	Key string
	// You'll have to add definitions here.
	Identifier int64
	Counter    int
}

type GetReply struct {
	WrongLeader bool
	Err         Err
	Value       string
}

type GetStateArgs struct {}

type GetStateReply struct {
	ConfigNum int
	State int
}

type GetShardArgs struct {
	ConfigNum int
	Shards []int
}

type GetShardReply struct {
	ConfigNum int
	State int
	Data map[string]string
	LastClerkAppliedOpIndex map[int64]int
	Valid bool
}
