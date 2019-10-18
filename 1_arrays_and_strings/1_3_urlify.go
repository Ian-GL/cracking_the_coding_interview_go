package main

import (
  "fmt"
  "strings"
)

func main() {
  ex := "Mr John Smith     "
  rlen := 13

  url := Urlify(ex, rlen)

  fmt.Println(url)
}

func Urlify (str string, len int) string {
  chars := strings.Split(str, "")
  url := ""
  new := ""

  for i := 0; i < len; i++ {
    current := chars[i]
    if(current == " ") {
      new = "%20"
    } else {
      new = current
    }
    url = url + new
  }

  return url
}
