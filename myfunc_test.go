package main

import (
"testing"
)

func TestSayName(t *testing.T) {

output:=SayName("")

if output!="hello stranger" {
t.Errorf("Failed")
}else {
t.Logf("Success")
}

}
