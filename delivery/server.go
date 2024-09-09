package delivery

import (
	"asetku-bukan-asetmu/config"
	"asetku-bukan-asetmu/delivery/controller"
	"asetku-bukan-asetmu/manager"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	usecaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func (a *appServer) initController() {
	controller.NewTestController(a.engine, a.usecaseManager.TestUsecase())
	controller.NewEmployeeController(a.engine, a.usecaseManager.EmployeeUseCase())
	controller.NewAssetController(a.engine, a.usecaseManager.AssetUsecase())
	controller.NewAssetLocationController(a.engine, a.usecaseManager.AssetLocationUsecase())
	controller.NewAssetCategoriesController(a.engine, a.usecaseManager.AssetCategoriesUseCase())
	controller.NewVendorController(a.engine, a.usecaseManager.VendorUseCase())
}

func (a *appServer) Run() {
	a.initController()

	err := a.engine.Run(a.host)
	if err != nil {
		panic(err.Error())
	}
}

func Server() *appServer {
	engine := gin.Default()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln("Error Config : ()", err.Error())
	}

	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		log.Fatalln("Error Conection : ", err.Error())
	}

	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	host := fmt.Sprintf("%s:%s", cfg.APIHost, cfg.APIPort)

	return &appServer{
		engine:         engine,
		host:           host,
		usecaseManager: useCaseManager,
	}
}
