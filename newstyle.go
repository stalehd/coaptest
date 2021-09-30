package main

import (
	"bytes"
	"fmt"

	"github.com/plgd-dev/go-coap/v2"
	"github.com/plgd-dev/go-coap/v2/message"
	"github.com/plgd-dev/go-coap/v2/message/codes"
	"github.com/plgd-dev/go-coap/v2/mux"
)

func newServer() {
	r := mux.NewRouter()
	if err := r.Handle("/fw", mux.HandlerFunc(func(w mux.ResponseWriter, r *mux.Message) {
		fmt.Println("Serving 1M bytes")

		var opts []message.Option
		// libcoap wants etag for blockwise transfers
		opts = append(opts, message.Option{ID: message.ETag, Value: etag})

		if err := w.SetResponse(codes.Content, message.AppOctets, bytes.NewReader(dummyBuf), opts...); err != nil {
			fmt.Printf("Error setting response: %v\n", err)
		}
	})); err != nil {
		fmt.Printf("Error assigning handler: %v\n", err)
	}

	fmt.Println("Running new (plgd-dev)")
	if err := coap.ListenAndServe("udp", "127.0.0.1:5683", r); err != nil {
		fmt.Printf("Error serving: %v\n", err)
	}
}
