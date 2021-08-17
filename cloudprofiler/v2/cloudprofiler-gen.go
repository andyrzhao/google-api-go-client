// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package cloudprofiler provides access to the Stackdriver Profiler API.
//
// For product documentation, see: https://cloud.google.com/profiler/
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/cloudprofiler/v2"
//   ...
//   ctx := context.Background()
//   cloudprofilerService, err := cloudprofiler.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// By default, all available scopes (see "Constants") are used to authenticate. To restrict scopes, use option.WithScopes:
//
//   cloudprofilerService, err := cloudprofiler.NewService(ctx, option.WithScopes(cloudprofiler.MonitoringWriteScope))
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   cloudprofilerService, err := cloudprofiler.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   cloudprofilerService, err := cloudprofiler.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package cloudprofiler // import "google.golang.org/api/cloudprofiler/v2"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint

const apiId = "cloudprofiler:v2"
const apiName = "cloudprofiler"
const apiVersion = "v2"
const basePath = "https://cloudprofiler.googleapis.com/"
const mtlsBasePath = "https://cloudprofiler.mtls.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// See, edit, configure, and delete your Google Cloud data and see the
	// email address for your Google Account.
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View and write monitoring data for all of your Google and third-party
	// Cloud and API projects
	MonitoringScope = "https://www.googleapis.com/auth/monitoring"

	// Publish metric data to your Google Cloud projects
	MonitoringWriteScope = "https://www.googleapis.com/auth/monitoring.write"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := option.WithScopes(
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/monitoring",
		"https://www.googleapis.com/auth/monitoring.write",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Profiles = NewProjectsProfilesService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Profiles *ProjectsProfilesService
}

func NewProjectsProfilesService(s *Service) *ProjectsProfilesService {
	rs := &ProjectsProfilesService{s: s}
	return rs
}

type ProjectsProfilesService struct {
	s *Service
}

