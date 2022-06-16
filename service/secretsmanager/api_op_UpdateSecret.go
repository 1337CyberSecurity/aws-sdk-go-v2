// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the details of a secret, including metadata and the secret value. To
// change the secret value, you can also use PutSecretValue. To change the rotation
// configuration of a secret, use RotateSecret instead. We recommend you avoid
// calling UpdateSecret at a sustained rate of more than once every 10 minutes.
// When you call UpdateSecret to update the secret value, Secrets Manager creates a
// new version of the secret. Secrets Manager removes outdated versions when there
// are more than 100, but it does not remove versions created less than 24 hours
// ago. If you update the secret value more than once every 10 minutes, you create
// more versions than Secrets Manager removes, and you will reach the quota for
// secret versions. If you include SecretString or SecretBinary to create a new
// secret version, Secrets Manager automatically attaches the staging label
// AWSCURRENT to the new version. If you call this operation with a
// ClientRequestToken that matches an existing version's VersionId, the operation
// results in an error. You can't modify an existing version, you can only create a
// new version. To remove a version, remove all staging labels from it. See
// UpdateSecretVersionStage. If you don't specify an KMS encryption key, Secrets
// Manager uses the Amazon Web Services managed key aws/secretsmanager. If this key
// doesn't already exist in your account, then Secrets Manager creates it for you
// automatically. All users and roles in the Amazon Web Services account
// automatically have access to use aws/secretsmanager. Creating aws/secretsmanager
// can result in a one-time significant delay in returning the result. If the
// secret is in a different Amazon Web Services account from the credentials
// calling the API, then you can't use aws/secretsmanager to encrypt the secret,
// and you must create and use a customer managed key. Required permissions:
// secretsmanager:UpdateSecret. For more information, see  IAM policy actions for
// Secrets Manager
// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions)
// and Authentication and access control in Secrets Manager
// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html).
// If you use a customer managed key, you must also have kms:GenerateDataKey and
// kms:Decrypt permissions on the key. For more information, see  Secret encryption
// and decryption
// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/security-encryption.html).
func (c *Client) UpdateSecret(ctx context.Context, params *UpdateSecretInput, optFns ...func(*Options)) (*UpdateSecretOutput, error) {
	if params == nil {
		params = &UpdateSecretInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSecret", params, optFns, c.addOperationUpdateSecretMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSecretOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSecretInput struct {

	// The ARN or name of the secret. For an ARN, we recommend that you specify a
	// complete ARN rather than a partial ARN. See Finding a secret from a partial ARN
	// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/troubleshoot.html#ARN_secretnamehyphen).
	//
	// This member is required.
	SecretId *string

	// If you include SecretString or SecretBinary, then Secrets Manager creates a new
	// version for the secret, and this parameter specifies the unique identifier for
	// the new version. If you use the Amazon Web Services CLI or one of the Amazon Web
	// Services SDKs to call this operation, then you can leave this parameter empty.
	// The CLI or SDK generates a random UUID for you and includes it as the value for
	// this parameter in the request. If you don't use the SDK and instead generate a
	// raw HTTP request to the Secrets Manager service endpoint, then you must generate
	// a ClientRequestToken yourself for the new version and include the value in the
	// request. This value becomes the VersionId of the new version.
	ClientRequestToken *string

	// The description of the secret.
	Description *string

	// The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt
	// new secret versions as well as any existing versions the staging labels
	// AWSCURRENT, AWSPENDING, or AWSPREVIOUS. For more information about versions and
	// staging labels, see Concepts: Version
	// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/getting-started.html#term_version).
	// You can only use the Amazon Web Services managed key aws/secretsmanager if you
	// call this operation using credentials from the same Amazon Web Services account
	// that owns the secret. If the secret is in a different account, then you must use
	// a customer managed key and provide the ARN of that KMS key in this field. The
	// user making the call must have permissions to both the secret and the KMS key in
	// their respective accounts.
	KmsKeyId *string

	// The binary data to encrypt and store in the new version of the secret. We
	// recommend that you store your binary data in a file and then pass the contents
	// of the file as a parameter. Either SecretBinary or SecretString must have a
	// value, but not both. You can't access this parameter in the Secrets Manager
	// console.
	SecretBinary []byte

	// The text data to encrypt and store in the new version of the secret. We
	// recommend you use a JSON structure of key/value pairs for your secret value.
	// Either SecretBinary or SecretString must have a value, but not both.
	SecretString *string

	noSmithyDocumentSerde
}

type UpdateSecretOutput struct {

	// The ARN of the secret that was updated.
	ARN *string

	// The name of the secret that was updated.
	Name *string

	// If Secrets Manager created a new version of the secret during this operation,
	// then VersionId contains the unique identifier of the new version.
	VersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSecretMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSecret{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSecret{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateSecretMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateSecretValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSecret(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateSecret struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateSecret) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateSecret) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateSecretInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateSecretInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateSecretMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateSecret{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateSecret(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "UpdateSecret",
	}
}
