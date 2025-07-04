// Copyright 2021-2025 Nordcloud Oy or its affiliates. All Rights Reserved.

package imagetemplate

import (
	"errors"

	"github.com/nordcloud/terraform-provider-imagefactory/pkg/graphql"
)

func expandTemplateComponents(in []interface{}) *[]graphql.NewTemplateComponent {
	out := []graphql.NewTemplateComponent{}
	for i := range in {
		m := in[i].(map[string]interface{})
		out = append(out, graphql.NewTemplateComponent{
			ID: graphql.String(m["id"].(string)),
		})
	}

	return &out
}

func expandTemplateNotifications(in []interface{}) *[]graphql.NewNotification {
	out := []graphql.NewNotification{}
	for i := range in {
		m := in[i].(map[string]interface{})
		out = append(out, graphql.NewNotification{
			Type: graphql.NotificationType(m["type"].(string)),
			Uri:  graphql.String(m["uri"].(string)),
		})
	}

	return &out
}

func expandTemplateTags(in []interface{}) *[]graphql.NewTag {
	out := []graphql.NewTag{}
	for i := range in {
		m := in[i].(map[string]interface{})
		out = append(out, graphql.NewTag{
			Key:   graphql.String(m["key"].(string)),
			Value: graphql.String(m["value"].(string)),
		})
	}

	return &out
}

func expandAdditionalEBSVolumes(in []interface{}) *[]graphql.NewAdditionalEBSVolumes {
	out := []graphql.NewAdditionalEBSVolumes{}
	for i := range in {
		m := in[i].(map[string]interface{})
		v := graphql.NewAdditionalEBSVolumes{
			Size:       graphql.Int(m["size"].(int)),
			DeviceName: graphql.String(m["device_name"].(string)),
		}
		if m["volume_type"] != nil {
			volumeType := graphql.EBSVolumeType(m["volume_type"].(string))
			if volumeType != "" {
				v.VolumeType = &volumeType
			}
		}
		out = append(out, v)
	}

	return &out
}

func expandTemplateAwsConfig(in []interface{}, scope graphql.Scope) (*graphql.NewTemplateAWSConfig, error) {
	if len(in) == 0 || scope == graphql.ScopeCHINA {
		return nil, nil
	}

	m := in[0].(map[string]interface{})
	if m["region"] == nil || m["region"].(string) == "" {
		return nil, errors.New("AWS region is required for the AWS template with PUBLIC scope")
	}

	region := graphql.String(m["region"].(string))
	tplConfig := &graphql.NewTemplateAWSConfig{
		Region: &region,
	}

	if m["custom_image_name"] != nil || m["custom_image_name"].(string) != "" {
		imageName := graphql.String(m["custom_image_name"].(string))
		if imageName != "" {
			tplConfig.CustomImageName = &imageName
		}
	}

	if m["additional_ebs_volumes"] != nil {
		tplConfig.AdditionalEbsVolumes = expandAdditionalEBSVolumes(m["additional_ebs_volumes"].([]interface{}))
	}

	if m["kms_key_id"] != nil {
		kmsKeyID := graphql.String(m["kms_key_id"].(string))
		if kmsKeyID != "" {
			tplConfig.KmsKeyId = &kmsKeyID
		}
	}

	return tplConfig, nil
}

func expandVMImageDefinitionTemplateAzureConfig(in []interface{}) *graphql.NewVMImageDefinition {
	if len(in) == 0 {
		return nil
	}

	m := in[0].(map[string]interface{})

	return &graphql.NewVMImageDefinition{
		Name:  graphql.String(m["name"].(string)),
		Offer: graphql.String(m["offer"].(string)),
		Sku:   graphql.String(m["sku"].(string)),
	}
}

func expandAdditionalDataDisks(in []interface{}) *[]graphql.NewAdditionalDataDisks {
	out := []graphql.NewAdditionalDataDisks{}
	for i := range in {
		m := in[i].(map[string]interface{})
		out = append(out, graphql.NewAdditionalDataDisks{
			Size: graphql.Int(m["size"].(int)),
		})
	}

	return &out
}

func expandAdditionalSignatures(in []interface{}) *[]graphql.NewUefiKey {
	out := []graphql.NewUefiKey{}
	for i := range in {
		m := in[i].(map[string]interface{})
		out = append(out, graphql.NewUefiKey{
			VariableName: graphql.String(m["variable_name"].(string)),
		})
	}

	return &out
}

