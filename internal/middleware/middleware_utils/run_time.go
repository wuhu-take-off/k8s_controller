package middleware_utils

import (
	"github.com/gin-gonic/gin"
	"time"
)

const (
	//ResponseBodyConst = "responseBodyWriter"
	StartTimeConst = "startTime"
)

type RunTime struct {
}

// 获取开始时间
func (RunTime) GetStartTime(c *gin.Context) time.Time {
	if tmp, exists := c.Get(StartTimeConst); exists {
		return tmp.(time.Time)
	}
	startTime := time.Now()
	c.Set(StartTimeConst, startTime)
	return startTime
}

// 计算运行时间
func (RunTime) CalculateRunTime(c *gin.Context) time.Duration {
	startTime := RunTime{}.GetStartTime(c)
	return time.Now().Sub(startTime)
}
