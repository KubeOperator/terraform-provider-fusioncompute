module github.com/KubeOpereator/terraform-provider-fusioncompute

go 1.14

require (
	github.com/KubeOperator/FusionComputeGolangSDK v0.0.2
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.0.1
)

replace github.com/KubeOperator/FusionComputeGolangSDK => ../FusionComputeGolangSDK
