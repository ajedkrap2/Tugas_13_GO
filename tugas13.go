package main

import "fmt"
import "encoding/base64"
import "crpyto/sha1"

func main() {
  var input string
  fmt.Print("Masukan Nama Panjang: ")
  fmt.Scanln(&input)
  fmt.Println()

  //base64
  var enBase = base64.StdEncoding.EncodeToString([]byte(input))

  //sha1
  var sha = sha1.New()
  sha.Write([]byte(input))
  var encrpyt = sha.Sum(nil)
  var enSha = fmt.Sprintf(encrpyt)

  //output
  fmt.Println("Base64 : \t", enBase)
  fmt.Println("Sha1 : \t", enSha)
}
