package gstrings

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

func TestCamelToSnake(t *testing.T) {
	camel := "camel_baby"
	snake := CamelToSnake(camel, false)
	if snake != "camelBaby" {
		t.Error("Wrong answer: " + snake)
	}

	snake = CamelToSnake(camel, true)
	if snake != "CamelBaby" {
		t.Error("Wrong answer: " + snake)
	}

	snake = CamelToSnake("___what_is__GOING_on_", false)
	if snake != "___whatIs__GOINGOn_" {
		t.Error("Wrong answer: " + snake)
	}
}

func TestSnakeToCamel(t *testing.T) {
	snake := "snakeBaby"
	camel := SnakeToCamel(snake)
	if camel != "snake_baby" {
		t.Error("Wrong answer: " + camel)
	}

	snake = "SnakeBabyyy"
	camel = SnakeToCamel(snake)
	if camel != "snake_babyyy" {
		t.Error("Wrong answer: " + camel)
	}

	camel = SnakeToCamel("WhATisGOiNGoN")
	if camel != "wh_at_is_go_i_ng_o_n" {
		t.Error("Wrong answer: " + camel)
	}
}
