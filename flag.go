package flags

import (
	"github.com/aurumbot/lib/dat"
	"strings"
)

// A Flag is used to store a single flags data.
//
// Fields:
//  - Type: "-" or "--"
//  - Name: Name of the flag.
//      Ex: --name gabe miller --> Name = "name"
//  - Values: Single string of values after flag.
//      Ex: --name gabe miller --> Value = "gabe miller"
//
type Flag struct {
	Name  string
	Value string
}

// Parse parses a message for flags.
//
// Parameters:
// - args (string) | A string message
//
// Returns:
// - ([]*Flag) | A slice of each flag type
//
func Parse(msg string) []*Flag {
	args := strings.Split(msg, " ")
	flags := []*Flag{}
	var cur *Flag
	flags = append(flags, *Flag{
		Name:  "unflagged",
		Value: ""})
	for _, arg := range args {
		switch {
		case arg[0] == '-':
			cur = &Flag{
				Name: arg,
			}
			flags = append(flags, cur)
		default:
			if len(flags) > 0 {
				flags[len(flags)-1].Value += arg + " "
			} else {
				flags[0].Value += arg
			}
		}
	}
	// removes whitespace from flag values
	for f := range flags {
		flags[f].Value = strings.Trim(flags[f].Value, " ")
	}
	return flags
}
