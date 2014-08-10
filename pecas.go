package main

import (
	"fmt"
	"strconv"
	"strings"
)

var mappecas = map[string]string{}

func buildNametonumber(intm map[string]string) map[string]string {
	m := map[string]string{}

	for key, value := range intm {
		m[value] = key
	}

	return m
}

type Pecas struct {
	Pecas           map[string]string
	Separador       string
	Calcados        []string     //base|obrigatorio
	Calcas          []string     //base|obrigatorio
	Cintos          []string     //só com calca operacional
	MangasCurtas    []string     //base|obrigatorio
	MangasCompridas []string     //opção pode ser usado em cima mangacurta|camisas
	Camisas         []string     //pode usar em cima de tshirt
	Divisas         string       //so camisa | opção
	Gravata         string       //so camisa|opcao
	PullOver        string       //so camisa | opção
	Casacos         []string     //pode ser usado em tudo
	Parkas          []string     //pode ser usado em tudo
	Coletes         []string     //pode ser usado em tudo
	Bone            string       //pode ser usado em tudo
	C               *Combinacoes `json:"-"`
}

func (p *Pecas) AdicionarOpcionais(vestuario ...string) {
	p.C.Add(p.Separador, vestuario...)
	p.C.Add(p.Separador, append(vestuario, p.Bone)...)
	p.C.Add(p.Separador, append(vestuario, p.Bone, p.Divisas)...)
	p.C.Add(p.Separador, append(vestuario, p.Divisas)...)

	for _, colete := range p.Coletes {
		p.C.Add(p.Separador, append(vestuario, p.Bone, colete)...)
		p.C.Add(p.Separador, append(vestuario, p.Bone, p.Divisas, colete)...)
		p.C.Add(p.Separador, append(vestuario, p.Divisas, colete)...)
	}
}

func (p *Pecas) Caiemcimadetudo(vestuario ...string) {
	p.AdicionarOpcionais(vestuario...)
	for _, mangacomprida := range p.MangasCompridas {
		p.AdicionarOpcionais(p.C.J(vestuario, p.Separador), mangacomprida)
		for _, casaco := range p.Casacos {
			p.AdicionarOpcionais(p.C.J(vestuario, p.Separador), casaco)
			p.AdicionarOpcionais(p.C.J(vestuario, p.Separador), mangacomprida, casaco)
			for _, parka := range p.Parkas {
				p.AdicionarOpcionais(p.C.J(vestuario, p.Separador), parka)
				p.AdicionarOpcionais(p.C.J(vestuario, p.Separador), mangacomprida, parka)
				p.AdicionarOpcionais(p.C.J(vestuario, p.Separador), mangacomprida, casaco, parka)
			}
		}
	}
}

func (p *Pecas) ApplyFilters(filters []filter) {
	for i := range filters {
		filters[i](p)
	}
}
func (p *Pecas) PrintFull() {
	for _, combinacao := range *p.C {
		fmt.Println(combinacao)
	}

}

func (p *Pecas) PrintShort() {
	for _, combinacao := range *p.C {
		tmpstr := combinacao
		for key, value := range buildNametonumber(p.Pecas) {
			tmpstr = strings.Replace(tmpstr, key, value, -1)
		}
		fmt.Println(tmpstr)
	}

}
func (p *Pecas) PrintLegenda() {
	for i := 1; i <= len(p.Pecas); i++ {

		fmt.Println(i, p.Pecas[strconv.Itoa(i)])
	}

}
