
package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tigrisdata/tigris-client-go/filter"
	"github.com/tigrisdata/tigris-client-go/search"
	"github.com/tigrisdata/tigris-client-go/tigris"
)

func SetupTaskCRUD[T interface{}](r *gin.Engine, db *tigris.Database, name string) {
	setupTaskReadRoute[T](r, db, name)
	setupTaskSearchRoute[T](r, db, name)

	r.POST(fmt.Sprintf("/%s", name), func(c *gin.Context) {
		coll := tigris.GetCollection[T](db)

		var u T
		if err := c.Bind(&u); err != nil {
			return
		}

		if _, err := coll.Insert(c, &u); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, u)
	})

	r.DELETE(fmt.Sprintf("/%s/:id", name), func(c *gin.Context) {
		coll := tigris.GetCollection[T](db)

		if _, err := coll.Delete(c,
		    filter.Eq("id", c.Param("id")),
		); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Status": "DELETED"})
	})
}

// setupReadRoute sets route for reading documents by id.
func setupTaskReadRoute[T interface{}](r *gin.Engine, db *tigris.Database, name string) {
	r.GET(fmt.Sprintf("/%s//:id", name), func(c *gin.Context) {
		coll := tigris.GetCollection[T](db)

		u, err := coll.ReadOne(c,
		    filter.Eq("id", c.Param("id")),
        )
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, u)
	})
}

// Create routes for searching data in a collection.
func setupTaskSearchRoute[T interface{}](r *gin.Engine, db *tigris.Database, name string) {
	r.POST(fmt.Sprintf("/%s/search", name), func(c *gin.Context) {
		coll := tigris.GetCollection[T](db)

		var u search.Request
		if err := c.Bind(&u); err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		it, err := coll.Search(c, &u)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		r := &search.Result[T]{}
		for it.Next(r) {
			c.JSON(http.StatusOK, r)
		}
		if err := it.Err(); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
}
