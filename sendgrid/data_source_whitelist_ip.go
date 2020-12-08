package sendgrid

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tatsuo48/terraform-provider-sendgrid/client"
)

func dataSourceWhitelistIP() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWhitelistIPRead,
		Schema: map[string]*schema.Schema{
			"rule_id": {
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
	client := m.(client.SendgridCLient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ipID := strconv.Itoa(d.Get("rule_id").(int))
	path := fmt.Sprintf("/v3/access_settings/whitelist/%s", ipID)
	r, err := client.Get(path)
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
	if err := d.Set("rule_id", w.Result.ID); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(ipID)
	return diags
}
