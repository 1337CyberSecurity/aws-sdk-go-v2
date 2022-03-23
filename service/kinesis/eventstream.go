// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream/eventstreamapi"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithysync "github.com/aws/smithy-go/sync"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"io/ioutil"
	"sync"
)

// SubscribeToShardEventStreamReader provides the interface for reading events from
// a stream.
//
// The writer's Close method must allow multiple concurrent calls.
type SubscribeToShardEventStreamReader interface {
	Events() <-chan types.SubscribeToShardEventStream
	Close() error
	Err() error
}

type subscribeToShardEventStreamReadEvent interface {
	isSubscribeToShardEventStreamReadEvent()
}

type subscribeToShardEventStreamReadEventMessage struct {
	Value types.SubscribeToShardEventStream
}

func (*subscribeToShardEventStreamReadEventMessage) isSubscribeToShardEventStreamReadEvent() {}

type subscribeToShardEventStreamReadEventInitialResponse struct {
	Value interface{}
}

func (*subscribeToShardEventStreamReadEventInitialResponse) isSubscribeToShardEventStreamReadEvent() {
}

type subscribeToShardEventStreamReader struct {
	stream                      chan types.SubscribeToShardEventStream
	decoder                     *eventstream.Decoder
	eventStream                 io.ReadCloser
	err                         *smithysync.OnceErr
	payloadBuf                  []byte
	done                        chan struct{}
	closeOnce                   sync.Once
	initialResponseDeserializer func(*eventstream.Message) (interface{}, error)
	initialResponse             chan interface{}
}

func newSubscribeToShardEventStreamWriter(readCloser io.ReadCloser, decoder *eventstream.Decoder, ird func(*eventstream.Message) (interface{}, error)) *subscribeToShardEventStreamReader {
	w := &subscribeToShardEventStreamReader{
		stream:                      make(chan types.SubscribeToShardEventStream),
		decoder:                     decoder,
		eventStream:                 readCloser,
		err:                         smithysync.NewOnceErr(),
		done:                        make(chan struct{}),
		payloadBuf:                  make([]byte, 10*1024),
		initialResponseDeserializer: ird,
		initialResponse:             make(chan interface{}, 1),
	}

	go w.readEventStream()

	return w
}

func (r *subscribeToShardEventStreamReader) Events() <-chan types.SubscribeToShardEventStream {
	return r.stream
}

func (r *subscribeToShardEventStreamReader) readEventStream() {
	defer r.Close()
	defer close(r.stream)

	defer close(r.initialResponse)

	for {
		r.payloadBuf = r.payloadBuf[0:0]
		decodedMessage, err := r.decoder.Decode(r.eventStream, r.payloadBuf)
		if err != nil {
			if err == io.EOF {
				return
			}
			select {
			case <-r.done:
				return
			default:
				r.err.SetError(err)
				return
			}
		}

		event, err := r.deserializeEventMessage(&decodedMessage)
		if err != nil {
			r.err.SetError(err)
			return
		}

		switch ev := event.(type) {
		case *subscribeToShardEventStreamReadEventInitialResponse:
			select {
			case r.initialResponse <- ev.Value:
			case <-r.done:
				return
			default:
			}
		case *subscribeToShardEventStreamReadEventMessage:
			select {
			case r.stream <- ev.Value:
			case <-r.done:
				return
			}
		default:
			r.err.SetError(fmt.Errorf("unexpected event wrapper: %T", event))
			return
		}

	}
}

func (r *subscribeToShardEventStreamReader) deserializeEventMessage(msg *eventstream.Message) (subscribeToShardEventStreamReadEvent, error) {
	messageType := msg.Headers.Get(eventstreamapi.MessageTypeHeader)
	if messageType == nil {
		return nil, fmt.Errorf("%s event header not present", eventstreamapi.MessageTypeHeader)
	}

	switch messageType.String() {
	case eventstreamapi.EventMessageType:
		eventType := msg.Headers.Get(eventstreamapi.EventTypeHeader)
		if eventType == nil {
			return nil, fmt.Errorf("%s event header not present", eventstreamapi.EventTypeHeader)
		}

		if eventType.String() == "initial-response" {
			v, err := r.initialResponseDeserializer(msg)
			if err != nil {
				return nil, err
			}
			return &subscribeToShardEventStreamReadEventInitialResponse{Value: v}, nil
		}

		var v types.SubscribeToShardEventStream
		if err := awsAwsjson11_deserializeEventStreamSubscribeToShardEventStream(&v, msg); err != nil {
			return nil, err
		}
		return &subscribeToShardEventStreamReadEventMessage{Value: v}, nil

	case eventstreamapi.ExceptionMessageType:
		return nil, awsAwsjson11_deserializeEventStreamExceptionSubscribeToShardEventStream(msg)

	case eventstreamapi.ErrorMessageType:
		errorCode := "UnknownError"
		errorMessage := errorCode
		if header := msg.Headers.Get(eventstreamapi.ErrorCodeHeader); header != nil {
			errorCode = header.String()
		}
		if header := msg.Headers.Get(eventstreamapi.ErrorMessageHeader); header != nil {
			errorMessage = header.String()
		}
		return nil, &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}

	default:
		mc := msg.Clone()
		return nil, &UnknownEventMessageError{
			Type:    messageType.String(),
			Message: &mc,
		}

	}
}

