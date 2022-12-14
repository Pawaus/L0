package nats

import stan "github.com/nats-io/stan.go"

func Connect(clusterId string, clientId string) (*stan.Conn, error) {
	sc, err := stan.Connect(clusterId, clientId)
	if err != nil {
		return nil, err
	}
	return &sc, nil
}
