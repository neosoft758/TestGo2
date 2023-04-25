package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "github.com/awslabs/aws-sdk-go/aws"
  "github.com/awslabs/aws-sdk-go/gen/ec2"
)

func main() {

  creds := aws.IAMCreds()
  cli := ec2.New(creds,getRegion(),nil)
  resp, err := cli.DescribeInstances(nil)
  if err != nil {
    panic(err)
  }

  if len(resp.Reservations) < 1 {
    return
  }

  for i := range resp.Reservations {
    fmt.Print(i)
    fmt.Print(" , ")
    fmt.Print(*resp.Reservations[i].Instances[0].VPCID)
    fmt.Print(" , ")
    fmt.Print(*resp.Reservations[i].Instances[0].SubnetID)
    fmt.Print(" , ")
    fmt.Print(*resp.Reservations[i].Instances[0].InstanceID)
    fmt.Print(" , ")
    fmt.Print(*resp.Reservations[i].Instances[0].InstanceType)
    fmt.Print(" , ")
    fmt.Println(*resp.Reservations[i].Instances[0].ImageID)
  }
}

func getRegion()(string){
  url := "http://169.254.169.254/latest/meta-data//placement/availability-zone"
  r,_ := http.Get(url)
  defer r.Body.Close()
  byteArray, _ := ioutil.ReadAll(r.Body)
  str := string(byteArray)
  size := len(str)
  return str[0:size-1]
}