package machine

import (
	Boxes "github.com/becker63/tricky/Boxes/Demo"
	"github.com/becker63/tricky/types"
)

type BoxConstructor func(types.CommonBox) types.Machine

/*
*
Pretty proud of this one.
puts all our vm impls in one enum

All implementations of vms are
contained in here as our enum types
its checked by the protobuf type
*/
func Machines_Enum() map[types.BoxConst]BoxConstructor {
	m := map[types.BoxConst]BoxConstructor{
		// I really wish I could just do: boxes.ConstructDemo as the value here but nooooo
		types.Demo: func(c types.CommonBox) types.Machine {
			return Boxes.ConstructDemo(c)
		},
	}

	return m
}

func Create_Machine(machine_type types.Machine) {
	k := Machines_Enum()
	k[types.Demo](types.CommonBox{})
}
