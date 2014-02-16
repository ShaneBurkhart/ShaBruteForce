package main

import (	
  "os"
  "fmt"
  "crypto/sha512"
  "encoding/hex"
  "strconv"
)

func hash(test []byte, length int, limit int){

  hasher := sha512.New()
  hasher.Reset()
  hasher.Write(test)
  sha := hex.EncodeToString(hasher.Sum(nil)) 

  i := 0
  for i = 0 ; sha[i] == byte(48) ; i ++ {}


  if i >= 10 {
    fmt.Printf("COUNT: %d, %s\n", i, string(test))
    fmt.Println(sha)
  }

  if length < limit {
    for i := 33; i < 125 ; i ++ {
      hash(append(test, byte(i)), length + 1, limit)
    }
  }
}

func main(){

  x, err := strconv.Atoi(os.Args[1])
  
  fmt.Printf("LIMIT %d\n", x)

  if err == nil {
    for i := 33; i < 125 ; i ++ {
      hash([]byte{byte(i)} , 1, x)
    }
  }

}

