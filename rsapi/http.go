package rsapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rightscale/rsc/metadata"
)

// FileUpload represents payload fields that correspond to multipart file uploads.
type FileUpload struct {
	Name     string    // Multipart part name
	Filename string    // Uploaded filename
	Reader   io.Reader // Backing reader
}

// SourceUpload represents a payload for /api/right_scripts/:id/update_source
// It's needed as this one API call has unique properties -- it's content type is
// text/plain only, and body of the HTTP call is the script body rather than a properly
// formatted form request :(
type SourceUpload struct {
	Reader io.Reader
}

// This will take a hash (or multi-level hash) of APIParams and extract
// out all the multipart file uploads and put them in an array. The recursion
// is needed as CM API 1.5 and SS both use this, but SS has all its params
// on the top level while CM API 1.5 has the resource name on the top level
// with the params being one level down the tree.
func extractUploads(payload *APIParams) (uploads []*FileUpload) {
	for k, v := range *payload {
		if mpart, ok := v.(*FileUpload); ok {
			uploads = append(uploads, mpart)
			delete(*payload, k)
		} else if child_params, ok := v.(APIParams); ok {
			child_uploads := extractUploads(&child_params)
			uploads = append(uploads, child_uploads...)
		}
	}
	return uploads
}

func extractSourceUpload(payload *APIParams) *SourceUpload {
	var sourceUpload *SourceUpload
	for k, v := range *payload {
		if mpart, ok := v.(*SourceUpload); ok {
			sourceUpload = mpart
			delete(*payload, k)
		}
	}
	return sourceUpload
}

// Handle payload params. Each payload param gets its own multipart
// form section with the section name being the variable name and
// section contents being the variable contents. Handle recursion as well.
func writeMultipartParams(w *multipart.Writer, payload APIParams, prefix string) error {
	for k, v := range payload {
		fieldName := k
		if prefix != "" {
			fieldName = fmt.Sprintf("%s[%s]", prefix, k)
		}
		// Add more types as needed. These two cover both CM15 and SS.
		switch v.(type) {
		case string:
			err := w.WriteField(fieldName, v.(string))
			if err != nil {
				return err
			}
		case APIParams:
			err := writeMultipartParams(w, v.(APIParams), fieldName)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("Unknown type for multipart form section %s: %#v", fieldName, v)
		}
	}
	return nil
}

