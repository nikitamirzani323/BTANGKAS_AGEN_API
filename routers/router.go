package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/controllers"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/middleware"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		// c.Set("Content-Security-Policy", "frame-ancestors 'none'")
		// c.Set("X-XSS-Protection", "1; mode=block")
		// c.Set("X-Content-Type-Options", "nosniff")
		// c.Set("X-Download-Options", "noopen")
		// c.Set("Strict-Transport-Security", "max-age=5184000")
		// c.Set("X-Frame-Options", "SAMEORIGIN")
		// c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		return c.Next()
	})
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Get("/ipaddress", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      "data",
			"BASEURL":     c.BaseURL(),
			"HOSTNAME":    c.Hostname(),
			"IP":          c.IP(),
			"IPS":         c.IPs(),
			"OriginalURL": c.OriginalURL(),
			"Path":        c.Path(),
			"Protocol":    c.Protocol(),
			"Subdomain":   c.Subdomains(),
		})
	})
	app.Get("/dashboard", monitor.New())

	app.Post("/api/login", controllers.CheckLogin)
	app.Post("/api/valid", middleware.JWTProtected(), controllers.Home)
	app.Post("/api/alladmin", middleware.JWTProtected(), controllers.Adminhome)
	app.Post("/api/detailadmin", middleware.JWTProtected(), controllers.AdminDetail)
	app.Post("/api/saveadmin", middleware.JWTProtected(), controllers.AdminSave)
	app.Post("/api/alladminrule", middleware.JWTProtected(), controllers.Adminrulehome)
	app.Post("/api/saveadminrule", middleware.JWTProtected(), controllers.AdminruleSave)
	app.Post("/api/listbet", middleware.JWTProtected(), controllers.Listbethome)
	app.Post("/api/listbetconf", middleware.JWTProtected(), controllers.Listbetconfpointhome)
	app.Post("/api/listbetsave", middleware.JWTProtected(), controllers.ListbetSave)
	app.Post("/api/listbetconfsave", middleware.JWTProtected(), controllers.ListbetconfpointSave)

	app.Post("/api/listpattern", middleware.JWTProtected(), controllers.Listpatternhome)
	app.Post("/api/listpatternsave", middleware.JWTProtected(), controllers.ListpatternSave)
	app.Post("/api/listpatterndetail", middleware.JWTProtected(), controllers.Listpatterndetailhome)
	app.Post("/api/listpatterndetailsave", middleware.JWTProtected(), controllers.ListpatterndetailSave)
	app.Post("/api/listpatterndetaildelete", middleware.JWTProtected(), controllers.ListpatterndetailDelete)
	app.Post("/api/pattern", middleware.JWTProtected(), controllers.Patternhome)
	app.Post("/api/patternbycode", middleware.JWTProtected(), controllers.PatternByPoin)
	app.Post("/api/patternsave", middleware.JWTProtected(), controllers.PatternSave)
	app.Post("/api/patternsavemanual", middleware.JWTProtected(), controllers.PatternSavemanual)
	app.Post("/api/listpoint", middleware.JWTProtected(), controllers.Listpointhome)
	app.Post("/api/listpointshare", middleware.JWTProtected(), controllers.Listpointsharehome)
	app.Post("/api/listpointsave", middleware.JWTProtected(), controllers.ListpointSave)

	return app
}
