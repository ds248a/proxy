package main

import (
	"log"
	"os"

	"github.com/ds248a/proxy/socks5"
)

type params struct {
	User            string
	Password        string
	Port            string
	Addr            string
	AllowedDestFqdn string
}

func main() {
	cfg := params{
		User:     "user",
		Password: "password",
		Addr:     "127.0.0.1",
		Port:     "1018",
	}

	s5 := &socks5.Config{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}

	if cfg.User+cfg.Password != "" {
		creds := socks5.StaticCredentials{
			cfg.User: cfg.Password,
		}
		cator := socks5.UserPassAuthenticator{Credentials: creds}
		s5.AuthMethods = []socks5.Authenticator{cator}
	}

	if cfg.AllowedDestFqdn != "" {
		s5.Rules = PermitDestAddrPattern(cfg.AllowedDestFqdn)
	}

	server, err := socks5.New(s5)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.ListenAndServe("tcp", cfg.Addr+":"+cfg.Port); err != nil {
		log.Fatal(err)
	}
}