// CreateProfileRequest: CreateProfileRequest describes a profile
// resource online creation request. The deployment field must be
// populated. The profile_type specifies the list of profile types
// supported by the agent. The creation call will hang until a profile
// of one of these types needs to be collected.
type CreateProfileRequest struct {
	// Deployment: Deployment details.
	Deployment *Deployment `json:"deployment,omitempty"`

	// ProfileType: One or more profile types that the agent is capable of
	// providing.
	//
	// Possible values:
	//   "PROFILE_TYPE_UNSPECIFIED" - Unspecified profile type.
	//   "CPU" - Thread CPU time sampling.
	//   "WALL" - Wallclock time sampling. More expensive as stops all
	// threads.
	//   "HEAP" - In-use heap profile. Represents a snapshot of the
	// allocations that are live at the time of the profiling.
	//   "THREADS" - Single-shot collection of all thread stacks.
	//   "CONTENTION" - Synchronization contention profile.
	//   "PEAK_HEAP" - Peak heap profile.
	//   "HEAP_ALLOC" - Heap allocation profile. It represents the
	// aggregation of all allocations made over the duration of the profile.
	// All allocations are included, including those that might have been
	// freed by the end of the profiling interval. The profile is in
	// particular useful for garbage collecting languages to understand
	// which parts of the code create most of the garbage collection
	// pressure to see if those can be optimized.
	ProfileType []string `json:"profileType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Deployment") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Deployment") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CreateProfileRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CreateProfileRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Deployment: Deployment contains the deployment identification
// information.
type Deployment struct {
	// Labels: Labels identify the deployment within the user universe and
	// same target. Validation regex for label names: `^a-z0-9
	// ([a-z0-9-]{0,61}[a-z0-9])?$`. Value for an individual label must be
	// <= 512 bytes, the total size of all label names and values must be <=
	// 1024 bytes. Label named "language" can be used to record the
	// programming language of the profiled deployment. The standard choices
	// for the value include "java", "go", "python", "ruby", "nodejs",
	// "php", "dotnet". For deployments running on Google Cloud Platform,
	// "zone" or "region" label should be present describing the deployment
	// location. An example of a zone is "us-central1-a", an example of a
	// region is "us-central1" or "us-central".
	Labels map[string]string `json:"labels,omitempty"`

	// ProjectId: Project ID is the ID of a cloud project. Validation regex:
	// `^a-z{4,61}[a-z0-9]$`.
	ProjectId string `json:"projectId,omitempty"`

	// Target: Target is the service name used to group related deployments:
	// * Service name for GAE Flex / Standard. * Cluster and container name
	// for GKE. * User-specified string for direct GCE profiling (e.g.
	// Java). * Job name for Dataflow. Validation regex: `^a-z
	// ([-a-z0-9_.]{0,253}[a-z0-9])?$`.
	Target string `json:"target,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Labels") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Labels") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Deployment) MarshalJSON() ([]byte, error) {
	type NoMethod Deployment
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Profile: Profile resource.
type Profile struct {
	// Deployment: Deployment this profile corresponds to.
	Deployment *Deployment `json:"deployment,omitempty"`

	// Duration: Duration of the profiling session. Input (for the offline
	// mode) or output (for the online mode). The field represents requested
	// profiling duration. It may slightly differ from the effective
	// profiling duration, which is recorded in the profile data, in case
	// the profiling can't be stopped immediately (e.g. in case stopping the
	// profiling is handled asynchronously).
	Duration string `json:"duration,omitempty"`

	// Labels: Input only. Labels associated to this specific profile. These
	// labels will get merged with the deployment labels for the final data
	// set. See documentation on deployment labels for validation rules and
	// limits.
	Labels map[string]string `json:"labels,omitempty"`

	// Name: Output only. Opaque, server-assigned, unique ID for this
	// profile.
	Name string `json:"name,omitempty"`

	// ProfileBytes: Input only. Profile bytes, as a gzip compressed
	// serialized proto, the format is
	// https://github.com/google/pprof/blob/master/proto/profile.proto.
	ProfileBytes string `json:"profileBytes,omitempty"`

	// ProfileType: Type of profile. For offline mode, this must be
	// specified when creating the profile. For online mode it is assigned
	// and returned by the server.
	//
	// Possible values:
	//   "PROFILE_TYPE_UNSPECIFIED" - Unspecified profile type.
	//   "CPU" - Thread CPU time sampling.
	//   "WALL" - Wallclock time sampling. More expensive as stops all
	// threads.
	//   "HEAP" - In-use heap profile. Represents a snapshot of the
	// allocations that are live at the time of the profiling.
	//   "THREADS" - Single-shot collection of all thread stacks.
	//   "CONTENTION" - Synchronization contention profile.
	//   "PEAK_HEAP" - Peak heap profile.
	//   "HEAP_ALLOC" - Heap allocation profile. It represents the
	// aggregation of all allocations made over the duration of the profile.
	// All allocations are included, including those that might have been
	// freed by the end of the profiling interval. The profile is in
	// particular useful for garbage collecting languages to understand
	// which parts of the code create most of the garbage collection
	// pressure to see if those can be optimized.
	ProfileType string `json:"profileType,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Deployment") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Deployment") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Profile) MarshalJSON() ([]byte, error) {
	type NoMethod Profile
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "cloudprofiler.projects.profiles.create":

type ProjectsProfilesCreateCall struct {
	s                    *Service
	parent               string
	createprofilerequest *CreateProfileRequest
	urlParams_           gensupport.URLParams
	ctx_                 context.Context
	header_              http.Header
}

// Create: CreateProfile creates a new profile resource in the online
// mode. The server ensures that the new profiles are created at a
// constant rate per deployment, so the creation request may hang for
// some time until the next profile session is available. The request
// may fail with ABORTED error if the creation is not available within
// ~1m, the response will indicate the duration of the backoff the
// client should take before attempting creating a profile again. The
// backoff duration is returned in google.rpc.RetryInfo extension on the
// response status. To a gRPC client, the extension will be return as a
// binary-serialized proto in the trailing metadata item named
// "google.rpc.retryinfo-bin".
//
// - parent: Parent project to create the profile in.
func (r *ProjectsProfilesService) Create(parent string, createprofilerequest *CreateProfileRequest) *ProjectsProfilesCreateCall {
	c := &ProjectsProfilesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.createprofilerequest = createprofilerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsProfilesCreateCall) Fields(s ...googleapi.Field) *ProjectsProfilesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsProfilesCreateCall) Context(ctx context.Context) *ProjectsProfilesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsProfilesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsProfilesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20210816")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createprofilerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+parent}/profiles")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudprofiler.projects.profiles.create" call.
