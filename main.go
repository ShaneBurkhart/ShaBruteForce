package main

import (
  "fmt"
  "crypto/sha512"
  "encoding/hex"
)

func hash(test []byte, length int){

  hasher := sha512.New()
  hasher.Reset()
  hasher.Write(test)
  sha := hex.EncodeToString(hasher.Sum(nil)) 

  i := 0
  for i = 0 ; sha[i] == byte(48) ; i ++ {}


  if i >= 8 {
    fmt.Printf("COUNT: %d, %s", i, string(test))
    fmt.Println(sha)
  }

  if length < 9 {
    for i := 100; i < 115 ; i ++ {
      hash(append(test, byte(i)), length + 1)
    }
  }
}

func main(){
  
  for i := 100; i < 115 ; i ++ {
    hash([]byte{byte(i)} , 1, )
  }

}

