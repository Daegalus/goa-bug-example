package design

import (
	// nolint:golint
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/zaplogger"
)

var _ = API("test", func() {
	Server("test", func() {
		Host("test.example.com", func() {
			URI("https://test.example.com/api/v1")
		})
	})
	HTTP(func() {
		Consumes("application/json")
		Produces("application/json")
	})
	Description("The API for the test")
	Title("The Test API")
	Version("1")

	Error("bad_request", ErrorResult, "Bad request")
	Error("internal_server_error", ErrorResult, "Internal server error")
	Error("not_found", ErrorResult, "Not found")
})

var IDPayload = Type("idPayload", func() {
	Description("Payload for when we need an ID and a Token")
	Attribute("id", String)
	Required("id")
})

var testPayload = Type("testPayload", func() {
	Description("testPayload is a payload")

	Attribute("id", String, "Unique resource ID", func() {
		Example("c0e9c01d-2a43-4998-81f4-b48c2ce76ca9")
	})

	Attribute("name", String, "Generated name", func() {
		Example("something-something-something")
		//ReadOnly()
	})
})

// testMedia is a test media type
var testMedia = ResultType("application/vnd.test+json", "testResponse", func() {
	Description("A test response")
	Reference(testPayload)

	Attributes(func() {
		Attribute("id")
		Attribute("name")

		Required("id", "name")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
	})
})

var _ = Service("test", func() {
	Description("test request")

	HTTP(func() {
		Path("/test")
	})

	Method("update", func() {
		Description("modifies an existing test")
		HTTP(func() {
			PUT("/{id}")
			Params(func() {
				Param("id", func() {
					Description("Unique resource ID generated from creation")
					Example("c0e9c01d-2a43-4998-81f4-b48c2ce76ca9")
					MinLength(8)
					MaxLength(256)
				})
				Required("id")
			})
			Response(StatusAccepted)
		})

		Payload(testPayload)
		Result(testMedia)
	})

	Method("delete", func() {
		Description("delete a test")
		HTTP(func() {
			DELETE("/{id}")
			Params(func() {
				Param("id", func() {
					Description("Unique resource ID generated from creation")
					Example("c0e9c01d-2a43-4998-81f4-b48c2ce76ca9")
					MinLength(8)
					MaxLength(256)
				})
				Required("id")
			})
			Response(StatusNotFound, "not_found")
			Response(StatusBadRequest, "bad_request")
			Response(StatusInternalServerError, "internal_server_error")
		})

		Payload(IDPayload)
		Result(Empty)
	})

	Method("show", func() {
		Description("shows a test")
		HTTP(func() {
			GET("/{id}")
			Params(func() {
				Param("id", func() {
					Description("Unique resource ID generated from creation")
					Example("c0e9c01d-2a43-4998-81f4-b48c2ce76ca9")
					MinLength(8)
					MaxLength(256)
				})
				Required("id")
			})
			Response(StatusNotFound, "not_found")
			Response(StatusBadRequest, "bad_request")
			Response(StatusInternalServerError, "internal_server_error")
		})
		Payload(testPayload)
		Result(testMedia)
	})
})
