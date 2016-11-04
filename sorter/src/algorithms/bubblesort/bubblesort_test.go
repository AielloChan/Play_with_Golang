package bubblesort

import (
	"testing"
)

// TestBubbleSort1 测试排序
func TestBubbleSort1(t *testing.T)  {
    values := []int {5, 4, 3, 2, 1}
    BubbleSort(values)

    if values[0] == 1 && values[1] == 2 && values[2] == 3 && values[3] == 4 && values[4] == 5{

    } else { 
        t.Error("BubbleSort() faild. Got ", values, "Expected 1 2 3 4 5")
    }
}

// TestBubbleSort2 测试存在相同元素的情况
func TestBubbleSort2(t *testing.T)  {
    values := []int {5, 5, 3, 2, 1}
    BubbleSort(values)

    if values[0] == 1 && values[1] == 2 && values[2] == 3 && values[3] == 5 && values[4] == 5{

    } else { 
        t.Error("BubbleSort() faild. Got ", values, "Expected 1 2 3 5 5")
    }
}

// TestBubbleSort3 测试单个元素的情况
func TestBubbleSort3(t *testing.T)  {
    values := []int {5}
    BubbleSort(values)

    if values[0] != 5{
        t.Error("BubbleSort() faild. Got ", values, "Expected 5")
    }
}

// TestBubbleSort4 测试负数情况
func TestBubbleSort4(t *testing.T)  {
    values := []int {5, 4, -1, 3, 2, 1}
    BubbleSort(values)

    if values[0] == -1 && values[1] == 1 && values[2] == 2 && values[3] == 3 && values[4] == 4 && values[5] == 5{

    } else { 
        t.Error("BubbleSort() faild. Got ", values, "Expected -1 1 2 3 4 5")
    }
}