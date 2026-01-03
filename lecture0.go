package main
import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	welcome := "welcome to Golang"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for our pizza")


	// comma ok sntax // error ok 
	input , _ := reader.ReadString('\n')
	fmt.Println("thanks for ur rating" , input)
	fmt.Printf("type of the rating is : %T \n", input)
	
}