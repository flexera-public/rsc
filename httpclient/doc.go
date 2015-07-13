/*
Package httpclient provide a HTTP client factory for all code that need to make API calls.
This makes it possible for auth and actual API client code to share properties such as whether the
connections should be made using HTTP instead of HTTPS (Insecure) or whether server
certificate checking should be bypassed (NoCertCheck).
The HTTP clients created with this package also have built-in support for dumping requests and
responses, the "Debug" dump format produces:

https://us-3.rightscale.com/api/clouds
GET /api/clouds HTTP/1.1

GET https://us-3.rightscale.com/api/clouds
Host: us-3.rightscale.com
User-Agent: rsc/dev-unknown-branch
Content-Type: application/json
X-Api-Version: 1.5
Accept-Encoding: gzip
==> HTTP/1.1 200 OK
Content-Length: 17224
Cache-Control: private, max-age=0, must-revalidate
Connection: keep-alive
Content-Type: application/vnd.rightscale.cloud+json;type=collection;charset=utf-8
Date: Fri, 10 Jul 2015 22:01:35 GMT
Status: 200 OK
Strict-Transport-Security: max-age=31536000; includeSubdomains;
X-Request-Uuid: 3b036a56a5b04b35a94dbaf8240f1016

*/
package httpclient
