//We learned struct today!
//Bugün struct öğrendik!

package main

import  (
	"fmt" )

//is the function from which the application starts.
func main(){


	var product1 Product 
	product1.name = "Leptop"
	product1.price = 4000


	var product2 Product
	product2.name = "Mouse"
	product2.price = 100

	fmt.Println("Product1 name is ",product1.name)
	fmt.Println("Product1 price is ",product1.price)

	fmt.Println("-----------------------")

	fmt.Println("Product2 name is ",product2.name)
	fmt.Println("Product2 price is ",product2.price)



}

//a struct is defined like this
//bir struct bu şekilde tanımlanır
type Product struct {
	name string
    price int
}