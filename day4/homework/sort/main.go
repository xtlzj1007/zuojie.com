package main

import (
	"fmt"
)

func BubbleSort(s []int) {
	/*
	冒泡排序
    它重复地走访过要排序的数列，
    一次比较两个元素，如果他们的顺序错误就把他们交换过来。
    走访数列的工作是重复地进行直到没有再需要交换，
    也就是说该数列已经排序完成。
	 */
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
		fmt.Println(s)
	}

}

func SelectSort(s []int) {
	/*
	选择排序
    1.先从整个序列中选择最小的数据放到第一位
    2.再从剩余的序列中选择最小的数据放在第二位
    3.如此循环，直到最后一位。
	 */
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
		fmt.Println(s)
	}
}

func InsertSort(s []int) {
	/*
	插入排序
    1.每次取一个列表元素与列表中已经排序好的列表段进行比较
    2.然后插入从而得到新的排序好的列表段，最终获得排序好的列表。
	 */
	for i := 1; i < len(s); i++ {
		value := s[i]
		j := i
		for j > 0 && s[j-1] > value {
			s[j] = s[j-1]
			j -= 1
		}
		s[j] = value
		fmt.Println(s)
	}

}

func QuickSort(arr []int, start, end int) {
	/*
	递归快排
	通过一趟排序将要排序的数据分割成独立的两部分，
	其中一部分的所有数据都比另外一部分的所有数据都要小，
	然后再按此方法对这两部分数据分别进行快速排序，
	整个排序过程可以递归进行，以此达到整个数据变成有序序列。
	 */
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2] //获取基准
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			QuickSort(arr, start, j)
		}
		if end > i {
			QuickSort(arr, i, end)
		}
	}
}

func main() {
	array := []int{6, 5, 4, 3, 2, 1}
	//BubbleSort(array)
	//SelectSort(array)
	//InsertSort(array)
	QuickSort(array, 0, len(array)-1)
	fmt.Println(array)
}



