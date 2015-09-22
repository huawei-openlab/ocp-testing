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
	"flag"
	"fmt"
	// _ "github.com/huawei-openlab/oct/tools/specsValidator/cases/linuxcapabilities"
	// _ "github.com/huawei-openlab/oct/tools/specsValidator/cases/linuxnamespace"
	// _ "github.com/huawei-openlab/oct/tools/specsValidator/cases/linuxresources"
	// _ "github.com/huawei-openlab/oct/tools/specsValidator/cases/linuxrlimits"
	// _ "github.com/huawei-openlab/oct/tools/specsValidator/cases/linuxselinuxlabel"
	// _ "github.com/huawei-openlab/oct/tools/specsValidator/cases/linuxsysctl"
	//_ "github.com/huawei-openlab/oct/tools/specsValidator/cases/specmount"
	// _ "github.com/huawei-openlab/oct/tools/specsValidator/cases/specplatform"
	// _ "github.com/huawei-openlab/oct/tools/specsValidator/cases/specprocess"
	_ "github.com/huawei-openlab/oct/tools/specsValidator/cases/specroot"
	_ "github.com/huawei-openlab/oct/tools/specsValidator/cases/specversion"
	"github.com/huawei-openlab/oct/tools/specsValidator/hostenv"
	"github.com/huawei-openlab/oct/tools/specsValidator/manager"
	"github.com/huawei-openlab/oct/tools/specsValidator/utils"
	"log"
)

var specsRev = flag.String("specs", "", "Specify specs Revision from opencontainers/specs as the benchmark, in the form of commit id, keep empty to using the newest commit of [opencontainers/specs](https://github.com/opencontainers/specs)")
var runcRev = flag.String("runc", "", "Specify runc Revision from opencontainers/specs to be tested, in the form of commit id, keep empty to using the newest commit of [opencontainers/runc](https://github.com/opencontainers/runc")
var output = flag.String("o", "./report/", "Specify filePath to install the test result linuxspec.json")
var runctags = flag.String("rtags", "seccomp", "Build tags for runc, should be one of seccomp/selinux/apparmor, keep empty to using seccomp")

func main() {

	flag.Parse()
	fmt.Println("Testing Revision:")
	var checkoutSpecsRev, checkoutRuncRev, runcBuildtags string
	if *specsRev == "predraft" {
		checkoutSpecsRev = "45ae53d4dba8e550942f7384914206103f6d2216"
		fmt.Printf("	Specs revision: %v \n", checkoutSpecsRev)
	} else if *specsRev == "" {
		checkoutSpecsRev = "origin/master"
	} else {
		checkoutSpecsRev = *specsRev
		fmt.Printf("	Specs revision: %v \n", checkoutSpecsRev)
	}

	if *runcRev == "" {
		checkoutRuncRev = "origin/master"
		fmt.Printf("	Runc revision :  %v\n", checkoutRuncRev)
	} else {
		checkoutRuncRev = *runcRev
		fmt.Printf("	Runc revision: %v \n", checkoutRuncRev)
	}

	if *runctags == "" {
		runcBuildtags = "seccomp"
	} else if *runctags != "seccomp" && *runctags != "selinux" && *runctags != "apparmor" {
		log.Fatalf("Parameter runctags=%v is the wrong value", *runctags)
	} else {
		runcBuildtags = *runctags
	}

	hostenv.UpateSpecsRev(checkoutSpecsRev)
	hostenv.UpateRuncRev(checkoutRuncRev, runcBuildtags)

	fmt.Println("Testing output: ")
	if *output == "" {
		*output = "./report/"
		fmt.Println("	oct/tools/specsValidator/report/linuxspec.json")
	} else {
		fmt.Printf("	%v\n", *output)
	}

	err := hostenv.CreateBoundle()
	if err != nil {
		log.Fatalf("Create boundle error, %v", err)
	}

	/*linuxnamespace.TestSuiteNP.Run()
	result := linuxnamespace.TestSuiteNP.GetResult()

	err = utils.StringOutput("namespace_out.json", result)
	if err != nil {
		log.Fatalf("Write namespace out file error,%v\n", err)
	}
	*/
	// // spec.version test
	/*specversion.TestSuiteVersion.Run()
	result = specversion.TestSuiteVersion.GetResult()

	err = utils.StringOutput("Version_out.json", result)
	if err != nil {
		log.Fatalf("Write version out file error,%v\n", err)
	}*/

	// // spec.mount test
	/*specmount.TestSuiteMount.Run()
	result = specmount.TestSuiteMount.GetResult()
	err = utils.StringOutput("Mount_out.json", result)
	if err != nil {
		log.Fatalf("Write mount out file error,%v\n", err)
	}*/
	// manager * TestManager = new(TestManager)

	for _, ts := range manager.Manager.TestSuite {
		ts.Run()
		result := ts.GetResult()
		outputJson := ts.Name + ".json"
		err := utils.StringOutput(outputJson, result)
		if err != nil {
			log.Fatalf("Write %v out file error,%v\n", ts.Name, err)
		}

	}
	result := manager.Manager.GetTotalResult()
	err = utils.SpecifyOutput(*output, result)
	if err != nil {
		log.Fatalf("Write %v out file error,%v\n", *output, err)
	}
	/*
		specroot.TestSuiteRoot.Run()
		result := specroot.TestSuiteRoot.GetResult()
		err = utils.StringOutput("Root_out.json", result)
		if err != nil {
			log.Fatalf("Write Root out file error,%v\n", err)
		}*/

	/*specplatform.TestSuitePlatform.Run()
	result = specplatform.TestSuitePlatform.GetResult()
	err = utils.StringOutput("Platform_out.json", result)
	if err != nil {
		log.Fatalf("Write Platform out file error,%v\n", err)
	}
	*/
	//linux resources test
	/*linuxresources.TestSuiteLinuxResources.Run()
	result := linuxresources.TestSuiteLinuxResources.GetResult()
	err = ioutil.WriteFile("LinuxResources_out.json", []byte(result), 0777)
	if err != nil {
		log.Fatalf("Write LinuxResources out file error,%v\n", err)
	}*/

}
