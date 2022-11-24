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
	Id         int    `json:"id"`
	ProvinceId int    `json:"provinceId"`
	Name       string `json:"name"`
}

var ArrRegencies []CsvDataRegencieLines
var ArrRegenciesFile []CsvDataRegencieLines

type CsvDataDistrictLines struct {
	Id         int    `json:"id"`
	RegencieId int    `json:"regencieId"`
	Name       string `json:"name"`
}

var ArrDistricts []CsvDataDistrictLines
var ArrDistrictsFile []CsvDataDistrictLines

type CsvDataVillageLines struct {
	Id         int    `json:"id"`
	DistrictId int    `json:"districtId"`
	Name       string `json:"name"`
}

var ArrVillages []CsvDataVillageLines
var ArrVillagesFile []CsvDataVillageLines

func Generate() {
	//Provices
	dataProvinces, err := readCsvFile("./data/provinces.csv")
	if err != nil {
		panic(err)
	}

	for _, line := range dataProvinces {
		id, _ := strconv.Atoi(line[0])
		ArrProvinces = append(ArrProvinces, CsvDataProvinceLines{Id: id, Name: line[1]})
	}

	//Regencies
	dataRegencies, err := readCsvFile("./data/regencies.csv")
	if err != nil {
		panic(err)
	}

	for _, line := range dataRegencies {
		id, _ := strconv.Atoi(line[0])
		provinceId, _ := strconv.Atoi(line[1])
		ArrRegencies = append(ArrRegencies, CsvDataRegencieLines{Id: id, ProvinceId: provinceId, Name: line[2]})
	}

	//Districts
	dataDistricts, err := readCsvFile("./data/districts.csv")
	if err != nil {
		panic(err)
	}

	for _, line := range dataDistricts {
		id, _ := strconv.Atoi(line[0])
		regencieId, _ := strconv.Atoi(line[1])
		ArrDistricts = append(ArrDistricts, CsvDataDistrictLines{Id: id, RegencieId: regencieId, Name: line[2]})
	}

	//Villages
	dataVillages, err := readCsvFile("./data/villages.csv")
	if err != nil {
		panic(err)
	}

	for _, line := range dataVillages {
		id, _ := strconv.Atoi(line[0])
		districtId, _ := strconv.Atoi(line[1])
		ArrVillages = append(ArrVillages, CsvDataVillageLines{Id: id, DistrictId: districtId, Name: line[2]})
	}

	//Create Json File Provinces
	file, _ := json.MarshalIndent(ArrProvinces, "", " ")
	_ = ioutil.WriteFile("./output/provinces.json", file, 0644)

	//Create json File Regencies
	for _, v1 := range ArrProvinces {
		ArrRegenciesFile = nil
		for _, v2 := range ArrRegencies {
			if v2.ProvinceId == v1.Id {
				ArrRegenciesFile = append(ArrRegenciesFile, CsvDataRegencieLines{Id: v2.Id, ProvinceId: v2.ProvinceId, Name: v2.Name})
			}
		}
		file, _ = json.MarshalIndent(ArrRegenciesFile, "", " ")
		_ = ioutil.WriteFile(fmt.Sprintf(`./output/%v.json`, v1.Id), file, 0644)

		time.Sleep(10 * time.Millisecond)
		log.Printf("Selesai sleep.....")
	}

	log.Printf("masuk 1")
	//Create json File Districts
	for _, v1 := range ArrRegencies {
		ArrDistrictsFile = nil
		for _, v2 := range ArrDistricts {
			if v2.RegencieId == v1.Id {
				ArrDistrictsFile = append(ArrDistrictsFile, CsvDataDistrictLines{Id: v2.Id, RegencieId: v2.RegencieId, Name: v2.Name})
			}
		}
		file, _ := json.MarshalIndent(ArrDistrictsFile, "", " ")
		_ = ioutil.WriteFile(fmt.Sprintf(`./output/%v.json`, v1.Id), file, 0644)
		time.Sleep(10 * time.Millisecond)
		log.Printf("Selesai sleep.....")
	}

	//Create json File Villages
	for _, v1 := range ArrDistricts {
		ArrVillagesFile = nil
		for _, v2 := range ArrVillages {
			if v2.DistrictId == v1.Id {
				ArrVillagesFile = append(ArrVillagesFile, CsvDataVillageLines{Id: v2.Id, DistrictId: v2.DistrictId, Name: v2.Name})
			}
		}
		file, _ := json.MarshalIndent(ArrVillagesFile, "", " ")
		_ = ioutil.WriteFile(fmt.Sprintf(`./output/%v.json`, v1.Id), file, 0644)
		time.Sleep(10 * time.Millisecond)
		log.Printf("Selesai sleep.....")
	}
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
