package main

type set struct {
	name string
	apps []appconfig
}

var allcfg set = *newSet("GLOBAL")

type cfgdir struct {
	path   string
	ignore []string
}

type appconfig struct {
	name   string
	config []cfgdir
}

func newSet(name string, apps ...appconfig) *set {
	s := set{name: name, apps: apps}
	return &s
}

func NewAppConfig(name string, configpath string, ignore ...string) appconfig {
	cfg := appconfig{
		name:   name,
		config: []cfgdir{{path: configpath, ignore: []string{}}},
	}

	return cfg
}
