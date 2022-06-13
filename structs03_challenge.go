package main

import (
    "fmt"
)

type astro struct {
  name string
  age int
  lastmission string
  isneeded bool
}


func main() {
      astro1 := astro{"Iuri Gagarin", 27, "Souz", false}
      astro2 := astro{"Alan Shepard", 38, "Gemeni", false}
      astro3 := astro{"Doug Hurley",  55, "SpaceX", true}

      fmt.Println(astro1)
      fmt.Println(astro2)
      fmt.Println(astro3)

  p := []astro{astro1, astro2, astro3}

    // display the slice
    fmt.Println(p)

    // select data from the struct
    fmt.Println(p[2].lastmission)

  }

