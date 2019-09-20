package query

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/pflag"
	"github.com/wyb1/promclient/pkg/common"
)

// Handler initializes a prometheus client and sends all queries to Prometheus
func Handler(args []string, flags *pflag.FlagSet) error {
	address, err := flags.GetString("address")
	if err != nil {
		return err
	}
	port, err := flags.GetString("port")
	if err != nil {
		return err
	}
	address = fmt.Sprintf("%s%s", address, port)
	fmt.Println(address)
	promClient, err := common.InitClient(address)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for _, arg := range args {
		if err := promClient.PrintQuery(ctx, arg); err != nil {
			return err
		}
	}
	return nil
}
