package shardctrler

import "time"

//
// Shard controler: assigns shards to replication groups.
//
// RPC interface:
// Join(servers) -- add a set of groups (gid -> server-list mapping).
// Leave(gids) -- delete a set of groups.
// Move(shard, gid) -- hand off one shard from current owner to gid.
// Query(num) -> fetch Config # num, or latest config if num==-1.
//
// A Config (configuration) describes a set of replica groups, and the
// replica group responsible for each shard. Configs are numbered. Config
// #0 is the initial configuration, with no groups and all shards
// assigned to group 0 (the invalid group).
//
// You will need to add fields to the RPC argument structs.
//

// The number of shards.
const NShards = 10

const Timeout = 200 * time.Millisecond

// A configuration -- an assignment of shards to groups.
// Please don't change this.
type Config struct {
	Num    int              // config number
	Shards [NShards]int     // shard -> gid
	Groups map[int][]string // gid -> servers[]
}

const (
	OK               = "OK"
	ErrNoKey         = "ErrNoKey"
	ErrWrongGroup    = "ErrWrongGroup"
	ErrWrongLeader   = "ErrWrongLeader"
	ErrConfigVersion = "ErrConfigVersion"
)

type Err string

type JoinArgs struct {
	Servers   map[int][]string // new GID -> servers mappings
	RequestId int64
	ClientId  int64
}

type JoinReply struct {
	Err Err
}

type LeaveArgs struct {
	GIDs      []int
	RequestId int64
	ClientId  int64
}

type LeaveReply struct {
	Err Err
}

type MoveArgs struct {
	Shard     int
	GID       int
	RequestId int64
	ClientId  int64
}

type MoveReply struct {
	Err Err
}

type QueryArgs struct {
	Num       int // desired config number
	RequestId int64
	ClientId  int64
}

type QueryReply struct {
	Err    Err
	Config Config
}

const (
	JOIN  = "Join"
	LEAVE = "Leave"
	MOVE  = "Move"
	QUERY = "Query"
)
