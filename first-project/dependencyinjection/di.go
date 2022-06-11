package dependencyinjection

import (
	"fmt"
	"io"
)

func GreetWithDI(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s!", name)
}
