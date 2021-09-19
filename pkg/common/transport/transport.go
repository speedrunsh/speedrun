package portal

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"

	"github.com/melbahja/goph"
	qnet "github.com/speedrunsh/grpc-quic"

	"github.com/speedrunsh/speedrun/pkg/common/key"
	"github.com/speedrunsh/speedrun/pkg/common/ssh"
	"google.golang.org/grpc"
)

type Transport struct {
	Conn    *grpc.ClientConn
	Address string
	opts    options
}

type options struct {
	insecure bool
	key      *key.Key
}

type TransportOption interface {
	apply(*options)
}

func defaultOptions() options {
	return options{
		insecure: false,
	}
}

type withInsecure bool

func (w withInsecure) apply(o *options) {
	o.insecure = bool(w)
}

func WithInsecure(enable bool) TransportOption {
	return withInsecure(enable)
}

type withSSH key.Key

func (w withSSH) apply(o *options) {
	a := key.Key(w)
	o.key = &a
}

func WithSSH(key key.Key) TransportOption {
	return withSSH(key)
}

func NewTransport(address string, opts ...TransportOption) (*grpc.ClientConn, error) {
	var err error

	t := &Transport{
		Address: address,
		opts:    defaultOptions(),
	}
	for _, opt := range opts {
		opt.apply(&t.opts)
	}

	if t.opts.key != nil {
		if t.opts.insecure {
			t.Conn, err = ssh2TransportInsecure(address, t.opts.key)
		} else {
			t.Conn, err = ssh2Transport(address, t.opts.key)
		}
		if err != nil {
			return nil, err
		}
	} else {
		t.Conn, err = http2TransportInsecure(address)
		if err != nil {
			return nil, err
		}
	}

	return t.Conn, nil
}

func ssh2TransportInsecure(address string, key *key.Key) (*grpc.ClientConn, error) {
	var sshclient *goph.Client

	sshclient, err := ssh.ConnectInsecure(address, key)
	if err != nil {
		return nil, err
	}

	dialer := func(ctx context.Context, addr string) (net.Conn, error) {
		return sshclient.Dial("tcp", "127.0.0.1:1337")
	}

	return grpc.Dial("127.0.0.1:1337", grpc.WithInsecure(), grpc.WithContextDialer(dialer))
}

func ssh2Transport(address string, key *key.Key) (*grpc.ClientConn, error) {
	var sshclient *goph.Client

	sshclient, err := ssh.Connect(address, key)
	if err != nil {
		return nil, err
	}

	dialer := func(ctx context.Context, addr string) (net.Conn, error) {
		return sshclient.Dial("tcp", "127.0.0.1:1337")
	}

	return grpc.Dial("127.0.0.1:1337", grpc.WithInsecure(), grpc.WithContextDialer(dialer))
}

func http2TransportInsecure(address string) (*grpc.ClientConn, error) {
	target := fmt.Sprintf("%s:%d", address, 1337)
	return grpc.Dial(target, grpc.WithInsecure())
}

func QUICTransport(address string) (*grpc.ClientConn, error) {
	tlsConf := &tls.Config{
		MinVersion:         tls.VersionTLS13,
		CurvePreferences:   []tls.CurveID{tls.X25519},
		CipherSuites:       []uint16{tls.TLS_CHACHA20_POLY1305_SHA256},
		InsecureSkipVerify: true,
		NextProtos:         []string{"speedrun"},
	}

	creds := qnet.NewCredentials(tlsConf)

	dialer := qnet.NewQuicDialer(tlsConf)
	grpcOpts := []grpc.DialOption{
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(creds),
	}

	target := fmt.Sprintf("%s:%d", address, 1337)
	return grpc.Dial(target, grpcOpts...)
}
