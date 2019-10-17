package main

import (
  "fmt"
  "strings"
  "sort"
)

// Given two strings, write a method to decide if one is a permutation of the other

func main() {
  perm1 := "iamlordvoldemort"
  perm2 := "tommarvoloriddle"

  notperm1 := "acb"
  notperm2 := "cde"

  truthy := ArePerm(perm1, perm2)
  falsy := ArePerm(notperm1, notperm2)

  fmt.Println(truthy) // => true
  fmt.Println(falsy) // => false
}

func ArePerm(word1 string, word2 string) bool {
  sorted1 := LowerAndSort(word1)
  sorted2 := LowerAndSort(word2)

  return sorted1 == sorted2
}

func LowerAndSort(word string) string {
  word = strings.ToLower(word)
  sliced := strings.Split(word, "")
  sort.Strings(sliced)
  word = strings.Join(sliced, "")

  return word
}
