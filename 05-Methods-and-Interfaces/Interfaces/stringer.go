package main
import "fmt"

type Me struct {
	Name string
	accu int
}

func (m Me) String() string {
	return fmt.Sprintf("%s your accuracy is %d", m.Name, m.accu)
}

func main() {
	m := Me{"Ashhh", 56}
	fmt.Println(m)
}
