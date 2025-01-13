package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/loft-sh/admin-apis/pkg/licenseapi"
)

const (
	fileTemplate = `package licenseapi
// This code was generated. Change features.yaml to add, remove, or edit features.

//go:generate go run ../../hack/gen-features/main.go

// Features
const (
%s
)

func GetFeatures() []FeatureName {
	return []FeatureName{
%s
	}
}
`
)

var (
	// the process that maps the hyphenated name to a camel cased name
	// assumes every hyphen delimited string should stay the same but with
	// the leading letter capitalized, e.g. custom-storage-driver, will be
	// CustomStorageDriver. Any exceptions to the aforementioned assumption
	// should be mapped here, e.g. vcluster-auth-sso can become VirtualClusterAuthSSO
	// by adding the mapping `"vcluster": "VirtualCluster"` and `"sso": "SSO"`.
	aliasLookup = map[string]string{
		"authentication": "Auth",
		"vcp":            "VirtualClusterPro",
		"vclusters":      "VirtualCluster",
		"vcluster":       "VirtualCluster",
		"ui":             "UI",
		"sso":            "SSO",
		"oidc":           "OIDC",
		"ha":             "HighAvailability",
		"coredns":        "CoreDNS",
		"cp":             "ControlPlane",
	}
	reg = regexp.MustCompile(`^([a-zA-Z]+)|(-[a-zA-Z]+)`)
)

func main() {
	featuresYaml, err := os.Open("../../pkg/licenseapi/features.yaml")
	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(featuresYaml)
	if err != nil {
		panic(err)
	}

	features := struct {
		Features []licenseapi.Feature `json:"features"`
	}{}
	err = yaml.Unmarshal(bytes, &features)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("../../pkg/licenseapi/features.go")
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(fmt.Sprintf(fileTemplate, generateFeatureConstantsBody(features.Features), generateFeatureSliceBody(features.Features)))
	if err != nil {
		panic(err)
	}
}

func generateFeatureConstantsBody(features []licenseapi.Feature) string {
	featureConstants := ""
	for _, feature := range features {
		featureConstants += fmt.Sprintf(`	%s FeatureName = "%s" // %s

`, hyphenatedToCamelCase(replaceAliasWithFull(feature.Name)), feature.Name, feature.DisplayName)
	}
	return strings.TrimSuffix(featureConstants, "\n")
}

func generateFeatureSliceBody(features []licenseapi.Feature) string {
	featuresList := ""
	for _, feature := range features {
		featuresList += fmt.Sprintf(`		%s,
`, hyphenatedToCamelCase(replaceAliasWithFull(feature.Name)))
	}
	return strings.TrimSuffix(featuresList, "\n")
}

func replaceAliasWithFull(feature string) string {
	for alias, full := range aliasLookup {
		if feature == alias {
			return full
		}
		cutFeature, ok := strings.CutPrefix(feature, alias+"-")
		if ok {
			feature = full + "-" + cutFeature
		}
		cutFeature, ok = strings.CutSuffix(feature, "-"+alias)
		if ok {
			feature = cutFeature + "-" + full
		}
		feature = strings.ReplaceAll(feature, "-"+alias+"-", "-"+full+"-")
	}
	return feature
}

func hyphenatedToCamelCase(name string) string {
	return reg.ReplaceAllStringFunc(name, func(s string) string {
		return strings.ToUpper(string(strings.TrimPrefix(s, "-")[0])) + strings.TrimPrefix(s, "-")[1:]
	})
}
