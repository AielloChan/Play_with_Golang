package bubblesort

// BubbleSort 冒泡排序
func BubbleSort(values []int) {

	for i := 0; i < len(values); i++ {

		for j := i + 1; j < len(values); j++ {

			if values[i] > values[j] {
				
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}
