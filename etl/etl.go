package etl

import "strings"

// Transform
func Transform(pre map[int][]string) map[string]int {
  post := make(map[string]int, 26)

  for k, v := range pre {
    for _, letter := range v {
      post[strings.ToLower(letter)] = k
    }
  }

  return post
}
