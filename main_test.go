package main

import "testing"

func TestGreet(t *testing.T) {
    expected := "Hello, Renovate!"
    actual := greet()
    if actual != expected {
        t.Errorf("greet() = %q, want %q", actual, expected)
    }
}

