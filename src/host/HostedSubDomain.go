package host

type HostedSubDomain interface {
	SubDomainName() string
	AddHost(Host) error
	NextHost() Host
}
