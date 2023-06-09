package main

import (
	"regexp"

	"github.com/ds248a/proxy/socks5"

	"golang.org/x/net/context"
)

func PermitDestAddrPattern(pattern string) socks5.RuleSet {
	return &PermitDestAddrPatternRuleSet{pattern}
}

type PermitDestAddrPatternRuleSet struct {
	AllowedFqdnPattern string
}

func (p *PermitDestAddrPatternRuleSet) Allow(ctx context.Context, req *socks5.Request) (context.Context, bool) {
	match, _ := regexp.MatchString(p.AllowedFqdnPattern, req.DestAddr.FQDN)
	return ctx, match
}
