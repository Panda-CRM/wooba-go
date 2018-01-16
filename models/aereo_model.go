package models

import "encoding/xml"

//XML Response 'VendasString'
type VendasList struct {
	XMLName xml.Name `xml:"vendas"`
	Vendas  []Venda  `xml:"venda"`
}

type Venda struct {
	XMLName                       xml.Name      `xml:"venda"`
	PontodeVenda                  string        `xml:"pontodeVenda"`
	Origem                        string        `xml:"origem"`
	IdVenda                       string        `xml:"idvenda"`
	TipoProduto                   string        `xml:"tipoproduto"`
	ClasseProduto                 string        `xml:"clasproduto"`
	IdEmissor                     string        `xml:"idemissor"`
	DataEmissao                   string        `xml:"dtemissao"`
	NomeAgencia                   string        `xml:"nmagencia"`
	IdCliente                     string        `xml:"idcliente"`
	IdBackoffice                  string        `xml:"IdBackoffice"`
	IdBackofficeIntegracao        string        `xml:"IdBackofficeIntegracao"`
	IdERPUnidadeDeNegocio         string        `xml:"IdERPUnidadeDeNegocio"`
	Loja                          string        `xml:"loja"`
	CodigoUnidade                 string        `xml:"codunidade"`
	IdOperador                    string        `xml:"idoperador"`
	IdFornecedor                  string        `xml:"idfornecedor"`
	FormaPagamentoAgencia         string        `xml:"formrec"`
	VencimentoPagamentoAgencia    string        `xml:"vencrec"`
	FormaPagamentoFornecedor      string        `xml:"formpagto"`
	VencimentoPagamentoFornecedor string        `xml:"vencpagto"`
	CartaoCia                     string        `xml:"cartaocia"`
	CartaoNumero                  string        `xml:"cartaonumero"`
	CartaoValidade                string        `xml:"cartaovalidade"`
	CartaoCodigoSeguranca         string        `xml:"cartaocodseguranca"`
	CartaoTitular                 string        `xml:"cartaotitular"`
	CentroCustoCliente            string        `xml:"centrocustocli"`
	CartaoCentroDeCusto           string        `xml:"ccrCentroDeCusto"`
	CartaoProjeto                 string        `xml:"ccrProjeto"`
	CartaoRequisicao              string        `xml:"ccrRequisicao"`
	CartaoEmpenho                 string        `xml:"ccrEmpenho"`
	Departamento                  string        `xml:"departamento"`
	UnidadeDeNegocio              string        `xml:"UnidadeDeNegocio"`
	Projeto                       string        `xml:"Projeto"`
	Matricula                     string        `xml:"Matricula"`
	Solicitante                   string        `xml:"solicitante"`
	Justificativa                 string        `xml:"justificativa"`
	MotivoViagem                  string        `xml:"MotivoViagem"`
	MenorTarifa                   string        `xml:"MenorTarifa"`
	TarifaCheia                   string        `xml:"TarifaCheia"`
	TarifaBase                    string        `xml:"TarifaBase"`
	MaiorTarifa                   string        `xml:"MaiorTarifa"`
	Status                        string        `xml:"Status"`
	DescricaoStatus               string        `xml:"dsStatus"`
	TaxaDeServico                 TaxaDeServico `xml:"taxadeservico"`
	Movimentos                    []Movimento   `xml:"movimentos"`
	Trechos                       []Trecho      `xml:"trechos"`
}

type Movimento struct {
	XMLName        xml.Name       `xml:"movimentos"`
	AeroRodoviario AeroRodoviario `xml:"aereorodoviario"`
}

