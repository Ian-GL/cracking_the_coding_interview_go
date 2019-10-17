package main

import (
  "fmt"
  "strings"
  )

// Implement an algorithm to determine if a string has all unique characters.
// What if you cannot use additional data structures?

func main() {
  unique := "abcde"
  dup := "aabcde"

  uniqueRes := UniqueUsingMap(unique)
  dupRes := UniqueUsingMap(dup)

  fmt.Println(uniqueRes) // => true
  fmt.Println(dupRes) // => false

  uniqueRes = UniqueUsingString(unique)
  dupRes = UniqueUsingString(dup)

  fmt.Println(uniqueRes) // => true
  fmt.Println(dupRes) // => false
}

func UniqueUsingMap(str string) bool {
  ccount := map[string]bool{}
  chars := strings.Split(str, "")

  for _, char := range chars {
    if(ccount[char]) {
      return false
    } else {
      ccount[char] = true
    }
  }

  return true
}

func UniqueUsingString(str string) bool {
  if len(str) == 1 { return true }

  split := strings.SplitAfterN(str, "", 2)

  if strings.Contains(split[1], split[0]) {
    return false
  } else {
    UniqueUsingString(split[1])
  }
  return true
}
