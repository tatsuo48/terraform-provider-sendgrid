package sendgrid

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sendgrid/sendgrid-go"
)

func dataSourceWhitelistIP() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWhitelistIPRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceWhitelistIPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiKey := m.(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ipID := strconv.Itoa(d.Get("id").(int))
	path := fmt.Sprintf("/v3/access_settings/whitelist/%s", ipID)
	request := sendgrid.GetRequest(apiKey, path, "https://api.sendgrid.com")
	request.Method = "GET"

	r, err := sendgrid.API(request)
	if err != nil {
		return diag.FromErr(err)
	}
	if r.StatusCode != 200 {
		return diag.Errorf("Request is failed\n path: %s\n body: %s\n status code: %d\n ", path, r.Body, r.StatusCode)
	}
	w := whitelistIPResponse{}
	err = json.Unmarshal([]byte(r.Body), &w)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ip", w.Result.IP); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", w.Result.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated_at", w.Result.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
