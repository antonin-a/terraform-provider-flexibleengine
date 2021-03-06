// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file at
//     https://www.github.com/huaweicloud/magic-modules
//
// ----------------------------------------------------------------------------

package flexibleengine

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/huaweicloud/golangsdk"
)

func resourceIdentityRoleV3() *schema.Resource {
	return &schema.Resource{
		Create: resourceIdentityRoleV3Create,
		Read:   resourceIdentityRoleV3Read,
		Update: resourceIdentityRoleV3Update,
		Delete: resourceIdentityRoleV3Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"description": {
				Type:     schema.TypeString,
				Required: true,
			},

			"scope": {
				Type:     schema.TypeString,
				Required: true,
			},

			"policy": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"effect": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"catalog": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"domain_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIdentityRoleV3UserInputParams(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"description":   d.Get("description"),
		"display_layer": d.Get("scope"),
		"display_name":  d.Get("name"),
		"statement":     d.Get("policy"),
	}
}

func resourceIdentityRoleV3Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.identityV3Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}
	client.Endpoint = strings.Replace(client.Endpoint, "v3", "v3.0", 1)

	opts := resourceIdentityRoleV3UserInputParams(d)

	r, err := sendIdentityRoleV3CreateRequest(d, opts, nil, client)
	if err != nil {
		return fmt.Errorf("Error creating IdentityRoleV3: %s", err)
	}

	id, err := navigateValue(r, []string{"role", "id"}, nil)
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id.(string))

	return resourceIdentityRoleV3Read(d, meta)
}

func resourceIdentityRoleV3Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.identityV3Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}
	client.Endpoint = strings.Replace(client.Endpoint, "v3", "v3.0", 1)

	res := make(map[string]interface{})

	err = readIdentityRoleV3Read(d, client, res)
	if err != nil {
		return err
	}

	return setIdentityRoleV3Properties(d, res)
}

func resourceIdentityRoleV3Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.identityV3Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}
	client.Endpoint = strings.Replace(client.Endpoint, "v3", "v3.0", 1)

	opts := resourceIdentityRoleV3UserInputParams(d)

	_, err = sendIdentityRoleV3UpdateRequest(d, opts, nil, client)
	if err != nil {
		return fmt.Errorf("Error updating (IdentityRoleV3: %v): %s", d.Id(), err)
	}

	return resourceIdentityRoleV3Read(d, meta)
}

func resourceIdentityRoleV3Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.identityV3Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}
	client.Endpoint = strings.Replace(client.Endpoint, "v3", "v3.0", 1)

	url, err := replaceVars(d, "OS-ROLE/roles/{id}", nil)
	if err != nil {
		return err
	}
	url = client.ServiceURL(url)

	log.Printf("[DEBUG] Deleting Role %q", d.Id())
	r := golangsdk.Result{}
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		OkCodes:      successHTTPCodes,
		JSONBody:     nil,
		JSONResponse: &r.Body,
		MoreHeaders:  map[string]string{"Content-Type": "application/json"},
	})
	if r.Err != nil {
		return fmt.Errorf("Error deleting Role %q: %s", d.Id(), r.Err)
	}

	return nil
}

func sendIdentityRoleV3CreateRequest(d *schema.ResourceData, opts map[string]interface{},
	arrayIndex map[string]int, client *golangsdk.ServiceClient) (interface{}, error) {
	url := client.ServiceURL("OS-ROLE/roles")

	params, err := buildIdentityRoleV3CreateParameters(opts, arrayIndex)
	if err != nil {
		return nil, fmt.Errorf("Error building the request body of api(create)")
	}

	r := golangsdk.Result{}
	_, r.Err = client.Post(url, params, &r.Body, &golangsdk.RequestOpts{
		OkCodes: successHTTPCodes,
	})
	if r.Err != nil {
		return nil, fmt.Errorf("Error run api(create): %s", r.Err)
	}
	return r.Body, nil
}

