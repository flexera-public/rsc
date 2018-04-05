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
				ghttp.RespondWith(204, ""),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/api/catalog/new_session"),
				ghttp.RespondWith(303, "", http.Header(map[string][]string{"Location": []string{"foo"}})),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/api/catalog/accounts/42/user_preferences"),
				ghttp.RespondWith(200, ""),
			),
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/api/manager/projects/42/executions"),
				ghttp.RespondWith(200, responseBody),
			),
		)
		main()
		Ω(out.String()).Should(MatchRegexp(output))
	})

})

// Additional handler that can be used to debug requests
func trace(rw http.ResponseWriter, req *http.Request) {
	d, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(d) + "\n")
}

const output = `Name                                                 State      By                                   Link
-----                                                -----      -----                                -----
RM_SS_API_testNotification- Self Launch 1 1428525219 terminated qa_test\+ss_admin_user@rightscale.com https://127.0.0.1:[0-9]+/api/manager/projects/#/executions/552590c173656c589b6a0e00
RM_SS_API_testNotification- Self Launch 2 1428530238 failed     qa_test\+ss_admin_user@rightscale.com https://127.0.0.1:[0-9]+/api/manager/projects/#/executions/5525a4e073656c567f82c800
exe_3_sch_action_tests__2015-04-29_20:11:36.335      running    resat-premium@rightscale.com         https://127.0.0.1:[0-9]+/api/manager/projects/#/executions/55413b0b73656c4399040000`

const responseBody = `[{
	"kind":"self_service#execution",
	"id":"552590c173656c589b6a0e00",
	"name":"RM_SS_API_testNotification- Self Launch 1 1428525219",
	"href":"/api/manager/projects/62656/executions/552590c173656c589b6a0e00",
	"description":"This is short",
	"status":"terminated",
	"cost":{"value":"0.00","unit":"$","updated_at":"2015-04-08T20:35:38.761+00:00"},
	"deployment":"/api/deployments/557773004",
	"created_by":{"id":82528,"name":"QA_Test SS_Admin_User","email":"qa_test+ss_admin_user@rightscale.com"},
	"timestamps":{
		"created_at":"2015-04-08T20:34:09.535+00:00",
		"launched_at":"2015-04-08T20:34:09.745+00:00",
		"terminated_at":"2015-04-08T21:06:12.400+00:00"
	},
	"links":{
		"running_operations":{"href":"/api/manager/projects/62656/operations?filter[]=execution_id==552590c173656c589b6a0e00\u0026filter[]=status==running"},
		"latest_notifications":{"href":"/api/manager/projects/62656/notifications?filter[]=execution_id==552590c173656c589b6a0e00"}
	}
},{
	"kind":"self_service#execution",
	"id":"55413b0b73656c4399040000",
	"name":"exe_3_sch_action_tests__2015-04-29_20:11:36.335",
	"href":"/api/manager/projects/62656/executions/55413b0b73656c4399040000",
	"status":"running",
	"cost":{"value":"0.00","unit":"$","updated_at":"2015-04-30T16:40:27.471+00:00"},
	"deployment":"/api/deployments/565645004",
	"created_by":{"id":10662,"name":"QA-Nightly ","email":"resat-premium@rightscale.com"},
	"timestamps":{
		"created_at":"2015-04-29T20:11:55.235+00:00",
		"launched_at":"2015-04-29T20:11:55.347+00:00",
		"terminated_at":null
	},
	"links":{
		"running_operations":{"href":"/api/manager/projects/62656/operations?filter[]=execution_id==55413b0b73656c4399040000\u0026filter[]=status==running"},
		"latest_notifications":{"href":"/api/manager/projects/62656/notifications?filter[]=execution_id==55413b0b73656c4399040000"}
	}
},{
	"kind":"self_service#execution",
	"id":"5525a4e073656c567f82c800",
	"name":"RM_SS_API_testNotification- Self Launch 2 1428530238",
	"href":"/api/manager/projects/62656/executions/5525a4e073656c567f82c800",
	"description":"Test for 269.",
	"status":"failed",
	"cost":{"value":"0.00","unit":"$","updated_at":"2015-04-30T15:59:57.806+00:00"},
	"deployment":"/api/deployments/557828004",
	"created_by":{"id":82528,"name":"QA_Test SS_Admin_User","email":"qa_test+ss_admin_user@rightscale.com"},
	"timestamps":{
		"created_at":"2015-04-08T22:00:00.518+00:00",
		"launched_at":"2015-04-08T22:00:00.796+00:00",
		"terminated_at":null
	},
	"links":{
		"running_operations":{"href":"/api/manager/projects/62656/operations?filter[]=execution_id==5525a4e073656c567f82c800\u0026filter[]=status==running"},
		"latest_notifications":{"href":"/api/manager/projects/62656/notifications?filter[]=execution_id==5525a4e073656c567f82c800"}
	}
}]`
