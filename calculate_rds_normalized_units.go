// ==============================================================================
// Author : Mohamed Tanveer (tanveer.munavar@gmail.com)
// Usage : go run rds_instance_list.go -region ap-south-1 -instance db.m5
// ==============================================================================
//

package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

func main() {

	// arguments passed in command line
	region_v := flag.String("region", "ap-south-1", "mention the region")
	instance_type := flag.String("instance", "db.m5", "mention the instance type, ex : db.m5.large")
	flag.Parse()

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(*region_v)},
	)

	// Create RDS service client
	svc := rds.New(sess)

	result, err := svc.DescribeDBInstances(nil)
	if err != nil {
		exitErrorf("Unable to list instances, %v", err)
	}

	var total_rds_normalized_units int

	for _, d := range result.DBInstances {

		// print only m5 series instances
		match, _ := regexp.MatchString(*instance_type, string(aws.StringValue(d.DBInstanceClass)))

		// to print
		// fmt.Printf("\n %v",match)

		if match {

			var maz int
			if aws.BoolValue(d.MultiAZ) != false {
				maz = 2
			} else {
				maz = 1
			}

			// fmt.Printf("maz : %v \n",maz)

			instance_size_arr := strings.SplitN(aws.StringValue(d.DBInstanceClass), ".", 3)
			instance_size := fmt.Sprint(instance_size_arr[2])

			var normalized_units int
			var total_normalized_units int

			switch instance_size {
			case "small":
				normalized_units = 1
			case "medium":
				normalized_units = 2
			case "large":
				normalized_units = 4
			case "xlarge":
				normalized_units = 8
			case "2xlarge":
				normalized_units = 16
			case "4xlarge":
				normalized_units = 32
			default:
				panic("unrecognized value")
			}

			total_normalized_units = maz * normalized_units

			total_rds_normalized_units += total_normalized_units

		}

	} // end of for loop

	fmt.Printf("total_rds_normalized_units : %v \n", total_rds_normalized_units)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

// end of script

