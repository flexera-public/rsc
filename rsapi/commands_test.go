package rsapi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rightscale/rsc/cm15"
	"github.com/rightscale/rsc/rsapi"
	"github.com/rightscale/rsc/ss/ssm"
)

var _ = Describe("normalize", func() {
	var payload rsapi.APIParams
	var name string
	var value interface{}

	var res rsapi.APIParams
	var resErr error

	JustBeforeEach(func() {
		res, resErr = rsapi.Normalize(payload, name, value)
	})

	BeforeEach(func() {
		payload = rsapi.APIParams{}
	})

	Describe("with a simple string", func() {
		BeforeEach(func() {
			name = "val"
			value = "foo"
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			Ω(res).Should(Equal(rsapi.APIParams{"val": "foo"}))
		})
	})

	Describe("with an int", func() {
		BeforeEach(func() {
			name = "val"
			value = 42
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			Ω(res).Should(Equal(rsapi.APIParams{"val": 42}))
		})
	})

	Describe("with a float", func() {
		BeforeEach(func() {
			name = "val"
			value = 42.5
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			Ω(res).Should(Equal(rsapi.APIParams{"val": 42.5}))
		})
	})

	Describe("with a bool", func() {
		BeforeEach(func() {
			name = "val"
			value = true
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			Ω(res).Should(Equal(rsapi.APIParams{"val": true}))
		})
	})

	Describe("with a simple array", func() {
		BeforeEach(func() {
			name = "val[]"
			value = "foo"
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			Ω(res).Should(Equal(rsapi.APIParams{"val": []interface{}{"foo"}}))
		})
	})

	Describe("with a simple map", func() {
		BeforeEach(func() {
			name = "val[a]"
			value = "foo"
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			Ω(res).Should(Equal(rsapi.APIParams{"val": rsapi.APIParams{"a": "foo"}}))
		})
	})

	Describe("with a map of arrays", func() {
		BeforeEach(func() {
			name = "val[a][]"
			value = "foo"
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			expected := rsapi.APIParams{
				"val": rsapi.APIParams{
					"a": []interface{}{"foo"},
				},
			}
			Ω(res).Should(Equal(expected))
		})
	})

	Describe("with a map of arrays with existing values", func() {
		BeforeEach(func() {
			name = "val[a][]"
			value = "foo"
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			expected := rsapi.APIParams{
				"val": rsapi.APIParams{
					"a": []interface{}{"foo"},
				},
			}
			Ω(res).Should(Equal(expected))
			name = "val[a][]"
			value = "bar"
			res, resErr = rsapi.Normalize(res, name, value)
			Ω(resErr).ShouldNot(HaveOccurred())
			expected = rsapi.APIParams{
				"val": rsapi.APIParams{
					"a": []interface{}{"foo", "bar"},
				},
			}
			Ω(res).Should(Equal(expected))
		})
	})

	Describe("with an array of maps", func() {
		BeforeEach(func() {
			name = "val[][a]"
			value = "foo"
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			expected := rsapi.APIParams{
				"val": []interface{}{rsapi.APIParams{"a": "foo"}},
			}
			Ω(res).Should(Equal(expected))
		})
	})

	Describe("with an array of maps of arrays", func() {
		BeforeEach(func() {
			name = "val[][a][]"
			value = "foo"
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			expected := rsapi.APIParams{
				"val": []interface{}{rsapi.APIParams{"a": []interface{}{"foo"}}},
			}
			Ω(res).Should(Equal(expected))
		})
	})

	Describe("with an array of maps with existing keys", func() {
		BeforeEach(func() {
			name = "val[][a]"
			value = "foo"
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			expected := rsapi.APIParams{
				"val": []interface{}{rsapi.APIParams{"a": "foo"}},
			}
			Ω(res).Should(Equal(expected))
			name = "val[][b]"
			value = "bar"
			res, resErr = rsapi.Normalize(res, name, value)
			Ω(resErr).ShouldNot(HaveOccurred())
			expected = rsapi.APIParams{
				"val": []interface{}{rsapi.APIParams{"a": "foo", "b": "bar"}},
			}
			Ω(res).Should(Equal(expected))
		})
	})

	Describe("with an array of maps with existing keys with more than one element", func() {
		BeforeEach(func() {
			name = "val[][b]"
			value = "baz"
			payload = rsapi.APIParams{
				"val": []interface{}{rsapi.APIParams{"a": "foo", "b": "bar"}},
			}
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			expected := rsapi.APIParams{
				"val": []interface{}{
					rsapi.APIParams{"a": "foo", "b": "bar"},
					rsapi.APIParams{"b": "baz"},
				},
			}
			Ω(res).Should(Equal(expected))
		})
	})
})

