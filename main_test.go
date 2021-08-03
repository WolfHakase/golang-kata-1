package main

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestWelcomeMessage(t *testing.T) {
	g := NewGomegaWithT(t)

	expected := "Hello world!"
	actual := welcomeMessage()

	g.Expect(actual).To(Equal(expected))
}
