package healthchecks

/*
	Usage:
	```
	provider "healthchecks" {
	  api_key = "[YOUR MAIN API KEY]"
	}
	```
*/

import (
	"github.com/drubin/terraform-provider-healthchecks/healthchecks/api"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HEALTHCHECKS_API_KEY", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap:   map[string]*schema.Resource{},
		ConfigureFunc: func(r *schema.ResourceData) (interface{}, error) {
			config := healthchecksapi.New(r.Get("api_key").(string))
			return config, nil
		},
	}
}
