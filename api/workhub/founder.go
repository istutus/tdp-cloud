package workhub

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/opentdp/go-helper/psutil"
)

// 主机信息

func host(c *gin.Context) {

	var rq *HostInfoParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	c.Set("Payload", gin.H{
		"Stat":         psutil.Detail(rq.WithAddr),
		"MemStat":      psutil.GoMemory(),
		"NumGoroutine": runtime.NumGoroutine(),
	})

}

// 主机IP

func hostIp(c *gin.Context) {

	var rq *HostInfoParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	ipv4, ipv6 := psutil.PublicAddress(rq.Force)

	c.Set("Payload", gin.H{"Ipv4": ipv4, "Ipv6": ipv6})

}

// 请求参数

type HostInfoParam struct {
	Force    bool
	WithAddr bool
}
