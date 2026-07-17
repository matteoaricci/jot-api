package docs

import (
	_ "embed"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed openapi.yaml
var openapiSpec []byte

const swaggerUI = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>jot-api docs</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui.css" />
  </head>
  <body>
    <div id="swagger-ui"></div>
    <script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
    <script>
      window.onload = () => {
        window.ui = SwaggerUIBundle({
          url: '/api/public/docs/openapi.yaml',
          dom_id: '#swagger-ui',
        });
      };
    </script>
  </body>
</html>`

// AddRoutes registers the OpenAPI spec and Swagger UI under /api/public/docs.
func AddRoutes(e *echo.Echo) {
	e.GET("/api/public/docs", func(c echo.Context) error {
		return c.HTML(http.StatusOK, swaggerUI)
	})

	e.GET("/api/public/docs/openapi.yaml", func(c echo.Context) error {
		return c.Blob(http.StatusOK, "application/yaml", openapiSpec)
	})
}
