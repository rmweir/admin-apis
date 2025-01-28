package main

import (
	"fmt"
	"os"

	"github.com/loft-sh/admin-apis/hack/internal/featuresyaml"
	"github.com/stripe/stripe-go/v81"
	stripefeatures "github.com/stripe/stripe-go/v81/entitlements/feature"
	stripeproducts "github.com/stripe/stripe-go/v81/product"
)

func main() {
	stripeToken := os.Getenv("STRIPE_API_KEY")
	if stripeToken == "" {
		panic("stripe token cannot be empty")
	}
	stripe.Key = stripeToken

	features, err := featuresyaml.ReadFeaturesYaml("pkg/licenseapi/features.yaml")
	if err != nil {
		panic(err)
	}

	productsWithAllFeatures := stripeproducts.List(&stripe.ProductListParams{})
	for _, f := range features {
		err = ensureFeatureExists(f.Name, f.DisplayName)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if !f.Preview {
			continue
		}

		err = ensureFeatureExists(f.Name+"-preview", "Preview: "+f.DisplayName)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func ensureFeatureExists(name, displayName string) error {
	exists, err := featureExists(name)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	_, err = stripefeatures.New(&stripe.EntitlementsFeatureParams{
		Name:      &displayName,
		LookupKey: &name,
	})
	if err != nil {
		return fmt.Errorf("failed to create Stripe feature from feature %s: %v\n", name, err)
	}
	return nil
}

func featureExists(id string) (bool, error) {
	list := stripefeatures.List(&stripe.EntitlementsFeatureListParams{
		LookupKey: &id,
	})
	if err := list.Err(); err != nil {
		return false, err
	}
	if list.Next() {
		return true, nil
	}
	return false, nil
}
