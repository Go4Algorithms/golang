package graph

import (
	"strconv"
	"os"
	"bufio"
	"os/exec"
	"strings"
	"syscall"
)

func  (g *Graph)Encode(filepath string)error{
	//convert graph to dot
	s := "graph "
	if g.Directed{
		s = "di"+s
	}
	s += g.Name
	filepath += ".dot"
	s+= " {\n"
	for _,e := range g.Edges {
		line := " -- "
		if g.Directed{
			line = " -> "
		}
		s+= "\t" +e.From.Name + line + e.To.Name
		if g.Weighted {
			s+="[label=\""+strconv.Itoa(e.Weight)+"\",weight=\""+strconv.Itoa(e.Weight)+"\"]"
		}
		s+=";\n"
	}
	s+="}"
	f,er := os.Create(filepath)
	if er != nil {
		panic(er)
	}
	w:=bufio.NewWriter(f)
	_,er = w.WriteString(s)
	if er!= nil{
		panic(er)
	}
	w.Flush()

	//convert dot to png
	binary, er := exec.LookPath("dot")
	if er != nil{
		panic(er)
	}
	args := []string {"dot","-Tpng",filepath,"-o",strings.TrimRight(filepath,".dot")+".png"}
	env := os.Environ()
	er = syscall.Exec(binary,args,env)
	if er !=nil{
		panic(er)
	}
	return er
}
