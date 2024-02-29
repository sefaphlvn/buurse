package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/sefaphlvn/suubar/fileService"
	"github.com/sefaphlvn/suubar/runner"
	"google.golang.org/grpc"
)

type InterfaceStatus struct {
	Eth0 struct {
		AdministrativeStatus string `json:"administrativeStatus"`
		OperationalStatus    string `json:"operationalStatus"`
		LinkDetection        bool   `json:"linkDetection"`
		LinkUps              int    `json:"linkUps"`
		LinkDowns            int    `json:"linkDowns"`
		VrfName              string `json:"vrfName"`
		PseudoInterface      bool   `json:"pseudoInterface"`
		Index                int    `json:"index"`
		Metric               int    `json:"metric"`
		Mtu                  int    `json:"mtu"`
		Speed                int64  `json:"speed"`
		Flags                string `json:"flags"`
		Type                 string `json:"type"`
		HardwareAddress      string `json:"hardwareAddress"`
		IPAddresses          []struct {
			Address    string `json:"address"`
			Secondary  bool   `json:"secondary"`
			Unnumbered bool   `json:"unnumbered"`
		} `json:"ipAddresses"`
		InterfaceType      string `json:"interfaceType"`
		InterfaceSlaveType string `json:"interfaceSlaveType"`
		LacpBypass         bool   `json:"lacpBypass"`
		EvpnMh             struct {
		} `json:"evpnMh"`
		Protodown string `json:"protodown"`
	} `json:"eth0"`
	Lo struct {
		AdministrativeStatus string `json:"administrativeStatus"`
		OperationalStatus    string `json:"operationalStatus"`
		LinkDetection        bool   `json:"linkDetection"`
		LinkUps              int    `json:"linkUps"`
		LinkDowns            int    `json:"linkDowns"`
		VrfName              string `json:"vrfName"`
		PseudoInterface      bool   `json:"pseudoInterface"`
		Index                int    `json:"index"`
		Metric               int    `json:"metric"`
		Mtu                  int    `json:"mtu"`
		Speed                int    `json:"speed"`
		Flags                string `json:"flags"`
		Type                 string `json:"type"`
		IPAddresses          []any  `json:"ipAddresses"`
		InterfaceType        string `json:"interfaceType"`
		InterfaceSlaveType   string `json:"interfaceSlaveType"`
		LacpBypass           bool   `json:"lacpBypass"`
		EvpnMh               struct {
		} `json:"evpnMh"`
		Protodown string `json:"protodown"`
	} `json:"lo"`
}

func main() {
	a, _ := runner.RunCommandAndDecode[InterfaceStatus]("show interface json", []string{})
	PrettyPrinter(a)

	lis, err := net.Listen("tcp", ":50041")
	if err != nil {
		log.Fatalf("Port not listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// Servisleri gRPC sunucusuna kaydet
	fileService.RegisterWithServer(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server not started: %v", err)
	}
}

func PrettyPrinter(data interface{}) {
	prettyJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling error: %v", err)
	}

	fmt.Println(string(prettyJSON))
}
