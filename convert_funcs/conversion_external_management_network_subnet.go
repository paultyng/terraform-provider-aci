

func createMgmtSubnet(attributes map[string]interface{}) map[string]interface{} {
	ctx := context.Background()
	var diags diag.Diagnostics
	data := &provider.MgmtSubnetResourceModel{}
	if v, ok := attributes["parent_dn"].(string); ok && v != "" {
		data.ParentDn = types.StringValue(v)
	}
	if v, ok := attributes["annotation"].(string); ok && v != "" {
		data.Annotation = types.StringValue(v)
	}
	if v, ok := attributes["descr"].(string); ok && v != "" {
		data.Descr = types.StringValue(v)
	}
	if v, ok := attributes["ip"].(string); ok && v != "" {
		data.Ip = types.StringValue(v)
	}
	if v, ok := attributes["name"].(string); ok && v != "" {
		data.Name = types.StringValue(v)
	}
	if v, ok := attributes["name_alias"].(string); ok && v != "" {
		data.NameAlias = types.StringValue(v)
	}
	planTagAnnotation := convertToTagAnnotationMgmtSubnet(attributes["annotations"])
	planTagTag := convertToTagTagMgmtSubnet(attributes["tags"])

	newAciMgmtSubnet := provider.GetMgmtSubnetCreateJsonPayload(ctx, &diags, data, planTagAnnotation, planTagAnnotation, planTagTag, planTagTag)

	jsonPayload := newAciMgmtSubnet.EncodeJSON(container.EncodeOptIndent("", "  "))
	payload, err := parseCustomJSON(jsonPayload)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v\n", err)
	}

	provider.SetMgmtSubnetId(ctx, data)
	attrs := payload["mgmtSubnet"].(map[string]interface{})["attributes"].(map[string]interface{})
	attrs["dn"] = data.Id.ValueString()

	if status, ok := attributes["status"].(string); ok && status != "" {
		attrs["status"] = status
	}

	return payload
}
func convertToTagAnnotationMgmtSubnet(resources interface{}) []provider.TagAnnotationMgmtSubnetResourceModel {
	var planResources []provider.TagAnnotationMgmtSubnetResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.TagAnnotationMgmtSubnetResourceModel{
				Key:   types.StringValue(resourceMap["key"].(string)),
				Value: types.StringValue(resourceMap["value"].(string)),
			})
		}
	}
	return planResources
}
func convertToTagTagMgmtSubnet(resources interface{}) []provider.TagTagMgmtSubnetResourceModel {
	var planResources []provider.TagTagMgmtSubnetResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.TagTagMgmtSubnetResourceModel{
				Key:   types.StringValue(resourceMap["key"].(string)),
				Value: types.StringValue(resourceMap["value"].(string)),
			})
		}
	}
	return planResources
}
