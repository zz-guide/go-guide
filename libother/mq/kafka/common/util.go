package common

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/satori/go.uuid"
)

type JsonTime time.Time

// MarshalJSON 实现它的json序列化方法
func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

type Message struct {
	Id        string   `json:"id"`
	Data      string   `json:"data"`
	CreatedAt JsonTime `json:"created_at"`
}

func GetMessage(msg string) []byte {
	u1 := uuid.Must(uuid.NewV4(), nil)
	message := Message{
		Id:        u1.String(),
		Data:      msg,
		CreatedAt: JsonTime(time.Now()),
	}
	jsonBytes, _ := json.Marshal(message)
	return jsonBytes
}
