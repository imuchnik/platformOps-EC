package main

import (
	"github.com/BurntSushi/toml"
	"os"
	"os/exec"
	"testing"
	"strings"
)

func TestCommandExecutionWithVariables(t *testing.T) {

	os.Setenv("BUILD_ID", "123")

	out, _ := exec.Command("echo", os.ExpandEnv("$BUILD_ID")).Output()

	if strings.Compare(strings.TrimSuffix(string(out), "\n"), "123") != 0 {

		t.Errorf("expected 123 got %s ", string(out))

	}
}

func TestGoWitnessFromInside(t *testing.T) {

	//var chrome chrm.Chrome
	//var db storage.Storage
	//var waitTimeout int
	//dbLocation := "gowitness.db"
	////var rootCmd = &cobra.Command{Use: "app",  TraverseChildren: true,}
	//
	//var rootCmd = &cobra.Command{
	//	Use:          "gowitness",
	//	SilenceUsage: true,
	//	Short:        "A commandline web screenshot and information gathering tool by @leonjza",
	//	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	//
	//		// Init Google Chrome
	//		chrome = chrm.Chrome{
	//			Resolution:     "80",
	//			ChromeTimeout:  3,
	//			Path:           "/usr/bin/google-chrome",
	//			ScreenshotPath: ".",
	//		}
	//		chrome.Setup()
	//
	//		// Setup the destination directory
	//
	//		// open the database
	//		db = storage.Storage{}
	//		db.Open(dbLocation)
	//	},
	//	TraverseChildren: true,
	//}
	//
	//var singleCmd = &cobra.Command{
	//	Use:   "single",
	//	Short: "Take a screenshot of a single URL",
	//	Long: `
	//			Takes a screenshot of a single given URL and saves it to a file.
	//			If no --destination is provided, a filename for the screenshot will
	//			be automatically generated based on the given URL.
	//
	//			For example:
	//
	//			$ gowitness single --url https://twitter.com
	//			$ gowitness single --destination tweeps_page.png --url https://twitter.com
	//			$ gowitness single -u https://twitter.com`,
	//
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println(args)
	//
	//		u, err := url.ParseRequestURI(args[1])
	//		if err != nil {
	//			fmt.Println("Invalid URL specified", string(err.Error()))
	//		}
	//
	//		// Process this URL
	//		utils.ProcessURL(u, &chrome, &db, waitTimeout)
	//
	//	},
	//}
	//
	//rootCmd.AddCommand(singleCmd)
	//screenshotURL := "https://google.com"
	////args := []string{"single", "https://google.com"}
	//singleCmd.Flags().StringVarP(&screenshotURL, "url", "u", "", "The URL to screenshot")
	//singleCmd.ExecuteC()

	exec.Command("gowitness", "single", "-u", "https://consumerfinance.gov").Output()
}

func TestLoadConfigIntoSession(t *testing.T) {

	var config map[string]string
	configFile := "test-data/ec-config.toml"

	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal(err)
	}

	loadConfig(configFile)

	for k, v := range config {
		if os.Getenv(k) != v {
			t.Errorf("expected session value %s for key %s, but got %s ", v, k, os.Getenv(k))

		}
	}

}
