package main

import (
"container/list"
"fmt"
)

func main() {
	dataList := list.New()

	dataList.PushBack(1)	// 插入末尾
	dataList.PushBack(2)
	dataList.PushFront(3)	 // 插入表头
	dataList.PushBack(4)
	dataList.PushBack(5)
	//31245
	m := dataList.PushBack(6)
	m1 := dataList.InsertBefore(7,m)	// 6 之前插入 7
	m2 := dataList.InsertAfter(8,m)	// 6 之后插入 8
	//31245768
	// 从链表头开始遍历
	for e := dataList.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value) // 打印值
	}

	fmt.Println("----------------------------------------")

	dataList.Remove(dataList.Front())	// 移除头部
	dataList.MoveBefore(m2, m)	// 将m2移动m之前
	dataList.MoveAfter(m1, m)
	dataList.Remove(m)	// 移除
	//124587

	//PushBackList	// 插入列表
	//PushFrontList	//

	// 从链表头开始遍历
	fmt.Println()
	for e := dataList.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value) // 打印值
	}

	fmt.Println("----------------------------------------")

	fmt.Println()
	//124587->785421
	// 从链表尾开始遍历
	for e := dataList.Back(); e != nil; e = e.Prev() {
		fmt.Printf("%v ", e.Value) // 打印值
	}

	fmt.Println("----------------------------------------")
	dataList.Init()	// 清空链表
	// 从链表头开始遍历
	for e := dataList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // 打印值
	}
}