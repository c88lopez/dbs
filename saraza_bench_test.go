package main

import (
	"testing"
	"fmt"
)

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

func BenchmarkSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprint("hello")
	}
}

func BenchmarkSprintln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintln("hello")
	}
}

func BenchmarkPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Print("")
	}
}

func BenchmarkPrintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Printf("")
	}
}