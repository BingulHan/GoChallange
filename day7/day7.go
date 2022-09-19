//Today we learned about functions!
//Bugün fonksiyonları öğrendik!

package main

import  (
	"fmt" )
/*

func function_name( [parameter list] ) [return_types]
{
   body of the function
}
*/

//is the function from which the application starts.
func main(){

	//Calling a Function
	fmt.Println("Min:",min(1,2))

	//Calling a Function
	fmt.Println("Max:",max(1,2))


}

func min(var1 int , var2 int) int {
	if var1<var2  {
		return var1
	}else {
		return var2
	}
}

func max(var1 int , var2 int) int {
	if var1>var2  {
		return var1
	}else {
		return var2
	}
}
