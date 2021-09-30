package main

import (
	"fmt"

	"github.com/go-ocf/go-coap"
	"github.com/go-ocf/go-coap/codes"
)

// This is broken for libcoap since the blocks should return type acknowledgement
// for all blocks transferred from the service. THe new pldg-go version works
// with libcoap.
func oldServer() {
	mux := coap.NewServeMux()

	mux.HandleFunc("/fw", func(w coap.ResponseWriter, r *coap.Request) {
		fmt.Println("Serving 1M bytes")
		msg := w.NewResponse(codes.Content)
		// libcoap checks etag for blockwise transfers
		msg.SetType(coap.Acknowledgement)
		msg.SetOption(coap.ETag, etag)
		msg.SetOption(coap.ContentFormat, coap.AppOctets)
		msg.SetType(coap.NonConfirmable)
		msg.SetPayload(dummyBuf)
		if err := w.WriteMsg(msg); err != nil {
			fmt.Printf("Error writing message: %v\n", err)
		}
	})

	server := coap.Server{
		Addr:    "127.0.0.1:5683",
		Net:     "udp",
		Handler: mux,
	}

	fmt.Println("Running old server (go-ocf)")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error serving: %v\n", err)
	}
}
