package httphelp

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestCreate(t *testing.T) {

	Describe("Create()", func() {

		chrome := "application/xml,application/xhtml+xml,text/html;q=0.9,text/plain;q=0.8,image/png,*/*;q=0.5"

		It("should return [image/png]", func() {
			alternatives := []string{"text/html", "image/png"}
			content_type := Negotiate(chrome, alternatives)
			AssertEqual(content_type, "image/png")
		})

		It("should return [text/html]", func() {
			alternatives := []string{"text/html", "text/plain", "text/n3"}
			content_type := Negotiate(chrome, alternatives)
			AssertEqual(content_type, "text/html")
		})

		It("should return [text/plain]", func() {
			alternatives := []string{"text/n3", "text/plain"}
			content_type := Negotiate(chrome, alternatives)
			AssertEqual(content_type, "text/plain")
		})

		It("should return [text/n3]", func() {
			alternatives := []string{"text/n3", "application/rdf+xml"}
			content_type := Negotiate(chrome, alternatives)
			AssertEqual(content_type, "text/n3")
		})
	})

	Report(t)
}
