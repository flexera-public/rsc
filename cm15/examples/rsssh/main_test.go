package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"testing"
)

func TestRSSSHExample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RSSSH Example Suite")
}

var _ = Describe("RSSSH Example", func() {
	var server *ghttp.Server
	var written []byte
	var tmpFile *os.File

	// Override default failure handler so it doesn't kill the process
	fail = func(format string, v ...interface{}) {
		Fail(fmt.Sprintf(format, v...))
	}

	// Override default writter to capture scripts content
	writeFile = func(outputFile string, bytes []byte, perm os.FileMode) {
		written = bytes
	}

	BeforeEach(func() {
		server = ghttp.NewServer()
		var err error
		tmpFile, err = ioutil.TempFile("", "rsc_rsssh_test")
		tmpFile.WriteString(config)
		tmpFile.Close()
		Ω(err).ShouldNot(HaveOccurred())
		os.Args = []string{"rsssh", "-insecure", "-c=" + tmpFile.Name(),
			"-e=dummy", "-p=dummy", "-h=" + server.URL()[7:]}
	})

	AfterEach(func() {
		server.Close()
		Ω(os.Remove(tmpFile.Name())).Should(Succeed())
	})

	It("creates SSH scripts", func() {
		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("POST", "/api/sessions"),
				ghttp.RespondWith(204, ""),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/api/sessions"),
				ghttp.RespondWith(200, ""),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/api/server_arrays"),
				ghttp.RespondWith(200, serverArraysResponseBody),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/api/server_arrays/1/current_instances"),
				ghttp.RespondWith(200, instancesResponseBody),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/api/servers"),
				ghttp.RespondWith(200, serversResponseBody),
			),
		)
		Ω(main).ShouldNot(Panic())
		Ω(string(written)).Should(Equal(output))
	})

})

// Additional handler that can be used to debug requests
func trace(rw http.ResponseWriter, req *http.Request) {
	d, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(d) + "\n")
}

const config = `{
  "SshUser": "rightscale",
  "SshOptions": "-i ~/.id_rsa_rightscale -o StrictHostKeyChecking=no",
  "OutputFile": "/Users/arangamani/.rightscale_aliases",
  "Environments": {
    "dev": {
      "Account": 12345,
      "Servers": {
        "lb": "Loadbalancer"
      },
      "ServerArrays": {
        "app": "Application Server"
      }
    }
  }
}`

const output = `alias dev_app#1='ssh -i ~/.id_rsa_rightscale -o StrictHostKeyChecking=no rightscale@54.158.82.198'
alias dev_app#2='ssh -i ~/.id_rsa_rightscale -o StrictHostKeyChecking=no rightscale@54.205.150.8'
alias dev_app#3='ssh -i ~/.id_rsa_rightscale -o StrictHostKeyChecking=no rightscale@54.162.129.89'
alias dev_lb='ssh -i ~/.id_rsa_rightscale -o StrictHostKeyChecking=no rightscale@54.144.183.128'
`

const serverArraysResponseBody = `[{
	"actions": [
		{"rel": "launch"},
		{"rel": "multi_terminate"},
		{"rel": "multi_run_executable"},
		{"rel": "clone"}
	],
        "array_type": "alert",
        "datacenter_policy": [{
                "datacenter_href": "/api/clouds/1/datacenters/1",
                "max": "0",
                "weight": "50.0"
        }],
        "description": "Application Server",
        "elasticity_params": {
            "alert_specific_params": {
                "decision_threshold": "51",
                "voters_tag_predicate": "app"
            },
            "bounds": {
                "max_count": "1",
                "min_count": "1"
            },
            "pacing": {
                "resize_calm_time": "5",
                "resize_down_by": "1",
                "resize_up_by": "1"
            },
            "schedule_entries": []
        },
        "instances_count": 1,
        "links": [
            {"href": "/api/server_arrays/1", "rel": "self"},
            {"href": "/api/deployments/1", "rel": "deployment"},
            {"href": "/api/server_arrays/1/current_instances", "rel": "current_instances"},
            {"href": "/api/clouds/1/instances/1", "rel": "next_instance"},
            {"href": "/api/server_arrays/1/alert_specs", "rel": "alert_specs"},
            {"href": "/api/server_arrays/1/alerts", "rel": "alerts"}
        ],
        "name": "Application Server",
        "state": "enabled"
}]`

