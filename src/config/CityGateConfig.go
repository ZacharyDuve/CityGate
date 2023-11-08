package config

type CityGateConfig interface {
	DNSServer() DNSConfig
	HTTPServer() HTTPConfig
}

type DNSConfig interface {
	Port() int16
}

type HTTPConfig interface {
	Port() int16
}
