package handlers

import (
	"api/interal/pkg/domain"
	"encoding/json"
	"log"

	"github.com/nats-io/stan.go"
)

type Usecase interface {
	Update(uid string, data string)
	GetAll()
}
type Handler struct {
	usecase Usecase
}

func NewHandler(uc Usecase) *Handler {
	return &Handler{usecase: uc}
}

func (h *Handler) UpdateChannel(m *stan.Msg) {
	log := log.Default()
	msg := string(m.Data)
	//log.Print("from channel msg: " + msg)
	d := domain.Order{}
	err := json.Unmarshal(m.Data, &d)
	if err != nil {
		log.Printf("json.Unmarshal: %v", err)
	}

	h.usecase.Update(d.Order_uid, msg)
	//fmt.Printf("%+v", d)
}

func (h *Handler) UpdateBD() {
	h.usecase.GetAll()
}
