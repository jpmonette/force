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

	var result Response

	c.Tooling.Query("SELECT FullName, Description FROM Profile ORDER BY Id ASC", &result)

	fmt.Println(result.Records[0].FullName)
	// Output: Admin
}

// Response is a Tooling API Query response structure
type Response struct {
	Records []struct {
		FullName string `json:"FullName"`
	} `json:"records"`
}
