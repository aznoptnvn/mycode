package main

import (
    "fmt"
)

func main() {
     const uri_string = "https://example.org:6001/v2/snacks?"

     var r, q, s string  = "req=snickers", "quatity=1", "size=king"

     res := fmt.Sprintf("%s%s&%s&%s", uri_string, r,q, s)
      fmt.Println(res)

  }
