package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"testing"
)

func TestBasicExample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Basic Example Suite")
}

var _ = Describe("Basic Example", func() {
	var server *ghttp.Server
	var out bytes.Buffer

	// Override default failure handler so it doesn't kill the process
	fail = func(format string, v ...interface{}) {
		Fail(fmt.Sprintf(format, v...))
	}

	BeforeEach(func() {
		server = ghttp.NewServer()
		os.Args = []string{"basic", "-insecure", "-e=dummy", "-p=dummy", "-a=42", "-h=" + server.URL()[7:]}
		osStdout = &out
	})

	AfterEach(func() {
		server.Close()
	})

	It("lists executions", func() {
		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("POST", "/api/sessions"),
				ghttp.VerifyJSONRepresenting(map[string]interface{}{
					"email":        "dummy",
					"password":     "dummy",
					"account_href": "/api/accounts/42",
				}),
				ghttp.RespondWith(204, ""),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/api/sessions"),
				ghttp.RespondWith(200, ""),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/api/clouds"),
				ghttp.RespondWith(200, responseBody),
			),
		)
		main()
		Ω(out.String()).Should(Equal(output))
	})

})

// Additional handler that can be used to debug requests
func trace(rw http.ResponseWriter, req *http.Request) {
	d, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(d) + "\n")
}

const output = `
1. EC2 us-east-1
supports_networks: true
supports_route_tables: true
supports_volume_attachments: true
supports_volume_snapshots: true

2. EC2 us-west-1
supports_networks: true
supports_route_tables: true
supports_volume_attachments: true
supports_volume_snapshots: true

3. EC2 eu-west-1
supports_networks: true
supports_route_tables: true
supports_volume_attachments: true
supports_volume_snapshots: true
`

const responseBody = `[
    {
        "capabilities": [
            { "name": "supports_networks", "value": "true" },
            { "name": "supports_route_tables", "value": "true" },
            { "name": "supports_volume_attachments", "value": "true" },
            { "name": "supports_volume_snapshots", "value": "true" }
        ],
        "cloud_type": "amazon",
        "description": "Amazon's US Cloud on the East Coast",
        "display_name": "AWS US-East",
        "links": [
            { "href": "/api/clouds/1", "rel": "self" },
            { "href": "/api/clouds/1/datacenters", "rel": "datacenters" },
            { "href": "/api/clouds/1/instance_types", "rel": "instance_types" },
            { "href": "/api/clouds/1/security_groups", "rel": "security_groups" },
            { "href": "/api/clouds/1/instances", "rel": "instances" },
            { "href": "/api/clouds/1/ssh_keys", "rel": "ssh_keys" },
            { "href": "/api/clouds/1/images", "rel": "images" },
            { "href": "/api/clouds/1/ip_addresses", "rel": "ip_addresses" },
            { "href": "/api/clouds/1/ip_address_bindings", "rel": "ip_address_bindings" },
            { "href": "/api/clouds/1/volume_attachments", "rel": "volume_attachments" },
            { "href": "/api/clouds/1/recurring_volume_attachments", "rel": "recurring_volume_attachments" },
            { "href": "/api/clouds/1/volume_snapshots", "rel": "volume_snapshots" },
            { "href": "/api/clouds/1/volume_types", "rel": "volume_types" },
            { "href": "/api/clouds/1/volumes", "rel": "volumes" },
            { "href": "/api/clouds/1/subnets", "rel": "subnets" }
        ],
        "name": "EC2 us-east-1"
    },
    {
        "capabilities": [
            { "name": "supports_networks", "value": "true" },
            { "name": "supports_route_tables", "value": "true" },
            { "name": "supports_volume_attachments", "value": "true" },
            { "name": "supports_volume_snapshots", "value": "true" }
        ],
        "cloud_type": "amazon",
        "description": "Amazon's US Cloud on the West Coast",
        "display_name": "AWS US-West",
        "links": [
            { "href": "/api/clouds/3", "rel": "self" },
            { "href": "/api/clouds/3/datacenters", "rel": "datacenters" },
            { "href": "/api/clouds/3/instance_types", "rel": "instance_types" },
            { "href": "/api/clouds/3/security_groups", "rel": "security_groups" },
            { "href": "/api/clouds/3/instances", "rel": "instances" },
            { "href": "/api/clouds/3/ssh_keys", "rel": "ssh_keys" },
            { "href": "/api/clouds/3/images", "rel": "images" },
            { "href": "/api/clouds/3/ip_addresses", "rel": "ip_addresses" },
            { "href": "/api/clouds/3/ip_address_bindings", "rel": "ip_address_bindings" },
            { "href": "/api/clouds/3/volume_attachments", "rel": "volume_attachments" },
            { "href": "/api/clouds/3/recurring_volume_attachments", "rel": "recurring_volume_attachments" },
            { "href": "/api/clouds/3/volume_snapshots", "rel": "volume_snapshots" },
            { "href": "/api/clouds/3/volume_types", "rel": "volume_types" },
            { "href": "/api/clouds/3/volumes", "rel": "volumes" },
            { "href": "/api/clouds/3/subnets", "rel": "subnets" }
        ],
        "name": "EC2 us-west-1"
    },
    {
        "capabilities": [
            { "name": "supports_networks", "value": "true" },
            { "name": "supports_route_tables", "value": "true" },
            { "name": "supports_volume_attachments", "value": "true" },
            { "name": "supports_volume_snapshots", "value": "true" }
        ],
        "cloud_type": "amazon",
        "description": "Amazon's Europe cloud",
        "display_name": "AWS EU-Ireland",
        "links": [
            { "href": "/api/clouds/2", "rel": "self" },
            { "href": "/api/clouds/2/datacenters", "rel": "datacenters" },
            { "href": "/api/clouds/2/instance_types", "rel": "instance_types" },
            { "href": "/api/clouds/2/security_groups", "rel": "security_groups" },
            { "href": "/api/clouds/2/instances", "rel": "instances" },
            { "href": "/api/clouds/2/ssh_keys", "rel": "ssh_keys" },
            { "href": "/api/clouds/2/images", "rel": "images" },
            { "href": "/api/clouds/2/ip_addresses", "rel": "ip_addresses" },
            { "href": "/api/clouds/2/ip_address_bindings", "rel": "ip_address_bindings" },
            { "href": "/api/clouds/2/volume_attachments", "rel": "volume_attachments" },
            { "href": "/api/clouds/2/recurring_volume_attachments", "rel": "recurring_volume_attachments" },
            { "href": "/api/clouds/2/volume_snapshots", "rel": "volume_snapshots" },
            { "href": "/api/clouds/2/volume_types", "rel": "volume_types" },
            { "href": "/api/clouds/2/volumes", "rel": "volumes" },
            { "href": "/api/clouds/2/subnets", "rel": "subnets" }
        ],
        "name": "EC2 eu-west-1"
    }
]`
