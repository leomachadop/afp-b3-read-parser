package processor

import (
	"b3-read-parser/model"
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func ProcessB3(file string) {

	log.Printf("[ProcessB3] Processando %s\r", file)

	b3File, err := os.Open(file)
	if err != nil {
		log.Println("[ProcessB3] Houve um erro ao abrir o arquivo. Erro: ", err.Error())
		return
	}

	defer b3File.Close()

	scanner := bufio.NewScanner(b3File)
	for scanner.Scan() {
		linha := scanner.Text()

		if !strings.HasPrefix(linha, model.RecordTypeDetail) {
			continue
		}

		stockQuote := model.StockQuote{
			DataPregao: linha[model.SDataPregao:model.EDataPregao],
			CODBDI:     model.GetCODBDI(linha[model.SCODBDI:model.ECODBDI]),
			CODNEG:     strings.TrimSpace(linha[model.SCODNEG:model.ECODNEG]),
			TPMERC:     model.GetTPMERC(linha[model.STPMERC:model.ETPMERC]),
			NOMRES:     strings.TrimSpace(linha[model.SNOMRES:model.ENOMRES]),
			ESPECI:     model.GetESPECI(linha[model.SESPECI:model.EESPECI]),
			PRAZOT:     strings.TrimSpace(linha[model.SPRAZOT:model.EPRAZOT]),
			MODREF:     strings.TrimSpace(linha[model.SMODREF:model.EMODREF]),
			PREABE:     model.ConvertToDouble(strings.TrimSpace(linha[model.SPREABE:model.EPREABE])),
			PREMAX:     model.ConvertToDouble(strings.TrimSpace(linha[model.SPREMAX:model.EPREMAX])),
			PREMIN:     model.ConvertToDouble(strings.TrimSpace(linha[model.SPREMIN:model.EPREMIN])),
			PREMED:     model.ConvertToDouble(strings.TrimSpace(linha[model.SPREMED:model.EPREMED])),
			PREULT:     model.ConvertToDouble(strings.TrimSpace(linha[model.SPREULT:model.EPREULT])),
			PREOFC:     model.ConvertToDouble(strings.TrimSpace(linha[model.SPREOFC:model.EPREOFC])),
			PREOFV:     model.ConvertToDouble(strings.TrimSpace(linha[model.SPREOFV:model.EPREOFV])),
			TOTNEG:     model.ConvertToInt(strings.TrimSpace(linha[model.STOTNEG:model.ETOTNEG])),
			QUATOT:     model.ConvertToInt(strings.TrimSpace(linha[model.SQUATOT:model.EQUATOT])),
			VOLTOT:     model.ConvertToInt(strings.TrimSpace(linha[model.SVOLTOT:model.EVOLTOT])),
			PREEXE:     model.ConvertToDouble(strings.TrimSpace(linha[model.SPREEXE:model.EPREEXE])),
			INDOPC:     model.GetINDOPC(linha[model.SINDOPC:model.EINDOPC]),
			DATVEN:     strings.TrimSpace(linha[model.SDATVEN:model.EDATVEN]),
			FATCOT:     model.GetFATCOT(strings.TrimSpace(linha[model.SFATCOT:model.EFATCOT])),
			PTOEXE:     strings.TrimSpace(linha[model.SPTOEXE:model.EPTOEXE]),
			CODISI:     strings.TrimSpace(linha[model.SCODISI:model.ECODISI]),
			DISMES:     strings.TrimSpace(linha[model.SDISMES:model.EDISMES]),
		}

		histologically, err := json.MarshalIndent(stockQuote, "", "\t")
		if err != nil {
			log.Println("[ProcessB3] Houve um erro na geracao do objeto JSON: ", err.Error())
			return
		}

		wg.Add(1)

		go SendToKafka(histologically, wg)
	}

	wg.Wait()
}
