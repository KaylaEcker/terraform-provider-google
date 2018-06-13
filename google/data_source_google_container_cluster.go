package google

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceGoogleContainerCluster() *schema.Resource {
	// Generate datasource schema from resource
	dsSchema := datasourceSchemaFromResourceSchema(resourceContainerCluster().Schema)

	// Set 'Required' schema elements
	addRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	addOptionalFieldsToSchema(dsSchema, "project", "zone", "region")

	return &schema.Resource{
		Read:   datasourceContainerClusterRead,
		Schema: dsSchema,
		
		"labels": &schema.Schema{
		Type:     schema.TypeMap,
		Optional: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
		Set:      schema.HashString,
		},

		"label_fingerprint": &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
		},
	}
}

func datasourceContainerClusterRead(d *schema.ResourceData, meta interface{}) error {
	clusterName := d.Get("name").(string)

	d.SetId(clusterName)
	d.Set("labels", image.Labels)
	d.Set("label_fingerprint", image.LabelFingerprint)

	return resourceContainerClusterRead(d, meta)
}
