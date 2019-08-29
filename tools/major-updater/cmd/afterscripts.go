// Copyright 2018 Microsoft Corporation
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

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var afterscriptsCmd = &cobra.Command{
	Use:   "afterscripts <SDK dir>",
	Short: "Run afterscripts for SDK",
	Long: `This command will run the afterscripts in SDK repo, 
	including generation of profiles, and formatting the generated code`,
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.ExactArgs(1)(cmd, args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		sdk := args[0]
		err := theAfterscriptsCommand(sdk)
		return err
	},
}

func init() {
	rootCmd.AddCommand(afterscriptsCmd)
}

func theAfterscriptsCommand(sdk string) error {
	println("Generating profiles...")
	absolutePathOfSDK, err := filepath.Abs(sdk)
	if err != nil {
		return fmt.Errorf("failed to get the directory of SDK: %v", err)
	}
	// read options from config file
	file, err := os.Open(filepath.Join(absolutePathOfSDK, configFileName))
	if err != nil {
		return fmt.Errorf("failed to open config file %s: %v", configFileName, err)
	}
	afterscripts, err := expandAfterScripts(file)
	if err != nil {
		return err
	}
	err = changeDir(absolutePathOfSDK)
	if err != nil {
		return fmt.Errorf("failed to enter directory for profiles: %v", err)
	}
	for _, script := range afterscripts {
		args := strings.Split(script, " ")
		c := exec.Command(args[0], args[1:]...)
		vprintf("Invoke after script with args %v", c.Args)
		if output, err := c.CombinedOutput(); err != nil {
			return fmt.Errorf("failed to execute after script %s, messages %s: %v", script, string(output), err)
		}
	}
	return nil
	// c := exec.Command("go", "generate", "./...")
	// err = c.Run()
	// if err != nil {
	// 	return fmt.Errorf("Error occurs when generating profiles: %v", err)
	// }
	// vprintln("Formatting the whole SDK folder...")
	// err = changeDir(absolutePathOfSDK)
	// if err != nil {
	// 	return fmt.Errorf("failed to enter directory for SDK: %v", err)
	// }
	// c = exec.Command("gofmt", "-w", "./profiles/")
	// err = c.Run()
	// if err != nil {
	// 	return fmt.Errorf("Error occurs when formatting profiles: %v", err)
	// }
	// c = exec.Command("gofmt", "-w", "./services/")
	// err = c.Run()
	// if err != nil {
	// 	return fmt.Errorf("Error occurs when formatting the SDK folder: %v", err)
	// }
	// return nil
}

func expandAfterScripts(file *os.File) ([]string, error) {
	b, _ := ioutil.ReadAll(file)
	var config map[string]*json.RawMessage
	if err := json.Unmarshal(b, &config); err != nil {
		return nil, fmt.Errorf("failed to resolve config file: %v", err)
	}
	var meta map[string]*json.RawMessage
	if err := json.Unmarshal(*config["meta"], &meta); err != nil {
		return nil, fmt.Errorf("failed to resolve config file: %v", err)
	}
	var afterscripts []string
	if err := json.Unmarshal(*meta["after_scripts"], &afterscripts); err != nil {
		return nil, fmt.Errorf("failed to resolve config file: %v", err)
	}
	return afterscripts, nil
}
