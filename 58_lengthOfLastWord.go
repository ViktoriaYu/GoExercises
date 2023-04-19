/*
Given a string s consisting of words and spaces,
return the length of the last word in the string.

A word is a maximal substring consisting of non-space characters only.
*/

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for string(s[i]) == " " {
		i--
	}
	k := 0
	for i >= 0 && string(s[i]) != " " {
		k++
		i--
	}
	return k
}