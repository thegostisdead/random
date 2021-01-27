package main


func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)
	router.GET("/login", showLoginPage)
	router.GET("/home", showHomePage)
}