package Sorting

func Insert_sort(data []int) {
	if len(data) <= 1 {
		return
	}

	for i := 1; i <= len(data) - 1; i++ {
		tmpt := data[i]
		j := i - 1

		for ; j >= 0; j-- {
			if tmpt >=  data[j] {
				break
			}else {
				data[j+1] = data[j]
			}
		}

		data[j+1] = tmpt
	}
}

//希尔排序， 重要的前提： 大步长有序->小步长有序
//优化了步长
func shell_sort(data []int) {
	if len(data) <= 1 {
		return
	}

	inc := len(data) / 2
	for inc >= 1 {
		for i := inc; i <= len(data) - 1; i++ {
			tmpt := data[i]
			j := i - inc

			for ; j >= 0;  {
				if tmpt >=  data[j] {
					break
				}else {
					data[j+inc] = data[j]
				}
				j = j - inc
			}

			data[j+inc] = tmpt
		}
		inc = inc / 2
	}
}