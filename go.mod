module github.com/KubeOpereator/terraform-provider-fusioncompute

go 1.14

require (
	github.com/KubeOperator/FusionComputeGolangSDK v0.0.0-20201014102018-fbfbaf81d7a8 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.0.1 // indirect
)

replace (
	github.com/KubeOperator/FusionComputeGolangSDK  => ../FusionComputeGolangSDK
)