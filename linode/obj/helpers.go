package obj

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/linode/terraform-provider-linode/v2/linode/helper"
)

func populateLogAttributes(ctx context.Context, d *schema.ResourceData) context.Context {
	return helper.SetLogFieldBulk(ctx, map[string]any{
		"bucket":     d.Get("bucket"),
		"cluster":    d.Get("cluster"),
		"object_key": d.Get("key"),
	})
}
