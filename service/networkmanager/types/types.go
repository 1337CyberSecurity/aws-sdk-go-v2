// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Specifies a location in Amazon Web Services.
type AWSLocation struct {

	// The Amazon Resource Name (ARN) of the subnet that the device is located in.
	SubnetArn *string

	// The Zone that the device is located in. Specify the ID of an Availability Zone,
	// Local Zone, Wavelength Zone, or an Outpost.
	Zone *string

	noSmithyDocumentSerde
}

// Describes bandwidth information.
type Bandwidth struct {

	// Download speed in Mbps.
	DownloadSpeed *int32

	// Upload speed in Mbps.
	UploadSpeed *int32

	noSmithyDocumentSerde
}

// Describes a connection.
type Connection struct {

	// The ID of the second device in the connection.
	ConnectedDeviceId *string

	// The ID of the link for the second device in the connection.
	ConnectedLinkId *string

	// The Amazon Resource Name (ARN) of the connection.
	ConnectionArn *string

	// The ID of the connection.
	ConnectionId *string

	// The date and time that the connection was created.
	CreatedAt *time.Time

	// The description of the connection.
	Description *string

	// The ID of the first device in the connection.
	DeviceId *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The ID of the link for the first device in the connection.
	LinkId *string

	// The state of the connection.
	State ConnectionState

	// The tags for the connection.
	Tags []Tag

	noSmithyDocumentSerde
}

// Describes connection health.
type ConnectionHealth struct {

	// The connection status.
	Status ConnectionStatus

	// The time the status was last updated.
	Timestamp *time.Time

	// The connection type.
	Type ConnectionType

	noSmithyDocumentSerde
}

// Describes the association between a customer gateway, a device, and a link.
type CustomerGatewayAssociation struct {

	// The Amazon Resource Name (ARN) of the customer gateway.
	CustomerGatewayArn *string

	// The ID of the device.
	DeviceId *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The ID of the link.
	LinkId *string

	// The association state.
	State CustomerGatewayAssociationState

	noSmithyDocumentSerde
}

// Describes a device.
type Device struct {

	// The Amazon Web Services location of the device.
	AWSLocation *AWSLocation

	// The date and time that the site was created.
	CreatedAt *time.Time

	// The description of the device.
	Description *string

	// The Amazon Resource Name (ARN) of the device.
	DeviceArn *string

	// The ID of the device.
	DeviceId *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The site location.
	Location *Location

	// The device model.
	Model *string

	// The device serial number.
	SerialNumber *string

	// The site ID.
	SiteId *string

	// The device state.
	State DeviceState

	// The tags for the device.
	Tags []Tag

	// The device type.
	Type *string

	// The device vendor.
	Vendor *string

	noSmithyDocumentSerde
}

// Describes a global network.
type GlobalNetwork struct {

	// The date and time that the global network was created.
	CreatedAt *time.Time

	// The description of the global network.
	Description *string

	// The Amazon Resource Name (ARN) of the global network.
	GlobalNetworkArn *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The state of the global network.
	State GlobalNetworkState

	// The tags for the global network.
	Tags []Tag

	noSmithyDocumentSerde
}

// Describes a link.
type Link struct {

	// The bandwidth for the link.
	Bandwidth *Bandwidth

	// The date and time that the link was created.
	CreatedAt *time.Time

	// The description of the link.
	Description *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The Amazon Resource Name (ARN) of the link.
	LinkArn *string

	// The ID of the link.
	LinkId *string

	// The provider of the link.
	Provider *string

	// The ID of the site.
	SiteId *string

	// The state of the link.
	State LinkState

	// The tags for the link.
	Tags []Tag

	// The type of the link.
	Type *string

	noSmithyDocumentSerde
}

// Describes the association between a device and a link.
type LinkAssociation struct {

	// The device ID for the link association.
	DeviceId *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The state of the association.
	LinkAssociationState LinkAssociationState

	// The ID of the link.
	LinkId *string

	noSmithyDocumentSerde
}

// Describes a location.
type Location struct {

	// The physical address.
	Address *string

	// The latitude.
	Latitude *string

	// The longitude.
	Longitude *string

	noSmithyDocumentSerde
}