var _ = Describe("ParseCommand", func() {
	var cmd, hrefPrefix string
	var values rsapi.ActionCommands
	var api *rsapi.API

	var parsed *rsapi.ParsedCommand
	var parseErr error

	BeforeEach(func() {
		values = nil
		ssm := ssm.New("", nil)
		api = ssm.API
	})

	JustBeforeEach(func() {
		parsed, parseErr = api.ParseCommand(cmd, hrefPrefix, values)
	})

	Describe("with array of maps with one element", func() {
		BeforeEach(func() {
			cmd = "run"
			runCmd := rsapi.ActionCommand{
				Href: "/api/manager/projects/42/executions/54",
				Params: []string{
					"name=Tip CWF",
					"configuration_options[][name]=environment_name",
					"configuration_options[][type]=string",
					"configuration_options[][value]=ss2",
				},
			}
			values = rsapi.ActionCommands{"run": &runCmd}
		})
		It("parses", func() {
			Ω(parseErr).ShouldNot(HaveOccurred())
			Ω(parsed).ShouldNot(BeNil())
			payload := rsapi.APIParams{
				"name": "Tip CWF",
				"configuration_options": []interface{}{rsapi.APIParams{
					"name":  "environment_name",
					"type":  "string",
					"value": "ss2",
				}},
			}
			expected := rsapi.ParsedCommand{
				HTTPMethod:    "POST",
				URI:           "/api/manager/projects/42/executions/54/actions/run",
				QueryParams:   rsapi.APIParams{},
				PayloadParams: payload,
			}
			Ω(*parsed).Should(Equal(expected))
		})
	})

	Describe("with array of maps with two elements", func() {
		BeforeEach(func() {
			cmd = "run"
			runCmd := rsapi.ActionCommand{
				Href: "/api/manager/projects/42/executions/54",
				Params: []string{
					"name=Tip CWF2",
					"configuration_options[][name]=environment_name",
					"configuration_options[][type]=string",
					"configuration_options[][value]=ss2",
					"configuration_options[][name]=environment_name2",
					"configuration_options[][type]=string",
					"configuration_options[][value]=ss2",
				},
			}
			values = rsapi.ActionCommands{"run": &runCmd}
		})
		It("parses", func() {
			Ω(parseErr).ShouldNot(HaveOccurred())
			Ω(parsed).ShouldNot(BeNil())
			payload := rsapi.APIParams{
				"name": "Tip CWF2",
				"configuration_options": []interface{}{
					rsapi.APIParams{
						"name":  "environment_name",
						"type":  "string",
						"value": "ss2",
					},
					rsapi.APIParams{
						"name":  "environment_name2",
						"type":  "string",
						"value": "ss2",
					},
				},
			}
			expected := rsapi.ParsedCommand{
				HTTPMethod:    "POST",
				URI:           "/api/manager/projects/42/executions/54/actions/run",
				QueryParams:   rsapi.APIParams{},
				PayloadParams: payload,
			}
			Ω(*parsed).Should(Equal(expected))
		})
	})

	Describe("with an array of query parameters", func() {
		BeforeEach(func() {
			cmd = "index"
			indexCmd := rsapi.ActionCommand{
				Href: "/api/manager/projects/42/executions",
				Params: []string{
					"filter[]=status==running",
					"filter[]=status==stopped",
				},
			}
			values = rsapi.ActionCommands{"index": &indexCmd}
		})
		It("parses", func() {
			Ω(parseErr).ShouldNot(HaveOccurred())
			Ω(parsed).ShouldNot(BeNil())
			query := rsapi.APIParams{
				"filter[]": []interface{}{"status==running", "status==stopped"},
			}
			expected := rsapi.ParsedCommand{
				HTTPMethod:    "GET",
				URI:           "/api/manager/projects/42/executions",
				QueryParams:   query,
				PayloadParams: rsapi.APIParams{},
			}
			Ω(*parsed).Should(Equal(expected))
		})
	})
})

var _ = Describe("ParseCommand with cm15", func() {
	var cmd, hrefPrefix string
	var values rsapi.ActionCommands
	var api *rsapi.API

	var parsed *rsapi.ParsedCommand
	var parseErr error

	BeforeEach(func() {
		values = nil
		cm := cm15.New("", nil)
		api = cm.API
	})

	JustBeforeEach(func() {
		parsed, parseErr = api.ParseCommand(cmd, hrefPrefix, values)
	})

	Describe("with a deep map of inputs", func() {
		BeforeEach(func() {
			cmd = "wrap_instance"
			wrapCmd := rsapi.ActionCommand{
				Href: "/api/servers",
				Params: []string{
					"server[name]=server name",
					"server[deployment_href]=/api/deployments/1",
					"server[instance][href]=/api/clouds/1/instances/42",
					"server[instance][server_template_href]=/api/server_templates/123",
					"server[instance][inputs][STRING_INPUT_1]=text:testing123",
					"server[instance][inputs][STRING_INPUT_2]=text:testing124",
				},
			}
			values = rsapi.ActionCommands{"wrap_instance": &wrapCmd}
		})
		It("parses", func() {
			Ω(parseErr).ShouldNot(HaveOccurred())
			Ω(parsed).ShouldNot(BeNil())
			payload := rsapi.APIParams{
				"server": rsapi.APIParams{
					"name":            "server name",
					"deployment_href": "/api/deployments/1",
					"instance": rsapi.APIParams{
						"href":                 "/api/clouds/1/instances/42",
						"server_template_href": "/api/server_templates/123",
						"inputs": rsapi.APIParams{
							"STRING_INPUT_1": "text:testing123",
							"STRING_INPUT_2": "text:testing124",
						},
					},
				},
			}
			expected := rsapi.ParsedCommand{
				HTTPMethod:    "POST",
				URI:           "/api/servers/wrap_instance",
				QueryParams:   rsapi.APIParams{},
				PayloadParams: payload,
			}
			Ω(*parsed).Should(Equal(expected))
		})
	})

})
