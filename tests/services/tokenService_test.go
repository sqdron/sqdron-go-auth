package app

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"../../services"
)

func TestJwtTockenStuff(t *testing.T) {
	Convey("Given some integer with a starting value", t, func() {
		token := services.GenerateJwtToken()
		So(token, ShouldNotBeEmpty)
	})
}