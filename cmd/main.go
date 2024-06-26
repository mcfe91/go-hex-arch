package main

import (
	"log"
	"os"

	"github.com/mmcferren/go-hex-arch/internal/adapters/app/api"
	"github.com/mmcferren/go-hex-arch/internal/adapters/core/arithmetic"
	gRPC "github.com/mmcferren/go-hex-arch/internal/adapters/framework/left/grpc"
	"github.com/mmcferren/go-hex-arch/internal/adapters/framework/right/db"
	"github.com/mmcferren/go-hex-arch/internal/ports"
)

func main() {
	var err error

	// ports
	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DB_NAME")

	// adapters
	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(core, dbaseAdapter)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)

	// run
	gRPCAdapter.Run()
}
