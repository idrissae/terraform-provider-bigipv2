---
layout: "bigip"
page_title: "BIG-IP: bigip_ltm_profile_http"
subcategory: "Local Traffic Manager(LTM)"
description: |-
  Provides details about bigip_ltm_profile_http resource
---

# bigip\_ltm\_profile_http

`bigip_ltm_profile_http` Configures a custom profile_http for use by health checks.

For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.

## Example Usage


```hcl
resource "bigip_ltm_profile_http" "sanjose-http" {
  name                  = "/Common/sanjose-http"
  defaults_from         = "/Common/http"
  fallback_host         = "titanic"
  fallback_status_codes = ["400", "500", "300"]
}

```      

## Argument Reference

* `name` (Required,type `string`) Specifies the name of the http profile,name of Profile should be full path. Full path is the combination of the `partition + profile name`,For example `/Common/test-http-profile`.

* `proxy_type` - (optional,type `string`) Specifies the proxy mode for this profile: reverse, explicit, or transparent. The default is `reverse`.

* `defaults_from` - (optional,type `string`) Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.

* `description` - (optional,type `string`) Specifies user-defined description.

* `basic_auth_realm` - (Optional) Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is `none`

* `fallback_host` - (Optional) Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number

* `fallback_status_codes` - (Optional,type `list`) Specifies one or more three-digit status codes that can be returned by an HTTP server,that should trigger a redirection to the fallback host.

* `head_erase` - (Optional) Specifies the header string that you want to erase from an HTTP request. Default is `none`.

* `head_insert` - (Optional) Specifies a quoted header string that you want to insert into an HTTP request.Default is `none`.

* `insert_xforwarded_for` - (Optional) When using connection pooling, which allows clients to make use of other client requests' server-side connections, you can insert the X-Forwarded-For header and specify a client IP address

* `response_headers_permitted` - (Optional,type `list`) Specifies headers that the BIG-IP system allows in an HTTP response.If you are specifying more than one header, separate the headers with a blank space.

* `request_chunking` - (Optional,type `string`) Specifies how the system handles HTTP content that is chunked by a client. The default is `preserve`.

* `response_chunking` - (Optional,type `string`) Specifies how the system handles HTTP content that is chunked by a server. The default is `selective`.

* `oneconnect_transformations` - (Optional) Enables the system to perform HTTP header transformations for the purpose of  keeping server-side connections open. This feature requires configuration of a OneConnect profile

* `redirect_rewrite` - (Optional) Specifies whether the system rewrites the URIs that are part of HTTP redirect (3XX) responses. The default is `none`.

* `request_chunking` - (Optional) Specifies how the system handles HTTP content that is chunked by a client. The default is `preserve`.

* `encrypt_cookies` - (Optional) Type the cookie names for the system to encrypt.

* `encrypt_cookie_secret` - (Optional) Type a passphrase for cookie encryption.

* `insert_xforwarded_for` - (Optional) Specifies, when enabled, that the system inserts an X-Forwarded-For header in an HTTP request with the client IP address, to use with connection pooling. The default is `Disabled`.

* `lws_width` - (Optional,type `int`) Specifies the maximum column width for any given line, when inserting an HTTP header in an HTTP request. The default is `80`.

* `lws_width` - (Optional,type `string`) Specifies the linear white space (LWS) separator that the system inserts when a header exceeds the maximum width you specify in the LWS Maximum Columns setting.

* `accept_xff` - (Optional) Enables or disables trusting the client IP address, and statistics from the client IP address, based on the request's XFF (X-forwarded-for) headers, if they exist.

* `xff_alternative_names` - (Optional) Specifies alternative XFF headers instead of the default X-forwarded-for header.


## Import

BIG-IP LTM http profiles can be imported using the `name`, e.g.

```bash
terraform import bigip_ltm_profile_http.test-http /Common/test-http
```
