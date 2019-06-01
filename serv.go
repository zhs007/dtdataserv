package main

import (
	"context"
	"fmt"

	dtdatabasedef "github.com/zhs007/dtdataserv/basedef"
	"github.com/zhs007/dtdataserv/dtdata"
	"github.com/zhs007/jarviscore"
)

func startServ() {
	fmt.Printf("dtdata server start...\n")
	fmt.Printf("dtdata server version is %v \n", dtdatabasedef.VERSION)

	cfg, err := jarviscore.LoadConfig("cfg/jarvisnode.yaml")
	if err != nil {
		fmt.Printf("load jarvisnode.yaml fail!\n")

		return
	}

	jarviscore.InitJarvisCore(cfg, dtdatabasedef.JARVISNODETYPE, dtdatabasedef.VERSION)
	defer jarviscore.ReleaseJarvisCore()

	dtd, err := dtdata.NewDTData("./cfg/config.yaml")
	if err != nil {
		fmt.Printf("NewDTData %v", err)

		return
	}

	// pprof
	jarviscore.InitPprof(cfg)

	node, err := jarviscore.NewNode(cfg)
	if err != nil {
		fmt.Printf("jarviscore.NewNode fail! %v \n", err)

		return
	}

	node.SetNodeTypeInfo(dtdatabasedef.JARVISNODETYPE, dtdatabasedef.VERSION)

	go dtd.Start(context.Background(), node)
	node.Start(context.Background())

	fmt.Printf("dtdata server end.\n")
}
