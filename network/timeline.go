package network

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
)

func ShowTimelines(c *gin.Context) {

	c.JSON(200, gin.H{"message": "pong"})
}

func mapIt(tokens []string) (key, val string) {
	if len(tokens) < 2 {
		return "", ""
	}
	return tokens[0], tokens[1]
}
func CreateTimeline(c *gin.Context) {
	defer c.Request.Body.Close()
	body, _ := ioutil.ReadAll(c.Request.Body)
	m := map[string]string{}
	for _, line := range strings.Split(string(body), "\n") {
		k, v := mapIt(strings.Split(line, "="))
		if k != "" {
			m[k] = v
		}
	}
	fmt.Println(m)

}
