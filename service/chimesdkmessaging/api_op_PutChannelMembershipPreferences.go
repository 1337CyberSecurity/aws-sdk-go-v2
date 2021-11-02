// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the membership preferences of an AppInstanceUser for the specified channel.
// The AppInstanceUser must be a member of the channel. Only the AppInstanceUser
// who owns the membership can set preferences. Users in the AppInstanceAdmin and
// channel moderator roles can't set preferences for other users. Banned users
// can't set membership preferences for the channel from which they are banned.
func (c *Client) PutChannelMembershipPreferences(ctx context.Context, params *PutChannelMembershipPreferencesInput, optFns ...func(*Options)) (*PutChannelMembershipPreferencesOutput, error) {
	if params == nil {
		params = &PutChannelMembershipPreferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutChannelMembershipPreferences", params, optFns, c.addOperationPutChannelMembershipPreferencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutChannelMembershipPreferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutChannelMembershipPreferencesInput struct {

	// The ARN of the channel.
	//
	// This member is required.
	ChannelArn *string

	// The AppInstanceUserARN of the user making the API call.
	//
	// This member is required.
	ChimeBearer *string

	// The AppInstanceUserArn of the member setting the preferences.
	//
	// This member is required.
	MemberArn *string

	// The channel membership preferences of an AppInstanceUser .
	//
	// This member is required.
	Preferences *types.ChannelMembershipPreferences

	noSmithyDocumentSerde
}

type PutChannelMembershipPreferencesOutput struct {

	// The ARN of the channel.
	ChannelArn *string

	// The details of a user.
	Member *types.Identity

	// The ARN and metadata of the member being added.
	Preferences *types.ChannelMembershipPreferences

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutChannelMembershipPreferencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutChannelMembershipPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutChannelMembershipPreferences{}, middleware.After)
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
	if err = addOpPutChannelMembershipPreferencesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutChannelMembershipPreferences(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutChannelMembershipPreferences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "PutChannelMembershipPreferences",
	}
}
