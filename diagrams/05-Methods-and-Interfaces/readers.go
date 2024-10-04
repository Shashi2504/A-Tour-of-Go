package main 
import ( 
	"fmt" 
	"strings" 
	"io"
)

func main() {
	r := strings.NewReader("Heeehehehee, Coliiiiiieeeee!!")
	sed := make([]byte, 8)

	for {
		n, err := r.Read(sed)
		if err == io.EOF {
			break
		}
		fmt.Printf("Read %d bytes: %s\n", n, sed[:n])
	}
}
