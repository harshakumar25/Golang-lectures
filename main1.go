package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	fmt.Println("welcome to our pizza app")
	fmt.Println("please rate our pizza between 1 and  5")

	reader := bufio.NewReader(os.Stdin)

    input, _ :=	reader.ReadString('\n')

	fmt.Println("thanks for ypurs rating ", input)

	fmt.Printf("type of input is. : %T \n",input)

	if err != nil {
		fmt.Println("error in conversion" , err)
	}else {
		fmt.Println("added 1 to your rating " , numRating+1)

	}


}
