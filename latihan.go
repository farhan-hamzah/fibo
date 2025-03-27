package main
import "fmt"

func fibo(n int)int{
	if n == 1{
		return 1
	}else if n == 2{
		return 1
	}else if n == 3{
		return 4
	}else {
		return fibo(n-1) + fibo(n-2) + fibo(n-3)
	}
}
func main(){
	var num int
	fmt.Scan(&num)
	fmt.Print(fibo(num))
}