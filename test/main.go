package main


import "fmt"

type Person struct {
	name string
	surname string
	age int
}

func main(){
	dim := 10

	array := [...]int{1,2,3,4,5,6,7,8,9,10}

	fmt.Println("Array length is: ", len(array))
	fmt.Println(array);

	for i:=0; i<dim; i++{
		array[i] = array[i] * 2
		fmt.Println(array[i])
	}

	var i int = 0;
	for ;i< dim; i++{
		fmt.Println(array[i])
	}

	for i<dim {
		fmt.Println(array[i])
		i++
	}

	arr2 := [...]int{1,3,4,5,87}

	ptr := &arr2
	fmt.Println("Pointer to array: ", ptr)
	fmt.Println("Pointer to array: ", *ptr)
	fmt.Println(ptr[0])
}