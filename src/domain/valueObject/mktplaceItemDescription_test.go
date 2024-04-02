package valueObject

import (
	"testing"

	testHelpers "github.com/speedianet/os/src/devUtils"
)

func TestMktplaceItemDescription(t *testing.T) {
	t.Run("ValidMktplaceItemDescription", func(t *testing.T) {
		validMktplaceItemDescriptions := []string{
			"Build and grow your website with the best way to WordPress. Lightning-fast hosting, intuitive, flexible editing, and everything you need to grow your site and audience, baked right in.",
			"It's comprised of Elasticsearch, Kibana, Beats, and Logstash (also known as the ELK Stack) and more. Reliably and securely take data from any source, in any format, then search, analyze, and visualize.",
			"RabbitMQ is a reliable and mature messaging and streaming broker, which is easy to deploy on cloud environments, on-premises, and on your local machine. It is currently used by millions worldwide.",
		}
		for _, mid := range validMktplaceItemDescriptions {
			_, err := NewMktplaceItemDescription(mid)
			if err != nil {
				t.Errorf("Expected no error for %s, got %s", mid, err.Error())
			}
		}
	})

	t.Run("InvalidMktplaceItemDescription", func(t *testing.T) {
		invalidLength := 600
		invalidMktplaceItemDescriptions := []string{
			"",
			"a",
			testHelpers.GenerateString(invalidLength),
		}
		for _, mid := range invalidMktplaceItemDescriptions {
			_, err := NewMktplaceItemDescription(mid)
			if err == nil {
				t.Errorf("Expected error for %s, got nil", mid)
			}
		}
	})
}