type AeroRodoviario struct {
	XMLName            xml.Name `xml:"aereorodoviario"`
	Pax                string   `xml:"pax"`
	Tipo               string   `xml:"tipo"`
	Basetarifaria      string   `xml:"basetarifaria"`
	Tourcode           string   `xml:"tourcode"`
	Tarifa             string   `xml:"tarifa"`
	InvTarifa          string   `xml:"invTarifa"`
	CcTarifa           string   `xml:"ccTarifa"`
	GrTarifa           string   `xml:"grTarifa"`
	TktTarifa          string   `xml:"tktTarifa"`
	Moeda              string   `xml:"moeda"`
	Cambio             string   `xml:"cambio"`
	Datacambio         string   `xml:"datacambio"`
	Taxaembarque       string   `xml:"taxaembarque"`
	Fee                string   `xml:"fee"`
	Taxaservico        string   `xml:"taxaservico"`
	Markup             string   `xml:"markup"`
	Repassedu          string   `xml:"repassedu"`
	Repassecartaodu    string   `xml:"repassecartaodu"`
	MaiorTarifa        string   `xml:"maiorTarifa"`
	MenorTarifa        string   `xml:"menorTarifa"`
	Comisrecforvalor   string   `xml:"comisrecforvalor"`
	ComissaoRepassada  string   `xml:"ComissaoRepassada"`
	IncentivoRecebido  string   `xml:"IncentivoRecebido"`
	IncentivoRepassado string   `xml:"IncentivoRepassado"`
	Descpagclivalor    string   `xml:"descpagclivalor"`
	Observacao         string   `xml:"observacao"`
	Numautorizacao     string   `xml:"numautorizacao"`
}

type Bilhetes struct {
	XMLName  xml.Name  `xml:"bilhetes"`
	Bilhetes []Bilhete `xml:"bilhete"`
}

type Bilhete struct {
	XMLName         xml.Name `xml:"bilhete"`
	Numero          string   `xml:"numero"`
	Reemissao       string   `xml:"reemissao"`
	Reemissaotipo   string   `xml:"reemissaotipo"`
	Bilheteoriginal string   `xml:"bilheteoriginal"`
	Mcos            string   `xml:"mcos"`
	Conjugados      string   `xml:"conjugados"`
}

type Trecho struct {
	XMLName   xml.Name   `xml:"trechos"`
	Segmentos []Segmento `xml:"segmento"`
}

type Segmento struct {
	XMLName     xml.Name `xml:"segmento"`
	Localizador string   `xml:"localizador"`
	De          string   `xml:"de"`
	Para        string   `xml:"para"`
	DataSaida   string   `xml:"datasaida"`
	DataChegada string   `xml:"datachegada"`
	HoraSaida   string   `xml:"horasaida"`
	HoraChegada string   `xml:"horachegada"`
	Voo         string   `xml:"voo"`
	Classe      string   `xml:"classe"`
	Status      string   `xml:"status"`
	Conexao     string   `xml:"conexao"`
	CiaAerea    string   `xml:"ciaaerea"`
}

type TaxaDeServico struct {
	XMLName                 xml.Name `xml:"taxadeservico"`
	Identificador           string   `xml:"identificador"`
	Statusid                string   `xml:"statusid"`
	Status                  string   `xml:"status"`
	Estabelecimentoid       string   `xml:"estabelecimentoid"`
	Estabelecimento         string   `xml:"estabelecimento"`
	Formaid                 string   `xml:"formaid"`
	Forma                   string   `xml:"forma"`
	Moeda                   string   `xml:"moeda"`
	Valor                   string   `xml:"valor"`
	Cartaobandeiraid        string   `xml:"cartaobandeiraid"`
	Cartaobandeira          string   `xml:"cartaobandeira"`
	Cartaonumero            string   `xml:"cartaonumero"`
	Cartaotitular           string   `xml:"cartaotitular"`
	Cartaotitularcpf        string   `xml:"cartaotitularcpf"`
	Cartaocodigodeseguranca string   `xml:"cartaocodigodeseguranca"`
	Cartaovalidade          string   `xml:"cartaovalidade"`
	Transacaoautorizacao    string   `xml:"transacaoautorizacao"`
	Transacaoidentificacao  string   `xml:"transacaoidentificacao"`
	Parcelas                string   `xml:"parcelas"`
	TaxaDescricao           string   `xml:"taxaDescricao"`
	Responsavelnome         string   `xml:"responsavelnome"`
	Responsavelcpf          string   `xml:"responsavelcpf"`
	Responsaveltelefone     string   `xml:"responsaveltelefone"`
	Responsavelemail        string   `xml:"responsavelemail"`
}

type VendasErros struct {
	XMLName xml.Name `xml:"Vendas"`
	Erros []Erro `xml:"Erros"`
}

type Erro struct {
	XMLName xml.Name `xml:"Erros"`
	Erro string `xml:"Erro"`
}