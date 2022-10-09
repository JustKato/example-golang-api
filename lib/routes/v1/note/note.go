package note

import "github.com/gin-gonic/gin"

// This simply is an example route that showcases how to use databases to basically insert crud actions,
// I've gone through with an example notes kind of dealio, you can of course adapt this to wahtever you need, this is mostly to show off the middleware and how it functions on in this regard.

func Handle(r *gin.RouterGroup) {

	// Handle the note creation
	r.PUT("/", handleCreate)

	// Handle the note update
	r.PUT("/:id", handleUpdate)

	// Handle fetching a note by ID
	r.GET("/:id", handleGet)

	// Handle the deletion of a note
	r.DELETE("/:id", handleDelete)

}

func handleCreate(c *gin.Context) {
	// TODO: Create a note
}

func handleGet(c *gin.Context) {
	// TODO: Get a note
}

func handleDelete(c *gin.Context) {
	// TODO: Delete a note
}

func handleUpdate(c *gin.Context) {
	// TODO: Delete a note
}
