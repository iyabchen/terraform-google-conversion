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

const ComputeRegionNetworkEndpointGroupAssetType string = "compute.googleapis.com/RegionNetworkEndpointGroup"

func ResourceConverterComputeRegionNetworkEndpointGroup() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: ComputeRegionNetworkEndpointGroupAssetType,
		Convert:   GetComputeRegionNetworkEndpointGroupCaiObject,
	}
}

func GetComputeRegionNetworkEndpointGroupCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetComputeRegionNetworkEndpointGroupApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: ComputeRegionNetworkEndpointGroupAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "RegionNetworkEndpointGroup",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetComputeRegionNetworkEndpointGroupApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeRegionNetworkEndpointGroupName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeRegionNetworkEndpointGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networkEndpointTypeProp, err := expandComputeRegionNetworkEndpointGroupNetworkEndpointType(d.Get("network_endpoint_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_endpoint_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkEndpointTypeProp)) && (ok || !reflect.DeepEqual(v, networkEndpointTypeProp)) {
		obj["networkEndpointType"] = networkEndpointTypeProp
	}
	pscTargetServiceProp, err := expandComputeRegionNetworkEndpointGroupPscTargetService(d.Get("psc_target_service"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("psc_target_service"); !tpgresource.IsEmptyValue(reflect.ValueOf(pscTargetServiceProp)) && (ok || !reflect.DeepEqual(v, pscTargetServiceProp)) {
		obj["pscTargetService"] = pscTargetServiceProp
	}
	networkProp, err := expandComputeRegionNetworkEndpointGroupNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	subnetworkProp, err := expandComputeRegionNetworkEndpointGroupSubnetwork(d.Get("subnetwork"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("subnetwork"); !tpgresource.IsEmptyValue(reflect.ValueOf(subnetworkProp)) && (ok || !reflect.DeepEqual(v, subnetworkProp)) {
		obj["subnetwork"] = subnetworkProp
	}
	cloudRunProp, err := expandComputeRegionNetworkEndpointGroupCloudRun(d.Get("cloud_run"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cloud_run"); !tpgresource.IsEmptyValue(reflect.ValueOf(cloudRunProp)) && (ok || !reflect.DeepEqual(v, cloudRunProp)) {
		obj["cloudRun"] = cloudRunProp
	}
	appEngineProp, err := expandComputeRegionNetworkEndpointGroupAppEngine(d.Get("app_engine"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("app_engine"); !tpgresource.IsEmptyValue(reflect.ValueOf(appEngineProp)) && (ok || !reflect.DeepEqual(v, appEngineProp)) {
		obj["appEngine"] = appEngineProp
	}
	cloudFunctionProp, err := expandComputeRegionNetworkEndpointGroupCloudFunction(d.Get("cloud_function"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cloud_function"); !tpgresource.IsEmptyValue(reflect.ValueOf(cloudFunctionProp)) && (ok || !reflect.DeepEqual(v, cloudFunctionProp)) {
		obj["cloudFunction"] = cloudFunctionProp
	}
	serverlessDeploymentProp, err := expandComputeRegionNetworkEndpointGroupServerlessDeployment(d.Get("serverless_deployment"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("serverless_deployment"); !tpgresource.IsEmptyValue(reflect.ValueOf(serverlessDeploymentProp)) && (ok || !reflect.DeepEqual(v, serverlessDeploymentProp)) {
		obj["serverlessDeployment"] = serverlessDeploymentProp
	}
	regionProp, err := expandComputeRegionNetworkEndpointGroupRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeRegionNetworkEndpointGroupName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupNetworkEndpointType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupPscTargetService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionNetworkEndpointGroupSubnetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseRegionalFieldValue("subnetworks", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for subnetwork: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionNetworkEndpointGroupCloudRun(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedService, err := expandComputeRegionNetworkEndpointGroupCloudRunService(original["service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedTag, err := expandComputeRegionNetworkEndpointGroupCloudRunTag(original["tag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTag); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tag"] = transformedTag
	}

	transformedUrlMask, err := expandComputeRegionNetworkEndpointGroupCloudRunUrlMask(original["url_mask"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrlMask); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["urlMask"] = transformedUrlMask
	}

	return transformed, nil
}

func expandComputeRegionNetworkEndpointGroupCloudRunService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudRunTag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudRunUrlMask(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupAppEngine(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedService, err := expandComputeRegionNetworkEndpointGroupAppEngineService(original["service"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["service"] = transformedService
	}

	transformedVersion, err := expandComputeRegionNetworkEndpointGroupAppEngineVersion(original["version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["version"] = transformedVersion
	}

	transformedUrlMask, err := expandComputeRegionNetworkEndpointGroupAppEngineUrlMask(original["url_mask"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrlMask); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["urlMask"] = transformedUrlMask
	}

	return transformed, nil
}

func expandComputeRegionNetworkEndpointGroupAppEngineService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupAppEngineVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupAppEngineUrlMask(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudFunction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFunction, err := expandComputeRegionNetworkEndpointGroupCloudFunctionFunction(original["function"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFunction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["function"] = transformedFunction
	}

	transformedUrlMask, err := expandComputeRegionNetworkEndpointGroupCloudFunctionUrlMask(original["url_mask"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrlMask); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["urlMask"] = transformedUrlMask
	}

	return transformed, nil
}

func expandComputeRegionNetworkEndpointGroupCloudFunctionFunction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupCloudFunctionUrlMask(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupServerlessDeployment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPlatform, err := expandComputeRegionNetworkEndpointGroupServerlessDeploymentPlatform(original["platform"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPlatform); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["platform"] = transformedPlatform
	}

	transformedResource, err := expandComputeRegionNetworkEndpointGroupServerlessDeploymentResource(original["resource"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["resource"] = transformedResource
	}

	transformedVersion, err := expandComputeRegionNetworkEndpointGroupServerlessDeploymentVersion(original["version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["version"] = transformedVersion
	}

	transformedUrlMask, err := expandComputeRegionNetworkEndpointGroupServerlessDeploymentUrlMask(original["url_mask"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUrlMask); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["urlMask"] = transformedUrlMask
	}

	return transformed, nil
}

func expandComputeRegionNetworkEndpointGroupServerlessDeploymentPlatform(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupServerlessDeploymentResource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupServerlessDeploymentVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupServerlessDeploymentUrlMask(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkEndpointGroupRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
