package model

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

/**
Os campos serão formatados com o seguinte layout
SNOMECAMPO - Posição inicial
ENOMECAMPO - Posição final
*/
const (
	RecordTypeDetail = "01"

	SDataPregao = 2
	EDataPregao = 10

	SCODBDI = 10
	ECODBDI = 12

	SCODNEG = 12
	ECODNEG = 24

	STPMERC = 24
	ETPMERC = 27

	SNOMRES = 27
	ENOMRES = 39

	SESPECI = 39
	EESPECI = 49

	SPRAZOT = 49
	EPRAZOT = 52

	SMODREF = 52
	EMODREF = 56

	SPREABE = 56
	EPREABE = 69

	SPREMAX = 69
	EPREMAX = 82

	SPREMIN = 82
	EPREMIN = 95

	SPREMED = 95
	EPREMED = 108

	SPREULT = 108
	EPREULT = 121

	SPREOFC = 121
	EPREOFC = 134

	SPREOFV = 134
	EPREOFV = 147

	STOTNEG = 148
	ETOTNEG = 152

	SQUATOT = 152
	EQUATOT = 170

	SVOLTOT = 171
	EVOLTOT = 188

	SPREEXE = 188
	EPREEXE = 201

	SINDOPC = 201
	EINDOPC = 202

	SDATVEN = 202
	EDATVEN = 210

	SFATCOT = 210
	EFATCOT = 217

	SPTOEXE = 217
	EPTOEXE = 230

	SCODISI = 230
	ECODISI = 242

	SDISMES = 242
	EDISMES = 245
)

func ConvertToDouble(value string) float64 {
	floatNumber, err := strconv.ParseFloat(value[:len(value)-2]+"."+value[len(value)-2:], 2)

	if err != nil {
		log.Printf("ConvertToDouble err str %s", value)
	}

	return floatNumber
}

func ConvertToInt(value string) int {
	intNumber, err := strconv.Atoi(value)

	if err != nil {
		log.Printf("ConvertToDouble err str %s", value)
	}

	return intNumber
}

func GetFATCOT(FATCOT string) string {
	switch FATCOT {
	case "1":
		return "COTAÇÃO UNITÁRIA"
	case "1000":
		return "COTAÇÃO POR LOTE DE MIL AÇÕES"
	default:
		return "N/A"
	}
}

func GetCODBDI(CODBDI string) string {
	switch CODBDI {
	case "02":
		return "LOTE PADRAO"
	case "05":
		return "SANCIONADAS PELOS REGULAMENTOS BMFBOVESPA"
	case "06":
		return "CONCORDATARIAS"
	case "07":
		return "RECUPERACAO EXTRAJUDICIAL"
	case "08":
		return "RECUPERAÇÃO JUDICIAL"
	case "09":
		return "RAET - REGIME DE ADMINISTRACAO ESPECIAL TEMPORARIA"
	case "10":
		return "DIREITOS E RECIBOS"
	case "11":
		return "INTERVENCAO"
	case "12":
		return "FUNDOS IMOBILIARIOS"
	case "14":
		return "CERT.INVEST/TIT.DIV.PUBLICA"
	case "18":
		return "OBRIGACÕES"
	case "22":
		return "BÔNUS (PRIVADOS)"
	case "26":
		return "APOLICES/BÔNUS/TITULOS PUBLICOS"
	case "32":
		return "EXERCICIO DE OPCOES DE COMPRA DE INDICES"
	case "33":
		return "EXERCICIO DE OPCOES DE VENDA DE INDICES"
	case "38":
		return "EXERCICIO DE OPCOES DE COMPRA"
	case "42":
		return "EXERCICIO DE OPCOES DE VENDA"
	case "46":
		return "LEILAO DE NAO COTADOS"
	case "48":
		return "LEILAO DE PRIVATIZACAO"
	case "49":
		return "LEILAO DO FUNDO RECUPERACAO ECONOMICA ESPIRITO SANTO"
	case "50":
		return "LEILAO"
	case "51":
		return "LEILAO FINOR"
	case "52":
		return "LEILAO FINAM"
	case "53":
		return "LEILAO FISET"
	case "54":
		return "LEILAO DE ACÕES EM MORA"
	case "56":
		return "VENDAS POR ALVARA JUDICIAL"
	case "58":
		return "OUTROS"
	case "60":
		return "PERMUTA POR ACÕES"
	case "61":
		return "META"
	case "62":
		return "MERCADO A TERMO"
	case "66":
		return "DEBENTURES COM DATA DE VENCIMENTO ATE 3 ANOS"
	case "68":
		return "DEBENTURES COM DATA DE VENCIMENTO MAIOR QUE 3 ANOS"
	case "70":
		return "FUTURO COM RETENCAO DE GANHOS"
	case "71":
		return "MERCADO DE FUTURO"
	case "74":
		return "OPCOES DE COMPRA DE INDICES"
	case "75":
		return "OPCOES DE VENDA DE INDICES"
	case "78":
		return "OPCOES DE COMPRA"
	case "82":
		return "OPCOES DE VENDA"
	case "83":
		return "BOVESPAFIX"
	case "84":
		return "SOMA FIX"
	case "90":
		return "TERMO VISTA REGISTRADO"
	case "96":
		return "MERCADO FRACIONARIO"
	case "99":
		return "TOTAL GERAL"
	default:
		return "N/A"
	}
}

