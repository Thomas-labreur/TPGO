package main
import (
	"fmt"
	"bufio"
	"os"
)

func input(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(message)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	rep := input("Entrer une phrase:")
	fmt.Printf("Votre phrase est :%v", rep)
}