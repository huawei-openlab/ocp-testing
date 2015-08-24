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
	"io/ioutil"
	"log"

	"github.com/huawei-openlab/oct/cases/specstest/cases/linuxnamespace"
	"github.com/huawei-openlab/oct/cases/specstest/hostenv"
)

func main() {

	err := hostenv.SetupEnv("", "")
	if err != nil {
		log.Fatalf(" Pull image error, %v", err)
	}
	linuxnamespace.TestSuiteNP.Run()
	result := linuxnamespace.TestSuiteNP.GetResult()

	err = ioutil.WriteFile("namespace_out.json", []byte(result), 0777)
	if err != nil {
		log.Fatalf("Write file error,%v\n", err)
	}

}