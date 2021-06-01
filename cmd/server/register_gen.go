// IM AUTO GENERATED, BUT CAN BE OVERRIDDEN

package main

import (
	"github.com/dnk90/adlive/internal/services"
	"github.com/dnk90/adlive/internal/stores"
	"github.com/dnk90/adlive/pb"
)

func registerService() pb.AdliveServer {
	mainStore := stores.NewMainStore()
	return services.New(mainStore)
}
