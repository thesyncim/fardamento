package main

import "strings"

type filter func(p *Pecas)

var Filters = []filter{
	FixDivisasSoCamisa,
	FixPolarEmCamisola,
	FixCalcaClassica,
	FixPullOverSemMAngas,
	FixDuplicated,
}

//so pode ter divisas se usar camisa
var FixDivisasSoCamisa = func(p *Pecas) {
	var ids []int
	for i, combinacao := range *p.C {
		if strings.Contains(combinacao, p.Divisas) {
			if !strings.Contains(combinacao, "camisa") {
				ids = append(ids, i)
			}
		}
	}

	p.C.removeindexs(ids)

}

//se for polar em camisola nao pode usar  calca classica nem casaco modelo blazer
var FixPolarEmCamisola = func(p *Pecas) {
	var ids []int
	for i, combinacao := range *p.C {
		if strings.Contains(combinacao, p.PullOver) {

			if strings.Contains(combinacao, "calça clasica") {
				ids = append(ids, i)
				continue
			}
			if strings.Contains(combinacao, "casaco modelo blazer") {
				ids = append(ids, i)
				continue
			}

		}
	}
	p.C.removeindexs(ids)

}

//se usar sapato classico nao pode usar calca operational nem cinto
var FixCalcaClassica = func(p *Pecas) {
	var ids []int

	for i := range *p.C {
		if strings.Contains((*p.C)[i], "sapato Classico") {

			if strings.Contains((*p.C)[i], "calça operacional") {
				ids = append(ids, i)
				continue
			}

			if strings.Contains((*p.C)[i], "cinturão velcro") {
				ids = append(ids, i)
				continue
			}

		}
	}
	p.C.removeindexs(ids)

}

//se usar Pullover sem mangas tem de usar Calça Clássica e uma camisa
var FixPullOverSemMAngas = func(p *Pecas) {
	var ids []int
	for i, combinacao := range *p.C {
		if strings.Contains(combinacao, "pullover sem mangas") {

			if !strings.Contains(combinacao, "calça clasica") {
				ids = append(ids, i)
				continue
			}

			if !strings.Contains(combinacao, "camisa") {
				ids = append(ids, i)
				continue

			}

		}
	}
	p.C.removeindexs(ids)

}

var FixDuplicated = func(p *Pecas) {
	result := []string{}
	seen := map[string]string{}
	for _, val := range *p.C {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	*p.C = result
}
