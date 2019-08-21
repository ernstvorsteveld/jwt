package jwt

import (
  "crypto/hmac"
  "crypto/sha256"
  "encoding/base64"
  "encoding/json"
  "fmt"
)

func Decode(jwt string, pkey string) string {
    key := []byte(pkey)
    h := hmac.New(sha256.New, key)
    h.Write([]byte(jwt))
    b := base64.URLEncoding.EncodeToString(h.Sum(nil))
    return string(b)
}

func decodeBase64(sEnc string) *JwtHeader {
  sDec, _ := base64.StdEncoding.DecodeString(sEnc)
  fmt.Printf("sDec is %s.\n", sDec)
  var data JwtHeader
  err := json.Unmarshal([]byte(string(sDec)), &data)
  if err != nil {
    fmt.Printf("Error while marshalling %s.\n", err)
  } else {
    fmt.Printf("Data is %s.\n", data)
  }
  return &data
}