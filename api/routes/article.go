package routes

import (
	"backend_test_sharing_vision/api/controller"
	"backend_test_sharing_vision/infrastructure"
)

// PostRoute -> Route for question module
type PostRoute struct {
	Controller controller.PostController
	Handler    infrastructure.GinRouter
}

// NewPostRoute -> initializes new choice rouets
func NewPostRoute(
	controller controller.PostController,
	handler infrastructure.GinRouter,

) PostRoute {
	return PostRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (p PostRoute) Setup() {
	post := p.Handler.Gin.Group("/article") //Router group
	{
		post.GET("/:id", p.Controller.GetPost)
		post.GET("lists/:limit/:offset", p.Controller.GetPosts)
		post.POST("/", p.Controller.AddPost)
		post.DELETE("/:id", p.Controller.DeletePost)
		post.PUT("/:id", p.Controller.UpdatePost)
	}
}
