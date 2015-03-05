// This file contains the "go generate" directives used to generate the code
// for all API clients. Keep this file up-to-date when adding new APIs.

//go:generate api15gen -metadata=rsapi15 -output=rsapi15
//go:generate praxisgen -metadata=rsapi16/api_docs -output=rsapi16 -pkg=rsapi16 -target=1.6 -client=Api16
//go:generate praxisgen -metadata=ss/ssd/restful_doc -output=ss/ssd -pkg=ssd -target=1.0 -client=Api
//go:generate praxisgen -metadata=ss/ssc/restful_doc -output=ss/ssc -pkg=ssc -target=1.0 -client=Api
//go:generate praxisgen -metadata=ss/ssm/restful_doc -output=ss/ssm -pkg=ssm -target=1.0 -client=Api
package main
