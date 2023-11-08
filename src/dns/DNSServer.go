package dns

import (
	"fmt"

	"github.com/miekg/dns"
)

func ListenAndServeDNS(port int16) {
	server := &dns.Server{Addr: fmt.Sprintf(":%d", port), Net: "udp"}
	// dns.HandleFunc(hostedDomain.String(), func(w dns.ResponseWriter, m *dns.Msg) {
	// 	//This is an internal domain name we need to handle this by looking to see if it is in the hosts
	// })
	// dns.HandleFunc(".", func(w dns.ResponseWriter, m *dns.Msg) {
	// 	//This is an external domain we need to export it
	// })
	server.ListenAndServe()
}
