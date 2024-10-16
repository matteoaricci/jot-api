package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

type RequestLog struct {
	Time          time.Time   `json:"time"`
	Level         string      `json:"level"`
	Msg           string      `json:"msg"`
	App           AppLogInfo  `json:"app"`
	Http          HttpLogInfo `json:"http"`
	CorrelationId string      `json:"correlationId"`
	UserId        string      `json:"userId"`
	Status        int         `json:"status"`
	Ms            int64       `json:"ms"`
	Bytes         int64       `json:"bytes"`
	Source        string      `json:"source"`
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

func Log(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		l := RequestLog{
			Time:  time.Time{},
			Level: "",
			Msg:   "",
			App: AppLogInfo{
				Name:    "",
				Version: "",
				Commit:  "",
			},
			Http: HttpLogInfo{
				Method:        c.Request().Method,
				Route:         c.Request().URL.Path,
				RemoteAddr:    c.Request().RemoteAddr,
				XForwardedFor: "",
			},
			CorrelationId: "",
			UserId:        "",
			Status:        0,
			Ms:            0,
			Bytes:         0,
			Source:        "",
		}
		log.Printf("%v", l)
		return next(c)
	}
}
