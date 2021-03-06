package provider

type AwsProvider struct {
	region string
}

func (a *AwsProvider) Endpoint() string {
	var endpoint string
	if a.region == "us-east-1" {
		endpoint = "s3.amazonaws.com"
	} else {
		endpoint = "s3-" + a.region + ".amazonaws.com"
	}
	return endpoint
}
