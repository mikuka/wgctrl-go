//go:build freebsd
// +build freebsd

package wgctrl

import (
	"github.com/mikuka/wgctrl-go/internal/wgfreebsd"
	"github.com/mikuka/wgctrl-go/internal/wginternal"
	"github.com/mikuka/wgctrl-go/internal/wguser"
)

// newClients configures wginternal.Clients for FreeBSD systems.
func newClients() ([]wginternal.Client, error) {
	var clients []wginternal.Client

	// FreeBSD has an in-kernel WireGuard implementation. Determine if it is
	// available and make use of it if so.
	kc, ok, err := wgfreebsd.New()
	if err != nil {
		return nil, err
	}
	if ok {
		clients = append(clients, kc)
	}

	uc, err := wguser.New()
	if err != nil {
		return nil, err
	}

	clients = append(clients, uc)
	return clients, nil
}
