package printer

import (
	"fmt"
	"runtime"
)

// Hello function of package printer
func Hello() {
	fmt.Println("expected hello")
}

// Version with input string
func Version() string {
	return runtime.Version()
}
