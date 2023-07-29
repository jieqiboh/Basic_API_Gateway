// Code generated by hertz generator.

package helloapi

import (
	"context"
	"gateway/biz/model/helloapi"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client/genericclient"
	"sync"
)

func HelloMethod(ctx context.Context, c *app.RequestContext) {
	var req helloapi.HelloReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	cli := myGenClientPool.Get().(genericclient.Client)
	defer myGenClientPool.Put(cli)
	st, err := c.Body()
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := DoHelloMethod(ctx, cli, string(st)) // Pass the generic client and requestContext
	if err != nil {
		c.String(consts.StatusInternalServerError, "Failed to perform generic call")
		return
	}

	c.JSON(consts.StatusOK, resp)
}

var myGenClientPool = sync.Pool{
	New: func() interface{} {
		// Create and return a new object of the type you want to pool.
		return HelloServiceGenericClient()
	},
}