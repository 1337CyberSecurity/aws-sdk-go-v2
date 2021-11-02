// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Describes a standalone resource or similarly grouped resources that the
// application is made up of.
type ApplicationComponent struct {

	// The name of the component.
	ComponentName *string

	// If logging is supported for the resource type, indicates whether the component
	// has configured logs to be monitored.
	ComponentRemarks *string

	// Workloads detected in the application component.
	DetectedWorkload map[string]map[string]string

	// Indicates whether the application component is monitored.
	Monitor *bool

	// The operating system of the component.
	OsType OsType

	// The resource type. Supported resource types include EC2 instances, Auto Scaling
	// group, Classic ELB, Application ELB, and SQS Queue.
	ResourceType *string

	// The stack tier of the application component.
	Tier Tier

	noSmithyDocumentSerde
}

// Describes the status of the application.
type ApplicationInfo struct {
	AutoConfigEnabled *bool

	// Indicates whether Application Insights can listen to CloudWatch events for the
	// application resources, such as instance terminated, failed deployment, and
	// others.
	CWEMonitorEnabled *bool

	DiscoveryType DiscoveryType

	// The lifecycle of the application.
	LifeCycle *string

	// Indicates whether Application Insights will create opsItems for any problem
	// detected by Application Insights for an application.
	OpsCenterEnabled *bool

	// The SNS topic provided to Application Insights that is associated to the created
	// opsItems to receive SNS notifications for opsItem updates.
	OpsItemSNSTopicArn *string

	// The issues on the user side that block Application Insights from successfully
	// monitoring an application. Example remarks include:
	//
	// * “Configuring application,
	// detected 1 Errors, 3 Warnings”
	//
	// * “Configuring application, detected 1
	// Unconfigured Components”
	Remarks *string

	// The name of the resource group used for the application.
	ResourceGroupName *string

	noSmithyDocumentSerde
}

// The event information.
type ConfigurationEvent struct {

	// The details of the event in plain text.
	EventDetail *string

	// The name of the resource Application Insights attempted to configure.
	EventResourceName *string

	// The resource type that Application Insights attempted to configure, for example,
	// CLOUDWATCH_ALARM.
	EventResourceType ConfigurationEventResourceType

	// The status of the configuration update event. Possible values include INFO,
	// WARN, and ERROR.
	EventStatus ConfigurationEventStatus

	// The timestamp of the event.
	EventTime *time.Time

	// The resource monitored by Application Insights.
	MonitoredResourceARN *string

	noSmithyDocumentSerde
}

// An object that defines the log patterns that belongs to a LogPatternSet.
type LogPattern struct {

	// A regular expression that defines the log pattern. A log pattern can contain as
	// many as 50 characters, and it cannot be empty. The pattern must be DFA
	// compatible. Patterns that utilize forward lookahead or backreference
	// constructions are not supported.
	Pattern *string

	// The name of the log pattern. A log pattern name can contain as many as 50
	// characters, and it cannot be empty. The characters can be Unicode letters,
	// digits, or one of the following symbols: period, dash, underscore.
	PatternName *string

	// The name of the log pattern. A log pattern name can contain as many as 30
	// characters, and it cannot be empty. The characters can be Unicode letters,
	// digits, or one of the following symbols: period, dash, underscore.
	PatternSetName *string

	// Rank of the log pattern. Must be a value between 1 and 1,000,000. The patterns
	// are sorted by rank, so we recommend that you set your highest priority patterns
	// with the lowest rank. A pattern of rank 1 will be the first to get matched to a
	// log line. A pattern of rank 1,000,000 will be last to get matched. When you
	// configure custom log patterns from the console, a Low severity pattern
	// translates to a 750,000 rank. A Medium severity pattern translates to a 500,000
	// rank. And a High severity pattern translates to a 250,000 rank. Rank values less
	// than 1 or greater than 1,000,000 are reserved for AWS-provided patterns.
	Rank int32

	noSmithyDocumentSerde
}

