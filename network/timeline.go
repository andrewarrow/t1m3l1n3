package network

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"t1m3l1n3/keys"
	"time"

	"github.com/gin-gonic/gin"
)

func ToggleFollowPost(c *gin.Context) {
	from := c.Request.Header["Username"]
	to := c.Param("username")
	UniverseLock.Lock()
	bin := universes[uids[uidIndex]].ToggleFollow(from[0], to)
	c.JSON(200, gin.H{"mask": bin})
	UniverseLock.Unlock()
}

func TlzIndex(c *gin.Context) byte {
	index := c.Request.Header["Tlz-Index"]
	i, _ := strconv.Atoi(index[0])
	return byte(i)
}
func ShowTimelines(c *gin.Context) {

	i := TlzIndex(c)
	username := c.Param("username")
	UniverseLock.Lock()
	fromIndex := universes[uids[i]].UsernameToIndex(username) - 1
	c.JSON(200, gin.H{"profile": universes[uids[i]].Profile[fromIndex]})
	UniverseLock.Unlock()
}

func mapIt(tokens []string) (key, val string) {
	if len(tokens) < 2 {
		return "", ""
	}
	return tokens[0], tokens[1]
}
func mapJsonBody(c *gin.Context) map[string]string {
	defer c.Request.Body.Close()
	body, _ := ioutil.ReadAll(c.Request.Body)
	var j map[string]interface{}
	json.Unmarshal(body, &j)
	m := map[string]string{}
	for k, v := range j {
		m[k] = fmt.Sprintf("%v", v)
	}
	return m
}

type Timeline struct {
	Text     string `json:"text"`
	From     string `json:"from"`
	PostedAt int64  `json:"posted_at"`
	Origin   string `json:"origin"`
}

func (t *Timeline) AsTime() time.Time {
	return time.Unix(t.PostedAt, 0)
}

func TimelineFromMap(m map[string]string) *Timeline {
	t := Timeline{}
	t.Text = m["text"]
	t.From = m["from"]
	t.PostedAt, _ = strconv.ParseInt(m["posted_at"], 10, 64)
	t.Origin = m["origin"]

	return &t
}

var UniverseLock sync.Mutex

func NotifyTimeline(c *gin.Context) {
	m := mapJsonBody(c)
	fmt.Println("mapJsonBody", m)
	//t := TimelineFromMap(m)
	//t.AddToByKey()
}

func CreateTimeline(c *gin.Context) {
	m := mapJsonBody(c)
	t := Timeline{}
	t.Text = m["text"]
	t.From = c.Request.Header["Username"][0]
	i := TlzIndex(c)
	s := m["s"]
	sDec, _ := b64.StdEncoding.DecodeString(s)

	UniverseLock.Lock()
	pubKey := universes[uids[i]].UsernameKeys[t.From]
	UniverseLock.Unlock()
	if keys.VerifySig(pubKey, t.Text, sDec) == false {
		c.JSON(422, gin.H{"ok": false, "sig": "failed"})
		return
	}

	t.PostedAt = time.Now().Unix()
	t.Origin = globalInOut.Name

	if t.AddToUniverse(i) == true {
		//TellOutAboutNewTimeline(&t, globalInOut.Out)
		c.JSON(200, gin.H{"ok": true})
		return
	}
	c.JSON(422, gin.H{"ok": false})
}

func (t *Timeline) AddToUniverse(i byte) bool {
	val := true
	UniverseLock.Lock()
	val = universes[uids[i]].BroadcastNewTimeline(t)
	UniverseLock.Unlock()
	if val {
		fmt.Println("Add User or Existing User")
		return true
	}
	return false
}

func TellOutAboutNewTimeline(t *Timeline, out string) {
	os.Setenv("CLT_HOST", fmt.Sprintf("http://%s/", out))
	asBytes, _ := json.Marshal(t)
	DoPost("", "timelines/notify", asBytes)
	os.Setenv("CLT_HOST", "")
}

func PostNewTimelineAs(text, username string) {
	m := map[string]string{"text": text, "username": username}
	asBytes, _ := json.Marshal(m)
	DoPost("", "timelines_as", asBytes)
}
func PostNewTimeline(username, text, s string) {
	m := map[string]string{"text": text, "s": s}
	asBytes, _ := json.Marshal(m)
	DoPost(username, "timelines", asBytes)
}
