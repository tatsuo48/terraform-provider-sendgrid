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

func resourceWhitelistIP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWhitelistIPCreate,
		ReadContext:   resourceWhitelistIPRead,
		DeleteContext: resourceWhitelistIPDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"rule_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"created_at": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWhitelistIPCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(client.SendgridCLient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	path := "/v3/access_settings/whitelist"
	body := fmt.Sprintf(`{"ips": [{"ip": "%s"}]}`, d.Get("ip").(string))
	r, err := client.Post(path, body)
	if err != nil {
		return diag.FromErr(err)
	}
	if r.StatusCode != 201 {
		return diag.Errorf("Request is failed\n path: %s\n body: %s\n status code: %d\n ", path, r.Body, r.StatusCode)
	}
	w := whitelistResponse{}
	err = json.Unmarshal([]byte(r.Body), &w)
	if err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(strconv.Itoa(w.Result[0].ID))
	resourceWhitelistIPRead(ctx, d, m)
	return diags
}

func resourceWhitelistIPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(client.SendgridCLient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ipID := d.Id()

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
	return diags
}

func resourceWhitelistIPDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(client.SendgridCLient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ipID := d.Id()

	path := fmt.Sprintf("/v3/access_settings/whitelist/%s", ipID)
	r, err := client.Delete(path)

	if err != nil {
		return diag.FromErr(err)
	}
	if r.StatusCode != 204 {
		return diag.Errorf("Request is failed\n path: %s\n body: %s\n status code: %d\n ", path, r.Body, r.StatusCode)
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
