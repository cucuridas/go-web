package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct {
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello foo")
}

func main() {
	//Request 요청이 왔을 떄 어떠한 행위를 할 지 handler를 정의하는 함수
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	//OOP 구조상 구조체를 먼저 정의를 한뒤 controller나 비즈니스 로직을 지원하도록 하는 형태
	//Service라는 비즈니스 로직 개발 후 API endpoint에서 사용하도록 분리하는 역할이라 판단됨
	http.Handle("/foo", &fooHandler{})
	// 해당 서버의 특정 포트를 열어 request를 대기하도록 하는 함수,
	// 정의되어진 endpoint와 연관되는 HandleFunc을 실행시키도록 하는 함수
	http.ListenAndServe(":3000", nil)
}
