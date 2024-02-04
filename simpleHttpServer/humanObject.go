package simpleHttpServer

import (
	"encoding/json"
	"fmt"
	"os"
)

const dataFilePath = "./simpleHttpServer/data/sample_return_data.json"

type Address struct {
	Country string `json:"country"`
	State   string `json:"state"`
}

type Human struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Birthday string  `json:"birthday"`
	Address  Address `json:"address"`
}

func readDataFile() ([]Human, error) {
	f, err := os.Open(dataFilePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v", err)
	}
	defer f.Close()
	var humans []Human
	err = json.NewDecoder(f).Decode(&humans)
	if err != nil {
		return nil, fmt.Errorf("Error parsing data: %v", err)
	}
	return humans, nil
}

func writeDataFile(humans []Human) error {
	f, err := os.Create(dataFilePath)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("Error opening file: %v", err)
	}
	jsonStr, err := json.Marshal(humans)
	_, err = f.Write([]byte(jsonStr))
	if err != nil {
		return fmt.Errorf("Can not write data to %v", dataFilePath)
	}
	return nil
}

func appendHumanToDataFile(human Human) error {
	createdHumans, err := readDataFile()
	if err != nil {
		return fmt.Errorf("Can not cuurent humandata by json file.")
	}

	for _, h := range createdHumans {
		if h.Id == human.Id {
			return fmt.Errorf("User id %v is already exists.", h.Id)
		}
	}
	createdHumans = append(createdHumans, human)
	err = writeDataFile(createdHumans)
	if err != nil {
		return fmt.Errorf("Append human error.: %v", err)
	}

	return nil
}

func updateHumanForDataFile(humanId int, human Human) error {
	createdHumans, err := readDataFile()
	if err != nil {
		return fmt.Errorf("Can not cuurent humandata by json file.")
	}

	targetHumanIndex := -1
	for i, h := range createdHumans {
		if h.Id == humanId {
			targetHumanIndex = i
		}
		// IDを変更する場合、すでに同じIDのHumanが存在してないこと
		if h.Id == human.Id && humanId != human.Id {
			return fmt.Errorf("User id %v is already exists.", h.Id)
		}
	}
	if targetHumanIndex == -1 {
		return fmt.Errorf("User id %v was not found.", humanId)
	}
	createdHumans[targetHumanIndex] = human

	err = writeDataFile(createdHumans)
	if err != nil {
		return fmt.Errorf("Update human error.: %v", err)
	}

	return nil
}

func deleteHumnFromDataFile(humanId int) error {
	createdHumans, err := readDataFile()
	if err != nil {
		return fmt.Errorf("Can not cuurent humandata by json file.")
	}

	for i, h := range createdHumans {
		if h.Id == humanId {
			// remove target human data from json file.
			createdHumans = append(createdHumans[:i], createdHumans[i+1:]...)

			err = writeDataFile(createdHumans)
			if err != nil {
				return fmt.Errorf("Delete human error.: %v", err)
			}
			return nil
		}
	}

	return fmt.Errorf("Not found human id is %v.", humanId)
}

func TestWriteHumanToDataFile() {
	h := Human{
		Id:       4,
		Name:     "yusuke",
		Birthday: "1999-10-23",
		Address: Address{
			Country: "Japan",
			State:   "Wakayama",
		},
	}

	appendHumanToDataFile(h)
}
