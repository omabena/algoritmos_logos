package main

import "fmt"

func product(a int, b int) int {
	sum := 0
	for i := 0; i < b; i++ {
		sum += a
	}
	return sum
}

func power(a int, b int) int {
	if b < 0 {
		return 0
	} else if b == 0 {
		return 1
	} else {
		return a * power(a, b-1)
	}
}

func mod(a int, b int) int {
	if b <= 0 {
		return -1
	}

	div := a / b
	return a - div*b
}

func div(a int, b int) int {
	count := 0

	for sum := b; sum <= a; sum = sum + b {
		fmt.Println(count)
		count++
	}
	return count
}

func sqrt(n int) int {
	var sqrtHelper func(int, int, int) int

	sqrtHelper = func(n int, min int, max int) int {
		if max < min {
			return -1
		}
		guess := (min + max) / 2
		fmt.Println(guess)
		if guess*guess == n {
			return guess
		} else if guess*guess < n {
			return sqrtHelper(n, guess+1, max)
		} else {
			return sqrtHelper(n, min, guess-1)
		}
	}

	return sqrtHelper(n, 1, n)
}

func sqrtBigNumber(n int) int {
	for guess := 1; guess*guess <= n; guess++ {
		if guess*guess == n {

			return guess
		}
	}
	return -1
}
