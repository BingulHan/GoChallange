//How to define a variable today and I've made an introduction to basic variables
//Bugün değişken nasıl tanımlanır ve temel değişkenlere giriş yaptım

package main

import "fmt"

//is the function from which the application starts.
func main(){

    //Variable definition
    var age int
    age = 17
    _ = age

    var name string
    name = "Mustafa"
    _ = name
    
    var discord string
    discord = "BingulHan#4912"
    _ = discord

    var married bool 
    married = false
    _ = married
    
    //print to screen
    fmt.Println(name)
    fmt.Println(age)
    fmt.Println(discord)
    fmt.Println(married)

    
	
}