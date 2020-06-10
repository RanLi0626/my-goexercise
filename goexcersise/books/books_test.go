package books

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	g := NewGomegaWithT(t)

	b, _ := NewBookFromJSON(`{
		"title":"Les Miserables",
		"author":"Victor Hugo",
		"pages":1488
		}`)
	g.Expect(b.Title).To(Equal("Les Miserables"))
}
