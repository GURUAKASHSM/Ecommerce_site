package controller

import (
	"ecommerce/constants"
	"ecommerce/models"
	"ecommerce/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Getinventorydata(c *gin.Context) {
	Inventorydata := service.Getinventorydata()
	c.JSON(http.StatusOK, gin.H{"Inventory": Inventorydata})
}
func Getalldata(c *gin.Context) {
	alltransaction := service.Getalldata()
	c.JSON(http.StatusOK, alltransaction)
}
func CreateSeller(c *gin.Context) {
	var seller models.Seller
	if err := c.BindJSON(&seller); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	createseller := service.CreateSeller(seller)
	c.JSON(http.StatusOK, createseller)

}
func Getallsellerdata(c *gin.Context) {
	Getallsellerdata := service.Getallsellerdata()
	c.JSON(http.StatusOK, gin.H{"seller": Getallsellerdata})

}
func UpdateCart(c *gin.Context) {
	var cart models.Cart
	if err := c.BindJSON(&cart); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	t1, err := service.ExtractCustomerID(cart.CustomerId, constants.SecretKey)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	cart.CustomerId = t1

	result := service.UpdateCart(cart)
	c.JSON(http.StatusOK, result)
}
func Login(c *gin.Context) {
	var request models.Login
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	token, err,no := service.Login(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if no == 1 {
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}else if(no == 0){
		c.JSON(http.StatusOK, gin.H{"message": token})
		return
	} 
}

func Products(c *gin.Context) {
	var cartproducts *models.Addtocart
	if err := c.BindJSON(&cartproducts); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	token, err := service.ExtractCustomerID(cartproducts.Token, constants.SecretKey)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	cart := service.Cart(token)
	c.JSON(http.StatusOK, cart)
}

func Addtocart(c *gin.Context) {
	var addtocart models.Addtocart
	if err := c.BindJSON(&addtocart); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	token, err := service.ExtractCustomerID(addtocart.Token, constants.SecretKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	var addtocart1 models.Addtocart1
	addtocart1.CustomerId = token
	addtocart1.Name = addtocart.Name
	addtocart1.Price = addtocart.Price
	addtocart1.SellerQuantity = addtocart.Sellerquantity
	result := service.Addtocart(addtocart1)
	c.JSON(http.StatusOK, result)

}

func CreateProfile(c *gin.Context) {
	var profile models.Customer
	if err := c.BindJSON(&profile); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.Insert(profile)
	c.JSON(http.StatusOK, result)
}
func Inventory(c *gin.Context) {
	var inventory models.Inventory
	if err := c.BindJSON(&inventory); err != nil {

		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println(inventory)
	sellerid, err := service.ExtractCustomerID(inventory.SellerId, constants.SecretKey)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token : while Extracting"})
	}
	inventory.SellerId = sellerid
	success, err := service.Inventory(inventory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if success {
		c.JSON(http.StatusOK, gin.H{"data": success})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}

}

var SearchName string

func Getallinventorydata(c *gin.Context) {
	fmt.Println(SearchName)
	result := service.Search(SearchName)
	c.JSON(http.StatusOK, result)
}

func ValidateToken(c *gin.Context) {
	var userdata models.Userdata
	if err := c.BindJSON(&userdata); err != nil {

		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.Validatetoken(userdata.Token)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func Search(c *gin.Context) {
	type Search struct {
		ProductName string `json:"productName" bson:"productName"`
	}
	var search Search
	if err := c.BindJSON(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	SearchName = search.ProductName
}

func Update(c *gin.Context) {
	var update models.Update
	if err := c.BindJSON(&update); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.Update(update)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func Delete(c *gin.Context) {
	var delete models.Delete
	if err := c.BindJSON(&delete); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.Delete(delete)
	c.JSON(http.StatusOK, result)

}

func CheckSeller(c *gin.Context) {
	var check models.Login
	if err := c.BindJSON(&check); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	token, success, err := service.CheckSeller(check)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if success {
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}

}

func DeleteProduct(c *gin.Context) {
	var delete models.DeleteProduct
	if err := c.BindJSON(&delete); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println(delete)
	result := service.DeleteProduct(delete)
	c.JSON(http.StatusOK, result)

}

func DeleteProductBySeller(c *gin.Context) {
	var delete models.DeleteBySeller
	if err := c.BindJSON(&delete); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.DeleteProductBySeller(delete)
	c.JSON(http.StatusOK, result)

}

func UpdateProductBySeller(c *gin.Context) {
	var update models.UpdateProduct
	if err := c.BindJSON(&update); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.UpdateProductBySeller(update)
	c.JSON(http.StatusOK, result)
}

func Feedback(c *gin.Context) {
	var feedback models.Feedback

	if err := c.BindJSON(&feedback); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.Feedback(feedback)
	c.JSON(http.StatusOK, result)
}

func SellerFeedback(c *gin.Context) {
	result := service.SellerFeedback()
	c.JSON(http.StatusOK, result)
}

func CustomerFeedback(c *gin.Context) {
	result := service.CustomerFeedback()
	c.JSON(http.StatusOK, result)
}

func Deletefeedback(c *gin.Context) {
	var feedback models.FeedbacktoAdmin

	if err := c.BindJSON(&feedback); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.Deletefeedback(feedback)
	c.JSON(http.StatusOK, result)

}

func GetUser(c *gin.Context) {
	var data models.Address
	var token models.Token
	var ItemsToBuy []models.Item
	if err := c.BindJSON(&token); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println(token.Token)
	id, _ := service.ExtractCustomerID(token.Token, constants.SecretKey)

	data = service.GetUser(id)
	ItemsToBuy = service.Itemstobuy(id)
	fmt.Println(ItemsToBuy)
	service.CustomerOrders(ItemsToBuy, data)

	service.Buynow(id)
	c.JSON(http.StatusOK, data)
}

func TotalAmount(c *gin.Context) {
	var data models.TotalAmount
	var token models.Token
	if err := c.BindJSON(&token); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println(token.Token)
	id, _ := service.ExtractCustomerID(token.Token, constants.SecretKey)
	data.TotalAmount = service.TotalAmount(id)
	c.JSON(http.StatusOK, data)
}

func Orders(c *gin.Context) {
	data := service.Orders()
	c.JSON(http.StatusOK, data)
}

func DeleteOrder(c *gin.Context) {
	var delete models.DeleteOrder
	if err := c.BindJSON(&delete); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	service.DeleteOrder(delete)
	c.JSON(http.StatusOK, gin.H{"success": true})

}
func CustomerOrder(c *gin.Context) {
	var token models.Token
	if err := c.BindJSON(&token); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	data := service.CustomerOrder(token.Token)
	c.JSON(http.StatusOK, data)

}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}

}

func AdminLogin(c *gin.Context) {
	var login models.AdminData
	if err := c.BindJSON(&login); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	token, result := service.AdminLoginCheck(&login)
	if result != 5 {
		c.JSON(http.StatusOK, gin.H{"result": result})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetAllDetailsForAdmin(c *gin.Context) {
	data := service.AdminNeededData()
	c.JSON(http.StatusOK, gin.H{"result": data})
}

func GetWorkers(c *gin.Context) {
	data := service.GetWorkerdata()
	c.JSON(http.StatusOK, gin.H{"result": data})

}

func GetFeedback(c *gin.Context) {
	data := service.GetFeedBacks()
	c.JSON(http.StatusOK, gin.H{"result": data})

}

func CreateWorker(c *gin.Context) {
	var worker models.Workers
	if err := c.BindJSON(&worker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	data := service.CreateWorker(worker)
	c.JSON(http.StatusOK, gin.H{"result": data})
}

func CreateAdmin(c *gin.Context) {
	var admin models.AdminSignup
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result, data := service.CreateAdmin(admin)
	if result == "Created Successfully" {
		c.JSON(http.StatusOK, gin.H{"result": data})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": result})
}

func GetData(c *gin.Context) {
	var data models.Getdata
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	log.Println(data)
	result, err := service.GetData(data)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})
}

func AddEvent(c *gin.Context) {
	var upload models.UploadCalender
	if err := c.BindJSON(&upload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	err := service.AddEvent(upload)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Added Successfully"})
}

func GetEvent(c *gin.Context) {
	var GetData models.GetCalender
	if err := c.BindJSON(&GetData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	data,err := service.GetEvent(GetData)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": data})
}

func VerifyEmail(c *gin.Context) {
	var Data models.VerifyEmail
	if err := c.BindJSON(&Data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	data,err := service.EmailVerification(Data)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": data})
}

func Block(c *gin.Context){
	var data  models.Block
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	log.Println(data)
	result,err := service.Block(data)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})
}

