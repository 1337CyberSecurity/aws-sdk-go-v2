// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Validates the integrity of an assessment report in Audit Manager.
func (c *Client) ValidateAssessmentReportIntegrity(ctx context.Context, params *ValidateAssessmentReportIntegrityInput, optFns ...func(*Options)) (*ValidateAssessmentReportIntegrityOutput, error) {
	if params == nil {
		params = &ValidateAssessmentReportIntegrityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ValidateAssessmentReportIntegrity", params, optFns, c.addOperationValidateAssessmentReportIntegrityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ValidateAssessmentReportIntegrityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ValidateAssessmentReportIntegrityInput struct {

	// The relative path of the Amazon S3 bucket that the assessment report is stored
	// in.
	//
	// This member is required.
	S3RelativePath *string

	noSmithyDocumentSerde
}

type ValidateAssessmentReportIntegrityOutput struct {

	// The signature algorithm that's used to code sign the assessment report file.
	SignatureAlgorithm *string

	// The date and time signature that specifies when the assessment report was
	// created.
	SignatureDateTime *string

	// The unique identifier for the validation signature key.
	SignatureKeyId *string

	// Specifies whether the signature key is valid.
	SignatureValid *bool

	// Represents any errors that occurred when validating the assessment report.
	ValidationErrors []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationValidateAssessmentReportIntegrityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpValidateAssessmentReportIntegrity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpValidateAssessmentReportIntegrity{}, middleware.After)
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
	if err = addOpValidateAssessmentReportIntegrityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opValidateAssessmentReportIntegrity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opValidateAssessmentReportIntegrity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "ValidateAssessmentReportIntegrity",
	}
}
