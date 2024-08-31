package types

// TODO: Create validation function for these types

import (
	"net/url"
)

type Internal struct {
	// std.hash Sum32
	groupId uint32
}

/*
*
man golang enums suck

https://github.com/golang/go/issues/28987#issuecomment-2053744601
*/
type Runner int

const (
	Binary Runner = iota
	Java
)

type CommonBox struct {
	RemoteZip  url.URL
	DstPath    string
	BinaryPath string
	Runner     Runner
}
