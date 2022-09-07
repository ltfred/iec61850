package worker

import (
	"fmt"
	"github.com/ltfred/iec61850"
)

//Task 数据处理任务
func Task(data *iec61850.APDU) {
	//TODO 自定义数据处理
	println("do task")
	fmt.Println(*data.ASDU)
	fmt.Println(*data.APCI)
	fmt.Println(*data)
}
