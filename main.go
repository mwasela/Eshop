package main


import (
	"Eshop/config"
	"Eshop/controllers"
	"Eshop/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {

	r := gin.Default()
	// Add CORS middleware with custom settings
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * 60 * 60,
	}))

	// init DB
	config.ConnectDatabase()

	//public routes
	r.GET("/", controllers.Home)	
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	//protected routes
	api := r.Group("/api")

	api.Use(middlewares.JWTAuth())


	{
		api.GET("/me", controllers.Me)

		api.POST("/categories", controllers.CreateCategory)
		api.GET("/categories", controllers.GetCategories)
		api.GET("/categories/:id", controllers.GetCategoryByID)
		api.PUT("/categories/:id", controllers.UpdateCategory)
		api.DELETE("/categories/:id", controllers.DeleteCategory)

		api.POST("/inventory", controllers.CreateInventoryItem)
		api.GET("/inventory", controllers.GetInventoryItems)
		api.GET("/inventory/:id", controllers.GetInventoryItemByID)
		api.PUT("/inventory/:id", controllers.UpdateInventoryItem)
		api.DELETE("/inventory/:id", controllers.DeleteInventoryItem)

		api.POST("/orderitems", controllers.CreateOrderItem)
		api.GET("/orderitems", controllers.GetOrderItems)
		api.GET("/orderitems/:id", controllers.GetOrderItemByID)
		api.PUT("/orderitems/:id", controllers.UpdateOrderItem)
		api.DELETE("/orderitems/:id", controllers.DeleteOrderItem)

		api.POST("/pricelists", controllers.CreatePriceListItem)
		api.GET("/pricelists", controllers.GetPriceListItems)
		api.GET("/pricelists/:id", controllers.GetPriceListItemByID)
		api.PUT("/pricelists/:id", controllers.UpdatePriceListItem)
		api.DELETE("/pricelists/:id", controllers.DeletePriceListItem)
		
		
		api.POST("/products", controllers.CreateProduct)
		api.GET("/products", controllers.GetProducts)
		api.GET("/products/:id", controllers.GetProductByID)
		api.PUT("/products/:id", controllers.UpdateProduct)
		api.DELETE("/products/:id", controllers.DeleteProduct)
		
		api.POST("/salesorders", controllers.CreateSalesOrder)
		api.GET("/salesorders", controllers.GetSalesOrders)
		api.GET("/salesorders/:id", controllers.GetSalesOrderByID)
		api.PUT("/salesorders/:id", controllers.UpdateSalesOrder)
		api.DELETE("/salesorders/:id", controllers.DeleteSalesOrder)

		api.POST("/suppliers", controllers.CreateSupplier)
		api.GET("/suppliers", controllers.GetSuppliers)
		api.GET("/suppliers/:id", controllers.GetSupplierByID)
		api.PUT("/suppliers/:id", controllers.UpdateSupplier)
		api.DELETE("/suppliers/:id", controllers.DeleteSupplier)
		
		api.POST("/customers", controllers.CreateCustomer)
		api.GET("/customers", controllers.GetCustomers)
		api.GET("/customers/:id", controllers.GetCustomerByID)
		api.PUT("/customers/:id", controllers.UpdateCustomer)
		api.DELETE("/customers/:id", controllers.DeleteCustomer)

		api.POST("/locations", controllers.CreateLocation)
		api.GET("/locations", controllers.GetLocations)
		api.GET("/locations/:id", controllers.GetLocationByID)
		api.PUT("/locations/:id", controllers.UpdateLocation)
		api.DELETE("/locations/:id", controllers.DeleteLocation)

		api.POST("/storetypes", controllers.CreateStoretype)
		api.GET("/storetypes", controllers.GetStoretypes)
		api.GET("/storetypes/:id", controllers.GetStoretypeByID)
		api.PUT("/storetypes/:id", controllers.UpdateStoretype)
		//api.DELETE("/storetypes/:id", controllers.DeleteStoreType)	


		api.POST("/stores", controllers.CreateStore)
		api.GET("/stores", controllers.GetStores)
		api.GET("/stores/:id", controllers.GetStoreByID)
		api.PUT("/stores/:id", controllers.UpdateStore)
		//api.DELETE("/stores/:id", controllers.DeleteStore)

		api.POST("/companies", controllers.CreateCompany)
		api.GET("/companies", controllers.GetCompanies)
		api.GET("/companies/:id", controllers.GetCompanyByID)
		api.PUT("/companies/:id", controllers.UpdateCompany)
		// api.DELETE("/companies/:id", controllers.DeleteCompany)

		api.POST("/users", controllers.CreateUser)
		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUserByID)
		api.PUT("/users/:id", controllers.UpdateUser)
		// api.DELETE("/users/:id", controllers.DeleteUser)

	}

	r.Run(":4000") // listen and serve on 0.0.0.0:4000 (for windows "localhost:4000")
}