// Describes a network resource.
type NetworkResource struct {

	// The Amazon Web Services account ID.
	AccountId *string

	// The Amazon Web Services Region.
	AwsRegion *string

	// Information about the resource, in JSON format. Network Manager gets this
	// information by describing the resource using its Describe API call.
	Definition *string

	// The time that the resource definition was retrieved.
	DefinitionTimestamp *time.Time

	// The resource metadata.
	Metadata map[string]string

	// The ARN of the gateway.
	RegisteredGatewayArn *string

	// The ARN of the resource.
	ResourceArn *string

	// The ID of the resource.
	ResourceId *string

	// The resource type. The following are the supported resource types for Direct
	// Connect:
	//
	// * dxcon
	//
	// * dx-gateway
	//
	// * dx-vif
	//
	// The following are the supported
	// resource types for Network Manager:
	//
	// * connection
	//
	// * device
	//
	// * link
	//
	// * site
	//
	// The
	// following are the supported resource types for Amazon VPC:
	//
	// *
	// customer-gateway
	//
	// * transit-gateway
	//
	// * transit-gateway-attachment
	//
	// *
	// transit-gateway-connect-peer
	//
	// * transit-gateway-route-table
	//
	// * vpn-connection
	ResourceType *string

	// The tags.
	Tags []Tag

	noSmithyDocumentSerde
}

// Describes a resource count.
type NetworkResourceCount struct {

	// The resource count.
	Count *int32

	// The resource type.
	ResourceType *string

	noSmithyDocumentSerde
}

// Describes a network resource.
type NetworkResourceSummary struct {

	// Information about the resource, in JSON format. Network Manager gets this
	// information by describing the resource using its Describe API call.
	Definition *string

	// Indicates whether this is a middlebox appliance.
	IsMiddlebox bool

	// The value for the Name tag.
	NameTag *string

	// The ARN of the gateway.
	RegisteredGatewayArn *string

	// The ARN of the resource.
	ResourceArn *string

	// The resource type.
	ResourceType *string

	noSmithyDocumentSerde
}

// Describes a network route.
type NetworkRoute struct {

	// A unique identifier for the route, such as a CIDR block.
	DestinationCidrBlock *string

	// The destinations.
	Destinations []NetworkRouteDestination

	// The ID of the prefix list.
	PrefixListId *string

	// The route state. The possible values are active and blackhole.
	State RouteState

	// The route type. The possible values are propagated and static.
	Type RouteType

	noSmithyDocumentSerde
}

// Describes the destination of a network route.
type NetworkRouteDestination struct {

	// The ID of the resource.
	ResourceId *string

	// The resource type.
	ResourceType *string

	// The ID of the transit gateway attachment.
	TransitGatewayAttachmentId *string

	noSmithyDocumentSerde
}

// Describes the telemetry information for a resource.
type NetworkTelemetry struct {

	// The Amazon Web Services account ID.
	AccountId *string

	// The address.
	Address *string

	// The Amazon Web Services Region.
	AwsRegion *string

	// The connection health.
	Health *ConnectionHealth

	// The ARN of the gateway.
	RegisteredGatewayArn *string

	// The ARN of the resource.
	ResourceArn *string

	// The ID of the resource.
	ResourceId *string

	// The resource type.
	ResourceType *string

	noSmithyDocumentSerde
}

// Describes a path component.
type PathComponent struct {

	// The destination CIDR block in the route table.
	DestinationCidrBlock *string

	// The resource.
	Resource *NetworkResourceSummary

	// The sequence number in the path. The destination is 0.
	Sequence *int32

	noSmithyDocumentSerde
}

// Describes a resource relationship.
type Relationship struct {

	// The ARN of the resource.
	From *string

	// The ARN of the resource.
	To *string

	noSmithyDocumentSerde
}

// Describes a route analysis.
type RouteAnalysis struct {

	// The destination.
	Destination *RouteAnalysisEndpointOptions

	// The forward path.
	ForwardPath *RouteAnalysisPath

	// The ID of the global network.
	GlobalNetworkId *string

	// Indicates whether to analyze the return path. The return path is not analyzed if
	// the forward path analysis does not succeed.
	IncludeReturnPath bool

	// The ID of the AWS account that created the route analysis.
	OwnerAccountId *string

	// The return path.
	ReturnPath *RouteAnalysisPath

	// The ID of the route analysis.
	RouteAnalysisId *string

	// The source.
	Source *RouteAnalysisEndpointOptions

	// The time that the analysis started.
	StartTimestamp *time.Time

	// The status of the route analysis.
	Status RouteAnalysisStatus

	// Indicates whether to include the location of middlebox appliances in the route
	// analysis.
	UseMiddleboxes bool

	noSmithyDocumentSerde
}

