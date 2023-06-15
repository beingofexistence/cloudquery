// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

//go:generate mockgen -package=mocks -destination=../mocks/sqs.go -source=sqs.go SqsClient
type SqsClient interface {
	GetQueueAttributes(context.Context, *sqs.GetQueueAttributesInput, ...func(*sqs.Options)) (*sqs.GetQueueAttributesOutput, error)
	GetQueueUrl(context.Context, *sqs.GetQueueUrlInput, ...func(*sqs.Options)) (*sqs.GetQueueUrlOutput, error)
	ListDeadLetterSourceQueues(context.Context, *sqs.ListDeadLetterSourceQueuesInput, ...func(*sqs.Options)) (*sqs.ListDeadLetterSourceQueuesOutput, error)
	ListMessageMoveTasks(context.Context, *sqs.ListMessageMoveTasksInput, ...func(*sqs.Options)) (*sqs.ListMessageMoveTasksOutput, error)
	ListQueueTags(context.Context, *sqs.ListQueueTagsInput, ...func(*sqs.Options)) (*sqs.ListQueueTagsOutput, error)
	ListQueues(context.Context, *sqs.ListQueuesInput, ...func(*sqs.Options)) (*sqs.ListQueuesOutput, error)
}
