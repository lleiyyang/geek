package leetcode

import "testing"

func Test_leastInterval(t *testing.T) {
	tasks := []byte{'A','A','A','B','B','B'}
	cnt := leastInterval(tasks,2)
	t.Logf("cnt:%v", cnt)
}
