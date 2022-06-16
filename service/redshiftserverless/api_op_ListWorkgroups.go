// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about a list of specified workgroups.
func (c *Client) ListWorkgroups(ctx context.Context, params *ListWorkgroupsInput, optFns ...func(*Options)) (*ListWorkgroupsOutput, error) {
	if params == nil {
		params = &ListWorkgroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkgroups", params, optFns, c.addOperationListWorkgroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkgroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkgroupsInput struct {

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	MaxResults *int32

	// If your initial ListWorkgroups operation returns a nextToken, you can include
	// the returned nextToken in subsequent ListNamespaces operations, which returns
	// results in the next page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWorkgroupsOutput struct {

	// The returned array of workgroups.
	//
	// This member is required.
	Workgroups []types.Workgroup

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. To retrieve the next page,
	// make the call again using the returned token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkgroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListWorkgroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWorkgroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkgroups(options.Region), middleware.Before); err != nil {
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

// ListWorkgroupsAPIClient is a client that implements the ListWorkgroups
// operation.
type ListWorkgroupsAPIClient interface {
	ListWorkgroups(context.Context, *ListWorkgroupsInput, ...func(*Options)) (*ListWorkgroupsOutput, error)
}

var _ ListWorkgroupsAPIClient = (*Client)(nil)

// ListWorkgroupsPaginatorOptions is the paginator options for ListWorkgroups
type ListWorkgroupsPaginatorOptions struct {
	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkgroupsPaginator is a paginator for ListWorkgroups
type ListWorkgroupsPaginator struct {
	options   ListWorkgroupsPaginatorOptions
	client    ListWorkgroupsAPIClient
	params    *ListWorkgroupsInput
	nextToken *string
	firstPage bool
}

// NewListWorkgroupsPaginator returns a new ListWorkgroupsPaginator
func NewListWorkgroupsPaginator(client ListWorkgroupsAPIClient, params *ListWorkgroupsInput, optFns ...func(*ListWorkgroupsPaginatorOptions)) *ListWorkgroupsPaginator {
	if params == nil {
		params = &ListWorkgroupsInput{}
	}

	options := ListWorkgroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkgroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkgroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkgroups page.
func (p *ListWorkgroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkgroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListWorkgroups(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListWorkgroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-serverless",
		OperationName: "ListWorkgroups",
	}
}