// Exactly one of *Profile or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Profile.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsProfilesCreateCall) Do(opts ...googleapi.CallOption) (*Profile, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Profile{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "CreateProfile creates a new profile resource in the online mode. The server ensures that the new profiles are created at a constant rate per deployment, so the creation request may hang for some time until the next profile session is available. The request may fail with ABORTED error if the creation is not available within ~1m, the response will indicate the duration of the backoff the client should take before attempting creating a profile again. The backoff duration is returned in google.rpc.RetryInfo extension on the response status. To a gRPC client, the extension will be return as a binary-serialized proto in the trailing metadata item named \"google.rpc.retryinfo-bin\".",
	//   "flatPath": "v2/projects/{projectsId}/profiles",
	//   "httpMethod": "POST",
	//   "id": "cloudprofiler.projects.profiles.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Parent project to create the profile in.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+parent}/profiles",
	//   "request": {
	//     "$ref": "CreateProfileRequest"
	//   },
	//   "response": {
	//     "$ref": "Profile"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.write"
	//   ]
	// }

}

// method id "cloudprofiler.projects.profiles.createOffline":

type ProjectsProfilesCreateOfflineCall struct {
	s          *Service
	parent     string
	profile    *Profile
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// CreateOffline: CreateOfflineProfile creates a new profile resource in
// the offline mode. The client provides the profile to create along
// with the profile bytes, the server records it.
//
// - parent: Parent project to create the profile in.
func (r *ProjectsProfilesService) CreateOffline(parent string, profile *Profile) *ProjectsProfilesCreateOfflineCall {
	c := &ProjectsProfilesCreateOfflineCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.profile = profile
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsProfilesCreateOfflineCall) Fields(s ...googleapi.Field) *ProjectsProfilesCreateOfflineCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsProfilesCreateOfflineCall) Context(ctx context.Context) *ProjectsProfilesCreateOfflineCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsProfilesCreateOfflineCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsProfilesCreateOfflineCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20210816")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.profile)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+parent}/profiles:createOffline")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudprofiler.projects.profiles.createOffline" call.
// Exactly one of *Profile or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Profile.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsProfilesCreateOfflineCall) Do(opts ...googleapi.CallOption) (*Profile, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Profile{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "CreateOfflineProfile creates a new profile resource in the offline mode. The client provides the profile to create along with the profile bytes, the server records it.",
	//   "flatPath": "v2/projects/{projectsId}/profiles:createOffline",
	//   "httpMethod": "POST",
	//   "id": "cloudprofiler.projects.profiles.createOffline",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Parent project to create the profile in.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+parent}/profiles:createOffline",
	//   "request": {
	//     "$ref": "Profile"
	//   },
	//   "response": {
	//     "$ref": "Profile"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.write"
	//   ]
	// }

}

// method id "cloudprofiler.projects.profiles.patch":

type ProjectsProfilesPatchCall struct {
	s          *Service
	name       string
	profile    *Profile
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: UpdateProfile updates the profile bytes and labels on the
// profile resource created in the online mode. Updating the bytes for
// profiles created in the offline mode is currently not supported: the
// profile content must be provided at the time of the profile creation.
//
// - name: Output only. Opaque, server-assigned, unique ID for this
//   profile.
func (r *ProjectsProfilesService) Patch(name string, profile *Profile) *ProjectsProfilesPatchCall {
	c := &ProjectsProfilesPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.profile = profile
	return c
}

// UpdateMask sets the optional parameter "updateMask": Field mask used
// to specify the fields to be overwritten. Currently only profile_bytes
// and labels fields are supported by UpdateProfile, so only those
// fields can be specified in the mask. When no mask is provided, all
// fields are overwritten.
func (c *ProjectsProfilesPatchCall) UpdateMask(updateMask string) *ProjectsProfilesPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsProfilesPatchCall) Fields(s ...googleapi.Field) *ProjectsProfilesPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsProfilesPatchCall) Context(ctx context.Context) *ProjectsProfilesPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsProfilesPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsProfilesPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20210816")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.profile)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PATCH", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudprofiler.projects.profiles.patch" call.
// Exactly one of *Profile or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Profile.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsProfilesPatchCall) Do(opts ...googleapi.CallOption) (*Profile, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Profile{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "UpdateProfile updates the profile bytes and labels on the profile resource created in the online mode. Updating the bytes for profiles created in the offline mode is currently not supported: the profile content must be provided at the time of the profile creation.",
	//   "flatPath": "v2/projects/{projectsId}/profiles/{profilesId}",
	//   "httpMethod": "PATCH",
	//   "id": "cloudprofiler.projects.profiles.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Output only. Opaque, server-assigned, unique ID for this profile.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/profiles/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Field mask used to specify the fields to be overwritten. Currently only profile_bytes and labels fields are supported by UpdateProfile, so only those fields can be specified in the mask. When no mask is provided, all fields are overwritten.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "request": {
	//     "$ref": "Profile"
	//   },
	//   "response": {
	//     "$ref": "Profile"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.write"
	//   ]
	// }

}
