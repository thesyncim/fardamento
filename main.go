package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("config.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	p := new(Pecas)
	err = dec.Decode(p)
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
	fmt.Println(p.C.Len())
	p.ApplyFilters(Filters)
	fmt.Println(p.C.Len())

	p.PrintLegenda()
	p.PrintShort()
	p.PrintFull()

}
