package main

import (
	"fmt"
	"net"
	"reflect"
	"strconv"
	"encoding/json"
)

// 접속한 클라이언트 아이디와 맵 정의
var lastConnId = 0;
var connById map[int]net.Conn = make(map[int]net.Conn)

// 연습 유틸리티: 화면에 오브젝트 속성 표시
func printObj(obj interface{}) {
    e := reflect.ValueOf(&obj).Elem()
    for i := 0; i < e.NumField(); i++ {
        fieldName := e.Type().Field(i).Name
        fmt.Printf("%v\n", fieldName)
    }
}

// 연습 유틸리티: 화면에 오브젝트를 JSNO 형식으로 표시
func printJson(obj interface{}) {
	s, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(string(s))
}

// 모든 클라이언트에 메시지 전송
func broadcast(senderId int, data []byte) {
	fmt.Println("Broadcast from Client " + strconv.Itoa(senderId) + " : " + string(data));

	// 모든 클라이언트 중
	for id, conn := range connById {
		// 보낸 클라이언트 제외하고
		if id == senderId {
			continue
		}

		// 메시지 전송
		_, err := conn.Write(data);
		if err != nil {
			fmt.Println("Error broadcast", err)
		}
	}
}

// 클라이언트 메시지 대기
func requestHandler(id int, c net.Conn) {
	data := make([]byte, 4096)

	fmt.Println("Client " + strconv.Itoa(id) + " connected")

	c.Write([]byte("Welcome to simple chat"))

	for {
		// 클라이언트 메시지 대기
		// 오류 발생시 클라이언트와 연결 종료
		n, err := c.Read(data)
		if err != nil {
			fmt.Println("Client " + strconv.Itoa(id) + " disconnected")

			c.Close()

			delete(connById, id)

			break
		}

		// 클라이언트 메시지 브로드캐스팅
		broadcast(id, data[:n])
	}
}

func main() {
	// 리스닝 설정
	ln, err := net.Listen("tcp", ":44444")
	if err != nil {
		fmt.Println("Error connect", err)
		return
	}
	defer ln.Close()

	for {
		// 클라이언트 접속 대기
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accept", err)
			continue
		}
		defer conn.Close()

		// 새 클라이언트 아이디 설정 및 맵에 추가
		lastConnId++
		connById[lastConnId] = conn

		// 새 클라이언트와 통신할 고루틴 함수 실행
		go requestHandler(lastConnId, conn)
	}
}

