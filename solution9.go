package main

import "fmt"

func main(){
   findThemTriplets(3)
}

func findThemTriplets(m int)bool{
	var set []int 
    for n:=2; n<999; n++{
        set = tripletGen(m, n)
    	if sum(set) == 1000{
    		fmt.Println(set)
    		fmt.Println(set[0] * set[1] * set[2])
    		return true
    	} 
    }
    findThemTriplets(m + 1)

    return false
}

func tripletGen(m, n int) []int{
	var a,b,c int
    rslice := make([]int,0,3)   

    if m > n{ 
        a = (m*m) - (n*n)
        b = 2 * (m * n)
        c = (m*m) + (n*n)
        rslice = append(rslice, a,b,c)
    }

    return rslice
}

func sum(set []int)int{
	total := 0
	for _, n := range set{
		total += n
	}

	return total
}
