package config

import (
	"flag"
	"net/url"
	"os"
	"strings"
)

type Config struct {
	apiKey     string
	domainName string
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.apiKey, "apiKey", os.Getenv("API_KEY"), "Freshservice API Key")
	flag.StringVar(&conf.domainName, "freshServiceDomainName", os.Getenv("DOMAIN_NAME"), "Freshservice Domain Name")

	flag.Parse()

	return conf
}

func (c *Config) GetApiKey() string {
	return c.apiKey
}

func (c *Config) GetDomainName() string {
	return c.domainName
}

func (c *Config) GetFreshServiceURL() string {
	c.domainName = strings.TrimSuffix(c.domainName, "/")
	u, err := url.Parse(c.domainName)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return "https://" + c.domainName + ".freshservice.com"
	} else if strings.Contains(c.domainName, "freshservice.com") && u.Scheme == "" {
		return "https://" + c.domainName
	} else {
		return c.domainName
	}
}
