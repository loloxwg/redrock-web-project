
package main

import (

"fmt"
)

func main()  {

	var a ,b ,c int
	var x string


	for i := 0; i >= 0; i++ {
		fmt.Println("input:")
		fmt.Scanf("%d\n%s\n%d\n",&a,&x,&b)
		switch x {
		case "+":
			c=a+b
			fmt.Print("output:\n",c)

		case "-":
			c=a-b
			fmt.Print("output:\n",c)
		case "*":
			c=a*b
			fmt.Print("output:\n",c)
		case "/":
			c=a/b
			fmt.Print("output:\n",c)
		}
		fmt.Printf("\n")
	}
}
