//package main
//
//import "fmt"
//
//func main() {
//	var arr = [5]int{2, 20, 15, 22, 16}
//	var max = arr[0]
//
//	for i := 1; i < len(arr); i++ {
//		if arr[i] > arr[i-1] {
//			max = arr[i]
//		}
//	}
//	fmt.Printf("%d", max)
//
//}


//选择排序
//package main
//
//import "fmt"
//
//func main() {
//	var arr = [5]int{2, 20, 15, 22, 16}
//	var temp int
//
//	for i := 0; i < len(arr)-1; i++ {
//		for j:=i+1;j<len(arr);j++ {
//			if arr[i]<arr[j] {
//				temp=arr[i]
//				arr[i]=arr[j]
//				arr[j]=temp
//			}
//		}
//	}
//	for i:=0;i<len(arr);i++ {
//		fmt.Printf("%d ",arr[i])
//	}
//
//}

package main

import "fmt"

func main() {
	var arr = [5]int{2, 20, 15, 22, 16}
	var temp int

	for i := 0; i < len(arr)-1; i++ {
		for j:=0;j<len(arr)-1-i;j++ {
			if arr[j]<arr[j+1] {
				temp=arr[j]
				arr[j]=arr[j+1]
				arr[j+1]=temp
			}
		}
	}
	for j:=0;j<len(arr);j++ {
		fmt.Printf("%d ",arr[j])
	}

}