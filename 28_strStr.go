/*
Given two strings needle and haystack,
return the index of the first occurrence of needle in haystack,
or -1 if needle is not part of haystack.
*/

func strStr(haystack string, needle string) int {
	k := -1
	for i := 0; i < len(haystack); i++ {
		if i+len(needle) <= len(haystack) && haystack[i:i+len(needle)] == needle {
			k = i
			break
		}
	}
	return k
}