---
page_title: "genesyscloud_integration_custom_auth_action Resource - terraform-provider-genesyscloud"
subcategory: ""
description: |-
  Genesys Cloud Integration Actions. See this page for detailed information on configuring Actions: https://help.mypurecloud.com/articles/add-configuration-custom-actions-integrations/
---
# genesyscloud_integration_custom_auth_action (Resource)

Genesys Cloud Integration Actions. See this page for detailed information on configuring Actions: https://help.mypurecloud.com/articles/add-configuration-custom-actions-integrations/

## API Usage
The following Genesys Cloud APIs are used by this resource. Ensure your OAuth Client has been granted the necessary scopes and permissions to perform these operations:

* [GET /api/v2/integrations/actions](https://developer.genesys.cloud/api/rest/v2/integrations/#get-api-v2-integrations-actions)
* [GET /api/v2/integrations/actions/{actionId}](https://developer.genesys.cloud/api/rest/v2/integrations/#get-api-v2-integrations-actions--actionId-)
* [GET /api/v2/integrations/actions/{actionId}/templates/{fileName}](https://developer.genesys.cloud/api/rest/v2/integrations/#get-api-v2-integrations-actions--actionId--templates--fileName-)
* [PATCH /api/v2/integrations/actions/{actionId}](https://developer.genesys.cloud/api/rest/v2/integrations/#patch-api-v2-integrations-actions--actionId-)
* [GET /api/v2/integrations/{integrationId}/config/current](https://developer.mypurecloud.com/api/rest/v2/integrations/#get-api-v2-integrations--integrationId--config-current)
* [GET /api/v2/integrations/credentials/{credentialId}](https://developer.genesys.cloud/api/rest/v2/integrations/#get-api-v2-integrations-credentials--credentialId-)
* [PATCH /api/v2/integrations/actions/{actionId}/draft](https://developer.genesys.cloud/platform/integrations/#patch-api-v2-integrations-actions--actionId--draft)
* [POST /api/v2/integrations/actions/{actionId}/draft/publish](https://developer.genesys.cloud/platform/integrations/#post-api-v2-integrations-actions--actionId--draft-publish)
* [GET /api/v2/integrations/actions/{actionId}/draft](https://developer.genesys.cloud/platform/integrations/#get-api-v2-integrations-actions--actionId--draft)


## Example Usage

```terraform
resource "genesyscloud_integration_custom_auth_action" "example-custom-auth-action" {
  integration_id = genesyscloud_integration.example_integ.id
  name           = "Example Custom Auth Action"
  config_request {
    # Use '$${' to indicate a literal '${' in template strings. Otherwise Terraform will attempt to interpolate the string
    # See https://www.terraform.io/docs/language/expressions/strings.html#escape-sequences
    request_url_template = "$${credentials.loginUrl}"
    request_type         = "POST"
    request_template     = "grant_type=client_credentials"
    headers = {
      Authorization = "Basic $encoding.base64(\"$${credentials.clientId}:$${credentials.clientSecret}\")"
      Content-Type  = "application/x-www-form-urlencoded"
    }
  }
  config_response {
    translation_map = {
      tokenValue = "$.token"
    }
    success_template = "{ \"token\": $${tokenValue} }"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `integration_id` (String) The ID of the integration this action is associated with. The integration is required to be of type `custom-rest-actions` and its credentials type set as `userDefinedOAuth`.

### Optional

- `config_request` (Block List, Max: 1) Configuration of outbound request. (see [below for nested schema](#nestedblock--config_request))
- `config_response` (Block List, Max: 1) Configuration of response processing. (see [below for nested schema](#nestedblock--config_response))
- `name` (String) Name of the action to override the default name. Can be up to 256 characters long

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--config_request"></a>
### Nested Schema for `config_request`

Required:

- `request_type` (String) HTTP method to use for request (GET | PUT | POST | PATCH | DELETE).
- `request_url_template` (String) URL that may include placeholders for requests to 3rd party service.

Optional:

- `headers` (Map of String) Map of headers in name, value pairs to include in request.
- `request_template` (String) Velocity template to define request body sent to 3rd party service. Any instances of '${' must be properly escaped as '$${'


<a id="nestedblock--config_response"></a>
### Nested Schema for `config_response`

Optional:

- `success_template` (String) Velocity template to build response to return from Action. Any instances of '${' must be properly escaped as '$${'.
- `translation_map` (Map of String) Map 'attribute name' and 'JSON path' pairs used to extract data from REST response.
- `translation_map_defaults` (Map of String) Map 'attribute name' and 'default value' pairs used as fallback values if JSON path extraction fails for specified key.
