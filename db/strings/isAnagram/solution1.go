package isAnagram

import "sort"

func isAnagram(s string, t string) bool {
    // Anagrams must be of the same length.
    if len(s) != len(t) {
        return false
    }

    s1 := SortedRuneArr(s)
    s2 := SortedRuneArr(t)

    return Compare(s1, s2) // Compare the sorted strings.
}

// SortedRuneArr converts a string to a rune slice, sorts it, and returns the sorted slice.
func SortedRuneArr(s string) []rune {
    r := []rune(s)
    sort.Slice(r, func(i, j int) bool{
        return r[i] < r[j]
    }) 
    return r 
}

// Compare checks if two rune slices are equal.
func Compare(a, b []rune) bool {
  for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true 
}