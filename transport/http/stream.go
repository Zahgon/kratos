package http

import (
	"bufio"
	"context"
	stdhttp "net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-kratos/kratos/v3/encoding"
)

const (
	sseContentType = "text/event-stream"

	websocketControlPrefix = "\x1e"
	websocketControlEnd    = websocketControlPrefix + "end"
	websocketControlError  = websocketControlPrefix + "error:"
)

type streamMode int

const (
	streamModeSSE streamMode = iota + 1
	streamModeWebSocket
)

type ServerStream interface {
	grpc.ServerStream
	Send(any) error
	Recv(any) error
	SendAndClose(any) error
	Close(error) error
	SetContext(context.Context)
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}

type ClientStream interface {
	grpc.ClientStream
	Send(any) error
	Recv(any) error
	CloseAndRecv(any) error
}

type serverStream struct {
	ctx       context.Context
	req       *stdhttp.Request
	res       stdhttp.ResponseWriter
	mode      streamMode
	conn      *websocket.Conn
	header    metadata.MD
	trailer   metadata.MD
	encoder   encoding.Codec
	decoder   encoding.Codec
	started   bool
	writeMu   sync.Mutex
	upgrader  websocket.Upgrader
	bodyField string
}

type ServerStreamOption func(*serverStream)

func WithStreamBodyField(name string) ServerStreamOption {
	_ = "STUB: not implemented"
	return *new(ServerStreamOption)
}

func NewServerSentEventServerStream(ctx Context) ServerStream {
	_ = "STUB: not implemented"
	return *new(ServerStream)
}

func NewWebSocketServerStream(ctx Context, opts ...ServerStreamOption) (ServerStream, error) {
	_ = "STUB: not implemented"
	return *new(ServerStream), nil
}

func (s *serverStream) SetContext(ctx context.Context) { _ = "STUB: not implemented"; return }

func detachStreamContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *serverStream) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) SetHeader(md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) SendHeader(md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) SetTrailer(md metadata.MD) { _ = "STUB: not implemented"; return }

func (s *serverStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *serverStream) Send(m any) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) Recv(m any) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) recvMessage(m any) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) SendAndClose(m any) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) SendMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) RecvMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) Close(err error) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) startSSE() { _ = "STUB: not implemented"; return }

func (s *serverStream) sendSSE(event string, v any) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) writeWebSocketMessage(m any) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) writeWebSocketControl(message string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *serverStream) writeWebSocketClose(code int, text string) error {
	_ = "STUB: not implemented"
	return nil
}

type sseClientStream struct {
	ctx       context.Context
	res       *stdhttp.Response
	scanner   *bufio.Scanner
	decoder   encoding.Codec
	closeOnce sync.Once
	closeErr  error
}

func newSSEClientStream(ctx context.Context, res *stdhttp.Response, decoder encoding.Codec) ClientStream {
	_ = "STUB: not implemented"
	return *new(ClientStream)
}

func (s *sseClientStream) Header() (metadata.MD, error) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), nil
}

func (s *sseClientStream) Trailer() metadata.MD {
	_ = "STUB: not implemented"
	return *new(metadata.MD)
}

func (s *sseClientStream) CloseSend() error { _ = "STUB: not implemented"; return nil }

func (s *sseClientStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *sseClientStream) Send(any) error { _ = "STUB: not implemented"; return nil }

func (s *sseClientStream) Recv(m any) error { _ = "STUB: not implemented"; return nil }

func (s *sseClientStream) CloseAndRecv(any) error { _ = "STUB: not implemented"; return nil }

func (s *sseClientStream) SendMsg(any) error { _ = "STUB: not implemented"; return nil }

func (s *sseClientStream) RecvMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (s *sseClientStream) closeBody() error { _ = "STUB: not implemented"; return nil }

func (s *sseClientStream) readEvent() (string, []byte, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

type websocketClientStream struct {
	ctx        context.Context
	conn       *websocket.Conn
	header     stdhttp.Header
	done       func(error)
	encoder    encoding.Codec
	decoder    encoding.Codec
	mu         sync.Mutex
	sendClosed bool
	closed     bool
	closeOnce  sync.Once
	closeErr   error
	writeMu    sync.Mutex
}

func (s *websocketClientStream) Header() (metadata.MD, error) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), nil
}

func (s *websocketClientStream) Trailer() metadata.MD {
	_ = "STUB: not implemented"
	return *new(metadata.MD)
}

func (s *websocketClientStream) CloseSend() error { _ = "STUB: not implemented"; return nil }

func (s *websocketClientStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *websocketClientStream) Send(m any) error { _ = "STUB: not implemented"; return nil }

func (s *websocketClientStream) Recv(m any) error { _ = "STUB: not implemented"; return nil }

func (s *websocketClientStream) CloseAndRecv(m any) error { _ = "STUB: not implemented"; return nil }

func (s *websocketClientStream) SendMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (s *websocketClientStream) RecvMsg(m any) error { _ = "STUB: not implemented"; return nil }

func (s *websocketClientStream) writeControl(message string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *websocketClientStream) close(err error) error { _ = "STUB: not implemented"; return nil }

func (s *websocketClientStream) checkSendOpen() error { _ = "STUB: not implemented"; return nil }

func (client *Client) ServerSentEvent(ctx context.Context, method, path string, args any, opts ...CallOption) (ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(ClientStream), nil
}

//nolint:bodyclose // newSSEClientStream owns and closes res.Body on success.

func (client *Client) WebSocket(ctx context.Context, path string, opts ...CallOption) (ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(ClientStream), nil
}

func clientStreamFromHandler(v any) (ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(ClientStream), nil
}

func prepareClientRequest(client *Client, req *stdhttp.Request, contentType string, c callInfo) {
	_ = "STUB: not implemented"
	return
}

func marshalStreamMessage(v any, codec encoding.Codec) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unmarshalStreamMessage(data []byte, v any, codec encoding.Codec) error {
	_ = "STUB: not implemented"
	return nil
}

func readWebSocketMessage(conn *websocket.Conn, m any, codec encoding.Codec) error {
	_ = "STUB: not implemented"
	return nil
}

func streamCodecFromCallInfo(c callInfo, names ...string) encoding.Codec {
	_ = "STUB: not implemented"
	return *new(encoding.Codec)
}

func streamCodecFromHeaders(header stdhttp.Header, names ...string) encoding.Codec {
	_ = "STUB: not implemented"
	return *new(encoding.Codec)
}

func defaultStreamCodec() encoding.Codec { _ = "STUB: not implemented"; return *new(encoding.Codec) }

func copyMetadataToHeader(h stdhttp.Header, md metadata.MD) { _ = "STUB: not implemented"; return }

func metadataFromHeader(h stdhttp.Header) metadata.MD {
	_ = "STUB: not implemented"
	return *new(metadata.MD)
}
