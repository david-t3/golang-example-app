// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package oauth

import (
	"github.com/aristat/golang-gin-oauth2-example-app/app/config"
	"github.com/aristat/golang-gin-oauth2-example-app/app/entrypoint"
	"github.com/aristat/golang-gin-oauth2-example-app/app/logger"
	"github.com/aristat/golang-gin-oauth2-example-app/app/session"
)

// Injectors from injector.go:

func Build() (*Manager, func(), error) {
	context, cleanup, err := entrypoint.ContextProvider()
	if err != nil {
		return nil, nil, err
	}
	viper, cleanup2, err := config.Provider()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	loggerConfig, cleanup3, err := logger.ProviderCfg(viper)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	zap, cleanup4, err := logger.Provider(context, loggerConfig)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	oauthConfig, cleanup5, err := Cfg(viper)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tokenStore, cleanup6, err := TokenStore(oauthConfig)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	sessionConfig, cleanup7, err := session.Cfg(viper)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	manager, cleanup8, err := session.Provider(context, sessionConfig)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	oauthManager, cleanup9, err := Provider(context, zap, tokenStore, manager)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return oauthManager, func() {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

func BuildTest() (*Manager, func(), error) {
	context, cleanup, err := entrypoint.ContextProviderTest()
	if err != nil {
		return nil, nil, err
	}
	loggerConfig, cleanup2, err := logger.ProviderCfgTest()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mock, cleanup3, err := logger.ProviderTest(context, loggerConfig)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tokenStore, cleanup4, err := TokenStoreTest()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	manager, cleanup5, err := session.ProviderTest()
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	oauthManager, cleanup6, err := Provider(context, mock, tokenStore, manager)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return oauthManager, func() {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
