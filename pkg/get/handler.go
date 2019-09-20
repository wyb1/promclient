package get

import (
	"fmt"

	"github.com/spf13/pflag"
)

// Handler handles the get command
func Handler(args []string, flags *pflag.FlagSet) error {
	fmt.Println("Get Handler")
	return nil
}
