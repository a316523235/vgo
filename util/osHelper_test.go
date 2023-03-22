package utils

import "testing"

func TestReadLine(t *testing.T) {
	expected := "hello"
	result := ReadLine()
	if result != expected {
		t.Errorf("ReadLine() failed, expected %v, got %v", expected, result)
	}
}

func TestReadInt(t *testing.T) {
	expected := 10
	result := ReadInt()
	if result != expected {
		t.Errorf("ReadInt() failed, expected %v, got %v", expected, result)
	}
}

func TestReadFloat(t *testing.T) {
	expected := 3.14
	result := ReadFloat()
	if result != expected {
		t.Errorf("ReadFloat() failed, expected %v, got %v", expected, result)
	}
}

func TestReadMultiLine(t *testing.T) {
	expected := "hello\nworld\n"
	result := ReadMultiLine()
	if result != expected {
		t.Errorf("ReadMultiLine() failed, expected %v, got %v", expected, result)
	}
}