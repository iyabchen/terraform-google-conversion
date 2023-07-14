// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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

package compute

import (
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ComputeNodeTemplateAssetType string = "compute.googleapis.com/NodeTemplate"

func ResourceConverterComputeNodeTemplate() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: ComputeNodeTemplateAssetType,
		Convert:   GetComputeNodeTemplateCaiObject,
	}
}

func GetComputeNodeTemplateCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetComputeNodeTemplateApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: ComputeNodeTemplateAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "NodeTemplate",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetComputeNodeTemplateApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeNodeTemplateDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeNodeTemplateName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	nodeAffinityLabelsProp, err := expandComputeNodeTemplateNodeAffinityLabels(d.Get("node_affinity_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_affinity_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeAffinityLabelsProp)) && (ok || !reflect.DeepEqual(v, nodeAffinityLabelsProp)) {
		obj["nodeAffinityLabels"] = nodeAffinityLabelsProp
	}
	nodeTypeProp, err := expandComputeNodeTemplateNodeType(d.Get("node_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeTypeProp)) && (ok || !reflect.DeepEqual(v, nodeTypeProp)) {
		obj["nodeType"] = nodeTypeProp
	}
	nodeTypeFlexibilityProp, err := expandComputeNodeTemplateNodeTypeFlexibility(d.Get("node_type_flexibility"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_type_flexibility"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeTypeFlexibilityProp)) && (ok || !reflect.DeepEqual(v, nodeTypeFlexibilityProp)) {
		obj["nodeTypeFlexibility"] = nodeTypeFlexibilityProp
	}
	serverBindingProp, err := expandComputeNodeTemplateServerBinding(d.Get("server_binding"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("server_binding"); !tpgresource.IsEmptyValue(reflect.ValueOf(serverBindingProp)) && (ok || !reflect.DeepEqual(v, serverBindingProp)) {
		obj["serverBinding"] = serverBindingProp
	}
	cpuOvercommitTypeProp, err := expandComputeNodeTemplateCpuOvercommitType(d.Get("cpu_overcommit_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cpu_overcommit_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(cpuOvercommitTypeProp)) && (ok || !reflect.DeepEqual(v, cpuOvercommitTypeProp)) {
		obj["cpuOvercommitType"] = cpuOvercommitTypeProp
	}
	regionProp, err := expandComputeNodeTemplateRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeNodeTemplateDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeTemplateName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeTemplateNodeAffinityLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeNodeTemplateNodeType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeTemplateNodeTypeFlexibility(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCpus, err := expandComputeNodeTemplateNodeTypeFlexibilityCpus(original["cpus"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCpus); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cpus"] = transformedCpus
	}

	transformedMemory, err := expandComputeNodeTemplateNodeTypeFlexibilityMemory(original["memory"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMemory); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["memory"] = transformedMemory
	}

	transformedLocalSsd, err := expandComputeNodeTemplateNodeTypeFlexibilityLocalSsd(original["local_ssd"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocalSsd); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["localSsd"] = transformedLocalSsd
	}

	return transformed, nil
}

func expandComputeNodeTemplateNodeTypeFlexibilityCpus(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeTemplateNodeTypeFlexibilityMemory(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeTemplateNodeTypeFlexibilityLocalSsd(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeTemplateServerBinding(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedType, err := expandComputeNodeTemplateServerBindingType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["type"] = transformedType
	}

	return transformed, nil
}

func expandComputeNodeTemplateServerBindingType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeTemplateCpuOvercommitType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeTemplateRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
