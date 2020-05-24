package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main() {
	// 서버에 접속
	client, err := net.Dial("tcp", "127.0.0.1:44444")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	// 서버 메시지 대기
	go func(c net.Conn) {
		data := make([]byte, 4096)

		for {
			// 서버 메시지 대기
			n, err := c.Read(data)
			if err != nil {
				fmt.Println(err)
				return
			}

			// 서버 메시지를 화면에 표시
			fmt.Println(string(data[:n]))
		}
	}(client)

	// 사용자 입력 인터페이스
	in := bufio.NewReader(os.Stdin)

	for {
		// 사용자 입력 대기
		line, errToRead := in.ReadString('\n')
		if errToRead != nil {
			fmt.Println(errToRead)
			continue
		}

		// 공백 문자 제거
		// 사용자 입력이 공백이면, 다시 사용자 입력 대기
		line = strings.Trim(line, " \n\t")
		if len(line) == 0 {
			continue
		}

		// 사용자 입력 메시지를 서버로 전송
		// 서버로 전송 실패시 프로그램 종료
		_, errToWrite := client.Write([]byte(line))
		if errToWrite != nil {
			fmt.Println(errToWrite)
			return
		}
	}
}
