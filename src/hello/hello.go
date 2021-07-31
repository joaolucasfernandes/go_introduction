package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 5
const delay = 3

func main() {
	exibeIntro()
	for {
		exibeMenu()
		comando := recebeComando()
		processaComando(comando)
	}
}

func iniciaMonitoramento() {
	fmt.Println("Iniciando o monitoramento...")
	sites := listaSitesTxT()
	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			monitoraSite(site)
		}
		time.Sleep(delay * time.Second)
	}
	fmt.Println("")
}

func monitoraSite(site string) {
	result, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um errro:", err)
	}
	if result.StatusCode == 200 {
		fmt.Println("Site ", site, " funcionando normalmente")
	} else {
		fmt.Println("Site", site, " com problemas. Status code:", result.StatusCode)
	}
}

func processaComando(comando int) {
	switch comando {
	case 1:
		iniciaMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs")
	case 3:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço esse comando")
		os.Exit(-1)
	}
}

func exibeIntro() {
	nome := "Joao"
	versao := 1.1
	fmt.Println("Olá", nome)
	fmt.Println("Versão do Software - ", versao)
}

func exibeMenu() {
	fmt.Println("O que deseja fazer?")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair")
}

func recebeComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi: ", comandoLido)
	return comandoLido
}

func listaSitesTxT() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	return sites
}
