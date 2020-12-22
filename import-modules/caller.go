package main

import (
	"fmt"
	// ▼ 引用自定义的包
	// 工程名开头（即 go.mod 里面的工程名）这里是【learning-go-01】
	// 工程名/文件夹/文件夹
	"learning-go-01/import-modules/module01"

	// ▼ 引用自定义的包（不良的包命名）
	// 这里包的目录是 folder02 但包名是 badPackageName
	// 不提倡如此给包命名
	"learning-go-01/import-modules/folder02"
)

func main() {
	module01.ToSayHello()
	sum := badpackagename.GetSumFn(6, 6, 6)
	fmt.Println("和：", sum)
	
	// badpackagename.Abc()
	biggest := badpackagename.GetBiggest(1, 2, 3)
	biggest02 := badpackagename.GetBiggest([]int{4,5,6}...)
	biggest03 := badpackagename.GetBiggest02([]interface{}{
		"abc", 4, 5, 6.789, float32(9.9),
	})
	fmt.Println("最大值：", biggest)
	fmt.Println("最大值02：", biggest02)
	fmt.Println("最大值03：", biggest03)
}

