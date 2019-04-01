package main



import "fmt"

func adder()     func(returnValue int) int      { //闭包
	sum := 0
	return func(parI int) int {
		sum += parI
		return sum
	}
}

func main()  {
	// 调用函数adder() 返回匿名函数 func(parI int) int {
	//		sum += parI
	//		return sum
	//	}
	a := adder()

	// 10 次调用都是调用a  10次操作的都是a的 sum := 0  ,操作同一个变量 sum
	for i:= 0;i < 10; i++{
		fmt.Printf("0 + 1 + 2 + ... + %d = %d \n",i,a(i))
	}
}
