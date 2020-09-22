package json_demo

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"name"`
	ServerIp  string `json:"ip"`
	ServerPort int `json:"port"`
}

//序列化结构体
func Serialize()  {
	server := new(Server)
	server.ServerName="Demo-for-json"
	server.ServerIp="127.0.0.1"
	server.ServerPort= 8080

	// 序列化成json 字节数组
	b,err :=json.Marshal(server)
	if err!=nil{
		fmt.Println("Marshal err:",err.Error())
		return
	}
	// 将json 字节数组 string
	fmt.Println("Marshal json:",string(b))
}



// 序列化map
func SerializeMap()  {
	server := make(map[string]interface{})
	server["ServerName"]="Demo-for-json-map"
	server["ServerIp"]="127.0.0.1"
	server["ServerPort"]= 9090

	// 序列化成json 字节数组
	b,err :=json.Marshal(server)
	if err!=nil{
		fmt.Println("Marshal err:",err.Error())
		return
	}
	// 将json 字节数组 string
	fmt.Println("Marshal json:",string(b))
}

func DeSerializeStruct()  {
	jsonString :=`{"name":"Demo-for-json","ip":"127.0.0.1","port":8080}`
	server := new(Server)
	err:=json.Unmarshal([]byte(jsonString),&server)
	if err!=err{
		fmt.Println("Unmarshal err:",err.Error())
	}
	fmt.Println("Unmarshal struct :",server)
	}

func DeSerializeMap()  {
	jsonString :=`{"name":"Demo-for-json","ip":"127.0.0.1","port":8080}`
    server :=make(map[string] interface{})
	err:=json.Unmarshal([]byte(jsonString),&server)
	if err!=err{
		fmt.Println("Unmarshal err:",err.Error())
	}
	fmt.Println("Unmarshal struct :",server)

}