package main

import (
"testing"
)

func TestSayName(t *testing.T) {

output:=SayName("")

if output!="Hello stranger" {
t.Errorf("Failed")
}else {
t.Logf("Success")
}

}
