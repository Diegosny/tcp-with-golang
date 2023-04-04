package main

import (
	"fmt"
	"net"
)

const DEFAULT_PORT = "8080"


func handleConnection(conn net.Conn) {
    buf := make([]byte, 1024)
    n, err := conn.Read(buf)
    if err != nil {
        fmt.Println("Erro ao ler mensagem do cliente:", err)
        return
    }
    msg := string(buf[:n])


    fmt.Println("Mensagem recebida:", msg)

    response := "Olá, cliente!"
    _, err = conn.Write([]byte(response))
    if err != nil {
        fmt.Println("Erro ao enviar resposta para o cliente:", err)
        return
    }

    conn.Close()
}

func main() {
	fmt.Println("******* RESPONSE ********")
	
    listener, err := net.Listen("tcp", DEFAULT_PORT)
	
	fmt.Println(listener)


    if err != nil {
        fmt.Println("Erro ao criar servidor:", err)
        return
    }
    fmt.Println("Servidor iniciado na porta 8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Erro ao aceitar conexão do cliente:", err)
            continue
        }

        go handleConnection(conn)
    }
}
