/*
Given an integer x, return true if x is a
palindrome, and false otherwise.
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	x_copy := x
	x_rev := 0
	for x_copy > 0 {
		x_rev *= 10
		x_rev += x_copy % 10
		x_copy /= 10
	}
	return x == x_rev
}