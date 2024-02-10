package app

//
// import (
// 	"context"
// 	"fmt"
//
// 	"github.com/Snorkin/auth_service/configs"
// 	"github.com/Snorkin/auth_service/internal/service"
// )

//
// type App struct {
// 	serviceProvider *service.ServiceProvider
// }
//
// func NewApp(ctx context.Context) (*App, error) {
// 	a := &App{}
//
// 	err := a.initDependencies(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return a, nil
// }
//
// func (a *App) initServiceProvider(_ context.Context) error {
// 	a.serviceProvider = service.NewServiceProvider()
// 	return nil
// }
//
// func (a *App) initDependencies(ctx context.Context) error {
// 	inits := []func(context.Context) error{
// 		a.initConfig,
// 		a.initServiceProvider,
// 		// a.connectDB,
// 	}
// 	fmt.Println(inits)
// 	return nil
// }
//
// func (a *App) initConfig(_ context.Context) error {
// 	err := configs.Load(".env")
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
