package config

import "strconv"

type Redis struct {
	Ip       string `yaml:"ip"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
}

func (r *Redis) Conn() string {
	return r.Ip + ":" + strconv.Itoa(r.Port)
}