// BuildHTTPRequest creates a http.Request given all its parts.
// If any member of the Payload field is of type io.Reader then the resulting request has a
// multipart body where each member of type io.Reader is mapped to a single part and all other
// members make up the first part.
func (a *API) BuildHTTPRequest(verb, path, version string, params, payload APIParams) (*http.Request, error) {
	u := url.URL{Host: a.Host, Path: path}
	if params != nil {
		var values = u.Query()
		for n, p := range params {
			switch t := p.(type) {
			case string:
				values.Set(n, t)
			case int:
				values.Set(n, strconv.Itoa(t))
			case bool:
				values.Set(n, strconv.FormatBool(t))
			case []string:
				for _, e := range t {
					values.Add(n, e)
				}
			case []int:
				for _, e := range t {
					values.Add(n, strconv.Itoa(e))
				}
			case []bool:
				for _, e := range t {
					values.Add(n, strconv.FormatBool(e))
				}
			case []interface{}:
				for _, e := range t {
					values.Add(n, fmt.Sprintf("%v", e))
				}
			case map[string]string:
				for pn, e := range t {
					values.Add(fmt.Sprintf("%s[%s]", n, pn), e)
				}
			default:
				return nil, fmt.Errorf("Invalid param value <%+v> for %s, value must be a string, an integer, a bool, an array of these types of a map of strings", p, n)
			}
		}
		u.RawQuery = values.Encode()
	}
	var jsonBytes []byte
	var body io.Reader
	var isMultipart bool
	var isSourceUpload bool
	var boundary string
	if payload != nil {
		var fields io.Reader
		uploads := extractUploads(&payload)
		sourceUpload := extractSourceUpload(&payload)

		if len(payload) > 0 {
			var err error
			if jsonBytes, err = json.Marshal(payload); err != nil {
				return nil, fmt.Errorf("failed to serialize request body: %s", err.Error())
			}
			fields = bytes.NewBuffer(jsonBytes)
		}
		if len(uploads) > 0 {
			isMultipart = true
			var buffer bytes.Buffer
			w := multipart.NewWriter(&buffer)
			if len(payload) > 0 {
				err := writeMultipartParams(w, payload, "")
				if err != nil {
					return nil, fmt.Errorf("failed to write multipart params: %s", err.Error())
				}
			}
			for _, u := range uploads {
				p, err := w.CreateFormFile(u.Name, u.Filename)
				if err != nil {
					return nil, fmt.Errorf("failed to create multipart file: %s", err.Error())
				}
				_, err = io.Copy(p, u.Reader)
				if err != nil {
					return nil, fmt.Errorf("failed to copy multipart file: %s", err.Error())
				}
			}
			if err := w.Close(); err != nil {
				return nil, fmt.Errorf("failed to close multipart body: %s", err.Error())
			}
			boundary = w.Boundary()
			body = &buffer
		} else if sourceUpload != nil {
			isSourceUpload = true
			body = sourceUpload.Reader
		} else {
			body = fields
		}
	}
	var req, err = http.NewRequest(verb, u.String(), body)
	if err != nil {
		return nil, err
	}
	if version != "" {
		req.Header.Set("X-API-Version", version)
	}
	if isMultipart {
		req.Header.Set("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", boundary))
	} else if isSourceUpload {
		req.Header.Set("Content-Type", "text/plain")
	} else {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// PerformRequest logs the request, dumping its content if required then makes the request and logs
// and dumps the corresponding response.
func (a *API) PerformRequest(req *http.Request) (*http.Response, error) {
	// Sign last so auth headers don't get printed or logged
	if a.Auth != nil {
		if err := a.Auth.Sign(req); err != nil {
			return nil, err
		}
	}
	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// PerformRequestWithContext performs everything the PerformRequest does but is also context-aware.
func (a *API) PerformRequestWithContext(ctx context.Context, req *http.Request) (*http.Response, error) {
	// Sign last so auth headers don't get printed or logged
	if a.Auth != nil {
		if err := a.Auth.Sign(req); err != nil {
			return nil, err
		}
	}
	resp, err := a.Client.DoWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// IdentifyParams organizes the given params in two groups: the payload params and the query params.
func IdentifyParams(a *metadata.Action, params APIParams) (payloadParams APIParams, queryParams APIParams) {
	payloadParamNames := a.PayloadParamNames()
	payloadParams = make(APIParams)
	for _, n := range payloadParamNames {
		if p, ok := params[n]; ok {
			payloadParams[n] = p
		}
	}
	queryParamNames := a.QueryParamNames()
	queryParams = make(APIParams)
	for _, n := range queryParamNames {
		if p, ok := params[n]; ok {
			queryParams[n] = p
		}
	}
	return payloadParams, queryParams
}

// LoadResponse deserializes the JSON response into a generic object.
// If the response has a "Location" header then the returned object is a map with one key "Location"
// containing the value of the header.
func (a *API) LoadResponse(resp *http.Response) (interface{}, error) {
	defer resp.Body.Close()
	var respBody interface{}
	jsonResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response (%s)", err)
	}
	if len(jsonResp) > 0 {
		err = json.Unmarshal(jsonResp, &respBody)
		if err != nil {
			return nil, fmt.Errorf("Failed to load response (%s)", err)
		}
	}
	// Special case for "Location" header, assume that if there is a location there is no body
	loc := resp.Header.Get("Location")
	if len(loc) > 0 {
		var bodyMap = make(map[string]interface{})
		bodyMap["Location"] = loc
		respBody = interface{}(bodyMap)
	}
	return respBody, err
}
