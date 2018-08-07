package main

import "testing"

func TestDecode(t *testing.T) {

	s := "d4:testi33ee"
	d, _ := decode(s)

	d1 := d.(map[string]interface{})

	t.Log(d1["test"])
	_ = d1

}
