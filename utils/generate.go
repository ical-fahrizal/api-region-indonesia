package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	conf "wec-region-indonesia/config"
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

		// for _, v2 := range ArrPostalCodes {
		// 	if v2.VillageName == line[2] {
		// 		postalCode = v2.PostalCode
		// 		break
		// 	}
		// }
		ArrVillages = append(ArrVillages, CsvDataVillageLines{Id: id, DistrictId: districtId, PostalCode: postalCode, Name: line[2]})
	}

	//Create Json File Provinces
	conf.LogInfo.Printf("Create json File Provinces")
	file, _ := json.MarshalIndent(ArrProvinces, "", " ")
	_ = ioutil.WriteFile("./output/provinces.json", file, 0644)

	//Create json File Regencies
	conf.LogInfo.Printf("Create json File Regencies")
	for _, v1 := range ArrProvinces {
		ArrRegenciesFile = nil
		for _, v2 := range ArrRegencies {
			if v2.ProvinceId == v1.Id {
				ArrRegenciesFile = append(ArrRegenciesFile, CsvDataRegencieLines{Id: v2.Id, ProvinceId: v2.ProvinceId, ProvinceName: v1.Name, Name: v2.Name})
			}
		}
		file, _ = json.MarshalIndent(ArrRegenciesFile, "", " ")
		//create file name by province id
		_ = ioutil.WriteFile(fmt.Sprintf(`./output/%v.json`, v1.Id), file, 0644)

		//create file name by province name
		// nameProvince := strings.ReplaceAll(v1.Name, " ", "-")
		_ = ioutil.WriteFile(fmt.Sprintf(`./output/%v.json`, v1.Name), file, 0644)

		time.Sleep(10 * time.Millisecond)
		// conf.LogInfo.Printf("Selesai sleep.....")
	}

	// conf.LogInfo.Printf("masuk 1")
	//Create json File Districts
	conf.LogInfo.Printf("Create json File Districts")
	for _, v1 := range ArrRegencies {
		ArrDistrictsFile = nil
		for _, v2 := range ArrDistricts {
			if v2.RegencieId == v1.Id {
				ArrDistrictsFile = append(ArrDistrictsFile, CsvDataDistrictLines{Id: v2.Id, RegencieId: v2.RegencieId, RegencieName: v1.Name, Name: v2.Name})
			}
		}
		file, _ := json.MarshalIndent(ArrDistrictsFile, "", " ")
		//create file name by regency id
		_ = ioutil.WriteFile(fmt.Sprintf(`./output/%v.json`, v1.Id), file, 0644)

		//create file name by regency name
		_ = ioutil.WriteFile(fmt.Sprintf(`./output/%v.json`, v1.Name), file, 0644)

		//create file name by province name - regency name
		// nameRegency := strings.ReplaceAll(v1.Name, " ", "-")
		// nameProvince := strings.ReplaceAll(v1.ProvinceName, " ", "-")
		_ = ioutil.WriteFile(fmt.Sprintf(`./output/%v %v.json`, v1.ProvinceName, v1.Name), file, 0644)

		time.Sleep(10 * time.Millisecond)
		// conf.LogInfo.Printf("Selesai sleep.....")
	}

	//Create json File Villages
	conf.LogInfo.Printf("Create json File Villages")
	hitung := 0
	for _, v1 := range ArrDistricts {
		ArrVillagesFile = nil
		for _, v2 := range ArrVillages {
			if v2.DistrictId == v1.Id {

				postalCode := ""
				for _, v3 := range ArrPostalCodes {
					// if v2.Name == "MUNCUL" && v3.VillageName == "MUNCUL" {
					// 	conf.LogInfo.Printf("v3.VillageName : %v", v3.VillageName)
					// 	conf.LogInfo.Printf("v2.Name : %v", v2.Name)
					// 	conf.LogInfo.Printf("v3.DistrictName : %v", v3.DistrictName)
					// 	conf.LogInfo.Printf("v1.Name : %v", v1.Name)
					// 	conf.LogInfo.Printf("v3.RegencieName : %v", v3.RegencieName)
					// 	conf.LogInfo.Printf("v1.RegencieName : %v", v1.RegencieName)
					// 	conf.LogInfo.Printf("v2.PostalCode : %v", v2.PostalCode)
					// }
					covV3VillageName := strings.ReplaceAll(v3.VillageName, " ", "")
					covV2Name := strings.ReplaceAll(v2.Name, " ", "")

					covV3DistrictName := strings.ReplaceAll(v3.DistrictName, " ", "")
					covV1Name := strings.ReplaceAll(v1.Name, " ", "")

					if covV3VillageName == covV2Name && covV3DistrictName == covV1Name {
						postalCode = v3.PostalCode
						break
					}
				}
				// ArrVillagesFile = append(ArrVillagesFile, CsvDataVillageLines{Id: v2.Id, DistrictId: v2.DistrictId, DistrictName: v1.Name, PostalCode: v2.PostalCode, Name: v2.Name})
				ArrVillagesFile = append(ArrVillagesFile, CsvDataVillageLines{Id: v2.Id, DistrictId: v2.DistrictId, DistrictName: v1.Name, PostalCode: postalCode, Name: v2.Name})
				if postalCode == "" {
					// conf.LogInfo.Printf("v1.RegencieId : %v", v1.RegencieId)
					// if (v1.RegencieId >= 3171 && v1.RegencieId <= 3175) || v1.RegencieId == 3101 ||
					//if (v1.RegencieId >= 3601 && v1.RegencieId <= 3604) || (v1.RegencieId >= 3671 && v1.RegencieId <= 3674) {
					if (v1.RegencieId >= 3201 && v1.RegencieId <= 3218) || (v1.RegencieId >= 3271 && v1.RegencieId <= 3279) {
						hitung += 1
						b1, err := json.Marshal(ArrVillagesFile)
						if err == nil {
							conf.LogInfo.Printf("b1 %v : %v", hitung, string(b1))
							log.Println()
						}
					}

				}
			}
		}
		file, _ := json.MarshalIndent(ArrVillagesFile, "", " ")
		//create file name by District id
		_ = ioutil.WriteFile(fmt.Sprintf(`./output/%v.json`, v1.Id), file, 0644)

		//create file name by District name
		// nameDistrip := strings.ReplaceAll(v1.Name, " ", "-")
		// nameRegency := strings.ReplaceAll(v1.RegencieName, " ", "-")
		_ = ioutil.WriteFile(fmt.Sprintf(`./output/%v.json`, v1.Name), file, 0644)

		//create file name by Regency name - District name
		// nameDistrip := strings.ReplaceAll(v1.Name, " ", "-")
		// nameRegency := strings.ReplaceAll(v1.RegencieName, " ", "-")
		_ = ioutil.WriteFile(fmt.Sprintf(`./output/%v %v.json`, v1.RegencieName, v1.Name), file, 0644)

		time.Sleep(10 * time.Millisecond)
		// conf.LogInfo.Printf("Selesai sleep.....")
	}
	conf.LogInfo.Printf("Finish")
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
