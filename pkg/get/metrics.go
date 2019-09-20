package get

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/pflag"
	"github.com/wyb1/promclient/pkg/common"
)

func getGet() {
	fmt.Println("Get")
}

// All gets all metrics
func All(args []string, flags *pflag.FlagSet) error {
	fmt.Println("Get All")
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
	result, _, err := promClient.Api.Series(context.Background(), []string{"{__name__=~'.+'}"}, time.Now(), time.Now())
	fmt.Printf("Result:\n%v\n", result)
	fmt.Println("-----------------------------------------")
	return nil
}
