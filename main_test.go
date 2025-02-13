// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package logger

import (
	"github.com/gofiber/fiber"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNew_withRoutePath(t *testing.T) {
	routePath := "/test/:param/sufix"
	format := "route=${route}"
	expectedOutput := "route=/test/:param/sufix"

	// fake output
	buf := &strings.Builder{}
	stdout := log.New(buf, "", 0)

	n := New(Config{
		Format: format,
		Output: stdout.Writer(),
	})
	app := fiber.New()
	app.Use(n)

	app.Get(routePath, func(ctx *fiber.Ctx) {
		ctx.SendStatus(200)
	})

	req := httptest.NewRequest(http.MethodGet, "/test/af593469-3133-4943-b193-31f02e6e82e9/sufix", nil)

	_, err := app.Test(req, 1000)
	if err != nil {
		t.Errorf("Has: %+v, expected: nil", err)
	}

	if buf.String() != expectedOutput {
		t.Errorf("Has: %s, expected: %s", buf.String(), expectedOutput)
	}
}
