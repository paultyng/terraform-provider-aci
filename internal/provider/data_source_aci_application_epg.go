// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"context"
	"fmt"

	customTypes "github.com/CiscoDevNet/terraform-provider-aci/v2/internal/custom_types"
	"github.com/ciscoecosystem/aci-go-client/v2/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &FvAEPgDataSource{}

func NewFvAEPgDataSource() datasource.DataSource {
	return &FvAEPgDataSource{}
}

// FvAEPgDataSource defines the data source implementation.
type FvAEPgDataSource struct {
	client *client.Client
}

func (d *FvAEPgDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of datasource: aci_application_epg")
	resp.TypeName = req.ProviderTypeName + "_application_epg"
	tflog.Debug(ctx, "End metadata of datasource: aci_application_epg")
}

func (d *FvAEPgDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of datasource: aci_application_epg")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The application_epg datasource for the 'fvAEPg' class",

		Attributes: map[string]schema.Attribute{
			// Deprecated attributes
			"exception_tag": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "Attribute 'exception_tag' will be deprecated soon, please refer to 'contract_exception_tag' instead.",
			},
			"flood_on_encap": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "Attribute 'flood_on_encap' will be deprecated soon, please refer to 'flood_in_encapsulation' instead.",
			},
			"fwd_ctrl": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "Attribute 'fwd_ctrl' will be deprecated soon, please refer to 'forwarding_control' instead.",
			},
			"has_mcast_source": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "Attribute 'has_mcast_source' will be deprecated soon, please refer to 'has_multicast_source' instead.",
			},
			"is_attr_based_epg": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "Attribute 'is_attr_based_epg' will be deprecated soon, please refer to 'useg_epg' instead.",
			},
			"match_t": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "Attribute 'match_t' will be deprecated soon, please refer to 'match_criteria' instead.",
			},
			"application_profile_dn": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "Attribute 'application_profile_dn' will be deprecated soon, please refer to 'parent_dn' instead.",
			},
			"pc_enf_pref": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "Attribute 'pc_enf_pref' will be deprecated soon, please refer to 'intra_epg_isolation' instead.",
			},
			"pref_gr_memb": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "Attribute 'pref_gr_memb' will be deprecated soon, please refer to 'preferred_group_member' instead.",
			},
			"prio": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "Attribute 'prio' will be deprecated soon, please refer to 'priority' instead.",
			},
			"shutdown": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "Attribute 'shutdown' will be deprecated soon, please refer to 'admin_state' instead.",
			},
			"relation_fv_rs_prov_def": schema.SetAttribute{
				Computed:           true,
				ElementType:        types.StringType,
				DeprecationMessage: "Attribute `relation_fv_rs_prov_def` is deprecated. The attribute will be removed in the next major version of the provider.",
			},
			"relation_fv_rs_aepg_mon_pol": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "Attribute 'relation_fv_rs_aepg_mon_pol' will be deprecated soon, please refer to 'relation_to_application_epg_monitoring_policy.monitoring_policy_name' instead.",
			},
			"relation_fv_rs_bd": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "Attribute 'relation_fv_rs_bd' will be deprecated soon, please refer to 'relation_to_bridge_domain.bridge_domain_name' instead.",
			},
			"relation_fv_rs_cons": schema.SetAttribute{
				Computed:           true,
				ElementType:        types.StringType,
				DeprecationMessage: "Attribute 'relation_fv_rs_cons' will be deprecated soon, please refer to 'relation_to_consumed_contracts.contract_name' instead.",
			},
			"relation_fv_rs_sec_inherited": schema.SetAttribute{
				Computed:           true,
				ElementType:        types.StringType,
				DeprecationMessage: "Attribute 'relation_fv_rs_sec_inherited' will be deprecated soon, please refer to 'relation_to_contract_masters.target_dn' instead.",
			},
			"relation_fv_rs_cust_qos_pol": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "Attribute 'relation_fv_rs_cust_qos_pol' will be deprecated soon, please refer to 'relation_to_custom_qos_policy.custom_qos_policy_name' instead.",
			},
			"relation_fv_rs_dpp_pol": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "Attribute 'relation_fv_rs_dpp_pol' will be deprecated soon, please refer to 'relation_to_data_plane_policing_policy.data_plane_policing_policy_name' instead.",
			},
			"relation_fv_rs_fc_path_att": schema.SetAttribute{
				Computed:           true,
				ElementType:        types.StringType,
				DeprecationMessage: "Attribute 'relation_fv_rs_fc_path_att' will be deprecated soon, please refer to 'relation_to_fibre_channel_paths.fibre_channel_path_name' instead.",
			},
			"relation_fv_rs_cons_if": schema.SetAttribute{
				Computed:           true,
				ElementType:        types.StringType,
				DeprecationMessage: "Attribute 'relation_fv_rs_cons_if' will be deprecated soon, please refer to 'relation_to_imported_contracts.imported_contract_name' instead.",
			},
			"relation_fv_rs_intra_epg": schema.SetAttribute{
				Computed:           true,
				ElementType:        types.StringType,
				DeprecationMessage: "Attribute 'relation_fv_rs_intra_epg' will be deprecated soon, please refer to 'relation_to_intra_epg_contracts.contract_name' instead.",
			},
			"relation_fv_rs_prov": schema.SetAttribute{
				Computed:           true,
				ElementType:        types.StringType,
				DeprecationMessage: "Attribute 'relation_fv_rs_prov' will be deprecated soon, please refer to 'relation_to_provided_contracts.contract_name' instead.",
			},
			"relation_fv_rs_path_att": schema.SetAttribute{
				Computed:           true,
				ElementType:        types.StringType,
				DeprecationMessage: "Attribute 'relation_fv_rs_path_att' will be deprecated soon, please refer to 'relation_to_static_paths.target_dn' instead.",
			},
			"relation_fv_rs_prot_by": schema.SetAttribute{
				Computed:           true,
				ElementType:        types.StringType,
				DeprecationMessage: "Attribute 'relation_fv_rs_prot_by' will be deprecated soon, please refer to 'relation_to_taboo_contracts.taboo_contract_name' instead.",
			},
			"relation_fv_rs_trust_ctrl": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "Attribute 'relation_fv_rs_trust_ctrl' will be deprecated soon, please refer to 'relation_to_trust_control_policy.trust_control_policy_name' instead.",
			},
			// End of deprecated attributes
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the Application EPG object.",
			},
			"parent_dn": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The distinguished name (DN) of the parent object.",
			},
			"annotation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The annotation of the Application EPG object.`,
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The description of the Application EPG object.`,
			},
			"contract_exception_tag": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The contract exception tag of the Application EPG object.`,
			},
			"flood_in_encapsulation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `Flood L2 Multicast/Broadcast and Link Local Layer based on encapsulation.`,
			},
			"forwarding_control": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The forwarding control of the Application EPG object.`,
			},
			"has_multicast_source": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The Application EPG object has a multicast source.`,
			},
			"useg_epg": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The Application EPG object is microsegmented (uSeg).`,
			},
			"match_criteria": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The provider label match criteria.`,
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: `The name of the Application EPG object.`,
			},
			"name_alias": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The name alias of the Application EPG object.`,
			},
			"intra_epg_isolation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `Parameter used to determine whether communication between endpoints within the EPG is blocked.`,
			},
			"pc_tag": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The classification tag used for policy enforcement and zoning.`,
			},
			"preferred_group_member": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `Parameter used to determine whether the EPG is part of the preferred group. Members of this group are allowed to communicate without contracts.`,
			},
			"priority": schema.StringAttribute{
				CustomType:          customTypes.FvAEPgPrioStringType{},
				Computed:            true,
				MarkdownDescription: `The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.`,
			},
			"admin_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `Withdraw AEPg Configuration from all Nodes in the Fabric.`,
			},
			"epg_useg_block_statement": schema.SingleNestedAttribute{
				MarkdownDescription: `A criterion.`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"annotation": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The annotation of the EPG uSeg Block Statement object.`,
					},
					"description": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The description of the EPG uSeg Block Statement object.`,
					},
					"match": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The Matching Rule Type of the EPG uSeg Block Statement object.`,
					},
					"name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The name of the EPG uSeg Block Statement object.`,
					},
					"name_alias": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The name alias of the EPG uSeg Block Statement object.`,
					},
					"owner_key": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The key for enabling clients to own their data for entity correlation.`,
					},
					"owner_tag": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `A tag for enabling clients to add their own data. For example, to indicate who created this object.`,
					},
					"precedence": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The precedence of the EPG uSeg Block Statement object.`,
					},
					"scope": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The scope of the EPG uSeg Block Statement object.`,
					},
				},
			},
			"relation_to_application_epg_monitoring_policy": schema.SingleNestedAttribute{
				MarkdownDescription: `A source relation to the monitoring policy model for the endpoint group semantic scope. This is an internal object.`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"annotation": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The annotation of the Relation To Application EPG Monitoring Policy object.`,
					},
					"monitoring_policy_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The name of the monitoring policy.`,
					},
				},
			},
			"relation_to_bridge_domain": schema.SingleNestedAttribute{
				MarkdownDescription: `A source relation to the bridge domain associated to this endpoint group. This is an internal object.`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"annotation": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The annotation of the Relation To Bridge Domain object.`,
					},
					"bridge_domain_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The name of the bridge domain associated with this EPG.`,
					},
				},
			},
			"relation_to_consumed_contracts": schema.SetNestedAttribute{
				MarkdownDescription: `The Consumer contract profile information.`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"annotation": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The annotation of the Relation To Consumed Contract object.`,
						},
						"priority": schema.StringAttribute{
							CustomType:          customTypes.FvRsConsPrioStringType{},
							Computed:            true,
							MarkdownDescription: `The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.`,
						},
						"contract_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The consumer contract name.`,
						},
					},
				},
			},
			"relation_to_imported_contracts": schema.SetNestedAttribute{
				MarkdownDescription: `A contract for which the EPG will be a consumer.`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"annotation": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The annotation of the Relation To Imported Contract object.`,
						},
						"priority": schema.StringAttribute{
							CustomType:          customTypes.FvRsConsIfPrioStringType{},
							Computed:            true,
							MarkdownDescription: `The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.`,
						},
						"imported_contract_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The contract interface name.`,
						},
					},
				},
			},
			"relation_to_custom_qos_policy": schema.SingleNestedAttribute{
				MarkdownDescription: `A source relation to a custom QoS policy that enables different levels of service to be assigned to network traffic, including specifications for the Differentiated Services Code Point (DSCP) value(s) and the 802.1p Dot1p priority. This is an internal object.`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"annotation": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The annotation of the Relation To Custom Qos Policy object.`,
					},
					"custom_qos_policy_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The Custom QoS traffic policy name.`,
					},
				},
			},
			"relation_to_domains": schema.SetNestedAttribute{
				MarkdownDescription: `An EPG can be linked to a domain profile via the Associated Domain Profiles. The domain profiles attached can be VMM, Physical, L2 External, or L3 External Domains.`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"annotation": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The annotation of the Relation To Domain object.`,
						},
						"binding_type": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The binding type of the Relation To Domain object.`,
						},
						"class_preference": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The class preference of the Relation To Domain object. Set 'useg' to allow microsegmentation.`,
						},
						"custom_epg_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The display name of the user configured port-group.`,
						},
						"delimiter": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The delimiter of the Relation To Domain object.`,
						},
						"encapsulation": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The encapsulation of the Relation To Domain object. The encapsulation refers to the EPG VLAN when class preference is set to 'encap, or to the Secondary VLAN when class preference is set to 'useg'.`,
						},
						"encapsulation_mode": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The encapsulation mode of the Relation To Domain object.`,
						},
						"epg_cos": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The class of service (CoS) of the Relation To Domain object.`,
						},
						"epg_cos_pref": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The class of service (CoS) preference of the Relation To Domain object.`,
						},
						"deployment_immediacy": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The deployment immediacy of the Relation To Domain object. Specifies when the policy is pushed into the hardware policy content-addressable memory (CAM).`,
						},
						"ipam_dhcp_override": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The IP address management (IPAM) DHCP override of the Relation To Domain object. Only applicable for Nutanix domains.`,
						},
						"ipam_enabled": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The IP address management (IPAM) enabled status of the Relation To Domain object. Only applicable for Nutanix domains.`,
						},
						"ipam_gateway": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The IP address management (IPAM) gateway of the Relation To Domain object. Only applicable for Nutanix domains.`,
						},
						"lag_policy_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The link aggregation group (LAG) policy name of the Relation To Domain object.`,
						},
						"netflow_direction": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The NetFlow monitoring direction of the Relation To Domain object.`,
						},
						"enable_netflow": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The Netflow enabled status for the Relation To Domain object.`,
						},
						"number_of_ports": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The number of ports of the Relation To Domain object.`,
						},
						"port_allocation": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `Port allocation for ports.`,
						},
						"primary_encapsulation": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The primary encapsulation of the Relation To Domain object. This is used when the class preference is set to 'useg'.`,
						},
						"primary_encapsulation_inner": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The primary inner encapsulation of the Relation To Domain object. This is used for the portgroup at the VMWare Distributed Virtual Switch (DVS). This VLAN is internal to the DVS and is used for communication between the other VMs and the AVE VM at a host. Traffic is not forwarded to the fabric over the VLAN. Only applicable for Cisco ACI Virtual Edge (AVE) domains.`,
						},
						"resolution_immediacy": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The resolution immediacy of the Relation To Domain object. Specifies if policies are resolved immmediately or when needed.`,
						},
						"secondary_encapsulation_inner": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The secondary inner encapsulation of the Relation To Domain object. This is used for the portgroup at the VMWare Distributed Virtual Switch (DVS). This VLAN is internal to the DVS and is used for communication between the other VMs and the AVE VM at a host. Traffic is not forwarded to the fabric over the VLAN. Only applicable for Cisco ACI Virtual Edge (AVE) domains.`,
						},
						"switching_mode": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The switching mode of the Relation To Domain object.`,
						},
						"target_dn": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The distinguished name of the target Domain object.`,
						},
						"untagged": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The untagged status of the Relation To Domain object.`,
						},
					},
				},
			},
			"relation_to_data_plane_policing_policy": schema.SingleNestedAttribute{
				MarkdownDescription: `Relationship for Dpp QOS policy`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"annotation": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The annotation of the Relation To Data Plane Policing Policy object.`,
					},
					"data_plane_policing_policy_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `Name.`,
					},
				},
			},
			"relation_to_fibre_channel_paths": schema.SetNestedAttribute{
				MarkdownDescription: `this object is used for creation of static association
                     with a Path for fcoe. Existence of this implies that the
                     corresponding set of policies will be resolved into the
                     node to which the relationship points.`,
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"annotation": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The annotation of the Relation To Fibre Channel Path object.`,
						},
						"description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The description of the Relation To Fibre Channel Path object.`,
						},
						"target_dn": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The distinguished name of the target.`,
						},
						"vsan": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The virtual storage area network (VSAN) of the Relation To Fibre Channel Path object.`,
						},
						"vsan_mode": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The virtual storage area network (VSAN) mode of the Relation To Fibre Channel Path object.`,
						},
					},
				},
			},
			"relation_to_intra_epg_contracts": schema.SetNestedAttribute{
				MarkdownDescription: `Intra EPg contract:
                      Represents that the EPg is moving from "allow all within epg" mode
                      to a "deny all within epg" mode.
                      The only type of traffic allowed between EPs in this EPg is the one
                      specified by contracts EPg associates to with this relation.`,
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"annotation": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The annotation of the Relation To Intra EPG Contract object.`,
						},
						"contract_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The contract name.`,
						},
					},
				},
			},
			"relation_to_static_leafs": schema.SetNestedAttribute{
				MarkdownDescription: `The static association with an access group is a bundled or unbundled group of ports. The existence of this object implies that the corresponding set of policies will be resolved into the node to which the relationship points.`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"annotation": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The annotation of the Relation To Static Leaf object.`,
						},
						"description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The description of the Relation To Static Leaf object.`,
						},
						"encapsulation": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The VLAN encapsulation of the Relation To Static Leaf object.`,
						},
						"deployment_immediacy": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The deployment immediacy of the Relation To Static Leaf object. Specifies when the policy is pushed into the hardware policy content-addressable memory (CAM).`,
						},
						"mode": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The static association mode with the path of the Relation To Static Leaf object.`,
						},
						"target_dn": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The distinguished name of the target of this static binding.`,
						},
					},
				},
			},
			"relation_to_static_paths": schema.SetNestedAttribute{
				MarkdownDescription: `A source relation to an abstraction of a path endpoint.`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"annotation": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The annotation of the Relation To Static Path object.`,
						},
						"description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The description of the Relation To Static Path object.`,
						},
						"encapsulation": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The VLAN encapsulation of the Relation To Static Path object.`,
						},
						"deployment_immediacy": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The deployment immediacy of the Relation To Static Path object. Specifies when the policy is pushed into the hardware policy content-addressable memory (CAM).`,
						},
						"mode": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The static association mode of the Relation To Static Path object.`,
						},
						"primary_encapsulation": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The primary VLAN encapsulation of the Relation To Static Path object.`,
						},
						"target_dn": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `null.`,
						},
					},
				},
			},
			"relation_to_taboo_contracts": schema.SetNestedAttribute{
				MarkdownDescription: `The taboo contract for which the EPG will be a provider and consumer.`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"annotation": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The annotation of the Relation To Taboo Contract object.`,
						},
						"taboo_contract_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `A contract for denying specific classes of traffic. Taboo rules are applied in the hardware before applying the rules of regular contracts. Without a contract, the default forwarding policy is to not allow any communication between EPGs.`,
						},
					},
				},
			},
			"relation_to_provided_contracts": schema.SetNestedAttribute{
				MarkdownDescription: `A contract for which the EPG will be a provider.`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"annotation": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The annotation of the Relation To Provided Contract object.`,
						},
						"match_criteria": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The provider label match criteria.`,
						},
						"priority": schema.StringAttribute{
							CustomType:          customTypes.FvRsProvPrioStringType{},
							Computed:            true,
							MarkdownDescription: `The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.`,
						},
						"contract_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The provider contract name.`,
						},
					},
				},
			},
			"relation_to_contract_masters": schema.SetNestedAttribute{
				MarkdownDescription: `Represents that the EPg is inheriting security configuration from another EPg`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"annotation": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The annotation of the Relation To Contract Master object.`,
						},
						"target_dn": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The distinguished name of the target.`,
						},
					},
				},
			},
			"relation_to_trust_control_policy": schema.SingleNestedAttribute{
				MarkdownDescription: `Relationship for FHS trust control`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"annotation": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The annotation of the Relation To Trust Control Policy object.`,
					},
					"trust_control_policy_name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `Name.`,
					},
				},
			},
			"annotations": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The key used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
			"tags": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The key used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
		},
		Blocks: map[string]schema.Block{
			"relation_fv_rs_node_att": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"deployment_immediacy": schema.StringAttribute{
							Computed:           true,
							DeprecationMessage: "Attribute 'deployment_immediacy' will be deprecated soon, please refer to 'relation_to_static_leafs.deployment_immediacy' instead",
						},
						"description": schema.StringAttribute{
							Computed:           true,
							DeprecationMessage: "Attribute 'description' will be deprecated soon, please refer to 'relation_to_static_leafs.description' instead",
						},
						"encap": schema.StringAttribute{
							Computed:           true,
							DeprecationMessage: "Attribute 'encap' will be deprecated soon, please refer to 'relation_to_static_leafs.encapsulation' instead",
						},
						"mode": schema.StringAttribute{
							Computed:           true,
							DeprecationMessage: "Attribute 'mode' will be deprecated soon, please refer to 'relation_to_static_leafs.mode' instead",
						},
						"node_dn": schema.StringAttribute{
							Computed:           true,
							DeprecationMessage: "Attribute 'node_dn' will be deprecated soon, please refer to 'relation_to_static_leafs.target_dn' instead",
						},
					},
				},
			},
		},
	}
	tflog.Debug(ctx, "End schema of datasource: aci_application_epg")
}

func (d *FvAEPgDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of datasource: aci_application_epg")
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
	tflog.Debug(ctx, "End configure of datasource: aci_application_epg")
}

func (d *FvAEPgDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Start read of datasource: aci_application_epg")
	var data *FvAEPgResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setFvAEPgId(ctx, data)

	// Create a copy of the Id for when not found during getAndSetFvAEPgAttributes
	cachedId := data.Id.ValueString()

	tflog.Debug(ctx, fmt.Sprintf("Read of datasource aci_application_epg with id '%s'", data.Id.ValueString()))

	getAndSetFvAEPgAttributes(ctx, &resp.Diagnostics, d.client, data)

	if data.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Failed to read aci_application_epg data source",
			fmt.Sprintf("The aci_application_epg data source with id '%s' has not been found", cachedId),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End read of datasource aci_application_epg with id '%s'", data.Id.ValueString()))
}
