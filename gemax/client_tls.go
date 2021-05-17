package gemax

import (
	"crypto/tls"
)

func tlsVerifyDomain(cs *tls.ConnectionState, domain string) (err error) {
	for _, cert := range cs.PeerCertificates {
		err = cert.VerifyHostname(domain)
		if err != nil {
			return err
		}
	}
	return nil
}
