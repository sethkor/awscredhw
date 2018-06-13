package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"fmt"
)

var (
	pProfile	= kingpin.Flag("profile", "AWS credentials/config file profile to use").Short('p').String()
)

func main () {

	var svc *sts.STS

	kingpin.Parse()

	if *pProfile != "" {

		awskeyprofile := *pProfile
		fmt.Println("using profile passed and aws credentials/config files")
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			Profile:           awskeyprofile,
			SharedConfigState: session.SharedConfigEnable,
		}))
		svc = sts.New(sess)

	} else {
		fmt.Println("using default provider chain")
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
		svc = sts.New(sess)
	}//else

	result,err := svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)


}