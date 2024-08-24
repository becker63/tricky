package types

// internal machine api
type Machine interface {
	Start_i() error
	Stop_i() error
	GetConfig_i() error
	ChangeConfig_i() error
}

// Man go enums suck
type BoxConst int

const (
	Demo BoxConst = iota
)
