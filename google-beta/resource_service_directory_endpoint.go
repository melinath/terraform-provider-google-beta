// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceServiceDirectoryEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceDirectoryEndpointCreate,
		Read:   resourceServiceDirectoryEndpointRead,
		Update: resourceServiceDirectoryEndpointUpdate,
		Delete: resourceServiceDirectoryEndpointDelete,

		Importer: &schema.ResourceImporter{
			State: resourceServiceDirectoryEndpointImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"endpoint_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateRFC1035Name(2, 63),
				Description: `The Resource ID must be 1-63 characters long, including digits,
lowercase letters or the hyphen character.`,
			},
			"service": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The resource name of the service that this endpoint provides.`,
			},
			"address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateIpAddress,
				Description:  `IPv4 or IPv6 address of the endpoint.`,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Metadata for the endpoint. This data can be consumed
by service clients. The entire metadata dictionary may contain
up to 512 characters, spread across all key-value pairs.
Metadata that goes beyond any these limits will be rejected.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"port": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(0, 65535),
				Description: `Port that the endpoint is running on, must be in the
range of [0, 65535]. If unspecified, the default is 0.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name for the endpoint in the format
'projects/*/locations/*/namespaces/*/services/*/endpoints/*'.`,
			},
		},
	}
}

func resourceServiceDirectoryEndpointCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	addressProp, err := expandServiceDirectoryEndpointAddress(d.Get("address"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("address"); !isEmptyValue(reflect.ValueOf(addressProp)) && (ok || !reflect.DeepEqual(v, addressProp)) {
		obj["address"] = addressProp
	}
	portProp, err := expandServiceDirectoryEndpointPort(d.Get("port"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("port"); !isEmptyValue(reflect.ValueOf(portProp)) && (ok || !reflect.DeepEqual(v, portProp)) {
		obj["port"] = portProp
	}
	metadataProp, err := expandServiceDirectoryEndpointMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(metadataProp)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}

	url, err := replaceVars(d, config, "{{ServiceDirectoryBasePath}}{{service}}/endpoints?endpointId={{endpoint_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Endpoint: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", "", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Endpoint: %s", err)
	}
	if err := d.Set("name", flattenServiceDirectoryEndpointName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Endpoint %q: %#v", d.Id(), res)

	return resourceServiceDirectoryEndpointRead(d, meta)
}

func resourceServiceDirectoryEndpointRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ServiceDirectoryBasePath}}{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", "", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ServiceDirectoryEndpoint %q", d.Id()))
	}

	if err := d.Set("name", flattenServiceDirectoryEndpointName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Endpoint: %s", err)
	}
	if err := d.Set("address", flattenServiceDirectoryEndpointAddress(res["address"], d, config)); err != nil {
		return fmt.Errorf("Error reading Endpoint: %s", err)
	}
	if err := d.Set("port", flattenServiceDirectoryEndpointPort(res["port"], d, config)); err != nil {
		return fmt.Errorf("Error reading Endpoint: %s", err)
	}
	if err := d.Set("metadata", flattenServiceDirectoryEndpointMetadata(res["metadata"], d, config)); err != nil {
		return fmt.Errorf("Error reading Endpoint: %s", err)
	}

	return nil
}

func resourceServiceDirectoryEndpointUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	addressProp, err := expandServiceDirectoryEndpointAddress(d.Get("address"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("address"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, addressProp)) {
		obj["address"] = addressProp
	}
	portProp, err := expandServiceDirectoryEndpointPort(d.Get("port"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("port"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, portProp)) {
		obj["port"] = portProp
	}
	metadataProp, err := expandServiceDirectoryEndpointMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}

	url, err := replaceVars(d, config, "{{ServiceDirectoryBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Endpoint %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("address") {
		updateMask = append(updateMask, "address")
	}

	if d.HasChange("port") {
		updateMask = append(updateMask, "port")
	}

	if d.HasChange("metadata") {
		updateMask = append(updateMask, "metadata")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	_, err = sendRequestWithTimeout(config, "PATCH", "", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Endpoint %q: %s", d.Id(), err)
	}

	return resourceServiceDirectoryEndpointRead(d, meta)
}

func resourceServiceDirectoryEndpointDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ServiceDirectoryBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Endpoint %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", "", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Endpoint")
	}

	log.Printf("[DEBUG] Finished deleting Endpoint %q: %#v", d.Id(), res)
	return nil
}

func resourceServiceDirectoryEndpointImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats cannot import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	nameParts := strings.Split(d.Get("name").(string), "/")
	if len(nameParts) == 10 {
		// `projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}/endpoints/{{endpoint_id}}`
		d.Set("service", fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s", nameParts[1], nameParts[3], nameParts[5], nameParts[7]))
		d.Set("endpoint_id", nameParts[9])
	} else if len(nameParts) == 5 {
		// `{{project}}/{{location}}/{{namespace_id}}/{{service_id}}/{{endpoint_id}}`
		d.Set("service", fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s", nameParts[0], nameParts[1], nameParts[2], nameParts[3]))
		d.Set("endpoint_id", nameParts[4])
		id := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s/endpoints/%s", nameParts[0], nameParts[1], nameParts[2], nameParts[3], nameParts[4])
		d.Set("name", id)
		d.SetId(id)
	} else if len(nameParts) == 4 {
		// `{{location}}/{{namespace_id}}/{{service_id}}/{{endpoint_id}}`
		project, err := getProject(d, config)
		if err != nil {
			return nil, err
		}
		d.Set("service", fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s", project, nameParts[0], nameParts[1], nameParts[2]))
		d.Set("endpoint_id", nameParts[3])
		id := fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s/endpoints/%s", project, nameParts[0], nameParts[1], nameParts[2], nameParts[3])
		d.Set("name", id)
		d.SetId(id)
	} else {
		return nil, fmt.Errorf(
			"Saw %s when the name is expected to have shape %s, %s or %s",
			d.Get("name"),
			"projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}/endpoints/{{endpoint_id}}",
			"{{project}}/{{location}}/{{namespace_id}}/{{service_id}}/{{endpoint_id}}",
			"{{location}}/{{namespace_id}}/{{service_id}}/{{endpoint_id}}")
	}
	return []*schema.ResourceData{d}, nil
}

func flattenServiceDirectoryEndpointName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenServiceDirectoryEndpointAddress(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenServiceDirectoryEndpointPort(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenServiceDirectoryEndpointMetadata(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandServiceDirectoryEndpointAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandServiceDirectoryEndpointPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandServiceDirectoryEndpointMetadata(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}