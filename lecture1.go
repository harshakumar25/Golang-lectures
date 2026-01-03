package main
import "fmt"

//outside 
var vwtoken = 3000000  //global varaible (value declare outside fucntion)

const LoginToken string = "harsha123"  //const is used to declare constant value which cannot be changed

func main() {
	// var declare string
	var username string = "harsha"
	fmt.Printf("variable is of type  :  %T  \n", username  )  //%T is used to know the type of variable 
//var declare bool
	var islooggedin bool = false
	fmt.Println(islooggedin)
	fmt.Printf("variable is of type : %T \n" , islooggedin )

//var declare for uint8 and float32
	var smallval uint8 = 255  //uint8 can store values from 0 to 255
    fmt.Println(smallval)
	fmt.Printf("variable is of type : %T \n", smallval )
	
	var smallFloat float64 = 255.45544511254451885
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type : %T\n", smallFloat ) // use printf to format the output


	//default values and aliases
	var anothervariable int 
	fmt.Println(anothervariable)
	fmt.Printf("variable is of type : %T\n" , anothervariable)
	
	//implicit type declare
	var website = "harsha.com"
	fmt.Println(website)
	fmt.Printf("variable is of type : %T\n" , website)

	//no var style declare
	numberofusers := 3000;   // valor operator :=
	fmt.Println(numberofusers)
	fmt.Printf("variable is of type : %T\n", numberofusers)

	//using const
	fmt.Println(LoginToken)
	fmt.Printf("variable is of type : %T \n", LoginToken)

	//using global varaiable 
	fmt.Println(vwtoken)
	fmt.Printf("varaible is of type : %T \n", vwtoken)

}