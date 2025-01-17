package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// album represents data about a record album.
type link struct {
	ID          string `json:"id"`
	ShortnedURL string `json:"shortnedUrl"`
	CallbackURL string `json:"callbackUrl"`
}

// albums slice to seed record album data.
var links = []link{
	{ID: "1", ShortnedURL: "asdf", CallbackURL: "https://google.com"},
	{ID: "1", ShortnedURL: "test", CallbackURL: "https://google.com"},
}

func main() {
	router := gin.Default()
	router.GET("/links", getLinks)
	router.GET("/:ShortnedURL", redirectToCallback)
  router.PUT("/createLink", createLink)
	router.Run("localhost:25565")

}

// getAlbums responds with the list of all albums as JSON.
func getLinks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, links)
}

func redirectToCallback(c *gin.Context) {
	shortnedURL := c.Param("ShortnedURL")

	for _, i := range links {
		if i.ShortnedURL == shortnedURL {
			c.Redirect(http.StatusFound, i.CallbackURL)
		}
	}
}

func createLink(c *gin.Context){
      var newLink link

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    
    if err := c.BindJSON(&newLink); err != nil {
        return
    }

    // Add the new album to the slice.
    links = append(links, newLink)
    c.IndentedJSON(http.StatusCreated, newLink)
}
