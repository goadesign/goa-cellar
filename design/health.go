package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Service("health", func() {
	Description("Health check endpoint")

	Endpoint("do", func() {
		Description("Perform health check.")
		Response(String)
		Error("unhealthy", String)

		HTTP(func() {
			GET("/health")
			GET("/_ah/health") // App Engine path
			Error("unhealthy", func() {
				Status(InternalServerError)
			})
		})
	})
})
