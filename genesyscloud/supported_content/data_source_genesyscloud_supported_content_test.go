package supported_content

import (
	"fmt"
	"testing"

	"terraform-provider-genesyscloud/genesyscloud/provider"
	"terraform-provider-genesyscloud/genesyscloud/util"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

/*
Test Class for the supported content Data Source
*/

func TestAccDataSourceSupportedContent(t *testing.T) {
	t.Parallel()
	var (
		resourceId   = "testSupportedContent"
		dataSourceId = "testSupportedContent_data"
		name         = "Terraform Supported Content - " + uuid.NewString()
		inboundType  = "*/*"
		outboundType = "image/*"
	)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { util.TestAccPreCheck(t) },
		ProviderFactories: provider.GetProviderFactories(providerResources, providerDataSources),
		Steps: []resource.TestStep{
			{
				Config: GenerateSupportedContentResource(
					resourceId,
					name,
					GenerateInboundTypeBlock(inboundType),
					GenerateOutboundTypeBlock(outboundType),
				) +
					GenerateDataSourceForSupportedContent(
						dataSourceId,
						name,
						"genesyscloud_supported_content."+resourceId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.genesyscloud_supported_content."+dataSourceId, "id", "genesyscloud_supported_content."+resourceId, "id"),
				),
			},
		},
	})
}

func GenerateDataSourceForSupportedContent(
	resourceId string,
	name string,
	dependsOn string,
) string {
	return fmt.Sprintf(`
	data "`+resourceName+`" "%s" {
		name = "%s"
		depends_on = [%s]
	}
	`, resourceId, name, dependsOn)
}