const instancesResponseBody = `[{
        "actions": [{"rel": "terminate"}, {"rel": "reboot"}, {"rel": "run_executable"}, {"rel": "lock"}, {"rel": "unlock"}],
        "associate_public_ip_address": true,
        "cloud_specific_attributes": { "ebs_optimized": false},
        "created_at": "2015/04/27 21:21:17 +0000",
        "ip_forwarding_enabled": false,
        "links": [
            {"href": "/api/clouds/1/instances/1", "rel": "self"},
            {"href": "/api/clouds/1", "rel": "cloud"},
            {"href": "/api/deployments/1", "rel": "deployment"},
            {"href": "/api/server_templates/1", "rel": "server_template"},
            {"href": "/api/multi_cloud_images/1", "rel": "multi_cloud_image"},
            {"href": "/api/server_arrays/1", "rel": "parent"},
            {"href": "/api/clouds/1/instances/1/volume_attachments", "rel": "volume_attachments"},
            {"href": "/api/clouds/1/instances/1/inputs", "rel": "inputs"},
            {"href": "/api/clouds/1/instances/1/monitoring_metrics", "rel": "monitoring_metrics"},
            {"href": "/api/clouds/1/instances/1/alerts", "rel": "alerts"}
        ],
        "locked": false,
        "name": "App #1",
        "pricing_type": "fixed",
        "private_ip_addresses": [ "10.31.237.56" ],
        "public_ip_addresses": [ "54.158.82.198" ],
        "resource_uid": "i-eeb47b38",
        "state": "operational",
        "updated_at": "2015/04/27 21:42:55 +0000"
   },
   {
        "actions": [{"rel": "run_executable"}, {"rel": "lock"}, {"rel": "unlock"}],
        "associate_public_ip_address": true,
        "cloud_specific_attributes": { "ebs_optimized": false},
        "created_at": "2015/04/30 23:52:55 +0000",
        "ip_forwarding_enabled": false,
        "links": [
            {"href": "/api/clouds/1/instances/2", "rel": "self"},
            {"href": "/api/clouds/1", "rel": "cloud"},
            {"href": "/api/deployments/2", "rel": "deployment"},
            {"href": "/api/server_templates/2", "rel": "server_template"},
            {"href": "/api/multi_cloud_images/2", "rel": "multi_cloud_image"},
            {"href": "/api/server_arrays/2", "rel": "parent"},
            {"href": "/api/clouds/1/instances/2/volume_attachments", "rel": "volume_attachments"},
            {"href": "/api/clouds/1/instances/2/inputs", "rel": "inputs"},
            {"href": "/api/clouds/1/instances/2/monitoring_metrics", "rel": "monitoring_metrics"},
            {"href": "/api/clouds/1/instances/2/alerts", "rel": "alerts"}
        ],
        "locked": false,
        "name": "App #2",
        "pricing_type": "fixed",
        "private_ip_addresses": [ "10.31.138.200" ],
        "public_ip_addresses": [ "54.205.150.8" ],
        "resource_uid": "i-bb5efa6d",
        "state": "terminated",
        "updated_at": "2015/04/30 23:56:48 +0000"
   },
   {
        "actions": [ { "rel": "terminate"}, { "rel": "run_executable"}, { "rel": "lock"}, { "rel": "unlock" } ],
        "associate_public_ip_address": true,
        "cloud_specific_attributes": { "automatic_instance_store_mapping": false, "ebs_optimized": false},
	"created_at": "2015/04/30 23:56:57 +0000",
        "ip_forwarding_enabled": false,
        "links": [
            {"href": "/api/clouds/1/instances/3", "rel": "self"},
            {"href": "/api/clouds/1", "rel": "cloud"},
            {"href": "/api/deployments/3", "rel": "deployment"},
            {"href": "/api/server_templates/3", "rel": "server_template"},
            {"href": "/api/multi_cloud_images/3", "rel": "multi_cloud_image"},
            {"href": "/api/server_arrays/3", "rel": "parent"},
            {"href": "/api/clouds/1/instances/3/volume_attachments", "rel": "volume_attachments"},
            {"href": "/api/clouds/1/instances/3/inputs", "rel": "inputs"},
            {"href": "/api/clouds/1/instances/3/monitoring_metrics", "rel": "monitoring_metrics"},
            {"href": "/api/clouds/1/instances/3/alerts", "rel": "alerts" }
        ],
        "locked": false,
        "name": "App #3",
        "pricing_type": "fixed",
        "private_ip_addresses": [ "10.37.135.162" ],
        "public_ip_addresses": [ "54.162.129.89" ],
        "resource_uid": "i-485441b4",
        "state": "terminating",
        "updated_at": "2015/05/01 00:11:05 +0000"
}]`

