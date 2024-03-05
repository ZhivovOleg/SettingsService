package api

func (c *controller) initV1settingsController() {
	c.routerGroup.GET("/", c.getAllOptions)
	c.routerGroup.POST("/", c.addServiceOptions)
    
	c.routerGroup.GET("/:serviceName", c.getServiceOptions)
	c.routerGroup.DELETE("/:serviceName", c.deleteServiceWithOptions)
	c.routerGroup.PUT("/:serviceName", c.replaceServiceOptions)

    c.routerGroup.GET("/:serviceName/:path", c.getSingleValue)
	c.routerGroup.DELETE("/:serviceName/:path", c.deleteSingleValue)
	c.routerGroup.PATCH("/:serviceName/:path", c.updateSingleValue)
}