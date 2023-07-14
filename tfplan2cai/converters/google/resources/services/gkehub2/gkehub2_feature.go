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

package gkehub2

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const GKEHub2FeatureAssetType string = "gkehub.googleapis.com/Feature"

func ResourceConverterGKEHub2Feature() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: GKEHub2FeatureAssetType,
		Convert:   GetGKEHub2FeatureCaiObject,
	}
}

func GetGKEHub2FeatureCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//gkehub.googleapis.com/projects/{{project}}/locations/{{location}}/features/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetGKEHub2FeatureApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: GKEHub2FeatureAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/gkehub/v1beta/rest",
				DiscoveryName:        "Feature",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetGKEHub2FeatureApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelsProp, err := expandGKEHub2FeatureLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	specProp, err := expandGKEHub2FeatureSpec(d.Get("spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(specProp)) && (ok || !reflect.DeepEqual(v, specProp)) {
		obj["spec"] = specProp
	}

	return obj, nil
}

func expandGKEHub2FeatureLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandGKEHub2FeatureSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMulticlusteringress, err := expandGKEHub2FeatureSpecMulticlusteringress(original["multiclusteringress"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMulticlusteringress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["multiclusteringress"] = transformedMulticlusteringress
	}

	transformedFleetobservability, err := expandGKEHub2FeatureSpecFleetobservability(original["fleetobservability"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFleetobservability); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fleetobservability"] = transformedFleetobservability
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecMulticlusteringress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedConfigMembership, err := expandGKEHub2FeatureSpecMulticlusteringressConfigMembership(original["config_membership"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConfigMembership); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["configMembership"] = transformedConfigMembership
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecMulticlusteringressConfigMembership(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureSpecFleetobservability(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLoggingConfig, err := expandGKEHub2FeatureSpecFleetobservabilityLoggingConfig(original["logging_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLoggingConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["loggingConfig"] = transformedLoggingConfig
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecFleetobservabilityLoggingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDefaultConfig, err := expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigDefaultConfig(original["default_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["defaultConfig"] = transformedDefaultConfig
	}

	transformedFleetScopeLogsConfig, err := expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(original["fleet_scope_logs_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFleetScopeLogsConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fleetScopeLogsConfig"] = transformedFleetScopeLogsConfig
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigDefaultConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMode, err := expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigDefaultConfigMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigDefaultConfigMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMode, err := expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