func GetTPMERC(TPMERC string) string {
	switch TPMERC {
	case "010":
		return "VISTA"
	case "012":
		return "EXERCÍCIO DE OPÇÕES DE COMPRA"
	case "013":
		return "EXERCÍCIO DE OPÇÕES DE VENDA"
	case "017":
		return "LEILÃO"
	case "020":
		return "FRACIONÁRIO"
	case "030":
		return "TERMO"
	case "050":
		return "FUTURO COM RETENÇÃO DE GANHO"
	case "060":
		return "FUTURO COM MOVIMENTAÇÃO CONTÍNUA"
	case "070":
		return "OPÇÕES DE COMPRA"
	case "080":
		return "OPÇÕES DE VENDA"
	default:
		return "N/A"
	}
}

var ReSpaces = regexp.MustCompile(`\s+`)

func GetESPECI(ESPECI string) string {
	var especiSpaces = ReSpaces.ReplaceAllString(ESPECI, " ")
	switch {
	case strings.HasPrefix(especiSpaces, "BDR"):
		return "BDR"
	case strings.HasPrefix(especiSpaces, "BNS"):
		return "BÔNUS DE SUBSCRIÇÃO EM ACÕES MISCELÂNEA"
	case strings.HasPrefix(especiSpaces, "BNS B/A"):
		return "BÔNUS DE SUBSCRIÇÃO EM ACÕES PREFERÊNCIA"
	case strings.HasPrefix(especiSpaces, "BNS ORD"):
		return "BÔNUS DE SUBSCRIÇÃO EM ACÕES ORDINÁRIAS"
	case strings.HasPrefix(especiSpaces, "BNS P/A"):
		return "BÔNUS DE SUBSCRIÇÃO EM ACÕES PREFERÊNCIA"
	case strings.HasPrefix(especiSpaces, "BNS P/B"):
		return "BÔNUS DE SUBSCRIÇÃO EM ACÕES PREFERÊNCIA"
	case strings.HasPrefix(especiSpaces, "BNS P/C"):
		return "BÔNUS DE SUBSCRIÇÃO EM ACÕES PREFERÊNCIA"
	case strings.HasPrefix(especiSpaces, "BNS P/D"):
		return "BÔNUS DE SUBSCRIÇÃO EM ACÕES PREFERÊNCIA"
	case strings.HasPrefix(especiSpaces, "BNS P/E"):
		return "BÔNUS DE SUBSCRIÇÃO EM ACÕES PREFERÊNCIA"
	case strings.HasPrefix(especiSpaces, "BNS P/F"):
		return "BÔNUS DE SUBSCRIÇÃO EM ACÕES PREFERÊNCIA"
	case strings.HasPrefix(especiSpaces, "BNS P/G"):
		return "BÔNUS DE SUBSCRIÇÃO EM ACÕES PREFERÊNCIA"
	case strings.HasPrefix(especiSpaces, "BNS P/H"):
		return "BÔNUS DE SUBSCRIÇÃO EM ACÕES PREFERÊNCIA"
	case strings.HasPrefix(especiSpaces, "BNS PRE"):
		return "BÔNUS DE SUBSCRIÇÃO EM ACÕES PREFERÊNCIA"
	case strings.HasPrefix(especiSpaces, "CDA"):
		return "CERTIFICADO DE DEPÓSITO DE ACÕES ORDINÁRIAS"
	case strings.HasPrefix(especiSpaces, "CI"):
		return "FUNDO DE INVESTIMENTO"
	case strings.HasPrefix(especiSpaces, "CI ATZ"):
		return "Fundo de Investimento Atualização"
	case strings.HasPrefix(especiSpaces, "CI EA"):
		return "Fundo de Investimento Ex-Atualização"
	case strings.HasPrefix(especiSpaces, "CI EBA"):
		return "Fundo de Investimento Ex-Bonificação e Ex-Atualização"
	case strings.HasPrefix(especiSpaces, "CI ED"):
		return "Fundo de Investimento Ex-dividendo"
	case strings.HasPrefix(especiSpaces, "CI ER"):
		return "Fundo de Investimento Ex-Rendimento"
	case strings.HasPrefix(especiSpaces, "CI ERA"):
		return "Fundo de Investimento Ex-rendimento e Ex-Atualização"
	case strings.HasPrefix(especiSpaces, "CI ERB"):
		return "Fundo de Investimento Ex-rendimento e Ex-Bonificação"
	case strings.HasPrefix(especiSpaces, "CI ERS"):
		return "Fundo de Investimento Ex-Rendimento e Ex-Subscrição"
	case strings.HasPrefix(especiSpaces, "CI ES"):
		return "Fundo de Investimento Ex-Subscrição"
	case strings.HasPrefix(especiSpaces, "CPA"):
		return "CERTIF. DE POTENCIAL ADIC. DE CONSTRUÇÃO"
	case strings.HasPrefix(especiSpaces, "DIR"):
		return "DIREITOS DE SUBSCRIÇÃO MISCELÂNEA (BÔNUS)"
	case strings.HasPrefix(especiSpaces, "DIR DEB"):
		return "Direito de Debênture"
	case strings.HasPrefix(especiSpaces, "DIR ORD"):
		return "DIREITOS DE SUBSCRIÇÃO EM ACÕES ORDINÁRIAS"
	case strings.HasPrefix(especiSpaces, "DIR P/A"):
		return "DIREITOS DE SUBSCRIÇÃO EM ACÕES PREFERENCIAIS"
	case strings.HasPrefix(especiSpaces, "DIR P/B"):
		return "DIREITOS DE SUBSCRIÇÃO EM ACÕES PREFERENCIAIS"
	case strings.HasPrefix(especiSpaces, "DIR P/C"):
		return "DIREITOS DE SUBSCRIÇÃO EM ACÕES PREFERENCIAIS"
	case strings.HasPrefix(especiSpaces, "DIR P/D"):
		return "DIREITOS DE SUBSCRIÇÃO EM ACÕES PREFERENCIAIS"
	case strings.HasPrefix(especiSpaces, "DIR P/E"):
		return "DIREITOS DE SUBSCRIÇÃO EM ACÕES PREFERENCIAIS"
	case strings.HasPrefix(especiSpaces, "DIR P/F"):
		return "DIREITOS DE SUBSCRIÇÃO EM ACÕES PREFERENCIAIS"
	case strings.HasPrefix(especiSpaces, "DIR P/G"):
		return "DIREITOS DE SUBSCRIÇÃO EM ACÕES PREFERENCIAIS"
	case strings.HasPrefix(especiSpaces, "DIR P/H"):
		return "DIREITOS DE SUBSCRIÇÃO EM ACÕES PREFERENCIAIS"
	case strings.HasPrefix(especiSpaces, "DIR PR"):
		return "DIREITOS DE SUBSCRIÇÃO EM ACÕES RESGATÁVEIS"
	case strings.HasPrefix(especiSpaces, "DIR PRA"):
		return "DIREITOS DE SUBSCRIÇÃO EM ACÕES RESGATÁVEIS"
	case strings.HasPrefix(especiSpaces, "DIR PRB"):
		return "DIREITOS DE SUBSCRIÇÃO EM ACÕES RESGATÁVEIS"
	case strings.HasPrefix(especiSpaces, "DIR PRC"):
		return "DIREITOS DE SUBSCRIÇÃO EM ACÕES RESGATÁVEIS"
	case strings.HasPrefix(especiSpaces, "DIR PRE"):
		return "DIREITOS DE SUBSCRIÇÃO EM ACÕES PREFERENCIAIS"
	case strings.HasPrefix(especiSpaces, "FIDC"):
		return "Fundo de Investimento em Direitos Creditórios"
	case strings.HasPrefix(especiSpaces, "LFT"):
		return "LETRA FINANCEIRA DO TESOURO"
	case strings.HasPrefix(especiSpaces, "M1 REC"):
		return "RECIBO DE SUBSCRIÇÃO DE MISCELÂNEAS"
	case strings.HasPrefix(especiSpaces, "ON"):
		return "ACÕES ORDINÁRIAS NOMINATIVAS"
	case strings.HasPrefix(especiSpaces, "ON ATZ"):
		return "Ações Ordinárias Atualização"
	case strings.HasPrefix(especiSpaces, "ON EB"):
		return "Ações Ordinárias Ex-Bonificação"
	case strings.HasPrefix(especiSpaces, "ON ED"):
		return "Ações Ordinárias Ex-Dividendo"
	case strings.HasPrefix(especiSpaces, "ON EDB"):
		return "Ações Ordinárias Ex-Dividendo e Ex-Bonificação"
	case strings.HasPrefix(especiSpaces, "ON EDJ"):
		return "Ações Ordinárias Ex-dividendo e Ex-Juros"
	case strings.HasPrefix(especiSpaces, "ON EDR"):
		return "Ações Ordinárias Ex-Dividendo e Ex-Rendimento"
	case strings.HasPrefix(especiSpaces, "ON EG"):
		return "Ações Ordinárias Ex-Grupamento"
	case strings.HasPrefix(especiSpaces, "ON EJ"):
		return "Ações Ordinárias Ex-juros"
	case strings.HasPrefix(especiSpaces, "ON EJB"):
		return "Ações Ordinárias Ex-juros e Ex-bonificação"
	case strings.HasPrefix(especiSpaces, "ON EJS"):
		return "Ações Ordinárias Ex-Juros e Ex-Subscrição"
	case strings.HasPrefix(especiSpaces, "ON ER"):
		return "Ações Ordinárias Ex-Rendimento"
	case strings.HasPrefix(especiSpaces, "ON ERJ"):
		return "Ações Ordinárias Ex-Rendimento e Ex-Juros"
	case strings.HasPrefix(especiSpaces, "ON ES"):
		return "Ações Ordinárias Ex-Subscrição"
	case strings.HasPrefix(especiSpaces, "ON P"):
		return "ACÕES ORDINÁRIAS NOMINATIVAS COM DIREITO"
	case strings.HasPrefix(especiSpaces, "ON REC"):
		return "RECIBO DE SUBSCRIÇÃO EM ACÕES ORDINÁRIAS"
	case strings.HasPrefix(especiSpaces, "OR"):
		return "ACÕES ORDINÁRIAS NOMINATIVAS RESGATÁVEIS"
	case strings.HasPrefix(especiSpaces, "OR P"):
		return "ACÕES ORDINÁRIAS NOMINATIVAS RESGATÁVEIS"
	case strings.HasPrefix(especiSpaces, "PCD"):
		return "POSIÇÃO CONSOLIDADA DA DIVIDA"
	case strings.HasPrefix(especiSpaces, "PN"):
		return "ACÕES PREFERÊNCIAIS NOMINATIVAS"
	case strings.HasPrefix(especiSpaces, "PN EB"):
		return "Ações Preferenciais Ex-Bonificação"
	case strings.HasPrefix(especiSpaces, "PN ED"):
		return "Ações Preferenciais Ex-Dividendo"
	case strings.HasPrefix(especiSpaces, "PN EDB"):
		return "Ações Preferenciais Ex-Dividendo e Ex-Bonificação"
	case strings.HasPrefix(especiSpaces, "PN EDJ"):
		return "Ações Preferenciais Ex-dividendo e Ex-Juros"
	case strings.HasPrefix(especiSpaces, "PN EDR"):
		return "Ações Preferenciais Ex-Dividendo e Ex-Rendimento"
	case strings.HasPrefix(especiSpaces, "PN EJ"):
		return "Ações Preferenciais Ex-Juros"
	case strings.HasPrefix(especiSpaces, "PN EJB"):
		return "Ações Preferenciais Ex-juros e Ex-bonificação"
	case strings.HasPrefix(especiSpaces, "PN EJS"):
		return "Ações Preferenciais Ex-Juros e Ex-Subscrição"
	case strings.HasPrefix(especiSpaces, "PN ES"):
		return "Ações Preferenciais Ex-Subscrição"
	case strings.HasPrefix(especiSpaces, "PN P"):
		return "ACÕES PREFERÊNCIAIS NOMINATIVAS COM DIREITO"
	case strings.HasPrefix(especiSpaces, "PN REC"):
		return "RECIBO DE SUBSCRIÇÃO EM ACÕES PREFERENCIAIS"
	case strings.HasPrefix(especiSpaces, "PNA"):
		return "ACÕES PREFERÊNCIAIS NOMINATIVAS CLASSE A"
	case strings.HasPrefix(especiSpaces, "PNA EB"):
		return "Ações Preferenciais Classe A Ex-BonificaçãoPreferencial"
	case strings.HasPrefix(especiSpaces, "PNA EDR"):
		return "Ações Preferenciais Classe A Ex-Dividendo e Ex-Rendimento"
	case strings.HasPrefix(especiSpaces, "PNA EJ"):
		return "Ações Preferenciais Classe A Ex-Juros"
	case strings.HasPrefix(especiSpaces, "PNA ES"):
		return "Ações Preferenciais Classe A Preferencial Ex-Subscrição"
	case strings.HasPrefix(especiSpaces, "PNA P"):
		return "ACÕES PREFERÊNCIAIS NOMINATIVAS CLASSE A"
	case strings.HasPrefix(especiSpaces, "PNA REC"):
		return "RECIBO DE SUBSCRIÇÃO EM ACÕES PREFERENCIAIS"
	case strings.HasPrefix(especiSpaces, "PNB"):
		return "ACÕES PREFERÊNCIAIS NOMINATIVAS CLASSE B"
	case strings.HasPrefix(especiSpaces, "PNB EB"):
		return "Ações Preferenciais Classe B Ex-Bonificação"
	case strings.HasPrefix(especiSpaces, "PNB ED"):
		return "Ações Preferenciais Classe B Ex-Dividendo"
	case strings.HasPrefix(especiSpaces, "PNB EDR"):
		return "Ações Preferenciais Classe B Ex-Dividendo e Ex-Rendimento"
	case strings.HasPrefix(especiSpaces, "PNB EJ"):
		return "Ações Preferenciais Classe B Ex-Juros"
	case strings.HasPrefix(especiSpaces, "PNB P"):
		return "ACÕES PREFERÊNCIAIS NOMINATIVAS CLASSE B"
	case strings.HasPrefix(especiSpaces, "PNB REC"):
		return "RECIBO DE SUBSCRIÇÃO EM ACÕES PREFERENCIAIS"
	case strings.HasPrefix(especiSpaces, "PNC"):
		return "ACÕES PREFERÊNCIAIS NOMINATIVAS CLASSE C"
	case strings.HasPrefix(especiSpaces, "PNC ED"):
		return "Ações Preferenciais Classe C Preferencial Classe C Ex-Dividendo"
	case strings.HasPrefix(especiSpaces, "PNC P"):
		return "ACÕES PREFERÊNCIAIS NOMINATIVAS CLASSE C"
	case strings.HasPrefix(especiSpaces, "PNC REC"):
		return "RECIBO DE SUBSCRIÇÃO EM ACÕES PREFERENCIAIS"
	case strings.HasPrefix(especiSpaces, "PND"):
		return "ACÕES PREFERÊNCIAIS NOMINATIVAS CLASSE D"
	case strings.HasPrefix(especiSpaces, "PND ED"):
		return "Ações Preferenciais Classe D Ex-Dividendo"
	case strings.HasPrefix(especiSpaces, "PND P"):
		return "ACÕES PREFERÊNCIAIS NOMINATIVAS CLASSE D"
	case strings.HasPrefix(especiSpaces, "PND REC"):
		return "RECIBO DE SUBSCRIÇÃO EM ACÕES PREFERENCIAIS"
	case strings.HasPrefix(especiSpaces, "PNE"):
		return "ACÕES PREFERÊNCIAIS NOMINATIVAS CLASSE E"
	case strings.HasPrefix(especiSpaces, "PNE ED"):
		return "Ações Preferenciais Classe E Ex-Dividendo"
	case strings.HasPrefix(especiSpaces, "PNE P"):
		return "ACÕES PREFERÊNCIAIS NOMINATIVAS CLASSE E"
	case strings.HasPrefix(especiSpaces, "PNE REC"):
		return "RECIBO DE SUBSCRIÇÃO EM ACÕES PREFERENCIAIS"
	default:
		return "N/A"
	}
}

func GetINDOPC(INDOPC string) string {
	switch INDOPC {
	case "1":
		return "US$ CORREÇÃO PELA TAXA DO DÓLAR 2 TJLP CORREÇÃO PELA TJLP"
	case "8":
		return "IGPM CORREÇÃO PELO IGP-M - OPÇÕES PROTEGIDAS"
	case "9":
		return "URV CORREÇÃO PELA URV"
	default:
		return "N/A"
	}
}
