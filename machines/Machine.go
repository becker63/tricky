package machine

import (
	"github.com/becker63/tricky/machines/Demo"
	"github.com/becker63/tricky/machines/Test"
	"github.com/becker63/tricky/types"
)

type ConfigOptions struct {
	base types.CommonBox
	test string
}

type BoxConstructor func(ConfigOptions) (types.Machine, error)

/*
*
Pretty proud of this one.
puts all our vm impls in one enum

All implementations of vms are
contained in here as our enum types
its checked by the protobuf type
*/
var Machines = map[types.BoxConst]BoxConstructor{
	// I really wish I could just do: boxes.ConstructDemo as the value here but nooooo
	types.Demo: func(c ConfigOptions) (types.Machine, error) {
		return Demo.ConstructDemo(c.base), nil
	},
	types.Test: func(c ConfigOptions) (types.Machine, error) {
		return Test.ConstructTest(c.test), nil
	},
}

func Create_Machine(machine_type types.Machine) {
	if f, err := Machines[types.Demo](ConfigOptions{
		base: types.CommonBox{
			DstPath: "string",
		},
	}); err != nil {
		f.ChangeConfig_i()
		return
	}

}
