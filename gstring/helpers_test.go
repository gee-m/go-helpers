package gstring

import "testing"

func TestSetIfExists(t *testing.T) {
	dst := "not done"
	src := "done"
	SetIfExists(&dst, src)

	if dst != src {
		t.Error("dst should have been equal to src")
	}

	SetIfExists(&dst, "")
	if dst != src {
		t.Error("dst should have been unchanged")
	}
}

func TestCamelToSnakeCase(t *testing.T) {
	camel := "camel_baby"
	snake := CamelToSnakeCase(camel, false)
	if snake != "camelBaby" {
		t.Error("Wrong answer: " + snake)
	}

	snake = CamelToSnakeCase(camel, true)
	if snake != "CamelBaby" {
		t.Error("Wrong answer: " + snake)
	}

	snake = CamelToSnakeCase("___what_is__GOING_on_", false)
	if snake != "___whatIs__GOINGOn_" {
		t.Error("Wrong answer: " + snake)
	}
}