func (r *subscribeToShardEventStreamReader) ErrorSet() <-chan struct{} {
	return r.err.ErrorSet()
}

func (r *subscribeToShardEventStreamReader) Close() error {
	r.closeOnce.Do(r.safeClose)
	return r.Err()
}

func (r *subscribeToShardEventStreamReader) safeClose() {
	close(r.done)
	r.eventStream.Close()

}

func (r *subscribeToShardEventStreamReader) Err() error {
	return r.err.Err()
}

func (r *subscribeToShardEventStreamReader) Closed() <-chan struct{} {
	return r.done
}

type awsAwsjson11_deserializeOpEventStreamSubscribeToShard struct {
	LogEventStreamWrites bool
	LogEventStreamReads  bool
}

func (*awsAwsjson11_deserializeOpEventStreamSubscribeToShard) ID() string {
	return "OperationEventStreamDeserializer"
}

func (m *awsAwsjson11_deserializeOpEventStreamSubscribeToShard) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	defer func() {
		if err == nil {
			return
		}
		m.closeResponseBody(out)
	}()

	logger := middleware.GetLogger(ctx)

	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type: %T", in.Request)
	}
	_ = request

	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	deserializeOutput, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type: %T", out.RawResponse)
	}
	_ = deserializeOutput

	output, ok := out.Result.(*SubscribeToShardOutput)
	if out.Result != nil && !ok {
		return out, metadata, fmt.Errorf("unexpected output result type: %T", out.Result)
	} else if out.Result == nil {
		output = &SubscribeToShardOutput{}
		out.Result = output
	}

	eventReader := newSubscribeToShardEventStreamWriter(
		deserializeOutput.Body,
		eventstream.NewDecoder(func(options *eventstream.DecoderOptions) {
			options.Logger = logger
			options.LogMessages = m.LogEventStreamReads

		}),
		awsAwsjson11_deserializeEventMessageResponseSubscribeToShardOutput,
	)
	defer func() {
		if err == nil {
			return
		}
		_ = eventReader.Close()
	}()

	ir := <-eventReader.initialResponse
	irv, ok := ir.(*SubscribeToShardOutput)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected output result type: %T", ir)
	}
	*output = *irv

	output.eventStream = NewSubscribeToShardEventStream(func(stream *SubscribeToShardEventStream) {
		stream.Reader = eventReader
	})

	go output.eventStream.waitStreamClose()

	return out, metadata, nil
}

func (*awsAwsjson11_deserializeOpEventStreamSubscribeToShard) closeResponseBody(out middleware.DeserializeOutput) {
	if resp, ok := out.RawResponse.(*smithyhttp.Response); ok && resp != nil && resp.Body != nil {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}
}

func addEventStreamSubscribeToShardMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Deserialize.Insert(&awsAwsjson11_deserializeOpEventStreamSubscribeToShard{
		LogEventStreamWrites: options.ClientLogMode.IsRequestEventMessage(),
		LogEventStreamReads:  options.ClientLogMode.IsResponseEventMessage(),
	}, "OperationDeserializer", middleware.Before)
}

// UnknownEventMessageError provides an error when a message is received from the stream,
// but the reader is unable to determine what kind of message it is.
type UnknownEventMessageError struct {
	Type    string
	Message *eventstream.Message
}

// Error retruns the error message string.
func (e *UnknownEventMessageError) Error() string {
	return "unknown event stream message type, " + e.Type
}

func setSafeEventStreamClientLogMode(o *Options, operation string) {
	switch operation {
	case "SubscribeToShard":
		toggleEventStreamClientLogMode(o, false, true)
		return

	default:
		return

	}
}
func toggleEventStreamClientLogMode(o *Options, request, response bool) {
	mode := o.ClientLogMode

	if request && mode.IsRequestWithBody() {
		mode.ClearRequestWithBody()
		mode |= aws.LogRequest
	}

	if response && mode.IsResponseWithBody() {
		mode.ClearResponseWithBody()
		mode |= aws.LogResponse
	}

	o.ClientLogMode = mode

}
