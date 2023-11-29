// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanroomsml

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cleanroomsml/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Defines the information necessary to create a configured audience model.
func (c *Client) CreateConfiguredAudienceModel(ctx context.Context, params *CreateConfiguredAudienceModelInput, optFns ...func(*Options)) (*CreateConfiguredAudienceModelOutput, error) {
	if params == nil {
		params = &CreateConfiguredAudienceModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConfiguredAudienceModel", params, optFns, c.addOperationCreateConfiguredAudienceModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConfiguredAudienceModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConfiguredAudienceModelInput struct {

	// The Amazon Resource Name (ARN) of the audience model to use for the configured
	// audience model.
	//
	// This member is required.
	AudienceModelArn *string

	// The name of the configured audience model.
	//
	// This member is required.
	Name *string

	// Configure the Amazon S3 location and IAM Role for audiences created using this
	// configured audience model. Each audience will have a unique location. The IAM
	// Role must have s3:PutObject permission on the destination Amazon S3 location.
	// If the destination is protected with Amazon S3 KMS-SSE, then the Role must also
	// have the required KMS permissions.
	//
	// This member is required.
	OutputConfig *types.ConfiguredAudienceModelOutputConfig

	// Whether audience metrics are shared.
	//
	// This member is required.
	SharedAudienceMetrics []types.SharedAudienceMetrics

	// Configure the list of output sizes of audiences that can be created using this
	// configured audience model. A request to StartAudienceGenerationJob that uses
	// this configured audience model must have an audienceSize selected from this
	// list. You can use the ABSOLUTE AudienceSize to configure out audience sizes
	// using the count of identifiers in the output. You can use the Percentage
	// AudienceSize to configure sizes in the range 1-100 percent.
	AudienceSizeConfig *types.AudienceSizeConfig

	// Configure how the service tags audience generation jobs created using this
	// configured audience model. If you specify NONE , the tags from the
	// StartAudienceGenerationJob request determine the tags of the audience generation
	// job. If you specify FROM_PARENT_RESOURCE , the audience generation job inherits
	// the tags from the configured audience model, by default. Tags in the
	// StartAudienceGenerationJob will override the default.
	ChildResourceTagOnCreatePolicy types.TagOnCreatePolicy

	// The description of the configured audience model.
	Description *string

	// The minimum number of users from the seed audience that must match with users
	// in the training data of the audience model.
	MinMatchingSeedSize *int32

	// The optional metadata that you apply to the resource to help you categorize and
	// organize them. Each tag consists of a key and an optional value, both of which
	// you define. The following basic restrictions apply to tags:
	//   - Maximum number of tags per resource - 50.
	//   - For each resource, each tag key must be unique, and each tag key can have
	//   only one value.
	//   - Maximum key length - 128 Unicode characters in UTF-8.
	//   - Maximum value length - 256 Unicode characters in UTF-8.
	//   - If your tagging schema is used across multiple services and resources,
	//   remember that other services may have restrictions on allowed characters.
	//   Generally allowed characters are: letters, numbers, and spaces representable in
	//   UTF-8, and the following characters: + - = . _ : / @.
	//   - Tag keys and values are case sensitive.
	//   - Do not use aws:, AWS:, or any upper or lowercase combination of such as a
	//   prefix for keys as it is reserved for AWS use. You cannot edit or delete tag
	//   keys with this prefix. Values can have this prefix. If a tag value has aws as
	//   its prefix but the key does not, then Forecast considers it to be a user tag and
	//   will count against the limit of 50 tags. Tags with only the key prefix of aws do
	//   not count against your tags per resource limit.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateConfiguredAudienceModelOutput struct {

	// The Amazon Resource Name (ARN) of the configured audience model.
	//
	// This member is required.
	ConfiguredAudienceModelArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConfiguredAudienceModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateConfiguredAudienceModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateConfiguredAudienceModel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateConfiguredAudienceModel"); err != nil {
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
	if err = addOpCreateConfiguredAudienceModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConfiguredAudienceModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConfiguredAudienceModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateConfiguredAudienceModel",
	}
}
