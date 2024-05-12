# go-utils

# Provides a `Flag` (`uint32`) type.
### Can store 32 boolean values.

*by chasecarlson1, MIT license*

`Flag` also provides variadic versions of its functions that end in "V", but these are not as efficient as their normal counterparts.

In other words,

*this*: `func (b *Flag) Clear(flag Flag)`

*is preferrable to this*: `func (b *Flag) ClearV(flags ...Flag)`.

But, you can use the variadic ones if you want and are okay with a little slice **heap allocation** overhead (which isn't much overhead because these are `uint32` anyways).

# Example:

```go
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
```

## todo

Figure out where to use `//go:` directives to make it even more efficient.