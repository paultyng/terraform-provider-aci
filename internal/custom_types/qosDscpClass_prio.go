package customTypes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// QosDscpClassPrio custom string type.

var _ basetypes.StringTypable = QosDscpClassPrioStringType{}

type QosDscpClassPrioStringType struct {
	basetypes.StringType
}

func (t QosDscpClassPrioStringType) Equal(o attr.Type) bool {
	other, ok := o.(QosDscpClassPrioStringType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t QosDscpClassPrioStringType) String() string {
	return "QosDscpClassPrioStringType"
}

func (t QosDscpClassPrioStringType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	value := QosDscpClassPrioStringValue{
		StringValue: in,
	}

	return value, nil
}

func (t QosDscpClassPrioStringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromString(ctx, stringValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return stringValuable, nil
}

func (t QosDscpClassPrioStringType) ValueType(ctx context.Context) attr.Value {
	return QosDscpClassPrioStringValue{}
}

// QosDscpClassPrio custom string value.

var _ basetypes.StringValuableWithSemanticEquals = QosDscpClassPrioStringValue{}

type QosDscpClassPrioStringValue struct {
	basetypes.StringValue
}

func (v QosDscpClassPrioStringValue) Equal(o attr.Value) bool {
	other, ok := o.(QosDscpClassPrioStringValue)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v QosDscpClassPrioStringValue) Type(ctx context.Context) attr.Type {
	return QosDscpClassPrioStringType{}
}

func (v QosDscpClassPrioStringValue) StringSemanticEquals(ctx context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(QosDscpClassPrioStringValue)

	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected value type was received while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Expected Value Type: "+fmt.Sprintf("%T", v)+"\n"+
				"Got Value Type: "+fmt.Sprintf("%T", newValuable),
		)

		return false, diags
	}

	priorMappedValue := QosDscpClassPrioValueMap(v.StringValue)

	newMappedValue := QosDscpClassPrioValueMap(newValue.StringValue)

	return priorMappedValue.Equal(newMappedValue), diags
}

func (v QosDscpClassPrioStringValue) NamedValueString() string {
	return QosDscpClassPrioValueMap(basetypes.NewStringValue(v.ValueString())).ValueString()
}

func QosDscpClassPrioValueMap(value basetypes.StringValue) basetypes.StringValue {
	matchMap := map[string]string{
		"0": "unspecified",
		"1": "level3",
		"2": "level2",
		"3": "level1",
		"7": "level6",
		"8": "level5",
		"9": "level4",
	}

	if val, ok := matchMap[value.ValueString()]; ok {
		return basetypes.NewStringValue(val)
	}

	return value
}

func NewQosDscpClassPrioStringNull() QosDscpClassPrioStringValue {
	return QosDscpClassPrioStringValue{
		StringValue: basetypes.NewStringNull(),
	}
}

func NewQosDscpClassPrioStringUnknown() QosDscpClassPrioStringValue {
	return QosDscpClassPrioStringValue{
		StringValue: basetypes.NewStringUnknown(),
	}
}

func NewQosDscpClassPrioStringValue(value string) QosDscpClassPrioStringValue {
	return QosDscpClassPrioStringValue{
		StringValue: basetypes.NewStringValue(value),
	}
}

func NewQosDscpClassPrioStringPointerValue(value *string) QosDscpClassPrioStringValue {
	return QosDscpClassPrioStringValue{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}