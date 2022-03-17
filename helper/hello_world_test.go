package helper

import (
	"fmt"
	"testing"
)

func TestHeloWorldDoni(t *testing.T) {
	result := HelloWorld("Doni")

	if result != "Hello Doni" {
		// jika error
		t.Error("Result Must Be 'Hello Doni")
	}
	fmt.Print("Test Hello WorldDoni Done")
}

func TestHelloWorldRamdoni(t *testing.T) {
	result := HelloWorld("Ramdoni")
	if result != "Hello Ramdoni" {
		//Jika Error
		t.Fatal("Result Must Be 'Hello Ramdoni'")
	}
	fmt.Println("Test Hello WorldRamdoni Done")
}
