package main

import "fmt"

//e = (1+1/n)^n

func main()  {
	var n float64
	var e float64
	var aux float64

	fmt.Scanln(&n)

	for i:= 1.0; i <= n; i++ {

		e = (1+(1/i))
		if i > 1 {
			aux = e
			for j:= 2.0; j <= i; j++{
				e = e * aux 
			}
		}

		//fmt.Println(e)
	}

	fmt.Println(e)
	
}