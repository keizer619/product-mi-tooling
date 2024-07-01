/*
*  Copyright (c) WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
*
*  WSO2 LLC. licenses this file to you under the Apache License,
*  Version 2.0 (the "License"); you may not use this file except
*  in compliance with the License.
*  You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
 */

package add

import (
	"github.com/spf13/cobra"
	"github.com/wso2/product-mi-tooling/cmd/utils"
)

const AddCmdLiteral = "add"
const addCmdShortDesc = "Add new users, roles or loggers to a Micro Integrator instance"

const addCmdLongDesc = "Add new users, roles or loggers to a Micro Integrator instance in the environment specified by the flag (--environment, -e)"

var addCmdExamples = utils.MiCmdLiteral + " " + AddCmdLiteral + " " + "user" + " capp-developer -e dev\n" +
	utils.MiCmdLiteral + " " + AddCmdLiteral + " " + "log-level" + " synapse-api org.apache.synapse.rest.API DEBUG -e dev"

// AddCmd represents the add command
var AddCmd = &cobra.Command{
	Use:     AddCmdLiteral,
	Short:   addCmdShortDesc,
	Long:    addCmdLongDesc,
	Example: addCmdExamples,
	Run: func(cmd *cobra.Command, args []string) {
		utils.Logln(utils.LogPrefixInfo + AddCmdLiteral + " called")
		cmd.Help()
	},
}

func printAddCmdVerboseLog(cmd string) {
	utils.Logln(utils.LogPrefixInfo + AddCmdLiteral + " " + cmd + " called")
}
