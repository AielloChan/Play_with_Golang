package qsort

import (
    "testing"
)

// TestQuikSort1 测试排序
func TestQuikSort1(t *testing.T)  {
    values := []int {5, 4, 3, 2, 1}
    QuikSort(values)

    if values[0] == 1 && values[1] == 2 && values[2] == 3 && values[3] == 4 && values[4] == 5{

    } else { 
        t.Error("QuikSort() faild. Got ", values, "Expected 1 2 3 4 5")
    }
}

// TestQuikSort2 测试存在相同元素的情况
func TestQuikSort2(t *testing.T)  {
    values := []int {5, 5, 3, 2, 1}
    QuikSort(values)

    if values[0] == 1 && values[1] == 2 && values[2] == 3 && values[3] == 5 && values[4] == 5{

    } else { 
        t.Error("QuikSort() faild. Got ", values, "Expected 1 2 3 5 5")
    }
}

// TestQuikSort3 测试单个元素的情况
func TestQuikSort3(t *testing.T)  {
    values := []int {5}
    QuikSort(values)

    if values[0] != 5{
        t.Error("QuikSort() faild. Got ", values, "Expected 5")
    }
}

// TestQuikSort4 测试负数情况
func TestQuikSort4(t *testing.T)  {
    values := []int {5, 4, -1, 3, 2, 1}
    QuikSort(values)

    if values[0] == -1 && values[1] == 1 && values[2] == 2 && values[3] == 3 && values[4] == 4 && values[5] == 5{

    } else { 
        t.Error("QuikSort() faild. Got ", values, "Expected -1 1 2 3 4 5")
    }
}