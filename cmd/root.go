/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/shigetaichi/prext/utils"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "prext",
	Short: "Prepare for Next.js",
	Long: `
This is a Command Line Interface for set up Next.js environment!
Next.js has a lot of options and ecosystem has many useful libraries.
So, you can use this CLI "PREXT"!

Moreover, my favorite environment can be completely recreated with a single option.
My favorite environments grow as the Next.js ecosystem changes, and as my likes and dislikes evolve. In other words, this is a command that grows 😂.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		yarnFlag := true
		if utils.YesOrNo(fmt.Sprintf("現在のディレクトリ(%s)にプロジェクトを展開しますか？[y/n]", utils.Pwd())) {
			fmt.Println()
			if utils.YesOrNo("yarn(recommended)[Y] or npm? [n]") {
				err := exec.Command("npx", "create-next-app", ".").Run()
				if err != nil {
					panic(err)
				}
			} else {
				yarnFlag = false
				err := exec.Command("npx", "create-next-app", ".", "--use=npm").Run()
				if err != nil {
					panic(err)
				}
			}
		} else {
			dirN := ""
			fmt.Printf("プロジェクトのディレクトリ名(%s/〇〇)を入力してください。", utils.Pwd())
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				if scanner.Text() != "" {
					dirN = scanner.Text()
					break
				}
			}
			if utils.YesOrNo("yarn(recommended)(y) or npm(n)? [y/n]") {
				err := exec.Command("npx", "create-next-app", dirN).Run()
				if err != nil {
					panic(err)
				}
			} else {
				yarnFlag = false
				err := exec.Command("npx", "create-next-app", dirN, "--use=npm").Run()
				if err != nil {
					panic(err)
				}
			}
		}

		err := exec.Command("mkdir", "src")
		if err != nil {
			panic(err)
		}
		err = exec.Command("mv", "pages/", "src/pages")
		if err != nil {
			panic(err)
		}
		err = exec.Command("mv", "styles/", "src/styles")
		if err != nil {
			panic(err)
		}

		err = exec.Command(utils.YarnOrNpm(yarnFlag), "add", "-D", "typescript", "@types/react", "@types/react-dom", "@types/node")
		if err != nil {
			panic(err)
		}

		//find, err = exec.Command("find","src/pages","-name","_app.js","-or -name","index.js").Output()
		//set, err = exec.Command( ""| sed 'p;s/.js$/.tsx/' | xargs -n2 mv & \\\n  find src/pages/api -name \"*.js\" | sed 'p;s/.js$/.ts/' | xargs -n2 mv")

		fmt.Println("終了します。")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.prext.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
