package main

import (
	"fmt"
	"log"
	"maps"
	"os"

	"github.com/loft-sh/admin-apis/hack/internal/featuresyaml"
	"github.com/loft-sh/admin-apis/pkg/licenseapi"
	"github.com/stripe/stripe-go/v81"
	stripefeatures "github.com/stripe/stripe-go/v81/entitlements/feature"
	stripeproducts "github.com/stripe/stripe-go/v81/product"
	"github.com/stripe/stripe-go/v81/productfeature"
)

const (
	metadataQueryFmt       = "metadata['%s']:'%s'"
	metadataFieldAttachAll = "attach_all_features"
)

// TODO possibly add displayName here either explicity or just embed licenseapi.Feature for all non-stripe fields
type syncedFeature struct {
	licenseapi.Feature
	name        string
	displayName string
	stripeID    string
}

func main() {
	stripeToken := os.Getenv("STRIPE_API_KEY")
	if stripeToken == "" {
		log.Println("stripe token cannot be empty")
		os.Exit(1)
	}
	stripe.Key = stripeToken

	features, err := featuresyaml.ReadFeaturesYaml("pkg/licenseapi/features.yaml")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	syncedFeatures := syncStripeFeatures(features)

	if err = ensureFeatureProducts(syncedFeatures); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if err = ensureAttachAll(syncedFeatures); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}

func syncStripeFeatures(features []*licenseapi.Feature) map[syncedFeature]bool {
	featureIDs := make(map[syncedFeature]bool, len(features))
	for _, f := range features {
		id, err := ensureFeatureExists(f.Name, f.DisplayName)
		if err != nil {
			log.Println(err)
			continue
		}
		featureIDs[syncedFeature{name: f.Name, stripeID: id}] = true

		if !f.Preview {
			continue
		}

		previewName := f.Name + "-preview"
		previewId, err := ensureFeatureExists(previewName, "Preview: "+f.DisplayName)
		if err != nil {
			log.Println(err)
			continue
		}
		featureIDs[syncedFeature{name: previewName, stripeID: previewId}] = true
	}
	return featureIDs
}

func ensureFeatureExists(name, displayName string) (string, error) {
	id, exists, err := featureExists(name)
	if err != nil {
		return "", err
	}

	if exists {
		return id, nil
	}

	feature, err := stripefeatures.New(&stripe.EntitlementsFeatureParams{
		Name:      &displayName,
		LookupKey: &name,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create Stripe feature from feature %s: %v\n", name, err)
	}
	return feature.ID, nil
}

func featureExists(id string) (string, bool, error) {
	list := stripefeatures.List(&stripe.EntitlementsFeatureListParams{
		LookupKey: &id,
	})
	if err := list.Err(); err != nil {
		return "", false, fmt.Errorf("failed to list features while check if feature [%s] exists: %w", id, err)
	}
	if !list.Next() {
		return "", false, nil
	}
	feature, ok := list.Current().(*stripe.EntitlementsFeature)
	if !ok {
		return "", false, fmt.Errorf("failed to ")
	}
	return feature.ID, true, nil
}

func ensureFeatureProducts(syncedFeatures map[syncedFeature]bool) error {
	for feature := range syncedFeatures {
		if err := ensureFeatureProduct(feature); err != nil {
			return err
		}
	}
	return nil
}

func ensureFeatureProduct(syncedFeature syncedFeature) error {
	search := stripeproducts.Search(&stripe.ProductSearchParams{
		SearchParams: stripe.SearchParams{
			Query: fmt.Sprintf(metadataQueryFmt, syncedFeature.name, "true"),
		},
	})
	if err := search.Err(); err != nil {
		return err
	}
	if search.Next() {
		// a product exists with features name
		return nil
	}
	product, err := stripeproducts.New(&stripe.ProductParams{
		Name: &syncedFeature.name, // TODO: probably going to change name to something like the displayName of feature
	})
	if err != nil {
		return err
	}
	_, err = productfeature.New(&stripe.ProductFeatureParams{
		Product:            &product.ID,
		EntitlementFeature: &syncedFeature.stripeID,
	})
	if err != nil {
		return err
	}
	return nil
}

func ensureAttachAll(featureIDs map[syncedFeature]bool) error {
	search := stripeproducts.Search(&stripe.ProductSearchParams{
		SearchParams: stripe.SearchParams{
			Query: fmt.Sprintf(metadataQueryFmt, metadataFieldAttachAll, "true"),
		},
	})
	if err := search.Err(); err != nil {
		return err
	}

	for search.Next() {
		prod := search.Product()
		featuresToCheck := maps.Clone[map[syncedFeature]bool](featureIDs)
		if err := SearchProductForFeatures(prod.ID, featuresToCheck); err != nil {
			return err
		}
		for feature := range featuresToCheck {
			_, err := productfeature.New(&stripe.ProductFeatureParams{
				Product:            &prod.ID,
				EntitlementFeature: &feature.stripeID,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func SearchProductForFeatures(productID string, featuresToCheck map[syncedFeature]bool) error {
	productsPage := productfeature.List(&stripe.ProductFeatureListParams{
		Product: &productID,
	})
	for productsPage.Next() {
		if err := productsPage.Err(); err != nil {
			return err
		}
		delete(featuresToCheck, syncedFeature{name: productsPage.ProductFeature().EntitlementFeature.LookupKey, stripeID: productsPage.ProductFeature().EntitlementFeature.ID})
	}
	return nil
}
