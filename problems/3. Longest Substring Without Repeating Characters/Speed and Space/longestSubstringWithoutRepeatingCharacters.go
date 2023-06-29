// The use of the constant is only for style, in order to achieve maximum speed it may have to be removed
const asciiSize = 256
func lengthOfLongestSubstring(s string) int {
    hist := make([]int, asciiSize)
    maxLen := 0
    start := 0
    
    for i := 0; i < len(s); i++ {
        // Checks if the current character has been seen before
        if start < hist[s[i]] {
            start = hist[s[i]]
        }
        if maxLen < i - start + 1 {
            maxLen = i - start + 1
        }
        hist[s[i]] = i + 1
    }
    return maxLen
}