// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkidentity/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the full details of an AppInstanceUserEndpoint.
func (c *Client) DescribeAppInstanceUserEndpoint(ctx context.Context, params *DescribeAppInstanceUserEndpointInput, optFns ...func(*Options)) (*DescribeAppInstanceUserEndpointOutput, error) {
	if params == nil {
		params = &DescribeAppInstanceUserEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAppInstanceUserEndpoint", params, optFns, c.addOperationDescribeAppInstanceUserEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAppInstanceUserEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAppInstanceUserEndpointInput struct {

	// The ARN of the AppInstanceUser.
	//
	// This member is required.
	AppInstanceUserArn *string

	// The unique identifier of the AppInstanceUserEndpoint.
	//
	// This member is required.
	EndpointId *string

	noSmithyDocumentSerde
}

type DescribeAppInstanceUserEndpointOutput struct {

	// The full details of an AppInstanceUserEndpoint: the AppInstanceUserArn, ID,
	// name, type, resource ARN, attributes, allow messages, state, and created and
	// last updated timestamps. All timestamps use epoch milliseconds.
	AppInstanceUserEndpoint *types.AppInstanceUserEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAppInstanceUserEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAppInstanceUserEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAppInstanceUserEndpoint{}, middleware.After)
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
	if err = addOpDescribeAppInstanceUserEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAppInstanceUserEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAppInstanceUserEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "DescribeAppInstanceUserEndpoint",
	}
}
