package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"time"
)

type RequestLog struct {
	Time          time.Time     `json:"time"`
	Level         string        `json:"level"`
	Msg           string        `json:"msg"`
	App           AppLogInfo    `json:"app"`
	Http          HttpLogInfo   `json:"http"`
	CorrelationId string        `json:"correlationId"`
	UserId        string        `json:"userId"`
	Status        int           `json:"status"`
	Ms            time.Duration `json:"ms"`
	Bytes         int64         `json:"bytes"`
	Source        string        `json:"source"`
}

type AppLogInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Commit  string `json:"commit"`
}

type HttpLogInfo struct {
	Method        string `json:"method"`
	Route         string `json:"route"`
	RemoteAddr    string `json:"remoteAddr"`
	XForwardedFor string `json:"X-Forwarded-For"`
}

func formatRequestLog(c echo.Context, v middleware.RequestLoggerValues) {
	l := RequestLog{
		Time:  v.StartTime,
		Level: "INFO",
		Msg:   "Incoming Request",
		App: AppLogInfo{
			Name:    "jot-api",
			Version: "",
			Commit:  "",
		},
		Http: HttpLogInfo{
			Method:        v.Method,
			Route:         v.RoutePath,
			RemoteAddr:    c.Request().RemoteAddr,
			XForwardedFor: "",
		},
		CorrelationId: v.RequestID,
		// we don't have users yet ://///////
		UserId: "",
		Status: v.Status,
		Ms:     v.Latency,
		Bytes:  v.ResponseSize,
		Source: "",
	}
	log.Printf("%v", l)
}
