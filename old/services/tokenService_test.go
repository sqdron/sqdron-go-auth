package services

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestJwtTockenStuff(t *testing.T) {
	Convey("Given some integer with a starting value", t, func() {
		token := GenerateJwtToken()
		So(token, ShouldNotBeEmpty)
	})
}