package requesthandler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-pg/pg"
	"github.com/nnti3n/voz-archive-plus/serviceWorker/vozscrape"
	"github.com/nnti3n/voz-archive-plus/utilities"

	"github.com/gin-gonic/gin"
)

// Env store db
type Env struct {
	Db *pg.DB
}

// FetchAllThread fetch all threads
func (e *Env) FetchAllThread(c *gin.Context) {
	boxID := c.Param("boxID")
	limit, offset := utilities.Pagination(c, 20)

	threads := []vozscrape.Thread{}
	err := e.Db.Model(&threads).Where("box_id = ?", boxID).
		Offset(offset).Limit(limit).Order("id ASC").Select()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data":   threads,
			"params": boxID,
		})
	} else {
		c.JSON(http.StatusNoContent, gin.H{
			"data":   []string{},
			"params": boxID,
		})
	}

}

// FetchSingleThread fetch all posts of thread
func (e *Env) FetchSingleThread(c *gin.Context) {
	threadID, _ := strconv.Atoi(c.Param("threadID"))
	log.Println(c.Param("threadID"))

	thread := vozscrape.Thread{ID: threadID}
	err := e.Db.Select(&thread)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": thread,
		})
	} else {
		c.JSON(http.StatusNoContent, gin.H{
			"err":  err,
			"data": []string{},
		})
	}
}

// FetchThreadPosts fetch all posts of thread
func (e *Env) FetchThreadPosts(c *gin.Context) {
	threadID, _ := strconv.Atoi(c.Param("threadID"))
	limit, offset := utilities.Pagination(c, 20)

	posts := []vozscrape.Post{}
	err := e.Db.Model(&posts).Where("thread_id = ?", threadID).
		Offset(offset).Limit(limit).Order("number ASC").Select()

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": posts,
		})
	} else {
		c.JSON(http.StatusNoContent, gin.H{
			"err":       err,
			"querydata": threadID,
			"data":      []string{},
		})
	}
}
