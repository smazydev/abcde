// internal/service/container.go
package services

type Container struct {
	userService     UserService
	authService     AuthService
	businessService BusinessService
	// Include other services here
}

func NewContainer(userService UserService, authService AuthService, businessService BusinessService) *Container {
	return &Container{
		userService:     userService,
		authService:     authService,
		businessService: businessService,
	}
}

func (c *Container) GetUserService() UserService {
	return c.userService
}

func (c *Container) GetAuthService() AuthService {
	return c.authService
}
func (c *Container) GetBusinessService() BusinessService {
	return c.businessService
}

// Implement methods to retrieve other services as needed
