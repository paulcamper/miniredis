//go:build int
// +build int

package main

import "testing"

func TestCommand(t *testing.T) {
	t.Skip("not sure about this one yet")
	testRaw(t, func(c *client) {
		c.DoLoosely("COMMAND")
	})
}

func TestInfo(t *testing.T) {
	t.Skip("not sure about this one yet")
	testRaw(t, func(c *client) {
		c.DoLoosely("INFO")
	})
}
