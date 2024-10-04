package main
import (
	"fmt"
	"strings"
	"io"
)

func main() {
	k := strings.NewReader("Heheheheh chinnnthaaa math")
	chill := make([]byte, 9)

	for {
		d, err := k.Read(chill)
		fmt.Printf("d = %v d = %v, err = %v")

		fmt.Printf("chill[:d] = %q\n", chill[:d])

		if err == io.EOF {
			break
		}
	}
}
