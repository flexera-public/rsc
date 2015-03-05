// This file contains the "go generate" directives used to generate the code
// for all API clients. Keep this file up-to-date when adding new APIs.

//go:generate api15gen -metadata=cm15 -output=cm15
//go:generate praxisgen -metadata=cm16/api_docs -output=cm16 -pkg=cm16 -target=1.6 -client=Api
//go:generate praxisgen -metadata=ss/ssd/restful_doc -output=ss/ssd -pkg=ssd -target=1.0 -client=Api
//go:generate praxisgen -metadata=ss/ssc/restful_doc -output=ss/ssc -pkg=ssc -target=1.0 -client=Api
//go:generate praxisgen -metadata=ss/ssm/restful_doc -output=ss/ssm -pkg=ssm -target=1.0 -client=Api
package main
