---
page_title: "genesyscloud_journey_outcome_predictor Resource - terraform-provider-genesyscloud"
subcategory: ""
description: |-
  Genesys Cloud journey outcome predictor
---
# genesyscloud_journey_outcome_predictor (Resource)

Genesys Cloud journey outcome predictor

## API Usage
The following Genesys Cloud APIs are used by this resource. Ensure your OAuth Client has been granted the necessary scopes and permissions to perform these operations:

* [GET /api/v2/journey/outcomes/predictors](https://apicentral.genesys.cloud/api-explorer#get-api-v2-journey-outcomes-predictors)
* [POST /api/v2/journey/outcomes/predictors](https://apicentral.genesys.cloud/api-explorer#post-api-v2-journey-outcomes-predictors)
* [GET /api/v2/journey/outcomes/predictors/{predictorId}](https://apicentral.genesys.cloud/api-explorer#get-api-v2-journey-outcomes-predictors--predictorId-)
* [DELETE /api/v2/journey/outcomes/predictors/{predictorId}](https://apicentral.genesys.cloud/api-explorer#delete-api-v2-journey-outcomes-predictors--predictorId-)

## Example Usage

```terraform
resource "genesyscloud_journey_outcome_predictor" "example_journey_outcome_predictor_resource" {
  outcome_id = data.genesyscloud_journey_outcome.example_outcome.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `outcome_id` (String) The outcome associated with this predictor

### Read-Only

- `id` (String) The ID of this resource.
