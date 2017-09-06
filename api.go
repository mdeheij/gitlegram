package main

import (
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mdeheij/gitlegram/gitlab"
	log "github.com/mdeheij/logwrap"
)

func main() {
	c := NewConfig()
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		body, ioerr := ioutil.ReadAll(c.Request.Body)
		if ioerr != nil {
			c.String(400, "Could not read request body")
			log.Critical(ioerr)
			return
		}

		//TODO: Request can be ambiguous
		request, err := gitlab.Parse(string(body))
		if err != nil {
			c.String(400, "Could not parse request body")
			log.Critical(err)
			return
		}
		c.JSON(200, getMessage(request))

	})
	address := c.Address + ":" + strconv.FormatInt(c.Port, 10)
	r.Run(address)
}
