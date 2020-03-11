package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/mrzon/jsonrpc2.0/model"
	"github.com/mrzon/jsonrpc2.0/util"
)

type RpcServerConnection struct {
	ServerList  map[int]*http.Server
	HandlerList map[int]*http.ServeMux
	Ports       []int
	isStarted   bool
	waitGroup   sync.WaitGroup
}

func NewRpcServerConnection() *RpcServerConnection {
	rpcServerConn := &RpcServerConnection{
		make(map[int]*http.Server),
		make(map[int]*http.ServeMux),
		make([]int, 0),
		false,
		sync.WaitGroup{},
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	go func() {
		select {
		case sig := <-c:
			log.Printf("Got %s signal. Shutting down...\n", sig)
			rpcServerConn.Shutdown()
		}
	}()

	return rpcServerConn
}

func (rpcServerConn *RpcServerConnection) Register(r *RpcServer) {
	if rpcServerConn.ServerList[r.Config.Port] == nil {
		var err error = nil

		rpcServerConn.HandlerList[r.Config.Port] = http.NewServeMux()
		rpcServerConn.ServerList[r.Config.Port] = &http.Server{
			Addr:           ":" + strconv.Itoa(r.Config.Port),
			Handler:        rpcServerConn.HandlerList[r.Config.Port],
			ReadTimeout:    r.Config.Timeout,
			WriteTimeout:   r.Config.Timeout,
			MaxHeaderBytes: 1 << 20,
		}
		rpcServerConn.Ports = append(rpcServerConn.Ports, r.Config.Port)
		if err != nil {
			panic(err)
		}
	}
	rpcServerConn.HandlerList[r.Config.Port].HandleFunc(r.Config.EndPoint, func(writer http.ResponseWriter,
		request *http.Request) {
		body, _ := ioutil.ReadAll(request.Body)

		jsonRequest := model.JsonRpcRequest{}
		json.Unmarshal(body, &jsonRequest)

		serviceType := reflect.TypeOf(r.Service)
		serviceValue := reflect.ValueOf(r.Service)
		if serviceType.Kind() == reflect.Ptr {
			serviceValue = serviceValue.Elem()
		}

		methodName := getMethodNameFromTagAndParam(r.Service, serviceValue, jsonRequest.Method, jsonRequest.Params)

		method := serviceValue.FieldByName(methodName)
		if !method.IsValid() {
			//if method name is not valid, the user might use tag to call the service
			jsonResponse := model.NewJsonRpcResponseWithError(map[string]string{
				"code": "-32601", "message": "Method not found",
			})

			responseBytes, _ := json.Marshal(jsonResponse)
			writer.Write(responseBytes)
			return
		}
		if r.Config.EnableLogging {
			log.Println("Method ", methodName, "is being called. Full Request:", string(body))
		}
		in := make([]reflect.Value, method.Type().NumIn())

		for i := 0; i < len(in); i++ {
			typeOf := method.Type().In(i)
			in[i] = util.GetVal(typeOf, jsonRequest.Params[i])
		}
		result := method.Call(in)
		var resultStr interface{}
		switch result[0].Kind() {
		case reflect.Int:
			resultStr = int(result[0].Int())
		case reflect.Int64:
			resultStr = result[0].Int
		case reflect.Float32:
			resultStr = float32(result[0].Float())
		case reflect.Float64:
			resultStr = result[0].Float()
		default:
			resultStr = result[0].Interface()
		}

		jsonResponse := model.NewJsonRpcResponseWithResult(resultStr)

		responseBytes, _ := json.Marshal(jsonResponse)
		writer.Header().Add("Content-Type", "application/json")
		writer.Write(responseBytes)
	})
}

/**
Non Blocking
*/
func (rpcServerConn *RpcServerConnection) Start() {
	if rpcServerConn.isStarted {
		return
	}
	for i := 0; i < len(rpcServerConn.Ports); i++ {
		go func(ind int) {
			log.Println("Listen and serve to port", rpcServerConn.Ports[ind])
			rpcServerConn.ServerList[rpcServerConn.Ports[ind]].ListenAndServe()
		}(i)
		rpcServerConn.waitGroup.Add(1)
	}
	rpcServerConn.isStarted = true
}

/**
Blocking
*/
func (rpcServerConn *RpcServerConnection) Serve() {
	if rpcServerConn.isStarted {
		rpcServerConn.waitGroup.Wait()
	}
}

/**
Blocking
*/
func (rpcServerConn *RpcServerConnection) StartAndServe() {
	rpcServerConn.Start()
	rpcServerConn.Serve()
}

func (rpcServerConn *RpcServerConnection) Shutdown() {
	for i := 0; i < len(rpcServerConn.Ports); i++ {
		rpcServerConn.ServerList[rpcServerConn.Ports[i]].Shutdown(context.Background())
		log.Println("Shutting down in port:", rpcServerConn.Ports[i])
		rpcServerConn.waitGroup.Done()
	}
}

func getMethodNameFromTagAndParam(service interface{}, serviceValue reflect.Value, tagName string, params []interface{}) string {
	tagNameDetail := fmt.Sprintf("%s,%d", tagName, len(params))
	fieldName := getFieldName(tagNameDetail, util.MetaTag, service)
	if fieldName != "" {
		return fieldName
	}

	fieldName = getFieldName(tagName, util.MetaTag, service)
	if fieldName != "" {
		return fieldName
	}

	return strings.Title(tagName)
}

func getFieldName(tag, key string, s interface{}) (fieldname string) {
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		panic("bad type")
	}
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		if f.Tag.Get(key) == tag {
			return f.Name
		}
	}
	return ""
}
