package identitystore

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/document"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
)

func flattenAddress(apiObject *types.Address) map[string]interface{} {
	if apiObject == nil {
		return nil
	}

	m := map[string]interface{}{}

	if v := apiObject.Country; v != nil {
		m["country"] = aws.ToString(v)
	}

	if v := apiObject.Formatted; v != nil {
		m["formatted"] = aws.ToString(v)
	}

	if v := apiObject.Locality; v != nil {
		m["locality"] = aws.ToString(v)
	}

	if v := apiObject.PostalCode; v != nil {
		m["postal_code"] = aws.ToString(v)
	}

	m["primary"] = apiObject.Primary

	if v := apiObject.Region; v != nil {
		m["region"] = aws.ToString(v)
	}

	if v := apiObject.StreetAddress; v != nil {
		m["street_address"] = aws.ToString(v)
	}

	if v := apiObject.Type; v != nil {
		m["type"] = aws.ToString(v)
	}

	return m
}

func expandAddress(tfMap map[string]interface{}) *types.Address {
	if tfMap == nil {
		return nil
	}

	a := &types.Address{}

	if v, ok := tfMap["country"].(string); ok && v != "" {
		a.Country = aws.String(v)
	}

	if v, ok := tfMap["formatted"].(string); ok && v != "" {
		a.Formatted = aws.String(v)
	}

	if v, ok := tfMap["locality"].(string); ok && v != "" {
		a.Locality = aws.String(v)
	}

	if v, ok := tfMap["postal_code"].(string); ok && v != "" {
		a.PostalCode = aws.String(v)
	}

	a.Primary = tfMap["primary"].(bool)

	if v, ok := tfMap["region"].(string); ok && v != "" {
		a.Region = aws.String(v)
	}

	if v, ok := tfMap["street_address"].(string); ok && v != "" {
		a.StreetAddress = aws.String(v)
	}

	if v, ok := tfMap["type"].(string); ok && v != "" {
		a.Type = aws.String(v)
	}

	return a
}

func flattenAddresses(apiObjects []types.Address) []interface{} {
	if len(apiObjects) == 0 {
		return nil
	}

	var l []interface{}

	for _, apiObject := range apiObjects {
		apiObject := apiObject
		l = append(l, flattenAddress(&apiObject))
	}

	return l
}

func expandAddresses(tfList []interface{}) []types.Address {
	s := make([]types.Address, 0, len(tfList))

	for _, r := range tfList {
		m, ok := r.(map[string]interface{})

		if !ok {
			continue
		}

		a := expandAddress(m)

		if a == nil {
			continue
		}

		s = append(s, *a)
	}

	return s
}

func expandAlternateIdentifier(tfMap map[string]interface{}) types.AlternateIdentifier {
	if tfMap == nil {
		return nil
	}

	if v, ok := tfMap["external_id"]; ok && len(v.([]interface{})) > 0 {
		return &types.AlternateIdentifierMemberExternalId{
			Value: *expandExternalId(v.([]interface{})[0].(map[string]interface{})),
		}
	} else if v, ok := tfMap["unique_attribute"]; ok && len(v.([]interface{})) > 0 {
		return &types.AlternateIdentifierMemberUniqueAttribute{
			Value: *expandUniqueAttribute(v.([]interface{})[0].(map[string]interface{})),
		}
	} else {
		return nil
	}
}

func expandExternalId(tfMap map[string]interface{}) *types.ExternalId {
	if tfMap == nil {
		return nil
	}

	a := &types.ExternalId{}

	if v, ok := tfMap["id"].(string); ok && v != "" {
		a.Id = aws.String(v)
	}

	if v, ok := tfMap["issuer"].(string); ok && v != "" {
		a.Issuer = aws.String(v)
	}

	return a
}

func flattenExternalId(apiObject *types.ExternalId) map[string]interface{} {
	if apiObject == nil {
		return nil
	}

	m := map[string]interface{}{}

	if v := apiObject.Id; v != nil {
		m["id"] = aws.ToString(v)
	}

	if v := apiObject.Issuer; v != nil {
		m["issuer"] = aws.ToString(v)
	}

	return m
}

func flattenExternalIds(apiObjects []types.ExternalId) []interface{} {
	if len(apiObjects) == 0 {
		return nil
	}

	var l []interface{}

	for _, apiObject := range apiObjects {
		apiObject := apiObject
		l = append(l, flattenExternalId(&apiObject))
	}

	return l
}

func expandUniqueAttribute(tfMap map[string]interface{}) *types.UniqueAttribute {
	if tfMap == nil {
		return nil
	}

	a := &types.UniqueAttribute{}

	if v, ok := tfMap["attribute_path"].(string); ok && v != "" {
		a.AttributePath = aws.String(v)
	}

	if v, ok := tfMap["attribute_value"].(string); ok && v != "" {
		a.AttributeValue = document.NewLazyDocument(v)
	}

	return a
}