// Describes the status of an analysis at completion.
type RouteAnalysisCompletion struct {

	// The reason code. Available only if a connection is not found.
	//
	// *
	// BLACKHOLE_ROUTE_FOR_DESTINATION_FOUND - Found a black hole route with the
	// destination CIDR block.
	//
	// * CYCLIC_PATH_DETECTED - Found the same resource
	// multiple times while traversing the path.
	//
	// *
	// INACTIVE_ROUTE_FOR_DESTINATION_FOUND - Found an inactive route with the
	// destination CIDR block.
	//
	// * MAX_HOPS_EXCEEDED - Analysis exceeded 64 hops without
	// finding the destination.
	//
	// * ROUTE_NOT_FOUND - Cannot find a route table with the
	// destination CIDR block.
	//
	// * TGW_ATTACH_ARN_NO_MATCH - Found an attachment, but
	// not with the correct destination ARN.
	//
	// * TGW_ATTACH_NOT_FOUND - Cannot find an
	// attachment.
	//
	// * TGW_ATTACH_NOT_IN_TGW - Found an attachment, but not to the
	// correct transit gateway.
	//
	// * TGW_ATTACH_STABLE_ROUTE_TABLE_NOT_FOUND - The state
	// of the route table association is not associated.
	ReasonCode RouteAnalysisCompletionReasonCode

	// Additional information about the path. Available only if a connection is not
	// found.
	ReasonContext map[string]string

	// The result of the analysis. If the status is NOT_CONNECTED, check the reason
	// code.
	ResultCode RouteAnalysisCompletionResultCode

	noSmithyDocumentSerde
}

// Describes a source or a destination.
type RouteAnalysisEndpointOptions struct {

	// The IP address.
	IpAddress *string

	// The ARN of the transit gateway.
	TransitGatewayArn *string

	// The ARN of the transit gateway attachment.
	TransitGatewayAttachmentArn *string

	noSmithyDocumentSerde
}

// Describes a source or a destination.
type RouteAnalysisEndpointOptionsSpecification struct {

	// The IP address.
	IpAddress *string

	// The ARN of the transit gateway attachment.
	TransitGatewayAttachmentArn *string

	noSmithyDocumentSerde
}

// Describes a route analysis path.
type RouteAnalysisPath struct {

	// The status of the analysis at completion.
	CompletionStatus *RouteAnalysisCompletion

	// The route analysis path.
	Path []PathComponent

	noSmithyDocumentSerde
}

// Describes a route table.
type RouteTableIdentifier struct {

	// The ARN of the transit gateway route table.
	TransitGatewayRouteTableArn *string

	noSmithyDocumentSerde
}

// Describes a site.
type Site struct {

	// The date and time that the site was created.
	CreatedAt *time.Time

	// The description of the site.
	Description *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The location of the site.
	Location *Location

	// The Amazon Resource Name (ARN) of the site.
	SiteArn *string

	// The ID of the site.
	SiteId *string

	// The state of the site.
	State SiteState

	// The tags for the site.
	Tags []Tag

	noSmithyDocumentSerde
}

// Describes a tag.
type Tag struct {

	// The tag key. Constraints: Maximum length of 128 characters.
	Key *string

	// The tag value. Constraints: Maximum length of 256 characters.
	Value *string

	noSmithyDocumentSerde
}

// Describes a transit gateway Connect peer association.
type TransitGatewayConnectPeerAssociation struct {

	// The ID of the device.
	DeviceId *string

	// The ID of the global network.
	GlobalNetworkId *string

	// The ID of the link.
	LinkId *string

	// The state of the association.
	State TransitGatewayConnectPeerAssociationState

	// The Amazon Resource Name (ARN) of the transit gateway Connect peer.
	TransitGatewayConnectPeerArn *string

	noSmithyDocumentSerde
}

// Describes the registration of a transit gateway to a global network.
type TransitGatewayRegistration struct {

	// The ID of the global network.
	GlobalNetworkId *string

	// The state of the transit gateway registration.
	State *TransitGatewayRegistrationStateReason

	// The Amazon Resource Name (ARN) of the transit gateway.
	TransitGatewayArn *string

	noSmithyDocumentSerde
}

// Describes the status of a transit gateway registration.
type TransitGatewayRegistrationStateReason struct {

	// The code for the state reason.
	Code TransitGatewayRegistrationState

	// The message for the state reason.
	Message *string

	noSmithyDocumentSerde
}

// Describes a validation exception for a field.
type ValidationExceptionField struct {

	// The message for the field.
	//
	// This member is required.
	Message *string

	// The name of the field.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
