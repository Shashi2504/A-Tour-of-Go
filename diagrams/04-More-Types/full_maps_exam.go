package main
import "fmt"


func main() {
	// Created the map using "make"
	HIMYM := make(map[string]int)
		HIMYM["Lily Aldrin"] = 5
		HIMYM["Marshall"] = 5
		HIMYM["Ted"] = 4

		fmt.Println("HIMYM started with:", HIMYM)


	// Updating the above map with adding one more person or we can update the value
	HIMYM["BarNey"] = 6
	HIMYM["Ted"] = 5
	fmt.Println("Here comes the awesome characters:", HIMYM)

	// Now retrieving the values

	sitcom := HIMYM["Lily Aldrin"]
	fmt.Println("This HIMYM sitcom is", sitcom)

	// Deleting the things
	delete(HIMYM, "Marshall")
	fmt.Println(HIMYM)

	// Doing the checklist and retrieving the value
	about_robin, exists := HIMYM["Robin"]

	if exists {
		fmt.Println("Robin is already here", about_robin)
	} else {
		fmt.Println("Start Entering", about_robin)
	}

	// Retreving the key which doesn't exist

	_, exists = HIMYM["Robin"]

	if exists {
		fmt.Println("Robin character has started")
	} else {
		fmt.Println("Preparing to add Ribin character")
	}
}
