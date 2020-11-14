// +build ignore

package main
import ("fmt")

func numToStr (num int) (string) {
	if num==0 {
		return "■■■ ■■■"
	}
	return "■■■■■■■"
}

func main() {
	var flag = 0
	for ii := 0; ii <= 1; ii++ {
		for jj := 0; jj<=1; jj++{
			for kk:=0; kk<=1; kk++{
				flag++
				myStr := numToStr(ii) + "\n" + numToStr(jj) + "\n" + numToStr(kk) + "\n"
				fmt.Println("--序号:", flag, "----------")
				fmt.Print(myStr)
			}
		}
	}
}

