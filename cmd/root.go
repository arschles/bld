package cmd

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var imageName string
var dockerFileLoc string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bld",
	Short: "Build Docker images locally or in the cloud",
	// 	Long: `A longer description that spans multiple lines and likely contains
	// examples and usage of using your application. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println(`Build context is missing. Please run like so: 
				bld -t my/image -f my/Dockerfile my/build/context`)
			os.Exit(1)
		}
		buildContextLoc := args[0]
		if imageName == "" {
			fmt.Println("Image name is missing. This is the value you pass to the '-t' argument")
			os.Exit(1)
		}
		if dockerCLIExists() {
			fmt.Println("'docker' CLI Exists, executing locally")
			ran, err := sh.Exec(
				map[string]string{},
				os.Stdout,
				os.Stderr,
				"docker",
				"build",
				"-t",
				imageName,
				"-f",
				dockerFileLoc,
				buildContextLoc,
			)
			if err != nil {
				fmt.Println("Error:", err)
				if sh.ExitStatus(err) != 0 {
					os.Exit(sh.ExitStatus(err))
				}
				os.Exit(1)
			}
			if !ran {
				fmt.Println("thing didn't run :(")
				os.Exit(1)
			}
		} else if azCLIExists() {
			fmt.Println("'az' CLI exists, executing in the cloud")
			ran, err := sh.Exec(
				map[string]string{},
				os.Stdout,
				os.Stderr,
				"az",
				"acr",
				"build",
				"-t",
				imageName,
				"-f",
				dockerFileLoc,
				buildContextLoc,
			)
			if err != nil {
				if err != nil {
					fmt.Println("Error:", err)
					if sh.ExitStatus(err) != 0 {
						os.Exit(sh.ExitStatus(err))
					}
					os.Exit(1)
				}
				if !ran {
					fmt.Println("thing didn't run :(")
					os.Exit(1)
				}
			}
		} else {
			fmt.Println("The 'docker' binary and the 'az' binary both don't exist :(")
			os.Exit(1)
		}
		fmt.Println("Image", imageName, "is built!")
		os.Exit(0)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bld.yaml)")
	rootCmd.PersistentFlags().StringVarP(&imageName, "image-name", "t", "./Dockerfile", "the name of the Docker image to create. this field is required")
	rootCmd.PersistentFlags().StringVarP(&dockerFileLoc, "dockerfile", "f", ".", "the location of the Dockerfile. Defaults to '.'")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".bld" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".bld")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
