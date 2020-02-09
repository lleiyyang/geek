package datastructure

import "testing"

func Test_queue(t *testing.T) {
	/** obj := Constructor(k);
	* param_1 := obj.EnQueue(value);
	* param_2 := obj.DeQueue();
	* param_3 := obj.Front();
	* param_4 := obj.Rear();
	* param_5 := obj.IsEmpty();
	* param_6 := obj.IsFull();*/
	obj := Constructor(3)
	datas := []int{1,2,3,4}
	for _, data := range datas {
		t.Log(obj.EnQueue(data))
	}
}