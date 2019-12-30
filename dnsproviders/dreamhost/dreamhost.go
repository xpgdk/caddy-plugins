// Package dnsimple adapts the lego DNS Made Easy DNS provider
// for Caddy. Importing this package plugs it in.
package dreamhost

import (
	"errors"

	"github.com/caddyserver/caddy/caddytls"
	"github.com/xenolf/lego/providers/dns/dreamhost"
)

func init() {
	caddytls.RegisterDNSProvider("dreamhost", NewDNSProvider)
}

// NewDNSProvider returns a new DNS Made Easy DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = API key
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return dreamhost.NewDNSProvider()
	case 1:
                config := dreamhost.NewDefaultConfig()
                config.APIKey = credentials[0]
		return dreamhost.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
