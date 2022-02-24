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
				role.PUT("/udpate", roleController.UpdateRole)
			}
		}
	}

	router.Run(os.Getenv("PORT"))
}
