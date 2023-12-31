package handler

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 테스트 코드 내에서 사용하게될 공용 변수인 mux를 전역변수 처리
var exam_mux = http.NewServeMux()

// 'exampleHandler.go' 패키지에서 정의한 FooHandler 객체를 테스트하는 함수
func TestExampleFooHandelr(t *testing.T) {
	test := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	exam_mux.Handle("/foo", &FooHandler{})
	exam_mux.ServeHTTP(res, req)

	data, _ := ioutil.ReadAll(res.Body)
	test.Equal("Hello foo", string(data))
}

/*
'exampleHandler.go' 패키지에서 정의한 barhandler 함수를 테스트하는 함수
*/
func TestExampleBarHanlder_WithOutName(t *testing.T) {
	// assert 객체를 사용하기위해 생성
	test := assert.New(t)
	// response 정보를 담아주는 response 작성 객체 생성
	res := httptest.NewRecorder()
	// request 정보가 담기는 객체 생성 'method', 'target(endpoint)', 'body' 정보를 파라미터로 전달받음
	req := httptest.NewRequest("GET", "/bar", nil)

	//전역 변수 mux에 handler 함수 초기화
	exam_mux.HandleFunc("/bar", BarHandler)
	exam_mux.ServeHTTP(res, req)

	// assertration 객체의 Equal 함수를 통해 원하는 값이 전달되었는지 확인
	test.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	test.Equal("Hello Unknown!", string(data))
}

/*
공통된 영역을 처리하기위한 테스트 생명주기 설정이 필요
생명 주기에 필요한 요소로 판단되는 부분
1) mux 객체에 테스트하고자하는 함수 정의
2) 공통 변수 생성 및 초기화 로직이 필요
3) req 데이터 res 데이터 정합성을 위한 table driven 정의
*/
func TestExampleBarHanlder_WithName(t *testing.T) {
	test := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=CHOI", nil)

	exam_mux.ServeHTTP(res, req)

	data, _ := ioutil.ReadAll(res.Body)
	test.Equal("Hello CHOI!", string(data))
}

func TestExampleBarHanlder_gorilla(t *testing.T) {
	test := assert.New(t)
	mockServer := httptest.NewServer(NewMux())
	defer mockServer.Close()

	resp, err := http.Get(mockServer.URL + "/bar/choi")
	if err != nil {
		t.Fatal(http.StatusBadRequest)
	}
	data, _ := ioutil.ReadAll(resp.Body)
	test.Equal("Hello choi user", string(data))

}