func expandTemplateAzureConfig(in []interface{}) *graphql.NewTemplateAZUREConfig {
	if len(in) == 0 {
		return nil
	}

	m := in[0].(map[string]interface{})

	e := graphql.Boolean(m["exclude_from_latest"].(bool))
	eol := graphql.Boolean(m["eol_date_option"].(bool))
	tl := graphql.Boolean(m["trusted_launch"].(bool))
	mi := graphql.Boolean(m["create_managed_image"].(bool))

	rr := []graphql.String{}
	for _, v := range m["replica_regions"].([]interface{}) {
		rr = append(rr, graphql.String(v.(string)))
	}

	out := &graphql.NewTemplateAZUREConfig{
		ExcludeFromLatest:  &e,
		EolDateOption:      &eol,
		ReplicaRegions:     &rr,
		TrustedLaunch:      &tl,
		CreateManagedImage: &mi,
		VmImageDefinition:  expandVMImageDefinitionTemplateAzureConfig(m["vm_image_definition"].([]interface{})),
	}

	if m["additional_data_disks"] != nil {
		out.AdditionalDataDisks = expandAdditionalDataDisks(m["additional_data_disks"].([]interface{}))
	}

	if m["additional_signatures"] != nil {
		out.AdditionalSignatures = expandAdditionalSignatures(m["additional_signatures"].([]interface{}))
	}

	if m["disable_vhd_cleanup"] != nil {
		disableVhdCleanup := graphql.Boolean(m["disable_vhd_cleanup"].(bool))
		out.DisableVhdCleanup = &disableVhdCleanup
	}

	return out
}

func expandTemplateExoscaleConfig(in []interface{}) *graphql.NewTemplateExoscaleConfig {
	if len(in) == 0 {
		return nil
	}

	m := in[0].(map[string]interface{})
	zone := graphql.String(m["zone"].(string))

	return &graphql.NewTemplateExoscaleConfig{
		Zone: &zone,
	}
}

func expandTemplateGcpConfig(in []interface{}) *graphql.NewTemplateGCPConfig {
	if len(in) == 0 {
		return nil
	}

	var imageName graphql.String

	m := in[0].(map[string]interface{})
	if m["custom_image_name"] != nil || m["custom_image_name"].(string) != "" {
		imageName = graphql.String(m["custom_image_name"].(string))
	}

	return &graphql.NewTemplateGCPConfig{
		CustomImageName: &imageName,
	}
}

func expandTemplateConfig(in []interface{}) (*graphql.NewTemplateConfig, error) {
	if len(in) == 0 {
		return &graphql.NewTemplateConfig{}, nil
	}

	m := in[0].(map[string]interface{})
	scope := graphql.Scope(m["scope"].(string))

	awsCfg, err := expandTemplateAwsConfig(m["aws"].([]interface{}), scope)
	if err != nil {
		return nil, err
	}

	templateConfig := graphql.NewTemplateConfig{
		Aws:             awsCfg,
		Azure:           expandTemplateAzureConfig(m["azure"].([]interface{})),
		Exoscale:        expandTemplateExoscaleConfig(m["exoscale"].([]interface{})),
		Gcp:             expandTemplateGcpConfig(m["gcp"].([]interface{})),
		BuildComponents: expandTemplateComponents(m["build_components"].([]interface{})),
		TestComponents:  expandTemplateComponents(m["test_components"].([]interface{})),
		Notifications:   expandTemplateNotifications(m["notifications"].([]interface{})),
		Tags:            expandTemplateTags(m["tags"].([]interface{})),
		Scope:           &scope,
	}

	if m["cloud_account_ids"] != nil {
		var cloudAccountIDs []graphql.String
		for _, i := range m["cloud_account_ids"].([]interface{}) {
			cloudAccountIDs = append(cloudAccountIDs, graphql.String(i.(string)))
		}
		templateConfig.CloudAccountIds = &cloudAccountIDs
	}

	if m["disable_cyclical_rebuilds"] != nil {
		disableCyclicalRebuilds := graphql.Boolean(m["disable_cyclical_rebuilds"].(bool))
		templateConfig.DisableCyclicalRebuilds = &disableCyclicalRebuilds
	}

	return &templateConfig, nil
}

func flattenTemplateState(in graphql.TemplateState) map[string]string {
	out := map[string]string{
		"status": string(in.Status),
	}
	if in.Error != nil {
		out["error"] = string(*in.Error)
	}

	return out
}
