package books_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "my-goexercise/goexcersise/books"
)

var _ = Describe("Book", func() {
	var (
		longBook  Book
		shortBook Book
		book      Book
		err       error
	)

	BeforeEach(func() {
		longBook = Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
		book, err = NewBookFromJSON(`{
			"title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":1488
			}`)
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})

	Describe("loading from JSON", func() {
		// 如果正常解析JSON
		Context("when the JSON parses succesfully", func() {
			It("should populate the fields correctly", func() {
				// 期望                相等
				Expect(book.Title).To(Equal("Les Miserables"))
				Expect(book.Author).To(Equal("Victor Hugo"))
				Expect(book.Pages).To(Equal(1488))
			})

			It("should not error", func() {
				// 期望      没有发生错误
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
