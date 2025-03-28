package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type MyProxy struct {
	*core.ProxyResponseWriterV2
}

func (m *MyProxy) Flush() {
	log.Info("in Flush")
}

func (m *MyProxy) FlushError() error {
	log.Info("in FlushError")
	return nil
}

func LambdaEchoProxy(e *echo.Echo) func(c echo.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	adapter := echoadapter.NewV2(e)
	return func(c echo.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		httpReq, err := adapter.EventToRequest(req)
		if err != nil {
			return core.GatewayTimeoutV2(), core.NewLoggedError("Error converting event to request: %v", err)
		}

		httpReq.Header.Add(echo.HeaderXRequestID, req.RequestContext.RequestID)

		respWriter := &MyProxy{
			core.NewProxyResponseWriterV2(),
		}
		adapter.Echo.ServeHTTP(http.ResponseWriter(respWriter), httpReq)

		proxyResponse, err := respWriter.GetProxyResponse()
		if err != nil {
			return core.GatewayTimeoutV2(), core.NewLoggedError("Error getting proxy response: %v", err)
		}

		return proxyResponse, nil
	}
}
