package main
import "fmt"

type bheem interface {
	mOEEE()
}

type raju struct {
	reeo string
}

func (j raju) mOEEE() {
	fmt.Println(j.reeo)
}

func main () {
	var chu bheem = raju{"Heheeeeeeeeeeeeeeeeeh"}
	chu.mOEEE()
}
