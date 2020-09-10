// +build ignore

package main
import (
	"fmt"
	"strings"
)

var dispensePlan = map[string]int{
	"赵":1,
	"钱":2,
}

func dispenseCoins(sum int, people []string)(
	resultMap map[string]int, rest int,
){
	rest = sum;
	resultMap = make(map[string]int)
	for _, curPersonName := range people {
		needToGive := 0
		for curFamilyName, numberToGive := range dispensePlan {
			if strings.HasPrefix(curPersonName, curFamilyName) {
				needToGive = numberToGive
				break
			}
		}
		rest -= needToGive
		resultMap[curPersonName] = needToGive
	}
	return
}

func main(){
	var plan, rest = dispenseCoins(10, []string{
		"赵一",
		"赵二",
		"钱大",
		"孙工",
		"孙国",
		"贝力",
		"安全",
	})
	fmt.Println(plan)
	fmt.Println("剩下：", rest)
}
