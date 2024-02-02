// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the default settings for new user profiles in the domain.
func (c *Client) UpdateDomain(ctx context.Context, params *UpdateDomainInput, optFns ...func(*Options)) (*UpdateDomainOutput, error) {
	if params == nil {
		params = &UpdateDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDomain", params, optFns, c.addOperationUpdateDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDomainInput struct {

	// The ID of the domain to be updated.
	//
	// This member is required.
	DomainId *string

	// Specifies the VPC used for non-EFS traffic.
	//   - PublicInternetOnly - Non-EFS traffic is through a VPC managed by Amazon
	//   SageMaker, which allows direct internet access.
	//   - VpcOnly - All Studio traffic is through the specified VPC and subnets.
	// This configuration can only be modified if there are no apps in the InService ,
	// Pending , or Deleting state. The configuration cannot be updated if
	// DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn is already
	// set or DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn is
	// provided as part of the same request.
	AppNetworkAccessType types.AppNetworkAccessType

	// The entity that creates and manages the required security groups for inter-app
	// communication in VPCOnly mode. Required when CreateDomain.AppNetworkAccessType
	// is VPCOnly and
	// DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn is
	// provided. If setting up the domain for use with RStudio, this value must be set
	// to Service .
	AppSecurityGroupManagement types.AppSecurityGroupManagement

	// The default settings used to create a space within the domain.
	DefaultSpaceSettings *types.DefaultSpaceSettings

	// A collection of settings.
	DefaultUserSettings *types.UserSettings

	// A collection of DomainSettings configuration values to update.
	DomainSettingsForUpdate *types.DomainSettingsForUpdate

	// The VPC subnets that Studio uses for communication. If removing subnets, ensure
	// there are no apps in the InService , Pending , or Deleting state.
	SubnetIds []string

	noSmithyDocumentSerde
}

type UpdateDomainOutput struct {

	// The Amazon Resource Name (ARN) of the domain.
	DomainArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDomain"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDomain(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDomain",
	}
}
