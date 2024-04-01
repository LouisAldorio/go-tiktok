package gotiktok

import (
	"context"
	"fmt"
	"strings"
)

func (c *Client) GetOrderList(ctx context.Context, p CommonParam, cp CursorPaginationParam, request OrderListRequest) (list OrderList, err error) {
	param, err := c.commonParam(p)
	if err != nil {
		return
	}

	if param, err = c.cursorPaginationParam(param, cp); err != nil {
		return
	}

	if err = c.validate.Struct(&request); err != nil {
		return
	}

	err = c.Post(
		ctx,
		fmt.Sprintf("/order/%s/order/search", c.version),
		param,
		request,
		&list,
	)
	return
}

// GetOrderDetail get order detail.
func (c *Client) GetOrderDetail(ctx context.Context, p CommonParam, orderIds []string) (list OrderDetailList, err error) {
	param, err := c.commonParam(p)
	if err != nil {
		return
	}

	param.Set("ids", strings.Join(orderIds, ","))

	err = c.Get(
		ctx,
		fmt.Sprintf("/order/%s/orders", c.version),
		param,
		&list,
	)
	return
}
