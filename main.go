package main

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

var (
	// raw url - https://www.viacaocometa.com.br/content/jca/cometa/pt-br/jcr:content.getServicos.json?origem=475&destino=467&data=2020-09-05
	baseURL           = "https://www.viacaocometa.com.br/content/jca/cometa/pt-br/jcr:content.getServicos.json"
	originID          = ""
	destinationID     = ""
	whatsappRecipient = ""
)

func main() {
	// load configurations
	LoadConfig()

	// find next weekend
	nextWeekend := GetNextWeekday(time.Now(), Sunday)
	searchDate := nextWeekend.Format("2006-01-02")
	log.Info("search date - ", searchDate)

	// build Search object
	search := &Search{
		BaseURL:       baseURL,
		Date:          searchDate,
		OriginID:      originID,
		DestinationID: destinationID,
	}
	log.Info("search url - ", search.GetURL())

	// get result
	searchResult := NewSearchResult(search)
	log.Info("search result", searchResult)

	message := "Viagens disponíveis para o dia " + searchDate
	// show results
	for _, result := range searchResult.ConsultaServicos.LsServicos {
		message += "\n========================="
		message += "\n Origem: " + result.Origem.Cidade
		message += "\n Destino: " + result.Destino.Cidade
		message += "\n Saída: " + result.Saida
		message += "\n Preço: " + result.Preco
		message += "\n Poltronas Livres: " + strconv.Itoa(result.PoltronasLivres)
	}
	log.Info(message)

	// message whatsapp
	SendMessage(message)
}
