package main

import (
	"algorithm/dataStruct/queue/qdata"
	"fmt"
)

//十进制转其它进制(小数部分)
//不断重复将十进制小数*目标进制，每次取乘积的整数部分入队列，循环结束条件：小数部分为0或达到指定精度
//最后将队列元素依次输出(先进先出)

func decimalConv(f float64, R int) {
	if f > -1 && f < 1 && R >= 2 && R <= 16 && R != 10 {
		if f == 0 {
			fmt.Print("0.0")
			return
		}
		//是否为负数
		isNegative := false
		if f < 0 {
			f = -f
			isNegative = true
		}
		n := 0 //精度最多保留8位
		mq := new(qdata.MyQueue)
		for n < 8 && f != 0 {
			result := f * float64(R)
			el := int(result) //会默认向下取整
			//将结果统一转换为字符
			var cel byte
			switch el {
			case 10:
				cel = 'A'
			case 11:
				cel = 'B'
			case 12:
				cel = 'C'
			case 13:
				cel = 'D'
			case 14:
				cel = 'E'
			case 15:
				cel = 'F'
			default:
				cel = byte(el + 48)
			}
			mq.Push(cel)
			f = result - float64(el)
			n += 1
		}
		if isNegative {
			fmt.Print("-")
		}
		fmt.Print("0.")
		for !mq.Empty() {
			fmt.Printf("%c", mq.Pop())
		}
		return
	}
	fmt.Print("参数错误")
}

func main() {
	decimalConv(-0.86, 16)
}
