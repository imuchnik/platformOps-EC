package converter

import (
	"flag"
	"fmt"
	"os"
	"platformOps-EC/services"
	"platformOps-EC/models"
	"log"
	"database/sql"
	"bytes"
	"io/ioutil"
	"encoding/json"
)
var excelFileName, output string
var excelFileName, configFile string


type Config struct {
	Dbname   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"-"`
	Sslmode  string `json:"sslmode"`
	Location string `json:"location"`
	Schema   string `json:"currentSchema"`
}
func getConnStr(config Config) string {
	var buffer bytes.Buffer
	buffer.WriteString("postgres://")
	buffer.WriteString(config.Username)
	buffer.WriteString(":")
	buffer.WriteString(config.Password)
	buffer.WriteString("@")
	buffer.WriteString(config.Location)
	buffer.WriteString("/")
	buffer.WriteString(config.Dbname)
	buffer.WriteString("?sslmode=")
	buffer.WriteString(config.Sslmode)

	fmt.Println(buffer.String())

	return buffer.String()
}

func setSearchPath(db *sql.DB, schema string) {

	sqlStatement := "SET search_path TO " + schema

	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func getConfig(configFile string) Config {
	raw, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Config
	errj := json.Unmarshal(raw, &c)
	if errj != nil {
		fmt.Println("error parsing json input", err)
	}
	return c[0]
}


func ToJson()  {



	if excelFileName == "" {
		fmt.Println("Missing input Excel file. Program will exit.")
		os.Exit(1)
	}

	if output == "manifest.json" {
		fmt.Println("Default to manifest.json")

	}

	fmt.Println("Loading Excel file ", excelFileName)

	baseline, controls := services.LoadFromExcel(excelFileName)
	var manifest []models.ECManifest

	fmt.Println("Converting to Json object")

	for _, c := range controls {

		m := models.ECManifest{ReqId: c.ReqId, Title: c.Category,
			Baseline: baseline.Name}
		manifest = append(manifest, m)

	}

	//fmt.Println(models.ToJson(manifest))

	file, err := os.Create(output)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Println("Writing Json Object to file")

	fmt.Fprintf(file, "%v", models.ToJson(manifest))

	fmt.Printf("Done writing to output file at [%v]\n", output)


}

func ToSql() {


	if excelFileName == "" {
		fmt.Println("Missing input excel baseline. Program will exit.")
		os.Exit(1)
	}

	if configFile == "" {
		fmt.Println("Missing configuration file. Program will exit.")
		os.Exit(1)
	}

	fmt.Println("Loading Excel file ", excelFileName)

	baseline, controls := services.LoadFromExcel(excelFileName)

	fmt.Println("Loading config file")

	config := getConfig(configFile)

	fmt.Println("Connecting to database ")

	connStr := getConnStr(config)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Set to schema [%v]\n", config.Schema)
	setSearchPath(db, config.Schema)

	fmt.Println("Inserting Baseline")

	baseline_id := services.InsertBaseline(db, baseline)

	services.ReadBaselineAll(db)

	fmt.Println("Inserting controls")
	for i := 0; i < len(controls); i++ {
		controls[i].SetBaselineId(baseline_id)
		services.InsertControl(db, controls[i])

	}

	//services.ReadControlByBaselineId(db, baseline_id)
	fmt.Println("Done inserting Baseline and Controls.  Check DB")
}
 func main(){

	 flag.StringVar(&excelFileName, "i", "", "Input Excel baseline file. If missing, program will exit.")
	 flag.StringVar(&output, "o", "manifest.json", "Execution output location.")

	 flag.StringVar(&excelFileName, "i", "", "Input excel baseline file. If missing, program will exit.")
	 flag.StringVar(&configFile, "c", "", "Configuration file. If missing, program will exit.")

	 flag.Parse()

 }