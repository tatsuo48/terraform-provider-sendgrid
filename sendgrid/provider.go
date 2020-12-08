package sendgrid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tatsuo48/terraform-provider-sendgrid/client"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("SENDGRID_API_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"sendgrid_whitelist_ip": resourceWhitelistIP(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"sendgrid_whitelist_ip": dataSourceWhitelistIP(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apikey := d.Get("api_key").(string)
	var diags diag.Diagnostics
	if apikey == "" {
		return nil, diag.Errorf("api key is not set, please see this document https://registry.terraform.io/providers/tatsuo48/sendgrid/latest/docs#authentication")
	}
	return client.SendgridCLient{
		APIKey: apikey,
		Host:   "https://api.sendgrid.com",
	}, diags
}
