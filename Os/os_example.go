package Os

import (
	"fmt"
	"os"
)

func PrintOsVariables() {
	fmt.Println(os.Environ())
}
