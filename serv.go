package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jarviscore"
)

func startServ() {
	fmt.Printf("dtdata server start...\n")
	fmt.Printf("dtdata server version is %v \n", VERSION)

	cfg, err := jarviscore.LoadConfig("cfg/config.yaml")
	if err != nil {
		fmt.Printf("load config.yaml fail!\n")

		return
	}

	jarviscore.InitJarvisCore(cfg)
	defer jarviscore.ReleaseJarvisCore()

	// pprof
	jarviscore.InitPprof(cfg)

	node, err := jarviscore.NewNode(cfg)
	if err != nil {
		fmt.Printf("jarviscore.NewNode fail! %v \n", err)

		return
	}

	node.SetNodeTypeInfo(DTDATASERVTYPE, VERSION)

	node.Start(context.Background())

	fmt.Printf("dtdata server end.\n")
}
