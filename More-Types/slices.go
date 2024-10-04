package main
import "fmt"

func main() {
	marvel_dc := [6] string {"Batman", "Cap America", "Hulk", "Hawkeyee", "Superman"}

	marvel := marvel_dc[1:4]

	fmt.Println("Identified as avengers members:", marvel)


	// Created another array and slice here
	HIMYM := [] string {"Ted", "Marshal", "Lilly", "BarNey", "Robin"}
	started := HIMYM[:3]
	fmt.Println("How I Met Your Mother is started with these three peeeple:", started)

	// Checking with append like adding the things if you forgot to add

	HIMYM = append(HIMYM, "tracy")

	fmt.Println("Here is the final character:", HIMYM)
}
