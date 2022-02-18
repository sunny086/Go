package main

import (
	"errors"
	"fmt"
)

func main() {
	// 创建错误方式
	createError()

	//recover处理panic
	recoverPanic()

}

func recoverPanic() {
	defer func() {
		if msg := recover(); msg != nil {
			//fmt.Println(msg)
			fmt.Println("recover")
		}
	}()
	funA()
}

func funA() {
	defer fmt.Println("panic1")
	defer fmt.Println("panic2")
	defer fmt.Println("panic3")
	for i := 1; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			//中断
			panic("funA-panic")
		}
	}
	defer fmt.Println("----------")
}

// createError 创建错误方式
func createError() {
	//1创建错误方式（1）
	err1 := errors.New("do by try")
	fmt.Println(err1)
	fmt.Printf("%T\n", err1)

	//2创建错误方式（2）
	err2 := fmt.Errorf("错误信息码%d", 404)
	fmt.Println(err2)
	fmt.Printf("%T\n", err2)

	err3 := checkAge(1)
	fmt.Println(err3)
}

func checkAge(age int) error {
	if age < 0 {
		err := fmt.Errorf("age must be positive")
		return err
	}
	fmt.Printf("age = %d", age)
	return nil
}
