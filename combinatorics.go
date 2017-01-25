package combinatorics

import "sort"

func Factorial (n int64) int64 {
  var result int64 = 1
  for i := result; i <= n; i++ {
    result *= i
  }
  return result
}

func StringFrequencies (xs []string) map[string]int {
  result := make(map[string]int)
  for _, s := range xs {
    result[s] += 1
  }
  return result
}

func NumCombinations(n, k int64) int64 {
  if k > n {
    return 0
  }
  return Factorial(n) / (Factorial(k) * Factorial(n - k))
}

func StringCombinations (xs []string, k int) [][]string {
  result := make([][]string, 0)
  if k > len(xs) || k == 0 {
    return result
  }
  if k == 1 {
    for _, x := range xs {
      result = append(result, []string{x})
    }
    return result
  }
  x := xs[0]
  for _, combo := range StringCombinations(xs[1:], k - 1) {
    combo = append(combo, x)
    result = append(result, combo)
  }
  for _, combo := range StringCombinations(xs[1:], k) {
    result = append(result, combo)
  }
  return result
}

func RuneFrequencies(xs []rune) map[rune]int {
  result := make(map[rune]int)
  for _, x := range xs {
    result[x] += 1
  }
  return result
}

func RuneCombinations(xs []rune, n int) [][]rune {
  result := make([][]rune, 0)
  if n > len(xs) {
    return result
  }
  freqs := RuneFrequencies(xs)
  for _, m := range runeCombinationsHelper(freqs, len(xs), n) {
    result = append(result, frequencyMapToSlice(m))
  }
  return result
}

func Combinations(s string, n int) []string {
  result := make([]string, 0)
  runes := []rune(s)
  if n > len(runes) {
    return result
  }
  freqs := RuneFrequencies(runes)
  for _, m := range runeCombinationsHelper(freqs, len(runes), n) {
    result = append(result, frequencyMapToString(m))
  }
  return result
}

type RuneSlice []rune

func (xs RuneSlice) Len() int {
  return len(xs)
}

func (xs RuneSlice) Less(i, j int) bool {
  return xs[i] < xs[j]
}

func (xs RuneSlice) Swap(i, j int) {
  xs[i], xs[j] = xs[j], xs[i]
}

// Given a map of rune => n, return a slice with each rune repeated n times.
func frequencyMapToSlice(m map[rune]int) []rune {
  result := make([]rune, 0)
  for x, v := range m {
    for i := 0; i < v; i++ {
      result = append(result, x)
    }
  }
  return result
}

// Given a map of rune => n, return a string with each rune repeated n times.
// Runes in the resultant string are sorted.
func frequencyMapToString(m map[rune]int) string {
  runes := frequencyMapToSlice(m)
  sort.Sort(RuneSlice(runes))
  return string(runes)
}

func max(n, m int) int {
  if n > m {
    return n
  }
  return m
}

func min(n, m int) int {
  if n < m {
    return n
  }
  return m
}

func runeCombinationsHelper(m map[rune]int, n, k int) []map[rune]int {
  first := true
  var first_rune rune
  var first_freq int
  remainder := make(map[rune]int)
  for r, f := range m {
    if first {
      first_rune = r
      first_freq = f
      first = false
    } else {
      remainder[r] = f
    }
  }
  min_occurrences := max(0, k - n + first_freq)
  max_occurrences := min(k, first_freq)
  result := make([]map[rune]int, 0)
  for i := min_occurrences; i <= max_occurrences; i++ {
    if i == k {
      combo := make(map[rune]int)
      combo[first_rune] = i
      result = append(result, combo)
      break
    }
    for _, combo := range runeCombinationsHelper(remainder, n - first_freq, k - i) {
      combo[first_rune] = i
      result = append(result, combo)
    }
  }
  return result
}
