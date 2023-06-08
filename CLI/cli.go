package CLI

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Option struct {
	ID        string
	Help      string
	ArgsRegex string
	Required  bool
	IsInt     bool
}

type OptionGot struct {
	ID    string
	Value any
}

type CLI struct {
	Options []Option
}

func (o *Option) GetRegex() *regexp.Regexp {
	if o.ArgsRegex != "" {
		return regexp.MustCompile(fmt.Sprintf("-%s %s", o.ID, o.ArgsRegex))
	}
	return regexp.MustCompile(fmt.Sprintf("-%s", o.ID))
}

func (o *Option) parse(cli string) (*OptionGot, error) {
	regex := o.GetRegex()
	un := regex.FindString(cli)
	if un == "" && o.Required {
		return nil, fmt.Errorf("option %s is required", o.ID)
	}
	if un == "" {
		return nil, nil
	}
	str := strings.Replace(un, fmt.Sprintf("-%s ", o.ID), "", 1)
	if !o.IsInt {
		return &OptionGot{ID: o.ID, Value: str}, nil
	}
	f, err := strconv.Atoi(str)
	if err != nil {
		return nil, err
	}
	return &OptionGot{ID: o.ID, Value: f}, nil
}

func (c *CLI) AddOption(o Option) *CLI {
	c.Options = append(c.Options, o)
	return c
}

func (c *CLI) Parse(cli string) ([]*OptionGot, error) {
	var got []*OptionGot
	for _, o := range c.Options {
		g, err := o.parse(cli)
		if err != nil {
			return nil, err
		}
		if g != nil {
			got = append(got, g)
		}
	}
	return got, nil
}
