// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native-sdk"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:network:AdminRule":
		r = &AdminRule{}
	case "azure-native:network:AdminRuleCollection":
		r = &AdminRuleCollection{}
	case "azure-native:network:ApplicationGateway":
		r = &ApplicationGateway{}
	case "azure-native:network:ApplicationGatewayPrivateEndpointConnection":
		r = &ApplicationGatewayPrivateEndpointConnection{}
	case "azure-native:network:ApplicationSecurityGroup":
		r = &ApplicationSecurityGroup{}
	case "azure-native:network:AzureFirewall":
		r = &AzureFirewall{}
	case "azure-native:network:BastionHost":
		r = &BastionHost{}
	case "azure-native:network:ConfigurationPolicyGroup":
		r = &ConfigurationPolicyGroup{}
	case "azure-native:network:ConnectionMonitor":
		r = &ConnectionMonitor{}
	case "azure-native:network:ConnectivityConfiguration":
		r = &ConnectivityConfiguration{}
	case "azure-native:network:CustomIPPrefix":
		r = &CustomIPPrefix{}
	case "azure-native:network:DdosCustomPolicy":
		r = &DdosCustomPolicy{}
	case "azure-native:network:DdosProtectionPlan":
		r = &DdosProtectionPlan{}
	case "azure-native:network:DefaultAdminRule":
		r = &DefaultAdminRule{}
	case "azure-native:network:DefaultUserRule":
		r = &DefaultUserRule{}
	case "azure-native:network:DnsForwardingRuleset":
		r = &DnsForwardingRuleset{}
	case "azure-native:network:DnsResolver":
		r = &DnsResolver{}
	case "azure-native:network:DscpConfiguration":
		r = &DscpConfiguration{}
	case "azure-native:network:Endpoint":
		r = &Endpoint{}
	case "azure-native:network:Experiment":
		r = &Experiment{}
	case "azure-native:network:ExpressRouteCircuit":
		r = &ExpressRouteCircuit{}
	case "azure-native:network:ExpressRouteCircuitAuthorization":
		r = &ExpressRouteCircuitAuthorization{}
	case "azure-native:network:ExpressRouteCircuitConnection":
		r = &ExpressRouteCircuitConnection{}
	case "azure-native:network:ExpressRouteCircuitPeering":
		r = &ExpressRouteCircuitPeering{}
	case "azure-native:network:ExpressRouteConnection":
		r = &ExpressRouteConnection{}
	case "azure-native:network:ExpressRouteCrossConnectionPeering":
		r = &ExpressRouteCrossConnectionPeering{}
	case "azure-native:network:ExpressRouteGateway":
		r = &ExpressRouteGateway{}
	case "azure-native:network:ExpressRoutePort":
		r = &ExpressRoutePort{}
	case "azure-native:network:ExpressRoutePortAuthorization":
		r = &ExpressRoutePortAuthorization{}
	case "azure-native:network:FirewallPolicy":
		r = &FirewallPolicy{}
	case "azure-native:network:FirewallPolicyRuleCollectionGroup":
		r = &FirewallPolicyRuleCollectionGroup{}
	case "azure-native:network:FirewallPolicyRuleGroup":
		r = &FirewallPolicyRuleGroup{}
	case "azure-native:network:FlowLog":
		r = &FlowLog{}
	case "azure-native:network:ForwardingRule":
		r = &ForwardingRule{}
	case "azure-native:network:FrontDoor":
		r = &FrontDoor{}
	case "azure-native:network:HubRouteTable":
		r = &HubRouteTable{}
	case "azure-native:network:HubVirtualNetworkConnection":
		r = &HubVirtualNetworkConnection{}
	case "azure-native:network:InboundEndpoint":
		r = &InboundEndpoint{}
	case "azure-native:network:InboundNatRule":
		r = &InboundNatRule{}
	case "azure-native:network:IpAllocation":
		r = &IpAllocation{}
	case "azure-native:network:IpGroup":
		r = &IpGroup{}
	case "azure-native:network:LoadBalancer":
		r = &LoadBalancer{}
	case "azure-native:network:LoadBalancerBackendAddressPool":
		r = &LoadBalancerBackendAddressPool{}
	case "azure-native:network:LocalNetworkGateway":
		r = &LocalNetworkGateway{}
	case "azure-native:network:ManagementGroupNetworkManagerConnection":
		r = &ManagementGroupNetworkManagerConnection{}
	case "azure-native:network:NatGateway":
		r = &NatGateway{}
	case "azure-native:network:NatRule":
		r = &NatRule{}
	case "azure-native:network:NetworkExperimentProfile":
		r = &NetworkExperimentProfile{}
	case "azure-native:network:NetworkGroup":
		r = &NetworkGroup{}
	case "azure-native:network:NetworkInterface":
		r = &NetworkInterface{}
	case "azure-native:network:NetworkInterfaceTapConfiguration":
		r = &NetworkInterfaceTapConfiguration{}
	case "azure-native:network:NetworkManager":
		r = &NetworkManager{}
	case "azure-native:network:NetworkProfile":
		r = &NetworkProfile{}
	case "azure-native:network:NetworkSecurityGroup":
		r = &NetworkSecurityGroup{}
	case "azure-native:network:NetworkSecurityPerimeter":
		r = &NetworkSecurityPerimeter{}
	case "azure-native:network:NetworkVirtualAppliance":
		r = &NetworkVirtualAppliance{}
	case "azure-native:network:NetworkWatcher":
		r = &NetworkWatcher{}
	case "azure-native:network:NspAccessRule":
		r = &NspAccessRule{}
	case "azure-native:network:NspAssociation":
		r = &NspAssociation{}
	case "azure-native:network:NspProfile":
		r = &NspProfile{}
	case "azure-native:network:OutboundEndpoint":
		r = &OutboundEndpoint{}
	case "azure-native:network:P2sVpnGateway":
		r = &P2sVpnGateway{}
	case "azure-native:network:P2sVpnServerConfiguration":
		r = &P2sVpnServerConfiguration{}
	case "azure-native:network:PacketCapture":
		r = &PacketCapture{}
	case "azure-native:network:Policy":
		r = &Policy{}
	case "azure-native:network:PrivateDnsZoneGroup":
		r = &PrivateDnsZoneGroup{}
	case "azure-native:network:PrivateEndpoint":
		r = &PrivateEndpoint{}
	case "azure-native:network:PrivateLinkService":
		r = &PrivateLinkService{}
	case "azure-native:network:PrivateLinkServicePrivateEndpointConnection":
		r = &PrivateLinkServicePrivateEndpointConnection{}
	case "azure-native:network:PrivateRecordSet":
		r = &PrivateRecordSet{}
	case "azure-native:network:PrivateZone":
		r = &PrivateZone{}
	case "azure-native:network:Profile":
		r = &Profile{}
	case "azure-native:network:PublicIPAddress":
		r = &PublicIPAddress{}
	case "azure-native:network:PublicIPPrefix":
		r = &PublicIPPrefix{}
	case "azure-native:network:RecordSet":
		r = &RecordSet{}
	case "azure-native:network:Route":
		r = &Route{}
	case "azure-native:network:RouteFilter":
		r = &RouteFilter{}
	case "azure-native:network:RouteFilterRule":
		r = &RouteFilterRule{}
	case "azure-native:network:RouteTable":
		r = &RouteTable{}
	case "azure-native:network:RoutingIntent":
		r = &RoutingIntent{}
	case "azure-native:network:RulesEngine":
		r = &RulesEngine{}
	case "azure-native:network:ScopeConnection":
		r = &ScopeConnection{}
	case "azure-native:network:SecurityAdminConfiguration":
		r = &SecurityAdminConfiguration{}
	case "azure-native:network:SecurityPartnerProvider":
		r = &SecurityPartnerProvider{}
	case "azure-native:network:SecurityRule":
		r = &SecurityRule{}
	case "azure-native:network:SecurityUserConfiguration":
		r = &SecurityUserConfiguration{}
	case "azure-native:network:ServiceEndpointPolicy":
		r = &ServiceEndpointPolicy{}
	case "azure-native:network:ServiceEndpointPolicyDefinition":
		r = &ServiceEndpointPolicyDefinition{}
	case "azure-native:network:StaticMember":
		r = &StaticMember{}
	case "azure-native:network:Subnet":
		r = &Subnet{}
	case "azure-native:network:SubscriptionNetworkManagerConnection":
		r = &SubscriptionNetworkManagerConnection{}
	case "azure-native:network:TrafficManagerUserMetricsKey":
		r = &TrafficManagerUserMetricsKey{}
	case "azure-native:network:UserRule":
		r = &UserRule{}
	case "azure-native:network:UserRuleCollection":
		r = &UserRuleCollection{}
	case "azure-native:network:VirtualApplianceSite":
		r = &VirtualApplianceSite{}
	case "azure-native:network:VirtualHub":
		r = &VirtualHub{}
	case "azure-native:network:VirtualHubBgpConnection":
		r = &VirtualHubBgpConnection{}
	case "azure-native:network:VirtualHubIpConfiguration":
		r = &VirtualHubIpConfiguration{}
	case "azure-native:network:VirtualHubRouteTableV2":
		r = &VirtualHubRouteTableV2{}
	case "azure-native:network:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure-native:network:VirtualNetworkGateway":
		r = &VirtualNetworkGateway{}
	case "azure-native:network:VirtualNetworkGatewayConnection":
		r = &VirtualNetworkGatewayConnection{}
	case "azure-native:network:VirtualNetworkGatewayNatRule":
		r = &VirtualNetworkGatewayNatRule{}
	case "azure-native:network:VirtualNetworkLink":
		r = &VirtualNetworkLink{}
	case "azure-native:network:VirtualNetworkPeering":
		r = &VirtualNetworkPeering{}
	case "azure-native:network:VirtualNetworkTap":
		r = &VirtualNetworkTap{}
	case "azure-native:network:VirtualRouter":
		r = &VirtualRouter{}
	case "azure-native:network:VirtualRouterPeering":
		r = &VirtualRouterPeering{}
	case "azure-native:network:VirtualWan":
		r = &VirtualWan{}
	case "azure-native:network:VpnConnection":
		r = &VpnConnection{}
	case "azure-native:network:VpnGateway":
		r = &VpnGateway{}
	case "azure-native:network:VpnServerConfiguration":
		r = &VpnServerConfiguration{}
	case "azure-native:network:VpnSite":
		r = &VpnSite{}
	case "azure-native:network:WebApplicationFirewallPolicy":
		r = &WebApplicationFirewallPolicy{}
	case "azure-native:network:Zone":
		r = &Zone{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := pulumiazurenativesdk.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"network",
		&module{version},
	)
}
