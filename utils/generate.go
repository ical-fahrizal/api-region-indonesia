package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type CsvDataProvinceLines struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var ArrProvinces []CsvDataProvinceLines

type CsvDataRegencieLines struct {
	Id           int    `json:"id"`
	ProvinceId   int    `json:"provinceId"`
	ProvinceName string `json:"provinceName"`

	Name string `json:"name"`
}

var ArrRegencies []CsvDataRegencieLines
var ArrRegenciesFile []CsvDataRegencieLines

type CsvDataDistrictLines struct {
	Id           int    `json:"id"`
	RegencieId   int    `json:"regencieId"`
	RegencieName string `json:"regencieName"`
	Name         string `json:"name"`
}

var ArrDistricts []CsvDataDistrictLines
var ArrDistrictsFile []CsvDataDistrictLines

type CsvDataVillageLines struct {
	Id           int    `json:"id"`
	DistrictId   int    `json:"districtId"`
	DistrictName string `json:"districtName"`
	Name         string `json:"name"`
	PostalCode   string `json:"postalCode"`
}

var ArrVillages []CsvDataVillageLines
var ArrVillagesFile []CsvDataVillageLines

type CsvDataPostalCodeLines struct {
	Id           int    `json:"id"`
	VillageName  string `json:"villageName"`
	DistrictName string `json:"districtName"`
	RegencieName string `json:"regencieName"`
	ProvinceId   int    `json:"provinceId"`
	PostalCode   string `json:"postalCode"`
}

var ArrPostalCodes []CsvDataPostalCodeLines
var ArrPostalCodesFile []CsvDataPostalCodeLines

func Generate() {
	//Provices
	dataProvinces, err := readCsvFile("./data/1/provinces.csv")
	if err != nil {
		panic(err)
	}

	for _, line := range dataProvinces {
		id, _ := strconv.Atoi(line[0])
		ArrProvinces = append(ArrProvinces, CsvDataProvinceLines{Id: id, Name: line[1]})
	}

	//Regencies
	dataRegencies, err := readCsvFile("./data/1/regencies.csv")
	if err != nil {
		panic(err)
	}

	for _, line := range dataRegencies {
		id, _ := strconv.Atoi(line[0])
		provinceId, _ := strconv.Atoi(line[1])
		ArrRegencies = append(ArrRegencies, CsvDataRegencieLines{Id: id, ProvinceId: provinceId, Name: line[2]})

	}

	//Districts
	dataDistricts, err := readCsvFile("./data/1/districts.csv")
	if err != nil {
		panic(err)
	}

	for _, line := range dataDistricts {
		id, _ := strconv.Atoi(line[0])
		regencieId, _ := strconv.Atoi(line[1])
		ArrDistricts = append(ArrDistricts, CsvDataDistrictLines{Id: id, RegencieId: regencieId, Name: line[2]})
	}

	//Postal Code
	dataPostalCode, err := readCsvFile("./data/1/postalcode.csv")
	if err != nil {
		panic(err)
	}

	for _, line := range dataPostalCode {
		id, _ := strconv.Atoi(line[0])
		provinceId2, _ := strconv.Atoi(line[4])
		ArrPostalCodes = append(ArrPostalCodes, CsvDataPostalCodeLines{Id: id, VillageName: line[1], DistrictName: line[2], RegencieName: line[3], ProvinceId: provinceId2, PostalCode: line[5]})
	}

	//Villages
	dataVillages, err := readCsvFile("./data/1/villages.csv")
	if err != nil {
		panic(err)
	}

	for _, line := range dataVillages {
		id, _ := strconv.Atoi(line[0])
		districtId, _ := strconv.Atoi(line[1])
		postalCode := ""

		for _, v2 := range ArrPostalCodes {
			if v2.VillageName == line[2] {
				postalCode = v2.PostalCode
				break
			}
		}
		ArrVillages = append(ArrVillages, CsvDataVillageLines{Id: id, DistrictId: districtId, PostalCode: postalCode, Name: line[2]})
	}

	//Create Json File Provinces
	log.Printf("Create json File Provinces")
	file, _ := json.MarshalIndent(ArrProvinces, "", " ")
	_ = ioutil.WriteFile("./output/provinces.json", file, 0644)

	//Create json File Regencies
	log.Printf("Create json File Regencies")
	for _, v1 := range ArrProvinces {
		ArrRegenciesFile = nil
		for _, v2 := range ArrRegencies {
			if v2.ProvinceId == v1.Id {
				ArrRegenciesFile = append(ArrRegenciesFile, CsvDataRegencieLines{Id: v2.Id, ProvinceId: v2.ProvinceId, ProvinceName: v1.Name, Name: v2.Name})
			}
		}
		file, _ = json.MarshalIndent(ArrRegenciesFile, "", " ")
		_ = ioutil.WriteFile(fmt.Sprintf(`./output/%v.json`, v1.Id), file, 0644)

		time.Sleep(10 * time.Millisecond)
		// log.Printf("Selesai sleep.....")
	}

	log.Printf("masuk 1")
	//Create json File Districts
	log.Printf("Create json File Districts")
	for _, v1 := range ArrRegencies {
		ArrDistrictsFile = nil
		for _, v2 := range ArrDistricts {
			if v2.RegencieId == v1.Id {
				ArrDistrictsFile = append(ArrDistrictsFile, CsvDataDistrictLines{Id: v2.Id, RegencieId: v2.RegencieId, RegencieName: v1.Name, Name: v2.Name})
			}
		}
		file, _ := json.MarshalIndent(ArrDistrictsFile, "", " ")
		_ = ioutil.WriteFile(fmt.Sprintf(`./output/%v.json`, v1.Id), file, 0644)
		time.Sleep(10 * time.Millisecond)
		// log.Printf("Selesai sleep.....")
	}

	//Create json File Villages
	log.Printf("Create json File Villages")
	for _, v1 := range ArrDistricts {
		ArrVillagesFile = nil
		for _, v2 := range ArrVillages {
			if v2.DistrictId == v1.Id {
				ArrVillagesFile = append(ArrVillagesFile, CsvDataVillageLines{Id: v2.Id, DistrictId: v2.DistrictId, DistrictName: v1.Name, PostalCode: v2.PostalCode, Name: v2.Name})
			}
		}
		file, _ := json.MarshalIndent(ArrVillagesFile, "", " ")
		_ = ioutil.WriteFile(fmt.Sprintf(`./output/%v.json`, v1.Id), file, 0644)
		time.Sleep(10 * time.Millisecond)
		// log.Printf("Selesai sleep.....")
	}
	log.Printf("Finish")
}

func readCsvFile(filename string) ([][]string, error) {
	// Open CSV file
	fileContent, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer fileContent.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(fileContent).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}
