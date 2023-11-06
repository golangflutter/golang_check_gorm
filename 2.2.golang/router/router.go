package router

import (
	"fmt"
	"golang-crud-gin/controller"
	"golang-crud-gin/middleware"
	"golang-crud-gin/repository"
	"log"
	"time"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(
    tagsController *controller.TagsController,
    userRepository repository.UsersRepository,
    authenticationController *controller.AuthenticationController,
    usersController *controller.UserController,
    refreshTokenController gin.HandlerFunc,
    imageUploadController *controller.ImageUploadController,
    categoryController *controller.CategoryController,
    sizeController *controller.SizeController,
    // Add color controller here
    colorController *controller.ColorController,
	billboardController *controller.BillboardController,productController *controller.ProductController,
) *gin.Engine {




	// corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"} // Change this to corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = 12 * time.Hour
	router := gin.Default()
	log.Println("Setting up CORS")  // Debugging log
	router.Use(cors.New(corsConfig))
	fmt.Println("after cors")
	router.Use(cors.New(corsConfig))
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home from golang")
	})
    fmt.Println("before cors")
		// Set CORS middleware

		// cors
	baseRouter := router.Group("/api")
	baseRouter.GET("users", usersController.GetUsers)
	// Authentication routes
	authenticationRouter := baseRouter.Group("/authentication")
	authenticationRouter.POST("/register", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.Login)
	// User routes, check register, refresh token
	// usersRouter := router.Group("/users")
	authenticationRouter.GET("users", middleware.DeserializeUser(userRepository), usersController.GetUsers)
	log.Println("Setting up CORS") 
	// tags
	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.GET("", tagsController.FindAll)
	tagsRouter.GET("/:tagId", tagsController.FindById)
	tagsRouter.POST("", tagsController.Create)
	tagsRouter.PATCH("/:tagId", tagsController.Update)
	tagsRouter.DELETE("/:tagId", tagsController.Delete)
    // Route for token refresh
    baseRouter.POST("/refresh-token", refreshTokenController)

	// image upload router

	router.POST("/upload",     imageUploadController.UploadImage)

	    // Add a route for serving the original image
		router.GET("/image", imageUploadController.ServeImage)

		// Add a route for serving a thumbnail
		router.GET("/thumbnail", imageUploadController.ServeThumbnail)
		router.Static("/uploads", "./uploads")
		router.DELETE("/deleteImage", imageUploadController.DeleteImage)
// categories
    // category routes
// categories
// category routes
categoryRouter := baseRouter.Group("/categories")
categoryRouter.GET("", categoryController.FindAll)
categoryRouter.GET("/findbyname", categoryController.FindByCategoryName)
categoryRouter.GET("/findbyid/:categoryId", categoryController.FindById)
categoryRouter.POST("", categoryController.Create)
categoryRouter.PATCH("/:categoryId", categoryController.Update)
categoryRouter.DELETE("/:categoryId", categoryController.Delete)
categoryRouter.GET("/findbypage", categoryController.FindCategoriesByPage)
categoryRouter.GET("/findbycharacteristic", categoryController.FindCategoriesByCharacteristicName)
// size 

sizeRouter := baseRouter.Group("/sizes")
sizeRouter.POST("", sizeController.CreateSize)
sizeRouter.PATCH("/:sizeId", sizeController.UpdateSize)
sizeRouter.DELETE("/:sizeId", sizeController.DeleteSize)
sizeRouter.GET("/:sizeId", sizeController.FindSizeByID)
sizeRouter.GET("", sizeController.FindAllSizes)
sizeRouter.GET("/findbyname", sizeController.FindSizeByName)
sizeRouter.GET("/findbypage", sizeController.FindSizesByPage)

// color

colorRouter := baseRouter.Group("/colors")
colorRouter.POST("", colorController.CreateColor)
colorRouter.PATCH("/:colorId", colorController.UpdateColor)
colorRouter.DELETE("/:colorId", colorController.DeleteColor)
colorRouter.GET("/:colorId", colorController.FindColorByID)
colorRouter.GET("", colorController.FindAllColors)
colorRouter.GET("/findbyname", colorController.FindColorByName)
colorRouter.GET("/findbypage", colorController.FindColorsByPage)
// bill board

billboardRouter := baseRouter.Group("/billboards")
billboardRouter.GET("", billboardController.FindAllBillboards)
billboardRouter.GET("/:billboardId", billboardController.FindBillboardByID)
billboardRouter.POST("", billboardController.CreateBillboard)
billboardRouter.PATCH("/:billboardId", billboardController.UpdateBillboard)
billboardRouter.DELETE("/:billboardId", billboardController.DeleteBillboard)
billboardRouter.GET("/findbylabel", billboardController.FindBillboardByLabel)
billboardRouter.GET("/findbypage", billboardController.FindBillboardsByPage)

// product 


    productRouter := baseRouter.Group("/products")
    productRouter.POST("", productController.CreateProduct)
    productRouter.PATCH("/:productId", productController.UpdateProduct)
    productRouter.DELETE("/:productId", productController.DeleteProduct)
    productRouter.GET("/findbyId/:productId", productController.FindProductByID)
    productRouter.GET("", productController.FindAllProducts)
    productRouter.GET("/findbycategory", productController.FindProductsByCategory)
    productRouter.GET("/findbycolor", productController.FindProductsByColor)
    productRouter.GET("/findbysize", productController.FindProductsBySize)
    productRouter.GET("/findbycharacteristic", productController.FindProductsByCharacteristic)
    productRouter.GET("/findbypage", productController.FindProductsByPage)
	productRouter.GET("/findbyproductname", productController.FindProductsByProductName)

return router

}
