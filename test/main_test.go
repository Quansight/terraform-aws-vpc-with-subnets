package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAwsVpcModule(t *testing.T) {
	t.Parallel()

	uniqueID := random.UniqueId()
	awsRegion := "us-east-1"
	vpcName := fmt.Sprintf("terratest-vpc-%s", uniqueID)

	terraformOptions := &terraform.Options{
		TerraformDir: "../",

		Vars: map[string]interface{}{
			"vpc_cidr_block":     "172.16.0.0/16",
			"region":             awsRegion,
			"availability_zones": []string{"us-east-1a", "us-east-1b"},
			"vpc_name":           vpcName,
		},

		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	vpcID := terraform.Output(t, terraformOptions, "vpc_id")
	assert.NotEmpty(t, vpcID)

	igwID := terraform.Output(t, terraformOptions, "igw_id")
	assert.NotEmpty(t, igwID)

	publicSubnetIDs := terraform.OutputList(t, terraformOptions, "public_subnet_ids")
	assert.GreaterOrEqual(t, len(publicSubnetIDs), 1)

	privateSubnetIDs := terraform.OutputList(t, terraformOptions, "private_subnet_ids")
	assert.GreaterOrEqual(t, len(privateSubnetIDs), 1)

	// Optionally, check VPC exists in AWS
	actualVpc := aws.GetVpcById(t, vpcID, awsRegion)
	assert.Equal(t, vpcID, actualVpc.Id)

	// Sleep to allow AWS eventual consistency
	time.Sleep(10 * time.Second)
}
