/*
Roman numerals are represented by seven different symbols: 
I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

Given a roman numeral, convert it to an integer.
*/

func romanToInt(s string) int {
    result := 0
    conv := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
    for i := 0; i < len(s) - 1; i++ {
        if conv[s[i]] >= conv[s[i+1]] {
            result += conv[s[i]]
        } else {
            result -= conv[s[i]]
        }
    }
    result += conv[s[len(s)-1]]
    return result
}