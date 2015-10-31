// Copyright 2015 Huawei Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	// "fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/huawei-openlab/oct/tools/runtimeValidator/config"
)

func init() {
	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum)
}

var Runtime string

func main() {
	app := cli.NewApp()
	app.Name = "oci-runtimeValidator"
	app.Version = "0.0.1"
	app.Usage = "Utilities for OCI runtime validation"
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "runtime",
			Value: "runc",
			Usage: "runtime to be validated",
		},
	}
	app.Action = func(c *cli.Context) {
		startTime := time.Now()
		Runtime = c.String("runtime")
		wg := new(sync.WaitGroup)
		wg.Add(config.ConfigLen)

		for key, value := range config.BundleMap {
			go validateRoutine(key, value, wg)

		}
		wg.Wait()
		logrus.Printf("Test runtime: %v, successed\n", Runtime)

		endTime := time.Now()
		dTime := endTime.Sub(startTime)
		logrus.Debugf("Cost time: %v\n", dTime.Nanoseconds())
	}

	logrus.SetLevel(logrus.InfoLevel)
	// logrus.SetLevel(logrus.DebugLevel)

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func validateRoutine(bundleName string, generateArgs string, wg *sync.WaitGroup) {
	logrus.Debugf("Test bundle name: %v, Test args: %v\n", bundleName, generateArgs)
	err := validate(bundleName, generateArgs)
	if err != nil {
		logrus.Fatal(err)
	} else {
		logrus.Printf("Test runtime: %v bundle: %v, args: %v, successed\n", Runtime, bundleName, generateArgs)
	}
	wg.Done()
	return
}
