package isAnagram 

func isAnagram2(s string, t string) bool {
    // Anagrams must be of the same length.
    if len(s) != len(t) {
        return false
    }

    // Create a map to count the occurrences of each rune in s.
    count := make(map[rune]int)
    for _, char := range s {
        count[char]++
    }

    // Decrement the count for each rune in t.
    for _, char := range t {
        count[char]--
        if count[char] < 0 {
            return false
        }
    }

    // If all counts are zero, s and t are anagrams.
    return true
}
