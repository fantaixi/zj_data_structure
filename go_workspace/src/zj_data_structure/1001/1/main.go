package main

import "fmt"

//复杂度 n^3
func MaxSubseqSum1(a []int, n int) int {
	var ThisSum, MaxSum int
	for i := 0; i < n; i++ { //i 是子列左端位置
		for j := i; j < n; j++ { //j 是子列右端位置
			ThisSum = 0 // a[i]到a[j]的子列和
			for k := i; k <= j; k++ {
				ThisSum += a[k]
			}
			if ThisSum > MaxSum { //如果得到的子列和更大则更新
				MaxSum = ThisSum
			}
		}
	}
	return MaxSum
}

//复杂度 n^2
func MaxSubseqSum2(a []int, n int) int {
	var ThisSum, MaxSum int
	for i := 0; i < n; i++ { //i 是子列左端位置
		ThisSum = 0
		for j := i; j < n; j++ { //j 是子列右端位置
			ThisSum += a[j]       // a[i]到a[j]的子列和
			if ThisSum > MaxSum { //如果得到的子列和更大则更新
				MaxSum = ThisSum
			}
		}
	}
	return MaxSum
}

//复杂度 n(log n)
func MaxSubseqSum3(a []int, n int) int {
	MaxSum := Divide(a, 0, n-1)
	return MaxSum
}

//分治
func Divide(a []int, left, right int) int {
	/*
		递归结束条件：子列只有一个数字
		该数为正数，最大子列和为其本身，否则为0
	*/
	if left == right {
		if 0 < a[left] {
			return a[left]
		}
		return 0
	}
	//分别递归找到左右最大子列和
	center := (left + right) / 2
	MaxLeftSum := Divide(a, left, center)
	MaxRightSum := Divide(a, center+1, right)

	//从边界向左找
	MaxLeftBorderSum := 0
	LeftBorderSum := 0
	for i := center; i >= left; i-- {
		LeftBorderSum += a[i]
		if MaxLeftBorderSum < LeftBorderSum {
			MaxLeftBorderSum = LeftBorderSum
		}
	}
	//从边界向右找
	MaxRightBorderSum := 0
	RightBorderSum := 0
	for i := center + 1; i <= right; i++ {
		RightBorderSum += a[i]
		if MaxRightBorderSum < RightBorderSum {
			MaxRightBorderSum = RightBorderSum
		}
	}
	return Max3(MaxLeftSum, MaxRightSum, MaxLeftBorderSum+MaxRightBorderSum)
}

//求三个数的最大
func Max3(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

//复杂度 n(最优解)
func MaxSubseqSum4(a []int, n int) int {
	var ThisSum, MaxSum int
	for i := 0; i < n; i++ { //i 是子列左端位置
		ThisSum += a[i]       // a[i]到a[j]的子列和
		if ThisSum > MaxSum { //如果得到的子列和更大则更新
			MaxSum = ThisSum
		} else if ThisSum < 0 { //如果当前子列和为负，则不可能使后面的部分和增大，直接抛弃
			ThisSum = 0
		}
	}
	return MaxSum
}
func main() {
	var a = []int{4, -3, 5, -2, -1, 2, 6, -2}
	n := 8
	//fmt.Println(MaxSubseqSum1(a,n))
	//fmt.Println(MaxSubseqSum2(a,n))
	fmt.Println(MaxSubseqSum3(a, n))
	//fmt.Println(MaxSubseqSum4(a,n))
}
