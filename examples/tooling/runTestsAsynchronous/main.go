package main

import (
	"fmt"
	"os"

	"github.com/jpmonette/force"
	"golang.org/x/oauth2"
)

func main() {

	conf := &oauth2.Config{
		ClientID:     os.Getenv("SFCLIENTID"),
		ClientSecret: os.Getenv("SFCLIENTSECRET"),
		Endpoint: oauth2.Endpoint{
			TokenURL: os.Getenv("SFTOKENURL"),
		},
	}

	token, _ := conf.PasswordCredentialsToken(oauth2.NoContext, os.Getenv("SFUSERNAME"), os.Getenv("SFPASSWORD"))

	instanceUrl, _ := token.Extra("instance_url").(string)
	client := conf.Client(oauth2.NoContext, token)

	c, _ := force.NewClient(client, instanceUrl)

	result, _ := c.Tooling.RunTestsAsynchronous(nil, nil, "", "RunLocalTests")

	// Output: 707E000004Btvxf
	fmt.Println(result)
}
