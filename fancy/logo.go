package fancy

import (
	"math/rand"
	"strings"
)

type LogoStr string
var (
	colors = []string{
		"\x1b[38;5;201m",
		"\x1b[38;5;165m",
		"\x1b[38;5;141m",
		"\x1b[38;5;105m",
		"\x1b[38;5;93m",
		"\x1b[38;5;212m",	
	}
)

func BuildLogo() *LogoStr {
	var out []string
	var logo LogoStr

	out = append(out, "╔═╗╦ ╦╦═╗╔═╗╔═╗╔╦╗  ╔═╗╔╗  ╦╔═╗╔═╗╔╦╗╔═╗")
	out = append(out, "║  ║ ║╠╦╝╚═╗║╣  ║║  ║ ║╠╩╗ ║║╣ ║   ║ ╚═╗")
	out = append(out, "╚═╝╚═╝╩╚═╚═╝╚═╝═╩╝  ╚═╝╚═╝╚╝╚═╝╚═╝ ╩ ╚═╝")


	logo = LogoStr(strings.Join(out, "\n"))
	return &logo
}

func(logo *LogoStr) Colorize(){
	chain := strings.Split(string(*logo), "\n")
	
	var out []string
	for _, v := range chain {
		var buff []string
		for _, v_ := range v {
			color := colors[rand.Intn(len(colors))]
			buff = append(buff, color+string(+v_)+"\x1b[0m")
		}
		
		out = append(out, strings.Join(buff, ""))
	}

	*logo = LogoStr(strings.Join(out,"\n"))
}