package main

import (
	"fmt"
	"time"
)

func getBigNumber(iTimes int) float64{
	iResult := 0.0
	for idx := 0; idx < iTimes; idx++ {
		iResult += 0.001
	}
	return iResult;
}

func main() {
	/* 
		console.time('测试1');
		var sum = 0;
		for (var idx=0; idx<300000000; idx++){
			sum+=0.001;
		}
		console.timeEnd('测试1');
		console.log('测试结果', sum);
		测试1: 2483.75390625 ms
		测试结果 299999.99833619915
	*/
	startTime := time.Now().UnixNano()
	iResult := getBigNumber(3 * 100000000)
	endTime := time.Now().UnixNano()
	// seconds:= float64((endTime - startTime) / 1e9)
	Milliseconds:= float64((endTime - startTime) / 1e6)
	// nanoSeconds:= float64(endTime - startTime)
	fmt.Println("结果", iResult)
	fmt.Println("耗时", Milliseconds)
	// 结果 299999.99833619915
	// 耗时 345
}