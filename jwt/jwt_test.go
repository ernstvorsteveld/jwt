package jwt

import "testing"

func Test_should_decode_header(t *testing.T) {
  header := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
  result := decodeBase64(header)
  
  if result.alg != "HS256" {
    t.Errorf("Header should contain alg equals to HS256, current value is %s", result.alg)
  }
}

func Test_should_decode_valid_jwt(t *testing.T) {
  message := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiYTFiMmMzIiwidXNlcm5hbWUiOiJuaWtvbGEifQ.8IVVnxXEgzr9VXIObNYfNM65PDeUd4HbXF1J40dn-z8"
  sKey := "42isTheAnswer"

  decoded := Decode(message, sKey)
  if decoded != "Allemaal goed" {
    t.Errorf("The jwt is invalid, %s", decoded) 
  }
}