package main
import "fmt"
var book [9999999]uint64
var maxone int
var i uint64

func main() {
 nums := []int{}//测试用例要手动修改 希望不要测大了(笑哭)
 
 x:=len(nums)
 for i:=0;i<x-1;i++{
	 if nums[i]<nums[i+1]{
          maxone=nums[i+1]
	 }
 }
	if x==0{
		fmt.Println("0")
	}else{
	for i:=0;i<x;i++{
		book[nums[i]]++
	}
	for i:=0;i<maxone+1;i++{
		if book[i]==1{
			fmt.Println(i)
          
		}
		
	}

	}
}