package api

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"

	"api/common/constants"
	"api/config"
	configWS "api/config/ws"
	"api/middlewares"
)

// StartGin Set up and start the API: set up API documentation,
// configure middlewares, and security measures.
func StartGin() {
	// Set up gin for your API
	gin.SetMode(config.Env.GinMode)
	gin.ForceConsoleColor()
	engine := gin.Default()
	engine.HandleMethodNotAllowed = true
	engine.ForwardedByClientIP = true
	err := engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}
	// Register gin middlewares
	engine.Use(middlewares.GinContextRegister())

	// Group API
	ginGroup := engine.Group(config.Env.ApiGroup)

	// OpenAPI documentation based on huma
	humaConfig := huma.DefaultConfig(constants.OpenApiTitle, constants.OpenApiVersion)
	// Custom hook to remove schema links
	humaConfig.CreateHooks = []func(huma.Config) huma.Config{
		func(c huma.Config) huma.Config {
			return c
		},
	}
	humaConfig.DocsPath = ""
	humaConfig.Servers = []*huma.Server{
		{URL: config.Env.ApiGroup},
	}
	humaConfig.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		constants.SecuritySchemeBearerToken: {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
			Description:  "Bearer token used to access some resources",
		},
		constants.SecuritySchemeSchoolToken: {
			Type:        "apiKey",
			In:          "header",
			Name:        "X-School-Api-Key",
			Description: "School token used to access school resources",
		},
		constants.SecuritySchemeSchoolID: {
			Type:        "apiKey",
			In:          "header",
			Name:        "X-School-Id",
			Description: "School ID used to access school resources",
		},
	}
	humaConfig.Info.Description = constants.OpenApiDescription
	humaApi := humagin.NewWithGroup(engine, ginGroup, humaConfig)
	// Register middlewares
	humaApi.UseMiddleware(
		middlewares.HeadersMiddleware(humaApi),
		middlewares.CorsMiddleware(humaApi),
		middlewares.SchoolMiddleware(humaApi),
		middlewares.AuthMiddleware(humaApi),
		middlewares.PermissionMiddleware(
			humaApi,
			AllControllers.UserController.Service.Repository,
			AllControllers.PermissionController.Service.Repository,
			AllControllers.DirectorController.Service.Repository,
			AllControllers.TeacherController.Service.Repository,
			AllControllers.StudentController.Service.Repository,
			AllControllers.ParentController.Service.Repository,
		),
	)

	// Serve public static files as favicon
	engine.StaticFS("/assets", http.Dir(constants.AssetPublicAppPath))
	engine.StaticFS("/openapi", http.Dir(constants.AssetPublicOpenApiPath))

	// Register websocket
	wsManager := configWS.SetupWebsocket()
	wsManager.SubscribeToNotifications()
	engine.GET("/ws/notifications", wsManager.GinWSListener)

	// Register API endpoints
	ginGroup.GET("/docs", func(ctx *gin.Context) {
		ctx.Data(200, "text/html", config.OpenAPITemplates.Docs)
	})
	ginGroup.GET("/docs/scalar", func(ctx *gin.Context) {
		ctx.Data(200, "text/html", config.OpenAPITemplates.Scalar)
	})
	ginGroup.GET("/docs/swagger", func(ctx *gin.Context) {
		ctx.Data(200, "text/html", config.OpenAPITemplates.Swagger)
	})
	ginGroup.GET("/docs/redocly", func(ctx *gin.Context) {
		ctx.Data(200, "text/html", config.OpenAPITemplates.Redocly)
	})
	registerEndpoints(&humaApi)

	// Start to listen
	formattedPort := fmt.Sprintf(":%d", config.Env.AppPort)
	err = engine.Run(formattedPort)
	if err != nil {
		panic(err)
	}
}
