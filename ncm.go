package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	for {

		var (
			numLoja int
			ref     string
			ncm     string
			dbIP    string
		)

		// Solicitar ao usuário as variáveis
		fmt.Print("Digite a Referencia: ")
		fmt.Scan(&ref)

		for len(ncm) != 8 {
			fmt.Print("Digite o NCM Correto: ")
			fmt.Scan(&ncm)
		}

		// Limpar o buffer antes de ler a descrição
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		fmt.Print("Digite o Numero da loja: ")
		fmt.Scan(&numLoja)

		switch numLoja {
		case 5:
			dbIP = "192.168.5.50"
			fmt.Print("5 - LIMEIRA\n\n")
		case 10:
			dbIP = "192.168.10.50"
			fmt.Print("10 - UBERLANDIA\n\n")
		case 11:
			dbIP = "192.168.11.50"
			fmt.Print("11 - UBERABA\n\n")
		case 13:
			dbIP = "192.168.13.50"
			fmt.Print("13 - JUNDIAI\n\n")
		case 14:
			dbIP = "192.168.14.50"
			fmt.Print("14 - ANAPOLIS\n\n")
		case 17:
			dbIP = "192.168.17.50"
			fmt.Print("17 - ARACATUBA\n\n")
		case 18:
			dbIP = "192.168.18.50"
			fmt.Print("18 - RIO CLARO\n\n")
		case 19:
			dbIP = "192.168.19.50"
			fmt.Print("19 - ARARAQUARA\n\n")
		case 20:
			dbIP = "192.168.20.50"
			fmt.Print("20 - RIO BONITO\n\n")
		case 24:
			dbIP = "192.168.24.50"
			fmt.Print("24 - RIO PRETO\n\n")
		default:
			fmt.Println("\n\nLoja Nao Encontrada\nVou encerrar o Sistema\n")
			time.Sleep(time.Second * 5)
			log.Fatal()
		}

		// DSN (Data Source Name) para conectar ao banco de dados MySQL
		dsn := fmt.Sprintf("root:260803@tcp(%s:3306)/db_spartakus_manager", dbIP)

		// Abrir conexão com o banco de dados
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// Preparar a declaração SQL
		stmt, err := db.Prepare("UPDATE cad_produto SET ncm=? WHERE codigoProduto=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		// Executar a declaração SQL com os valores fornecidos
		_, err = stmt.Exec(ref, ncm)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Ncm Atualizado com sucesso", err)
		}
		time.Sleep(time.Second * 3)
	}
}
