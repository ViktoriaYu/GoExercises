/*
Write a function to find the longest common prefix string 
amongst an array of strings.

If there is no common prefix, return an empty string "".
*/

/*
func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    if len(strs) == 1 {
        return strs[0]
    }

    compref := ""
    for i := 0; i < Min(len(strs[0]), len(strs[1])); i++ {
        if strs[0][i] == strs[1][i]{
            compref += string(strs[0][i])
        } else {
            break
        }
    }

    for i := 2; i < len(strs); i++ {
        if len(strs[i]) == 0 {
            return ""
        }

        if len(strs[i]) < len(compref) {
            compref = compref[:len(strs[i])]
        }

        for j := 0 ; j < len(compref); j++ {
            if strs[i][j] != compref[j] {
                compref = compref[:j]
            }
        }
    }

    return compref
}
*/

func longestCommonPrefix(strs []string) string {
    if strs == nil || len(strs) == 0 {
        return ""
    }
    for i := 0; i < len(strs[0]); i++ {
        for j := 1; j < len(strs); j++ {
            if i == len(strs[j]) || strs[j][i] != strs[0][i] {
                return strs[0][:i]
            }
        }
    }
    return strs[0]
}