// Describes an anomaly or error with the application.
type Observation struct {

	// The detail type of the CloudWatch Event-based observation, for example, EC2
	// Instance State-change Notification.
	CloudWatchEventDetailType *string

	// The ID of the CloudWatch Event-based observation related to the detected
	// problem.
	CloudWatchEventId *string

	// The source of the CloudWatch Event.
	CloudWatchEventSource CloudWatchEventSource

	// The CodeDeploy application to which the deployment belongs.
	CodeDeployApplication *string

	// The deployment group to which the CodeDeploy deployment belongs.
	CodeDeployDeploymentGroup *string

	// The deployment ID of the CodeDeploy-based observation related to the detected
	// problem.
	CodeDeployDeploymentId *string

	// The instance group to which the CodeDeploy instance belongs.
	CodeDeployInstanceGroupId *string

	// The status of the CodeDeploy deployment, for example SUCCESS or  FAILURE.
	CodeDeployState *string

	// The cause of an EBS CloudWatch event.
	EbsCause *string

	// The type of EBS CloudWatch event, such as createVolume, deleteVolume or
	// attachVolume.
	EbsEvent *string

	// The request ID of an EBS CloudWatch event.
	EbsRequestId *string

	// The result of an EBS CloudWatch event, such as failed or succeeded.
	EbsResult *string

	// The state of the instance, such as STOPPING or TERMINATING.
	Ec2State *string

	// The time when the observation ended, in epoch seconds.
	EndTime *time.Time

	// The Amazon Resource Name (ARN) of the AWS Health Event-based observation.
	HealthEventArn *string

	// The description of the AWS Health event provided by the service, such as Amazon
	// EC2.
	HealthEventDescription *string

	// The category of the AWS Health event, such as issue.
	HealthEventTypeCategory *string

	// The type of the AWS Health event, for example, AWS_EC2_POWER_CONNECTIVITY_ISSUE.
	HealthEventTypeCode *string

	// The service to which the AWS Health Event belongs, such as EC2.
	HealthService *string

	// The ID of the observation type.
	Id *string

	// The timestamp in the CloudWatch Logs that specifies when the matched line
	// occurred.
	LineTime *time.Time

	// The log filter of the observation.
	LogFilter LogFilter

	// The log group name.
	LogGroup *string

	// The log text of the observation.
	LogText *string

	// The name of the observation metric.
	MetricName *string

	// The namespace of the observation metric.
	MetricNamespace *string

	// The category of an RDS event.
	RdsEventCategories *string

	// The message of an RDS event.
	RdsEventMessage *string

	// The name of the S3 CloudWatch Event-based observation.
	S3EventName *string

	// The source resource ARN of the observation.
	SourceARN *string

	// The source type of the observation.
	SourceType *string

	// The time when the observation was first detected, in epoch seconds.
	StartTime *time.Time

	// The Amazon Resource Name (ARN) of the step function-based observation.
	StatesArn *string

	// The Amazon Resource Name (ARN) of the step function execution-based observation.
	StatesExecutionArn *string

	// The input to the step function-based observation.
	StatesInput *string

	// The status of the step function-related observation.
	StatesStatus *string

	// The unit of the source observation metric.
	Unit *string

	// The value of the source observation metric.
	Value *float64

	// The X-Ray request error percentage for this node.
	XRayErrorPercent *int32

	// The X-Ray request fault percentage for this node.
	XRayFaultPercent *int32

	// The name of the X-Ray node.
	XRayNodeName *string

	// The type of the X-Ray node.
	XRayNodeType *string

	// The X-Ray node request average latency for this node.
	XRayRequestAverageLatency *int64

	// The X-Ray request count for this node.
	XRayRequestCount *int32

	// The X-Ray request throttle percentage for this node.
	XRayThrottlePercent *int32

	noSmithyDocumentSerde
}

// Describes a problem that is detected by correlating observations.
type Problem struct {

	// The resource affected by the problem.
	AffectedResource *string

	// The time when the problem ended, in epoch seconds.
	EndTime *time.Time

	// Feedback provided by the user about the problem.
	Feedback map[string]FeedbackValue

	// The ID of the problem.
	Id *string

	// A detailed analysis of the problem using machine learning.
	Insights *string

	LastRecurrenceTime *time.Time

	RecurringCount *int64

	// The name of the resource group affected by the problem.
	ResourceGroupName *string

	// A measure of the level of impact of the problem.
	SeverityLevel SeverityLevel

	// The time when the problem started, in epoch seconds.
	StartTime *time.Time

	// The status of the problem.
	Status Status

	// The name of the problem.
	Title *string

	noSmithyDocumentSerde
}

// Describes observations related to the problem.
type RelatedObservations struct {

	// The list of observations related to the problem.
	ObservationList []Observation

	noSmithyDocumentSerde
}

// An object that defines the tags associated with an application. A tag is a label
// that you optionally define and associate with an application. Tags can help you
// categorize and manage resources in different ways, such as by purpose, owner,
// environment, or other criteria. Each tag consists of a required tag key and an
// associated tag value, both of which you define. A tag key is a general label
// that acts as a category for a more specific tag value. A tag value acts as a
// descriptor within a tag key. A tag key can contain as many as 128 characters. A
// tag value can contain as many as 256 characters. The characters can be Unicode
// letters, digits, white space, or one of the following symbols: _ . : / = + -.
// The following additional restrictions apply to tags:
//
// * Tag keys and values are
// case sensitive.
//
// * For each associated resource, each tag key must be unique and
// it can have only one value.
//
// * The aws: prefix is reserved for use by AWS; you
// can’t use it in any tag keys or values that you define. In addition, you can't
// edit or remove tag keys or values that use this prefix.
type Tag struct {

	// One part of a key-value pair that defines a tag. The maximum length of a tag key
	// is 128 characters. The minimum length is 1 character.
	//
	// This member is required.
	Key *string

	// The optional part of a key-value pair that defines a tag. The maximum length of
	// a tag value is 256 characters. The minimum length is 0 characters. If you don't
	// want an application to have a specific tag value, don't specify a value for this
	// parameter.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
