package genesyscloud

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/mypurecloud/platform-client-sdk-go/v91/platformclientv2"
)

func dataSourceKnowledgeLabel() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Genesys Cloud Knowledge Base Label. Select a label by name.",
		ReadContext: readWithPooledClient(dataSourceKnowledgeLabelRead),
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Knowledge base label name",
				Type:        schema.TypeString,
				Required:    true,
			},
			"knowledge_base_name": {
				Description: "Knowledge base name",
				Type:        schema.TypeString,
				Required:    true,
			},
			"core_language": {
				Description:  "Core language for knowledge base in which initial content must be created, language codes [en-US, en-UK, en-AU, de-DE] are supported currently, however the new DX knowledge will support all these language codes",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"en-US", "en-UK", "en-AU", "de-DE", "es-US", "es-ES", "fr-FR", "pt-BR", "nl-NL", "it-IT", "fr-CA"}, false),
			},
		},
	}
}

func dataSourceKnowledgeLabelRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sdkConfig := m.(*providerMeta).ClientConfig
	knowledgeAPI := platformclientv2.NewKnowledgeApiWithConfig(sdkConfig)

	name := d.Get("name").(string)
	knowledgeBaseName := d.Get("knowledge_base_name").(string)
	coreLanguage := d.Get("core_language").(string)

	// Find first non-deleted knowledge base by name. Retry in case new knowledge base is not yet indexed by search
	return withRetries(ctx, 15*time.Second, func() *resource.RetryError {
		for pageNum := 1; ; pageNum++ {
			const pageSize = 100
			publishedKnowledgeBases, _, getPublishedErr := knowledgeAPI.GetKnowledgeKnowledgebases("", "", "", fmt.Sprintf("%v", pageSize), knowledgeBaseName, coreLanguage, true, "", "")
			unpublishedKnowledgeBases, _, getUnpublishedErr := knowledgeAPI.GetKnowledgeKnowledgebases("", "", "", fmt.Sprintf("%v", pageSize), knowledgeBaseName, coreLanguage, false, "", "")

			if getPublishedErr != nil {
				return resource.NonRetryableError(fmt.Errorf("Failed to get knowledge base %s: %s", knowledgeBaseName, getPublishedErr))
			}
			if getUnpublishedErr != nil {
				return resource.NonRetryableError(fmt.Errorf("Failed to get knowledge base %s: %s", knowledgeBaseName, getPublishedErr))
			}

			noPublishedEntities := publishedKnowledgeBases.Entities == nil || len(*publishedKnowledgeBases.Entities) == 0
			noUnpublishedEntities := unpublishedKnowledgeBases.Entities == nil || len(*unpublishedKnowledgeBases.Entities) == 0
			if noPublishedEntities && noUnpublishedEntities {
				return resource.RetryableError(fmt.Errorf("no knowledge bases found with name %s", knowledgeBaseName))
			}

			// prefer published knowledge base
			for _, knowledgeBase := range *publishedKnowledgeBases.Entities {
				if knowledgeBase.Name != nil && *knowledgeBase.Name == knowledgeBaseName && *knowledgeBase.CoreLanguage == coreLanguage {
					knowledgeLabels, _, getErr := knowledgeAPI.GetKnowledgeKnowledgebaseLabels(*knowledgeBase.Id, "", "", fmt.Sprintf("%v", pageSize), name, false)

					if getErr != nil {
						return resource.NonRetryableError(fmt.Errorf("Failed to get knowledge category %s: %s", name, getErr))
					}

					for _, knowledgeLabel := range *knowledgeLabels.Entities {
						if *knowledgeLabel.Name == name {
							id := fmt.Sprintf("%s,%s", *knowledgeLabel.Id, *knowledgeBase.Id)
							d.SetId(id)
							return nil
						}
					}
				}
			}
			// use unpublished knowledge base if unpublished doesn't exist
			for _, knowledgeBase := range *unpublishedKnowledgeBases.Entities {
				if knowledgeBase.Name != nil && *knowledgeBase.Name == knowledgeBaseName && *knowledgeBase.CoreLanguage == coreLanguage {
					knowledgeLabels, _, getErr := knowledgeAPI.GetKnowledgeKnowledgebaseLabels(*knowledgeBase.Id, "", "", fmt.Sprintf("%v", pageSize), name, false)

					if getErr != nil {
						return resource.NonRetryableError(fmt.Errorf("Failed to get knowledge category %s: %s", name, getErr))
					}

					for _, knowledgeLabel := range *knowledgeLabels.Entities {
						if *knowledgeLabel.Name == name {
							id := fmt.Sprintf("%s,%s", *knowledgeLabel.Id, *knowledgeBase.Id)
							d.SetId(id)
							return nil
						}
					}
				}
			}
		}
	})
}