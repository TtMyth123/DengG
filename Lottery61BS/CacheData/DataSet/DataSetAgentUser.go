package DataSet

import (
	"encoding/json"
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/CacheData/Bo"
	"github.com/TtMyth123/kit/fileKit"
	"io/ioutil"
)

const FileName_AgentUser = "AgentUser.txt"

type DataAgentUser struct {
	FileName string
	DataSet  map[string]Bo.AgentUser
}

func (this *DataAgentUser) Reload(fileName string) error {
	if fileName == "" {
		fileName = fmt.Sprintf("conf/%s", FileName_AgentUser)
	}

	this.FileName = fileName
	b, _ := fileKit.IsFileExist(fileName)
	if !b {
		this.DataSet = make(map[string]Bo.AgentUser)
		return nil
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, this)
	return err
}

func (this *DataAgentUser) Save() error {
	data, e := json.Marshal(this)
	if e != nil {
		return e
	}
	e = fileKit.CreateFile(data, this.FileName)
	return e
}
