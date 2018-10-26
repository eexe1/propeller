package main

import (
  "flag"
  "fmt"
  "github.com/eexe1/propeller/api"
)

func main()  {

  // parsing arguments & flags

  emailPtr := flag.String("email", "", "Apple Store Connect Email")
  passwordPtr := flag.String("password", "", "Password")
  isVerbosePtr := flag.Bool("verbose", false, "Enable verbose logging")

  flag.Parse()

  if *isVerbosePtr {
    fmt.Println(*emailPtr, *passwordPtr, "verbose_mode")
  } else {
    fmt.Println(*emailPtr, *passwordPtr)
  }

  body, err := api.Get("https://www.google.com")
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(body)

}
