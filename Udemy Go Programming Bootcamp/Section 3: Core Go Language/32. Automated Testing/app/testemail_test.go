package app

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err != nil {
		t.Error("hello is not an email")
	} else {
		t.Error("hello is an email")
	}

	_, err = IsEmail("derek@aol.com")
	if err != nil {
		t.Error("derek@aol.com is not an email")
	} else {
		t.Error("derek@aol.com is an email")
	}

	_, err = IsEmail("derek@aol")
	if err != nil {
		t.Error("derek@aol is not an email")
	} else {
		t.Error("derek@aol is an email")
	}

	_, err = IsEmail("nols.smit@gmail.com")
	if err != nil {
		t.Error("nols.smit@gmail.com is not an email")
	} else {
		t.Error("nols.smit@gmail.com is an email")
	}

	_, err = IsEmail("piet.smit@gmail")
	if err != nil {
		t.Error("piet.smit@gmail is not an email")
	} else {
		t.Error("piet.smit@gmail is an email")
	}	
}	
