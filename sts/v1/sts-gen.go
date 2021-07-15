// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package sts provides access to the Security Token Service API.
//
// For product documentation, see: http://cloud.google.com/iam/docs/workload-identity-federation
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/sts/v1"
//   ...
//   ctx := context.Background()
//   stsService, err := sts.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   stsService, err := sts.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   stsService, err := sts.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package sts // import "google.golang.org/api/sts/v1"

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

const apiId = "sts:v1"
const apiName = "sts"
const apiVersion = "v1"
const basePath = "https://sts.googleapis.com/"
const mtlsBasePath = "https://sts.mtls.googleapis.com/"

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
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
	s.V1 = NewV1Service(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	V1 *V1Service
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewV1Service(s *Service) *V1Service {
	rs := &V1Service{s: s}
	return rs
}

type V1Service struct {
	s *Service
}

// GoogleIdentityStsV1ExchangeTokenRequest: Request message for
// ExchangeToken.
type GoogleIdentityStsV1ExchangeTokenRequest struct {
	// Audience: The full resource name of the identity provider; for
	// example:
	// `//iam.googleapis.com/projects//workloadIdentityPools//providers/`.
	// Required when exchanging an external credential for a Google access
	// token.
	Audience string `json:"audience,omitempty"`

	// GrantType: Required. The grant type. Must be
	// `urn:ietf:params:oauth:grant-type:token-exchange`, which indicates a
	// token exchange.
	GrantType string `json:"grantType,omitempty"`

	// Options: A set of features that Security Token Service supports, in
	// addition to the standard OAuth 2.0 token exchange, formatted as a
	// serialized JSON object of Options.
	Options string `json:"options,omitempty"`

	// RequestedTokenType: Required. An identifier for the type of requested
	// security token. Must be
	// `urn:ietf:params:oauth:token-type:access_token`.
	RequestedTokenType string `json:"requestedTokenType,omitempty"`

	// Scope: The OAuth 2.0 scopes to include on the resulting access token,
	// formatted as a list of space-delimited, case-sensitive strings.
	// Required when exchanging an external credential for a Google access
	// token.
	Scope string `json:"scope,omitempty"`

	// SubjectToken: Required. The input token. This token is either an
	// external credential issued by a workload identity pool provider, or a
	// short-lived access token issued by Google. If the token is an OIDC
	// JWT, it must use the JWT format defined in RFC 7523
	// (https://tools.ietf.org/html/rfc7523), and the `subject_token_type`
	// must be either `urn:ietf:params:oauth:token-type:jwt` or
	// `urn:ietf:params:oauth:token-type:id_token`. The following headers
	// are required: - `kid`: The identifier of the signing key securing the
	// JWT. - `alg`: The cryptographic algorithm securing the JWT. Must be
	// `RS256` or `ES256`. The following payload fields are required. For
	// more information, see RFC 7523, Section 3
	// (https://tools.ietf.org/html/rfc7523#section-3): - `iss`: The issuer
	// of the token. The issuer must provide a discovery document at the URL
	// `/.well-known/openid-configuration`, where `` is the value of this
	// field. The document must be formatted according to section 4.2 of the
	// OIDC 1.0 Discovery specification
	// (https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderConfigurationResponse).
	// - `iat`: The issue time, in seconds, since the Unix epoch. Must be in
	// the past. - `exp`: The expiration time, in seconds, since the Unix
	// epoch. Must be less than 48 hours after `iat`. Shorter expiration
	// times are more secure. If possible, we recommend setting an
	// expiration time less than 6 hours. - `sub`: The identity asserted in
	// the JWT. - `aud`: For workload identity pools, this must be a value
	// specified in the allowed audiences for the workload identity pool
	// provider, or one of the audiences allowed by default if no audiences
	// were specified. See
	// https://cloud.google.com/iam/docs/reference/rest/v1/projects.locations.workloadIdentityPools.providers#oidc
	// Example header: ``` { "alg": "RS256", "kid": "us-east-11" } ```
	// Example payload: ``` { "iss": "https://accounts.google.com", "iat":
	// 1517963104, "exp": 1517966704, "aud":
	// "//iam.googleapis.com/projects/1234567890123/locations/global/workload
	// IdentityPools/my-pool/providers/my-provider", "sub":
	// "113475438248934895348", "my_claims": { "additional_claim": "value" }
	// } ``` If `subject_token` is for AWS, it must be a serialized
	// `GetCallerIdentity` token. This token contains the same information
	// as a request to the AWS `GetCallerIdentity()`
	// (https://docs.aws.amazon.com/STS/latest/APIReference/API_GetCallerIdentity)
	// method, as well as the AWS signature
	// (https://docs.aws.amazon.com/general/latest/gr/signing_aws_api_requests.html)
	// for the request information. Use Signature Version 4. Format the
	// request as URL-encoded JSON, and set the `subject_token_type`
	// parameter to `urn:ietf:params:aws:token-type:aws4_request`. The
	// following parameters are required: - `url`: The URL of the AWS STS
	// endpoint for `GetCallerIdentity()`, such as
	// `https://sts.amazonaws.com?Action=GetCallerIdentity&Version=2011-06-15
	// `. Regional endpoints are also supported. - `method`: The HTTP
	// request method: `POST`. - `headers`: The HTTP request headers, which
	// must include: - `Authorization`: The request signature. -
	// `x-amz-date`: The time you will send the request, formatted as an
	// ISO8601 Basic
	// (https://docs.aws.amazon.com/general/latest/gr/sigv4_elements.html#sigv4_elements_date)
	// string. This value is typically set to the current time and is used
	// to help prevent replay attacks. - `host`: The hostname of the `url`
	// field; for example, `sts.amazonaws.com`. -
	// `x-goog-cloud-target-resource`: The full, canonical resource name of
	// the workload identity pool provider, with or without an `https:`
	// prefix. To help ensure data integrity, we recommend including this
	// header in the `SignedHeaders` field of the signed request. For
	// example:
	// //iam.googleapis.com/projects//locations//workloadIdentityPools//provi
	// ders/
	// https://iam.googleapis.com/projects//locations//workloadIdentityPools//providers/
	// If you are using temporary security credentials provided by AWS, you
	// must also include the header `x-amz-security-token`, with the value
	// set to the session token. The following example shows a
	// `GetCallerIdentity` token: ``` { "headers": [ {"key": "x-amz-date",
	// "value": "20200815T015049Z"}, {"key": "Authorization", "value":
	// "AWS4-HMAC-SHA256+Credential=$credential,+SignedHeaders=host;x-amz-dat
	// e;x-goog-cloud-target-resource,+Signature=$signature"}, {"key":
	// "x-goog-cloud-target-resource", "value":
	// "//iam.googleapis.com/projects//locations//workloadIdentityPools//prov
	// iders/"}, {"key": "host", "value": "sts.amazonaws.com"} . ],
	// "method": "POST", "url":
	// "https://sts.amazonaws.com?Action=GetCallerIdentity&Version=2011-06-15
	// " } ``` You can also use a Google-issued OAuth 2.0 access token with
	// this field to obtain an access token with new security attributes
	// applied, such as a Credential Access Boundary. In this case, set
	// `subject_token_type` to
	// `urn:ietf:params:oauth:token-type:access_token`. If an access token
	// already contains security attributes, you cannot apply additional
	// security attributes.
	SubjectToken string `json:"subjectToken,omitempty"`

	// SubjectTokenType: Required. An identifier that indicates the type of
	// the security token in the `subject_token` parameter. Supported values
	// are `urn:ietf:params:oauth:token-type:jwt`,
	// `urn:ietf:params:oauth:token-type:id_token`,
	// `urn:ietf:params:aws:token-type:aws4_request`, and
	// `urn:ietf:params:oauth:token-type:access_token`.
	SubjectTokenType string `json:"subjectTokenType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Audience") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Audience") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIdentityStsV1ExchangeTokenRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIdentityStsV1ExchangeTokenRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIdentityStsV1ExchangeTokenResponse: Response message for
// ExchangeToken.
type GoogleIdentityStsV1ExchangeTokenResponse struct {
	// AccessToken: An OAuth 2.0 security token, issued by Google, in
	// response to the token exchange request. Tokens can vary in size,
	// depending in part on the size of mapped claims, up to a maximum of
	// 12288 bytes (12 KB). Google reserves the right to change the token
	// size and the maximum length at any time.
	AccessToken string `json:"access_token,omitempty"`

	// ExpiresIn: The amount of time, in seconds, between the time when the
	// access token was issued and the time when the access token will
	// expire. This field is absent when the `subject_token` in the request
	// is a Google-issued, short-lived access token. In this case, the
	// access token has the same expiration time as the `subject_token`.
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// IssuedTokenType: The token type. Always matches the value of
	// `requested_token_type` from the request.
	IssuedTokenType string `json:"issued_token_type,omitempty"`

	// TokenType: The type of access token. Always has the value `Bearer`.
	TokenType string `json:"token_type,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AccessToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccessToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIdentityStsV1ExchangeTokenResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIdentityStsV1ExchangeTokenResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIdentityStsV1IntrospectTokenRequest: Request message for
// IntrospectToken.
type GoogleIdentityStsV1IntrospectTokenRequest struct {
	// Token: Required. The OAuth 2.0 security token issued by the Security
	// Token Service API.
	Token string `json:"token,omitempty"`

	// TokenTypeHint: Optional. The type of the given token. Supported
	// values are `urn:ietf:params:oauth:token-type:access_token` and
	// `access_token`.
	TokenTypeHint string `json:"tokenTypeHint,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Token") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Token") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIdentityStsV1IntrospectTokenRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIdentityStsV1IntrospectTokenRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIdentityStsV1IntrospectTokenResponse: Response message for
// IntrospectToken.
type GoogleIdentityStsV1IntrospectTokenResponse struct {
	// Active: A boolean value that indicates whether the provided access
	// token is currently active.
	Active bool `json:"active,omitempty"`

	// ClientId: The client identifier for the OAuth 2.0 client that
	// requested the provided token.
	ClientId string `json:"client_id,omitempty"`

	// Exp: The expiration timestamp, measured in the number of seconds
	// since January 1 1970 UTC, indicating when this token will expire.
	Exp int64 `json:"exp,omitempty,string"`

	// Iat: The issued timestamp, measured in the number of seconds since
	// January 1 1970 UTC, indicating when this token was originally issued.
	Iat int64 `json:"iat,omitempty,string"`

	// Iss: The issuer of the provided token.
	Iss string `json:"iss,omitempty"`

	// Scope: A list of scopes associated with the provided token.
	Scope string `json:"scope,omitempty"`

	// Sub: The unique user ID associated with the provided token. For
	// Google Accounts, this value is based on the Google Account's user ID.
	// For federated identities, this value is based on the identity pool ID
	// and the value of the mapped `google.subject` attribute.
	Sub string `json:"sub,omitempty"`

	// Username: The human-readable identifier for the token principal
	// subject. For example, if the provided token is associated with a
	// workload identity pool, this field contains a value in the following
	// format:
	// `principal://iam.googleapis.com/projects//locations//workloadIdentityP
	// ools//subject/`
	Username string `json:"username,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Active") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Active") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIdentityStsV1IntrospectTokenResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIdentityStsV1IntrospectTokenResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "sts.introspect":

type V1IntrospectCall struct {
	s                                         *Service
	googleidentitystsv1introspecttokenrequest *GoogleIdentityStsV1IntrospectTokenRequest
	urlParams_                                gensupport.URLParams
	ctx_                                      context.Context
	header_                                   http.Header
}

// Introspect: Gets information about a Google OAuth 2.0 access token
// issued by the Google Cloud Security Token Service API
// (https://cloud.google.com/iam/docs/reference/sts/rest).
func (r *V1Service) Introspect(googleidentitystsv1introspecttokenrequest *GoogleIdentityStsV1IntrospectTokenRequest) *V1IntrospectCall {
	c := &V1IntrospectCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.googleidentitystsv1introspecttokenrequest = googleidentitystsv1introspecttokenrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1IntrospectCall) Fields(s ...googleapi.Field) *V1IntrospectCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V1IntrospectCall) Context(ctx context.Context) *V1IntrospectCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *V1IntrospectCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *V1IntrospectCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20210714")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleidentitystsv1introspecttokenrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/introspect")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "sts.introspect" call.
// Exactly one of *GoogleIdentityStsV1IntrospectTokenResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleIdentityStsV1IntrospectTokenResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *V1IntrospectCall) Do(opts ...googleapi.CallOption) (*GoogleIdentityStsV1IntrospectTokenResponse, error) {
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
	ret := &GoogleIdentityStsV1IntrospectTokenResponse{
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
	//   "description": "Gets information about a Google OAuth 2.0 access token issued by the Google Cloud [Security Token Service API](https://cloud.google.com/iam/docs/reference/sts/rest).",
	//   "flatPath": "v1/introspect",
	//   "httpMethod": "POST",
	//   "id": "sts.introspect",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/introspect",
	//   "request": {
	//     "$ref": "GoogleIdentityStsV1IntrospectTokenRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleIdentityStsV1IntrospectTokenResponse"
	//   }
	// }

}

// method id "sts.token":

type V1TokenCall struct {
	s                                       *Service
	googleidentitystsv1exchangetokenrequest *GoogleIdentityStsV1ExchangeTokenRequest
	urlParams_                              gensupport.URLParams
	ctx_                                    context.Context
	header_                                 http.Header
}

// Token: Exchanges a credential for a Google OAuth 2.0 access token.
// The token asserts an external identity within a workload identity
// pool, or it applies a Credential Access Boundary to a Google access
// token. When you call this method, do not send the `Authorization`
// HTTP header in the request. This method does not require the
// `Authorization` header, and using the header can cause the request to
// fail.
func (r *V1Service) Token(googleidentitystsv1exchangetokenrequest *GoogleIdentityStsV1ExchangeTokenRequest) *V1TokenCall {
	c := &V1TokenCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.googleidentitystsv1exchangetokenrequest = googleidentitystsv1exchangetokenrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1TokenCall) Fields(s ...googleapi.Field) *V1TokenCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V1TokenCall) Context(ctx context.Context) *V1TokenCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *V1TokenCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *V1TokenCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20210714")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleidentitystsv1exchangetokenrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/token")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "sts.token" call.
// Exactly one of *GoogleIdentityStsV1ExchangeTokenResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleIdentityStsV1ExchangeTokenResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *V1TokenCall) Do(opts ...googleapi.CallOption) (*GoogleIdentityStsV1ExchangeTokenResponse, error) {
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
	ret := &GoogleIdentityStsV1ExchangeTokenResponse{
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
	//   "description": "Exchanges a credential for a Google OAuth 2.0 access token. The token asserts an external identity within a workload identity pool, or it applies a Credential Access Boundary to a Google access token. When you call this method, do not send the `Authorization` HTTP header in the request. This method does not require the `Authorization` header, and using the header can cause the request to fail.",
	//   "flatPath": "v1/token",
	//   "httpMethod": "POST",
	//   "id": "sts.token",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/token",
	//   "request": {
	//     "$ref": "GoogleIdentityStsV1ExchangeTokenRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleIdentityStsV1ExchangeTokenResponse"
	//   }
	// }

}
