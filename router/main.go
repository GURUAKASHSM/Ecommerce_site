package router

import (
	"ecommerce/controller"

	"github.com/gin-gonic/gin"
)

// Router creates and configures the Gin router.
func Router() *gin.Engine {
	router := gin.Default()

	// Serve static files for specific routes
	router.Static("/index", "./frontend/index")
	router.Static("/home", "./frontend/home")
	router.Static("/signup", "./frontend/signup")
	router.Static("/signin", "./frontend/signin")
	//router.Static("/additems", "./frontend/inventory")
	router.Static("/admin", "./frontend/admin")
	router.Static("/edit", "./frontend/edit")
	//router.Static("/delete", "./frontend/delete")
	router.Static("/seller", "./frontend/seller")
	router.Static("/feedback","./frontend/feedback")
	//router.Static("/feedbacks","./frontend/feedbackReview")



	// Define your routes
	router.GET("/getallcustomerdata", controller.Getalldata)
	router.GET("/getallinventorydata", controller.Getinventorydata)
	router.GET("/getallsellerdata", controller.Getallsellerdata)
	router.POST("/createseller", controller.CreateSeller)
	router.POST("/create", controller.CreateProfile)
	router.POST("/deletedata", controller.Delete)
	router.POST("/addtocart", controller.Addtocart)
	router.POST("/login", controller.Login)
	router.POST("/products", controller.Products)
	router.POST("/updatecart", controller.UpdateCart)
	router.POST("/inventory", controller.Inventory)
	router.POST("/search", controller.Search)
	router.POST("/update", controller.Update)
	router.GET("/inventorydata", controller.Getallinventorydata)
	router.POST("/sellercheck",controller.CheckSeller)
	router.POST("/deleteproduct",controller.DeleteProduct)
	router.POST("/deleteproductbyseller",controller.DeleteProductBySeller)
	router.POST("/updateproductbyseller",controller.UpdateProductBySeller)
	router.POST("/sitefeedback",controller.Feedback)
	router.GET("sellerfeedback",controller.SellerFeedback)
	router.GET("customerfeedback",controller.CustomerFeedback)
	router.POST("/deletefeedback",controller.Deletefeedback)


	return router
}

