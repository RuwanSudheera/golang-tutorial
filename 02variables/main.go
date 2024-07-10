package main

import "fmt"

const LoginToken string = "werwer1432sds"

func main() {
	var username string = "ruwan"
	fmt.Println(username)
	fmt.Printf("Variable is type of : %T \n", username)

	var isLogged bool = true
	fmt.Println(isLogged)
	fmt.Printf("Variable is type of : %T \n", isLogged)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is type of : %T \n", smallVal)

	var smallFloat float32 = 255.3434345
	fmt.Println(smallFloat)
	fmt.Printf("Variable is type of : %T \n", smallFloat)

	//implicit type
	var website = "trekinsight.com"
	fmt.Println(website)

	//no var style (this allowd only within a function)
	numOfVal := 3000
	fmt.Println(numOfVal)

	//print a const variable
	fmt.Println(LoginToken)
	fmt.Printf("Variable is type of : %T \n", LoginToken)
}