func buildIdentityRoleV3CreateParameters(opts map[string]interface{}, arrayIndex map[string]int) (interface{}, error) {
	params := make(map[string]interface{})

	descriptionProp, err := navigateValue(opts, []string{"description"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err := isEmptyValue(reflect.ValueOf(descriptionProp))
	if err != nil {
		return nil, err
	}
	if !e {
		params["description"] = descriptionProp
	}

	displayNameProp, err := navigateValue(opts, []string{"display_name"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err = isEmptyValue(reflect.ValueOf(displayNameProp))
	if err != nil {
		return nil, err
	}
	if !e {
		params["display_name"] = displayNameProp
	}

	policyProp, err := expandIdentityRoleV3CreatePolicy(opts, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err = isEmptyValue(reflect.ValueOf(policyProp))
	if err != nil {
		return nil, err
	}
	if !e {
		params["policy"] = policyProp
	}

	typeProp, err := expandIdentityRoleV3CreateType(opts, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err = isEmptyValue(reflect.ValueOf(typeProp))
	if err != nil {
		return nil, err
	}
	if !e {
		params["type"] = typeProp
	}

	params = map[string]interface{}{"role": params}

	return params, nil
}

func expandIdentityRoleV3CreatePolicy(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	req := make(map[string]interface{})

	statementProp, err := expandIdentityRoleV3CreatePolicyStatement(d, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err := isEmptyValue(reflect.ValueOf(statementProp))
	if err != nil {
		return nil, err
	}
	if !e {
		req["Statement"] = statementProp
	}

	req["Version"] = "1.1"

	return req, nil
}

func expandIdentityRoleV3CreatePolicyStatement(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	newArrayIndex := make(map[string]int)
	if arrayIndex != nil {
		for k, v := range arrayIndex {
			newArrayIndex[k] = v
		}
	}

	v, err := navigateValue(d, []string{"statement"}, newArrayIndex)
	if err != nil {
		return nil, err
	}

	n := len(v.([]interface{}))
	req := make([]interface{}, 0, n)
	for i := 0; i < n; i++ {
		newArrayIndex["statement"] = i
		transformed := make(map[string]interface{})

		actionProp, err := navigateValue(d, []string{"statement", "action"}, newArrayIndex)
		if err != nil {
			return nil, err
		}
		e, err := isEmptyValue(reflect.ValueOf(actionProp))
		if err != nil {
			return nil, err
		}
		if !e {
			transformed["Action"] = actionProp
		}

		effectProp, err := navigateValue(d, []string{"statement", "effect"}, newArrayIndex)
		if err != nil {
			return nil, err
		}
		e, err = isEmptyValue(reflect.ValueOf(effectProp))
		if err != nil {
			return nil, err
		}
		if !e {
			transformed["Effect"] = effectProp
		}

		req = append(req, transformed)
	}

	return req, nil
}

func expandIdentityRoleV3CreateType(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	v, err := navigateValue(d, []string{"display_layer"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if v == "domain" {
		return "AX", nil
	} else if v == "project" {
		return "XA", nil
	}
	return nil, fmt.Errorf("unknown display layer:%v", v)
}

func sendIdentityRoleV3UpdateRequest(d *schema.ResourceData, opts map[string]interface{},
	arrayIndex map[string]int, client *golangsdk.ServiceClient) (interface{}, error) {
	url, err := replaceVars(d, "OS-ROLE/roles/{id}", nil)
	if err != nil {
		return nil, err
	}
	url = client.ServiceURL(url)

	params, err := buildIdentityRoleV3UpdateParameters(opts, arrayIndex)
	if err != nil {
		return nil, fmt.Errorf("Error building the request body of api(update)")
	}

	r := golangsdk.Result{}
	_, r.Err = client.Patch(url, params, &r.Body, &golangsdk.RequestOpts{
		OkCodes: successHTTPCodes,
	})
	if r.Err != nil {
		return nil, fmt.Errorf("Error run api(update): %s", r.Err)
	}
	return r.Body, nil
}

func buildIdentityRoleV3UpdateParameters(opts map[string]interface{}, arrayIndex map[string]int) (interface{}, error) {
	params := make(map[string]interface{})

	descriptionProp, err := navigateValue(opts, []string{"description"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err := isEmptyValue(reflect.ValueOf(descriptionProp))
	if err != nil {
		return nil, err
	}
	if !e {
		params["description"] = descriptionProp
	}

	displayNameProp, err := navigateValue(opts, []string{"display_name"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err = isEmptyValue(reflect.ValueOf(displayNameProp))
	if err != nil {
		return nil, err
	}
	if !e {
		params["display_name"] = displayNameProp
	}

	policyProp, err := expandIdentityRoleV3UpdatePolicy(opts, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err = isEmptyValue(reflect.ValueOf(policyProp))
	if err != nil {
		return nil, err
	}
	if !e {
		params["policy"] = policyProp
	}

	typeProp, err := expandIdentityRoleV3UpdateType(opts, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err = isEmptyValue(reflect.ValueOf(typeProp))
	if err != nil {
		return nil, err
	}
	if !e {
		params["type"] = typeProp
	}

	params = map[string]interface{}{"role": params}

	return params, nil
}

func expandIdentityRoleV3UpdatePolicy(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	req := make(map[string]interface{})

	statementProp, err := expandIdentityRoleV3UpdatePolicyStatement(d, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err := isEmptyValue(reflect.ValueOf(statementProp))
	if err != nil {
		return nil, err
	}
	if !e {
		req["Statement"] = statementProp
	}

	req["Version"] = "1.1"

	return req, nil
}

func expandIdentityRoleV3UpdatePolicyStatement(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	newArrayIndex := make(map[string]int)
	if arrayIndex != nil {
		for k, v := range arrayIndex {
			newArrayIndex[k] = v
		}
	}

	v, err := navigateValue(d, []string{"statement"}, newArrayIndex)
	if err != nil {
		return nil, err
	}

	n := len(v.([]interface{}))
	req := make([]interface{}, 0, n)
	for i := 0; i < n; i++ {
		newArrayIndex["statement"] = i
		transformed := make(map[string]interface{})

		actionProp, err := navigateValue(d, []string{"statement", "action"}, newArrayIndex)
		if err != nil {
			return nil, err
		}
		e, err := isEmptyValue(reflect.ValueOf(actionProp))
		if err != nil {
			return nil, err
		}
		if !e {
			transformed["Action"] = actionProp
		}

		effectProp, err := navigateValue(d, []string{"statement", "effect"}, newArrayIndex)
		if err != nil {
			return nil, err
		}
		e, err = isEmptyValue(reflect.ValueOf(effectProp))
		if err != nil {
			return nil, err
		}
		if !e {
			transformed["Effect"] = effectProp
		}

		req = append(req, transformed)
	}

	return req, nil
}

func expandIdentityRoleV3UpdateType(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	v, err := navigateValue(d, []string{"display_layer"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if v == "domain" {
		return "AX", nil
	} else if v == "project" {
		return "XA", nil
	}
	return nil, fmt.Errorf("unknown display layer:%v", v)
}

func readIdentityRoleV3Read(d *schema.ResourceData, client *golangsdk.ServiceClient, result map[string]interface{}) error {
	url, err := replaceVars(d, "OS-ROLE/roles/{id}", nil)
	if err != nil {
		return err
	}
	url = client.ServiceURL(url)

	r := golangsdk.Result{}
	_, r.Err = client.Get(
		url, &r.Body,
		&golangsdk.RequestOpts{MoreHeaders: map[string]string{"Content-Type": "application/json"}})
	if r.Err != nil {
		return fmt.Errorf("Error running api(read) for resource(IdentityRoleV3: %v), error: %s", d.Id(), r.Err)
	}

	v, err := navigateValue(r.Body, []string{"role"}, nil)
	if err != nil {
		return err
	}
	result["read"] = v

	return nil
}

func setIdentityRoleV3Properties(d *schema.ResourceData, response map[string]interface{}) error {
	opts := resourceIdentityRoleV3UserInputParams(d)

	statementProp, _ := opts["statement"]
	statementProp, err := flattenIdentityRoleV3Statement(response, nil, statementProp)
	if err != nil {
		return fmt.Errorf("Error reading Role:statement, err: %s", err)
	}
	if err = d.Set("policy", statementProp); err != nil {
		return fmt.Errorf("Error setting Role:policy, err: %s", err)
	}

	catalogProp, err := navigateValue(response, []string{"read", "catalog"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Role:catalog, err: %s", err)
	}
	if err = d.Set("catalog", catalogProp); err != nil {
		return fmt.Errorf("Error setting Role:catalog, err: %s", err)
	}

	descriptionProp, err := navigateValue(response, []string{"read", "description"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Role:description, err: %s", err)
	}
	if err = d.Set("description", descriptionProp); err != nil {
		return fmt.Errorf("Error setting Role:description, err: %s", err)
	}

	displayLayerProp, _ := opts["display_layer"]
	displayLayerProp, err = flattenIdentityRoleV3DisplayLayer(response, nil, displayLayerProp)
	if err != nil {
		return fmt.Errorf("Error reading Role:display_layer, err: %s", err)
	}
	if err = d.Set("scope", displayLayerProp); err != nil {
		return fmt.Errorf("Error setting Role:scope, err: %s", err)
	}

	displayNameProp, err := navigateValue(response, []string{"read", "display_name"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Role:display_name, err: %s", err)
	}
	if err = d.Set("name", displayNameProp); err != nil {
		return fmt.Errorf("Error setting Role:name, err: %s", err)
	}

	domainIDProp, err := navigateValue(response, []string{"read", "domain_id"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Role:domain_id, err: %s", err)
	}
	if err = d.Set("domain_id", domainIDProp); err != nil {
		return fmt.Errorf("Error setting Role:domain_id, err: %s", err)
	}

	return nil
}

func flattenIdentityRoleV3Statement(d interface{}, arrayIndex map[string]int, currentValue interface{}) (interface{}, error) {
	result, ok := currentValue.([]interface{})
	if !ok || len(result) == 0 {
		v, err := navigateValue(d, []string{"read", "policy", "Statement"}, arrayIndex)
		if err != nil {
			return nil, err
		}
		n := len(v.([]interface{}))
		result = make([]interface{}, n, n)
	}

	newArrayIndex := make(map[string]int)
	if arrayIndex != nil {
		for k, v := range arrayIndex {
			newArrayIndex[k] = v
		}
	}

	for i := 0; i < len(result); i++ {
		newArrayIndex["read.policy.Statement"] = i
		if result[i] == nil {
			result[i] = make(map[string]interface{})
		}
		r := result[i].(map[string]interface{})

		actionProp, err := navigateValue(d, []string{"read", "policy", "Statement", "Action"}, newArrayIndex)
		if err != nil {
			return nil, fmt.Errorf("Error reading Role:action, err: %s", err)
		}
		r["action"] = actionProp

		effectProp, err := navigateValue(d, []string{"read", "policy", "Statement", "Effect"}, newArrayIndex)
		if err != nil {
			return nil, fmt.Errorf("Error reading Role:effect, err: %s", err)
		}
		r["effect"] = effectProp
	}

	return result, nil
}

func flattenIdentityRoleV3DisplayLayer(d interface{}, arrayIndex map[string]int, currentValue interface{}) (interface{}, error) {
	v, err := navigateValue(d, []string{"read", "type"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if v == "AX" {
		return "domain", nil
	} else if v == "XA" {
		return "project", nil
	}
	return nil, fmt.Errorf("unknown display type:%v", v)
}
