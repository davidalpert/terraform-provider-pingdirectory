package apiclient

import (
	"context"
	"fmt"
)

func (c *DataSyncService) GetBackendConfig(ctx context.Context) interface{} {
	c.HTTPClient.Get(fmt.Sprintf("%s/config/backends", c.BaseUrl))
	return "TODO"
}
