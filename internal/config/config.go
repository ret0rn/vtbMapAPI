package config

import (
	"os"
	"strings"
)

var config AppConfig

type AppConfig struct {
	listenerPort        string
	forwardedByClientIP string
	trustedProxies      string
	masterPool          string
}

func Load() {
	config.listenerPort = os.Getenv("LISTENER_PORT")
	config.forwardedByClientIP = os.Getenv("FORWARDED_BY_CLIENT_IP")
	config.trustedProxies = os.Getenv("TRUSTED_PROXIES")
	config.masterPool = os.Getenv("MASTER_POOL")
}

func GetListenerPort() string {
	return config.listenerPort
}

func GetForwardedByClientIP() bool {
	return strings.ToLower(config.forwardedByClientIP) == "true"
}

func GetTrustedProxies() []string {
	if config.trustedProxies == "" {
		return nil
	}
	return strings.Split(config.trustedProxies, ";")
}

func GetMasterPool() string {
	return config.masterPool
}
