// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmincidents

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a custom timeline event on the incident details page of an incident
// record. Timeline events are automatically created by Incident Manager, marking
// key moment during an incident. You can create custom timeline events to mark
// important events that are automatically detected by Incident Manager.
func (c *Client) CreateTimelineEvent(ctx context.Context, params *CreateTimelineEventInput, optFns ...func(*Options)) (*CreateTimelineEventOutput, error) {
	if params == nil {
		params = &CreateTimelineEventInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTimelineEvent", params, optFns, c.addOperationCreateTimelineEventMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTimelineEventOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTimelineEventInput struct {

	// A short description of the event as a valid JSON string. There is no other
	// schema imposed.
	//
	// This member is required.
	EventData *string

	// The time that the event occurred.
	//
	// This member is required.
	EventTime *time.Time

	// The type of the event. You can create timeline events of type Custom Event.
	//
	// This member is required.
	EventType *string

	// The Amazon Resource Name (ARN) of the incident record to which the event will be
	// added.
	//
	// This member is required.
	IncidentRecordArn *string

	// A token ensuring that the action is called only once with the specified details.
	ClientToken *string

	noSmithyDocumentSerde
}

type CreateTimelineEventOutput struct {

	// The ID of the event for easy reference later.
	//
	// This member is required.
	EventId *string

	// The ARN of the incident record that you added the event to.
	//
	// This member is required.
	IncidentRecordArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTimelineEventMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateTimelineEvent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateTimelineEvent{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateTimelineEventMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateTimelineEventValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTimelineEvent(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateTimelineEvent struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateTimelineEvent) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateTimelineEvent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateTimelineEventInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateTimelineEventInput ")
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
func addIdempotencyToken_opCreateTimelineEventMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateTimelineEvent{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateTimelineEvent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm-incidents",
		OperationName: "CreateTimelineEvent",
	}
}
