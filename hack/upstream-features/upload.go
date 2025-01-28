package main

import (
	"fmt"
	"os"

	"github.com/loft-sh/admin-apis/hack/internal/featuresyaml"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/entitlements/feature"
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

	for _, f := range features {
		exists, err := featureExists(f.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if !exists {
			_, err = feature.New(&stripe.EntitlementsFeatureParams{
				Name:      &f.DisplayName,
				LookupKey: &f.Name,
			})
			if err != nil {
				fmt.Printf("failed to create Stripe feature from feature %s: %v\n", f.Name, err)
				continue
			}
		}
		if
		_, err = feature.New(&stripe.EntitlementsFeatureParams{
			Name:      &f.DisplayName,
			LookupKey: &f.Name,
		})
		if err != nil {
			fmt.Printf("failed to create Stripe feature from feature %s: %v\n", f.Name, err)
			continue
		}
	}
	feature.Update()
}

func featureExists(id string) (bool, error) {
	_, err := feature.Get(id, &stripe.EntitlementsFeatureParams{})
	if err != nil {
		stripeErr, ok := err.(*stripe.Error)
		if !ok {
			return false, err
		}
		if stripeErr.Type != stripe.ErrorTypeInvalidRequest {
			return false, err
		}
		// going to assume here that the request is invalid because the id was not found, we may find
		// this check is not sufficient
		return false, nil
	}
	return true, nil
}
