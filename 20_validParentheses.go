/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', 
determine if the input string is valid.

An input string is valid if:
    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.

*/

import "strings"

func isValid(s string) bool {
    for strings.Contains(s, "()") || strings.Contains(s, "[]") || strings.Contains(s, "{}") {
        s = strings.Replace(s, "()", "", -1)
        s = strings.Replace(s, "[]", "", -1)
        s = strings.Replace(s, "{}", "", -1)
    }
    return s == ""
}