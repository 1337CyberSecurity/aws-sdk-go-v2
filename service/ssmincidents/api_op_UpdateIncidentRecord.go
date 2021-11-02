// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmincidents

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssmincidents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update the details of an incident record. You can use this operation to update
// an incident record from the defined chat channel. For more information about
// using actions in chat channels, see Interacting through chat
// (https://docs.aws.amazon.com/incident-manager/latest/userguide/chat.html#chat-interact).
func (c *Client) UpdateIncidentRecord(ctx context.Context, params *UpdateIncidentRecordInput, optFns ...func(*Options)) (*UpdateIncidentRecordOutput, error) {
	if params == nil {
		params = &UpdateIncidentRecordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateIncidentRecord", params, optFns, c.addOperationUpdateIncidentRecordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateIncidentRecordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateIncidentRecordInput struct {

	// The Amazon Resource Name (ARN) of the incident record you are updating.
	//
	// This member is required.
	Arn *string

	// The Chatbot chat channel where responders can collaborate.
	ChatChannel types.ChatChannel

	// A token that ensures that the operation is called only once with the specified
	// details.
	ClientToken *string

	// Defines the impact of the incident to customers and applications. Providing an
	// impact overwrites the impact provided by the response plan. Possible impacts:
	//
	// *
	// 1 - Critical impact, full application failure that impacts many to all
	// customers.
	//
	// * 2 - High impact, partial application failure with impact to many
	// customers.
	//
	// * 3 - Medium impact, the application is providing reduced service to
	// customers.
	//
	// * 4 - Low impact, customer aren't impacted by the problem yet.
	//
	// * 5
	// - No impact, customers aren't currently impacted but urgent action is needed to
	// avoid impact.
	Impact *int32

	// The Amazon SNS targets that are notified when updates are made to an incident.
	// Using multiple SNS topics creates redundancy in the event that a Region is down
	// during the incident.
	NotificationTargets []types.NotificationTargetItem

	// The status of the incident. An incident can be Open or Resolved.
	Status types.IncidentRecordStatus

	// A longer description of what occurred during the incident.
	Summary *string

	// A brief description of the incident.
	Title *string

	noSmithyDocumentSerde
}

type UpdateIncidentRecordOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateIncidentRecordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateIncidentRecord{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateIncidentRecord{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateIncidentRecordMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateIncidentRecordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIncidentRecord(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateIncidentRecord struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateIncidentRecord) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateIncidentRecord) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateIncidentRecordInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateIncidentRecordInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateIncidentRecordMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateIncidentRecord{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateIncidentRecord(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm-incidents",
		OperationName: "UpdateIncidentRecord",
	}
}
