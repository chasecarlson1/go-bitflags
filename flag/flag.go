/*
# Provides a simple `Flag` (uint32) type that can store 32 true-false values.

Provides variadic forms of functions that end in "V", but these are not as efficient as their normal counterparts.

In other words, `func (b *Flag) Clear(flag Flag)` is preferrable to `func (b *Flag) ClearV(flags ...Flag)`.

But, you can use the variadic ones if you want and are okay with a little slice allocation overhead (which isn't much because these are uint32 anyways).

# Example:

	import (
		"github.com/chasecarlson1/go-bitflags/flag"
	)

	const (
		FlagA flag.Flag = 1 << iota // 0001
		FlagB 			   // 0010
		FlagC 			   // 0100
	)

	var f flag.Flag = 0
	f.Set(FlagB)
	fmt.Print(f.String()) // prints binary "0010"
	f.ClearAll() // flag == 0 (no flags set on the "f" Flag variable)
	f.ToggleV(FlagC, FlagA) // variadic forms of functions end in "V". These have performance overhead though but allow for multiple arguments.
	f.IsSet(FlagA) // now returns true because of the line above
*/
package flag

import "fmt"

/*
`Flag` can store 32 true/false (or on/off) values.

# Example:

	import (
		"github.com/chasecarlson1/go-bitflags/flag"
	)

	const (
		FlagA flag.Flag = 1 << iota // 0001
		FlagB 			   // 0010
		FlagC 			   // 0100
	)

	var f flag.Flag = 0
	f.Set(FlagB)
	fmt.Print(f.String()) // prints binary "0010"
	f.ClearAll() // flag == 0 (no flags set on the "f" Flag variable)
	f.ToggleV(FlagC, FlagA) // variadic forms of functions end in "V". These have performance overhead though but allow for multiple arguments.
	f.IsSet(FlagA) // now returns true because of the line above
*/
type Flag uint32

//# New returns a Flag variable initialized to zero.
func New() Flag {
	return Flag(0)
}

//# NewV returns a Flag with all the provided `flags` set to true/on
//
//some heap overhead compared to `New()` because it is variadic and uses a slice of Flag (`uint32`s).
func NewV(flags ...Flag) Flag {
	var flag = New()
	flag.SetV(flags...)
	return flag
}

//String returns the binary formatted string
//
//implements the fmt.Stringer interface
func (b *Flag) String() string {
	return fmt.Sprintf("%b", b)
}

//Set sets a given flag to be true/on
func (b *Flag) Set(flag Flag) *Flag {
	*b |= flag
	return b
}

//SetAll sets all bits to `1` (on/true)
//
//flag will equal `0xFFFFFFFF` (max value of uint32)
func (b *Flag) SetAll() *Flag {
	*b = 0xFFFFFFFF
	return b
}

//SetV sets each flag provided to `1` (on/true)
func (b *Flag) SetV(flags ...Flag) *Flag {
	for _, flag := range flags {
		*b |= flag
	}
	return b
}

/*
# Toggle toggles the provided flag

if the provided flag is off, Toggle turns it on.

if it was off, Toggle it turns it on.
*/
func (b *Flag) Toggle(flag Flag) *Flag {
	*b ^= flag
	return b
}

//# ToggleAll toggles every bit/flag
func (b *Flag) ToggleAll() *Flag {
	*b = ^*b
	return b
}

//# ToggleV toggles each flag provided
func (b *Flag) ToggleV(flags ...Flag) *Flag {
	for _, flag := range flags {
		*b ^= flag
	}
	return b
}

//# Clear sets a provided flag to `0` (false/off)
func (b *Flag) Clear(flag Flag) *Flag {
	*b &^= flag
	return b
}

//# ClearAll sets all bits back to `0` (all false/off)
func (b *Flag) ClearAll() *Flag {
	*b = 0
	return b
}

//# ClearV sets all the provided flags to zero.
//
//Variadic version of `Clear(flag Flag)`.
func (b *Flag) ClearV(flags ...Flag) *Flag {
	for _, flag := range flags {
		*b &^= flag
	}
	return b
}

//# Has returns `true` if the provided flag is set
//
//Example:
//	const (
//		FlagA bitflag.Flag = 1 << iota // 0001
//		FlagB 			   // 0010
//		FlagC 			   // 0100
//	)
//	var flag bitflag.Flag = 0
//	flag.Set(FlagB).Has(FlagB) // true
func (b Flag) Has(flag Flag) bool {
	return b&flag == flag
}

//# HasV returns `true` if all the provided flags are set.
//
//Variadic version of `Has(flag Flag)`.
func (b Flag) HasV(flags ...Flag) bool {
	for _, flag := range flags {
		if b&flag != flag {
			return false
		}
	}
	return true
}
