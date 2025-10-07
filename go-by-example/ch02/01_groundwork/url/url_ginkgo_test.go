package url_test

import (
	. "github.com/cynicdog/gopherdojo/go-by-example/01_groundwork/url"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("URL Parsing", func() {
	Describe("Parse", func() {
		It("Parses a full URL with path", func() {
			u, err := Parse("https://github.com/cynicdog")
			Expect(err).ToNot(HaveOccurred())
			Expect(u.Scheme).To(Equal("https"))
			Expect(u.Host).To(Equal("github.com"))
			Expect(u.Path).To(Equal("cynicdog"))
		})
		It("parses a URL without path", func() {
			u, err := Parse("https://github.com")
			Expect(err).ToNot(HaveOccurred())
			Expect(u.Scheme).To(Equal("https"))
			Expect(u.Host).To(Equal("github.com"))
			Expect(u.Path).To(Equal(""))
		})

		It("parses a data URL", func() {
			u, err := Parse("data:text/plain;base64,R28gYnkgRXhhbXBsZQ==")
			Expect(err).ToNot(HaveOccurred())
			Expect(u.Scheme).To(Equal("data"))
		})
	})

	DescribeTable("table-driven Parse tests",
		func(uri string, expected *URL) {
			u, err := Parse(uri)
			Expect(err).ToNot(HaveOccurred())
			Expect(u).To(Equal(expected))
		},
		Entry(
			"full URL",
			"https://github.com/cynicdog",
			&URL{Scheme: "https", Host: "github.com", Path: "cynicdog"},
		),
		Entry(
			"without path",
			"https://github.com",
			&URL{Scheme: "https", Host: "github.com", Path: ""},
		),
		Entry(
			"data scheme",
			"data:text/plain;base64,R28gYnkgRXhhbXBsZQ==",
			&URL{Scheme: "data"},
		),
	)
})
