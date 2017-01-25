package combinatorics

import (
  "fmt"
  "reflect"
  "testing"
)

var factorial_tests []int64 = []int64{1,1,2,6,24,120,720}

func TestFactorial(t *testing.T) {
  for n, expected := range factorial_tests {
    observed := Factorial(int64(n))
    if observed != expected {
      t.Errorf("factorial(%d) => %d, expected %d", n, observed, expected)
    }
  }
}

var string_frequency_tests = []struct{
  input []string
  expected map[string]int
}{
  {[]string{"the", "cat", "sat", "on", "the", "mat"},
    map[string]int{"the": 2, "cat": 1, "sat": 1, "on": 1, "mat": 1}},

  {[]string{"the", "cat", "is", "the", "cat", "is", "the", "cat"},
    map[string]int{"the": 3, "cat": 3, "is": 2}},

  {[]string{}, map[string]int{}}}

func TestStringFrequencies(t *testing.T) {
  for _, tt := range string_frequency_tests {
    observed := StringFrequencies(tt.input)
    if ! reflect.DeepEqual(observed, tt.expected) {
      t.Errorf("StringFrequecies(%v) => %v, expected %v", tt.input, observed, tt.expected)
    }
  }
}

var num_combinations_tests = []struct{n int64; k int64; expected int64}{{3, 0, 1}, {3, 1, 3}, {3, 2, 3}, {3, 3, 1}, {3, 4, 0}}

func TestNumCombinations(t *testing.T) {
  for _, tt := range num_combinations_tests {
    observed := NumCombinations(tt.n, tt.k)
    if observed != tt.expected {
      t.Errorf("NumCombinations(%d,%d) => %d, expected %d", tt.n, tt.k, observed, tt.expected)
    }
  }
}

func TestStringCombinations(t *testing.T) {
  fmt.Println(StringCombinations([]string{"A", "B", "C", "D"}, 3))
}

func TestRuneCombinationsHelper(t *testing.T) {
  //s := "aaabbc"
  s := "abcdef"
  for n := 0; n < 8; n++ {
    fmt.Println(Combinations(s, n))
  }
}
