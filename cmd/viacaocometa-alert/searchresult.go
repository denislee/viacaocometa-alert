package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SearchResult struct {
	ConsultaServicos struct {
		LsServicos []struct {
			Servico         string `json:"servico"`
			Grupo           string `json:"grupo"`
			Saida           string `json:"saida"`
			Chegada         string `json:"chegada"`
			PoltronasLivres int    `json:"poltronasLivres"`
			PoltronasTotal  int    `json:"poltronasTotal"`
			Preco           string `json:"preco"`
			Tarifa          string `json:"tarifa"`
			VlrTaxaEmbarque string `json:"vlrTaxaEmbarque"`
			VlrPedagio      string `json:"vlrPedagio"`
			VlrSeguro       string `json:"vlrSeguro"`
			VlrOutros       string `json:"vlrOutros"`
			Classe          string `json:"classe"`
			Empresa         string `json:"empresa"`
			MensagemServico string `json:"mensagemServico"`
			Vende           bool   `json:"vende"`
			Origem          struct {
				ID     int    `json:"id"`
				Cidade string `json:"cidade"`
			} `json:"origem"`
			Destino struct {
				ID     int    `json:"id"`
				Cidade string `json:"cidade"`
			} `json:"destino"`
		} `json:"lsServicos"`
		Erro struct {
			OcorreuErro     bool   `json:"ocorreuErro"`
			DescricaoErro   string `json:"descricaoErro"`
			OcorreuAlerta   bool   `json:"ocorreuAlerta"`
			DescricaoAlerta string `json:"descricaoAlerta"`
			SessionValid    bool   `json:"sessionValid"`
		} `json:"erro"`
	} `json:"consultaServicos"`
}

func (searchResult *SearchResult) Count() int {
	return len(searchResult.ConsultaServicos.LsServicos)
}

func NewSearchResult(search *Search) SearchResult {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", search.GetURL(), nil)
	resp, _ := client.Do(req)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var searchResult SearchResult
	if err := json.Unmarshal(bytes, &searchResult); err != nil {
		panic(err)
	}
	return searchResult
}
