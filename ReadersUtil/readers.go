package ReadersUtil

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func StringStream() error {
	str := strings.NewReader("Richesh")

	b := make([]byte, 8)

	for {
		n, err := str.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.New("Something Awkward Occurred")
		}
		fmt.Printf(" Read b[:n] = %q\n", b[:n])
	}
	return nil

}
