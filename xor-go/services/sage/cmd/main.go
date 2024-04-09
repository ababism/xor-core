package main

import "fmt"

const configPath = "./services/sage/configs/config.dev.yaml"

func main() {
	fmt.Println("ok")
	//cfg, err := config.NewConfig(configPath, ServiceCapsName)
	//if err != nil {
	//	log.Fatalf("Fail to parse %s config: %v", ServiceName, err)
	//}
	//var appConfig config.Config
	//if err := cleanenv.ReadConfig(configPath, &appConfig); err != nil {
	//	log.Fatalf("failed to read config from path: %s, error: %v", configPath, err)
	//}
	//
	//application, err := app.NewApp(&appConfig)
	//if err != nil {
	//	log.Fatalf("failed to create app: %v", err)
	//}
	//application.Start()
}
