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

  p := true

  for i := 0 ; i < 7 ; i ++ {
    if sha[i] != byte(48) {
      p = false
    }
  }

  if p {
    fmt.Println(string(test))
    fmt.Println(sha)
  }

  if length < 10 {
    for i := 33; i < 125 ; i ++ {
      hash(append(test, byte(i)), length + 1)
    }
  }
}

func main(){
  
  for i := 33; i < 125 ; i ++ {
    hash([]byte{byte(i)} , 1, )
  }

}

