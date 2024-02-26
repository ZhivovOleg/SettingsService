package api

func (c *Controller) initV1settingsController() {
	c.routerGroup.GET("/", c.GetAllOptions)
	c.routerGroup.POST("/", c.NewOption)
    
	c.routerGroup.GET("/:serviceName", c.GetOptions)
	c.routerGroup.DELETE("/:serviceName", c.RemoveOptions)
	c.routerGroup.PUT("/:serviceName", c.ReplaceOptions)

    c.routerGroup.GET("/:serviceName/:path", c.GetConcreteOption)
	c.routerGroup.DELETE("/:serviceName/:path", c.DeleteConcreteOption)
	c.routerGroup.PATCH("/:serviceName/:path", c.UpdateOption)
}