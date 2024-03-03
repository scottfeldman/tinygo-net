package net

import (
	"context"
	"fmt"
)

func (sd *sysDialer) dialMPTCP(ctx context.Context, laddr, raddr *TCPAddr) (*TCPConn, error) {
	return nil, fmt.Errorf("MPTCP not supported")
}

func (sl *sysListener) listenMPTCP(ctx context.Context, laddr *TCPAddr) (*TCPListener, error) {
	return nil, fmt.Errorf("MPTCP not supported")
}
