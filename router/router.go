package router

import (
	"os"
	"svc-myg-ticketing/controller"
	"svc-myg-ticketing/repository"
	"svc-myg-ticketing/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AllRouter(db *gorm.DB) {

	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "token", "request-by", "signature-key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           86400,
	}))
	repository := repository.Repository(db)

	logService := service.LogService(repository)

	categoryService := service.CategoryService(repository)
	categoryController := controller.CategoryController(categoryService, logService)

	roleService := service.RoleService(repository, repository)
	roleController := controller.RoleController(roleService, logService)

	permissionService := service.PermissionService(repository)
	permissionController := controller.PermissionController(permissionService, logService)

	userService := service.UserService(repository, repository)
	userController := controller.UserController(userService, logService)

	ticketService := service.TicketService(repository, repository)
	ticketController := controller.TicketController(ticketService, logService)

	reportService := service.ReportService(repository)
	reportController := controller.ReportController(reportService, logService)

	authService := service.AuthService(repository, repository)
	authController := controller.AuthController(authService, logService)

	captchaService := service.CapthcaService()
	captchaController := controller.CaptchaController(captchaService, logService)

	dir := os.Getenv("FILE_DIR")
	router.Static("/assets", dir)

	myg_ticketing := router.Group("/myg-ticketing")
	{
		v1 := myg_ticketing.Group("/v1")
		{
			category := v1.Group("/category")
			{
				category.Use(authService.Authentication(), authService.Authorization())
				category.GET("/get/:size/:page_no/:sort_by/:order_by", categoryController.GetCategory)
				category.GET("/get-detail/:code-level", categoryController.GetDetailCategory)
				category.POST("/add", categoryController.CreateCategory)
				category.PUT("/update", categoryController.UpdateCategory)
				category.DELETE("/delete/:category-id", categoryController.DeleteCategory)
			}

			role := v1.Group("/role")
			{
				role.Use(authService.Authentication(), authService.Authorization())
				role.GET("/get", roleController.GetRole)
				role.POST("/add", roleController.CreateRole)
				role.PUT("/update", roleController.UpdateRole)
				role.DELETE("/delete/:role-id", roleController.DeleteRole)
			}

			permission := v1.Group("/permission")
			{
				permission.Use(authService.Authentication(), authService.Authorization())
				permission.GET("/get", permissionController.GetPermission)
			}

			user := v1.Group("/user")
			{
				user.Use(authService.Authentication(), authService.Authorization())
				user.GET("/get/:search/:size/:page_no", userController.GetUser)
				user.GET("/get-detail/:username", userController.GetUserDetail)
				user.DELETE("/delete/:user-id", userController.DeleteUser)
				user.POST("/add", userController.CreateUser)
				user.PUT("/update", userController.UpdateUser)
				user.POST("/change-pass", userController.ChangePassword)
				user.POST("/reset-pass", userController.ResetPassword)
				user.PUT("/update-profile", userController.UpdateProfile)
				user.PUT("/update-status", userController.UpdateUserStatus)
			}

			ticket := v1.Group("/ticket")
			{
				ticket.Use(authService.Authentication(), authService.Authorization())
				ticket.POST("/get", ticketController.GetTicket)
				ticket.GET("/get-detail/:ticket-code", ticketController.GetDetailTicket)
				ticket.POST("/add", ticketController.CreateTicket)
				ticket.PUT("/update", ticketController.UpdateTicket)
				ticket.POST("/reply", ticketController.ReplyTicket)
			}

			report := v1.Group("/report")
			{
				report.Use(authService.Authentication(), authService.Authorization())
				report.POST("/get", reportController.GetReport)
			}

			auth := v1.Group("/auth")
			{
				auth.POST("/login", authController.Login)
				auth.GET("/refresh-token", authController.RefreshToken)
			}

			captcha := v1.Group("/captcha")
			{
				captcha.POST("/generate", captchaController.GenerateCaptcha)
				captcha.POST("/verify", captchaController.CaptchaVerify)
			}
		}
	}

	router.Run(os.Getenv("PORT"))
}
