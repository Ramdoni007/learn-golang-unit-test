package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T) {
	test := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Doni",
			request:  "Doni",
			expected: "Hello Doni",
		},

		{
			name:     "Ahmad",
			request:  "Ahmad",
			expected: "Hello Ahmad",
		},
	}

	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}

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
