package router

import (
	"os"
	"svc-myg-ticketing/controller"
	"svc-myg-ticketing/repository"
	"svc-myg-ticketing/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AllRouter(db *gorm.DB) {
	router := gin.Default()
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

	myg_ticketing := router.Group("/myg-ticketing")
	{
		v1 := myg_ticketing.Group("/v1")
		{
			category := v1.Group("/category")
			{
				category.GET("/get/:size/:page_no/:sort_by/:order_by", categoryController.GetCategory)
				category.GET("/get-detail/:code-level", categoryController.GetDetailCategory)
				category.POST("/add", categoryController.CreateCategory)
				category.PUT("/update", categoryController.UpdateCategory)
				category.DELETE("/delete/:category-id", categoryController.DeleteCategory)
			}

			role := v1.Group("/role")
			{
				role.GET("/get", roleController.GetRole)
				role.POST("/add", roleController.CreateRole)
				role.PUT("/update", roleController.UpdateRole)
				role.DELETE("/delete/:role-id", roleController.DeleteRole)
			}

			permission := v1.Group("/permission")
			{
				permission.GET("/get", permissionController.GetPermission)
			}

			user := v1.Group("/user")
			{
				user.GET("/get/:search/:size/:page_no", userController.GetUser)
				user.GET("/get-detail/:username", userController.GetUserDetail)
				user.DELETE("/delete/:user-id", userController.DeleteUser)
				user.POST("/add", userController.CreateUser)
				user.PUT("/update", userController.UpdateUser)
				user.POST("/change-pass", userController.ChangePassword)
			}

			ticket := v1.Group("/ticket")
			{
				ticket.POST("/get", ticketController.GetTicket)
				ticket.GET("/get-detail/:ticket-code", ticketController.GetDetailTicket)
			}

			report := v1.Group("/report")
			{
				report.POST("/get", reportController.GetReport)
			}
		}
	}

	router.Run(os.Getenv("PORT"))
}
