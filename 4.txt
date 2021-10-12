package main
import "fmt"
var book [40]int
var i int
var q int
var booknew []string
var tongji [200009]int
var p int
func Compression(ch string) []string {
 // insert your code
 x:=len(ch)
 

 for i:=0;i<x-1;i++{
	 if ch[i]==ch[i+1]{
          	     
	 }else{
		 booknew=append(booknew,string(ch[i]))
	 }
	
}
p:=0
for i:=x-1;i>=1;i--{
	if ch[i]==ch[i-1]{
		p++
	}else{

		break
	}

}

p=p+1

if ch[x-p-1]!=ch[x-p]{
	booknew=append(booknew,string(ch[x-p]))
}
return booknew
} 

func main() {
 ch := "qqqqqqqqqqqqq11111111111111111wwwwwwwertttttttttrrr"
 fmt.Println(Compression(ch))
}