const serversResponseBody = `[{
        "actions": [ { "rel": "terminate"}, { "rel": "clone" } ],
        "created_at": "2015/04/27 21:17:24 +0000",
        "current_instance": {
            "actions": [ { "rel": "terminate"}, { "rel": "reboot"}, { "rel": "run_executable"}, { "rel": "lock"}, { "rel": "unlock" } ],
            "associate_public_ip_address": true,
            "cloud_specific_attributes": { "ebs_optimized": false},
            "created_at": "2015/04/27 21:18:22 +0000",
            "ip_forwarding_enabled": false,
            "links": [
                {"href": "/api/clouds/1/instances/1", "rel": "self"},
                {"href": "/api/clouds/1", "rel": "cloud"},
                {"href": "/api/deployments/1", "rel": "deployment"},
                {"href": "/api/server_templates/1", "rel": "server_template"},
                {"href": "/api/multi_cloud_images/1", "rel": "multi_cloud_image"},
                {"href": "/api/servers/1100716003", "rel": "parent"},
                {"href": "/api/clouds/1/instances/1/volume_attachments", "rel": "volume_attachments"},
                {"href": "/api/clouds/1/instances/1/inputs", "rel": "inputs"},
                {"href": "/api/clouds/1/instances/1/monitoring_metrics", "rel": "monitoring_metrics"},
                {"href": "/api/clouds/1/instances/1/alerts", "rel": "alerts" }
            ],
            "locked": false,
            "name": "Server",
            "pricing_type": "fixed",
            "private_ip_addresses": [ "10.122.110.41" ],
            "public_ip_addresses": [ "54.144.183.128" ],
            "resource_uid": "i-c5d01b2d",
            "state": "operational",
            "updated_at": "2015/04/27 21:33:16 +0000"
       },
        "description": "Test servers",
        "links": [
            {"href": "/api/servers/2", "rel": "self"},
            {"href": "/api/deployments/2", "rel": "deployment"},
            {"href": "/api/clouds/1/instances/2", "rel": "current_instance"},
            {"href": "/api/clouds/1/instances/2", "rel": "next_instance"},
            {"href": "/api/servers/2/alert_specs", "rel": "alert_specs"},
            {"href": "/api/servers/2/alerts", "rel": "alerts" }
        ],
        "name": "Loadbalancer",
        "next_instance": {
            "actions": [],
            "associate_public_ip_address": true,
            "cloud_specific_attributes": {},
            "created_at": "2015/04/27 21:17:24 +0000",
            "ip_forwarding_enabled": false,
            "links": [
                {"href": "/api/clouds/1/instances/3", "rel": "self"},
                {"href": "/api/clouds/1", "rel": "cloud"},
                {"href": "/api/deployments/3", "rel": "deployment"},
                {"href": "/api/server_templates/3", "rel": "server_template"},
                {"href": "/api/multi_cloud_images/3", "inherited_source": "server_template", "rel": "multi_cloud_image"},
		{"href": "/api/servers/3", "rel": "parent"},
		{"href": "/api/clouds/1/instances/3/volume_attachments", "rel": "volume_attachments"},
		{"href": "/api/clouds/1/instances/3/inputs", "rel": "inputs"},
		{"href": "/api/clouds/1/instances/3/monitoring_metrics", "rel": "monitoring_metrics"},
		{"href": "/api/clouds/1/instances/3/alerts", "rel": "alerts" }
            ],
            "locked": false,
            "name": "Server 2",
            "pricing_type": "fixed",
            "private_ip_addresses": [],
            "public_ip_addresses": [],
            "resource_uid": "ccd7a614-ed22-11e4-b1bc-22000b048d88",
            "state": "inactive",
            "updated_at": "2015/04/27 21:17:25 +0000"
       },
        "state": "operational",
        "updated_at": "2015/04/27 21:18:23 +0000"
}]`
