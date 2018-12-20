package hc

import (
	"github.com/brutella/dnssd"
	"net"
	"strings"
)

func newService(config *Config) dnssd.Service {
	// 2016-03-14(brutella): Replace whitespaces (" ") from service name
	// with underscores ("_")to fix invalid http host header field value
	// produces by iOS.
	//
	// [Radar] http://openradar.appspot.com/radar?id=4931940373233664
	stripped := strings.Replace(config.name, " ", "_", -1)

	var ips []net.IP
	if ip := net.ParseIP(config.IP); ip != nil {
		ips = append(ips, ip)
	}

	service := dnssd.NewService(stripped, "_hap._tcp.", "local.", "", ips, config.servePort)
	service.Text = config.txtRecords()

	return service
}
