/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pvillela/go-tryout/arch/errx"
	"github.com/pvillela/go-tryout/web"
	"net/http"
	"os"
	"time"
)

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"-"`
}

func processRequests() {
	callServer := func(path string, payload []byte) {
		fmt.Println()
		fmt.Println("payload:", string(payload))
		resp, err := http.Post("http://localhost:8080"+path, "application/json",
			bytes.NewBuffer(payload))
		errx.PanicOnError(err)

		var res map[string]any
		err = json.NewDecoder(resp.Body).Decode(&res)
		errx.PanicOnError(err)
		fmt.Println((resp.Status))
		fmt.Println("response:", res)
	}

	{
		path := "/loginJSON"
		payload := []byte(`{ "user": "manu"}`)
		callServer(path, payload)
	}

	{
		path := "/loginJSON"
		payload := []byte(`{ "user": "manu", "password": "123" }`)
		callServer(path, payload)
	}
}

func main() {
	router := gin.Default()

	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var input Login
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if input.User != "manu" || input.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding XML (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>manu</user>
	//		<password>123</password>
	//	</root>)
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding a HTML form (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Launch the server
	go func() {
		// Listen and serve on 0.0.0.0:8080
		err := router.Run(":8080")
		errx.PanicOnError(err)
	}()

	// Wait for server to be ready
	err := web.WaitForHttpServer("http://localhost:8080/", 100*time.Millisecond)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("***** Server ready")

	processRequests()

	fmt.Println("Exiting")
}
