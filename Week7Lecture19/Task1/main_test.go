package main

import (
	"fmt"
	"io"
	"testing"
)

func TestReader(t *testing.T) {

	//Arrange
	expectedRes := "elppa"

	//Act
	ad := NewReverseStringReader("apple")
	pom, err := io.ReadAll(ad)
	if err != nil {
		fmt.Println(err)
	}
	myString := string(pom[:])
	result := myString
	//Assertion
	if result != expectedRes {
		t.Errorf("Expected %s got %s ", expectedRes, result)
	}
}
