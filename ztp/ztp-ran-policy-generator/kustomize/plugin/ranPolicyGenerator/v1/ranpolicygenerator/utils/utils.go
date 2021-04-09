package utils

const NotApplicable = "N/A"
const FileExt = ".yaml"
const CommonNS = "common-ran-subscriptions"
const GroupNS = "group-ran-subscriptions"
const SiteNS = "site-ran-subscriptions"

type RanGenConfig struct {
	SourcePoliciesPath string
	RanGenTempPath string
	OutPath string
	GenACM bool
	GenK8sRes bool
	Printstdout bool
}


type RanGenTemplate struct {
	ApiVersion string  `yaml:"apiVersion"`
	Kind string `yaml:"kind"`
	Metadata metaData `yaml:"metadata"`
	SourceFiles []sourceFile `yaml:"sourceFiles"`

}

type metaData struct {
	Name string `yaml:"name"`
	Labels labels `yaml:"labels"`
	Namespace string `yaml:"namespace"`
	Annotations []string `yaml:"annotations"`
}

type labels struct {
	Common bool  `yaml:"common"`
	GroupName string  `yaml:"groupName"`
	SiteName string  `yaml:"siteName"`
	Mcp string  `yaml:"mcp"`
}

type sourceFile struct {
	FileName string `yaml:"fileName"`
	Name string  `yaml:"name"`
	Spec map[string]interface{} `yaml:"spec"`
}

type AcmPolicy struct {
	ApiVersion string  `yaml:"apiVersion"`
	Kind string `yaml:"kind"`
	Metadata metaData `yaml:"metadata"`
	Spec acmPolicySpec `yaml:"spec"`
}

type acmPolicySpec struct {
	RemediationAction string `yaml:"remediationAction"`
	Disabled bool `yaml:"disabled`
	PolicyTemplates []PolicyObjectDefinition `yaml:"policy-templates"`
}

type PolicyObjectDefinition struct {
	ObjDef []AcmConfigurationPolicy `yaml:"objectDefinition"`
}

type AcmConfigurationPolicy struct {
	ApiVersion string  `yaml:"apiVersion"`
	Kind string `yaml:"kind"`
	Metadata metaData `yaml:"metadata"`
	Spec acmConfigPolicySpec `yaml:"spec"`
}

type acmConfigPolicySpec struct {
	RemediationAction string `yaml:"remediationAction"`
	Severity string `yaml:"severity"`
	NamespaceSelector struct {
		Exclude []string `yaml:"exclude"`
		Include []string `yaml:"include"`
	}
	ObjectTemplates []ObjectTemplates `yaml:"object-templates"`
}

type ObjectTemplates struct {
	ComplianceType string `yaml:"complianceType"`
	ObjectDefinition map[string]interface{} `yaml:"objectDefinition"`
}
