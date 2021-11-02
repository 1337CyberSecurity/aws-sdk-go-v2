// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the behavior configuration of Route 53 Resolver behavior for a single
// VPC from Amazon Virtual Private Cloud.
func (c *Client) UpdateResolverConfig(ctx context.Context, params *UpdateResolverConfigInput, optFns ...func(*Options)) (*UpdateResolverConfigOutput, error) {
	if params == nil {
		params = &UpdateResolverConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateResolverConfig", params, optFns, c.addOperationUpdateResolverConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateResolverConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateResolverConfigInput struct {

	// Indicates whether or not the Resolver will create autodefined rules for reverse
	// DNS lookups. This is enabled by default. Disabling this option will also affect
	// EC2-Classic instances using ClassicLink. For more information, see ClassicLink
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in
	// the Amazon EC2 guide. It can take some time for the status change to be
	// completed.
	//
	// This member is required.
	AutodefinedReverseFlag types.AutodefinedReverseFlag

	// Resource ID of the Amazon VPC that you want to update the Resolver configuration
	// for.
	//
	// This member is required.
	ResourceId *string

	noSmithyDocumentSerde
}

type UpdateResolverConfigOutput struct {

	// An array that contains settings for the specified Resolver configuration.
	ResolverConfig *types.ResolverConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateResolverConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateResolverConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateResolverConfig{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateResolverConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResolverConfig(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opUpdateResolverConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "UpdateResolverConfig",
	}
}
