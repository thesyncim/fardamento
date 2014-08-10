package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"io"
	"log"
	"os"
)

func CombinaPecas(f io.Reader) *Pecas {
	dec := json.NewDecoder(f)
	p := new(Pecas)
	err := dec.Decode(p)
	if err != nil {
		log.Fatalln(err)
	}
	p.C = NewCombination()

	//comecar pelos pes
	for _, calcado := range p.Calcados {
		//percorrer as calcas
		for _, calca := range p.Calcas {
			//percorrer todas as mangas curtas
			for _, mangacurta := range p.MangasCurtas {
				//se for a ultima das opçoes
				p.Caiemcimadetudo(calcado, calca, mangacurta)

				if mangacurta == "t'shirt" {
					//se camisa
					for _, camisa := range p.Camisas {
						p.Caiemcimadetudo(calcado, calca, mangacurta, camisa, p.Gravata)
						p.Caiemcimadetudo(calcado, calca, mangacurta, camisa, p.PullOver)
						p.Caiemcimadetudo(calcado, calca, mangacurta, camisa, p.Gravata, p.PullOver)
						p.Caiemcimadetudo(calcado, calca, mangacurta, camisa)

					}
				}
			}

			for _, camisa := range p.Camisas {
				p.Caiemcimadetudo(calcado, calca, camisa, p.Gravata)
				p.Caiemcimadetudo(calcado, calca, camisa, p.PullOver)
				p.Caiemcimadetudo(calcado, calca, camisa, p.Gravata, p.PullOver)
				p.Caiemcimadetudo(calcado, calca, camisa)
			}
		}
	}
	p.ApplyFilters(Filters)
	return p
}

func main() {

	app := cli.NewApp()
	app.Name = "Fardamento"
	app.Usage = "Combinações"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "short",

			Usage: "mostrar mapeamento numerico",
		},
		cli.BoolFlag{
			Name: "full",

			Usage: "mostrar texto completo",
		},
		cli.BoolFlag{
			Name:  "tabela",
			Usage: "mostrar a legenda (mapeamento numeros pecas)",
		},
		cli.BoolFlag{
			Name:  "len",
			Usage: "mostrar o numero de resultados",
		},
		cli.StringFlag{
			Name:  "config",
			Usage: "define a localização do ficheiro de configuração",
			Value: "config.json",
		},
	}

	app.Action = func(c *cli.Context) {

		f, err := os.Open(c.String("config"))
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		p := CombinaPecas(f)

		if c.Bool("short") {
			p.PrintShort()

		}

		if c.Bool("full") {
			p.PrintFull()

		}

		if c.Bool("len") {
			fmt.Println(p.C.Len())

		}

		if c.Bool("tabela") {
			p.PrintLegenda()

		}

	}

	app.Run(os.Args)

}
