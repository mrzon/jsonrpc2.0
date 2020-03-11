package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/mrzon/jsonrpc2.0/model"
	"github.com/mrzon/jsonrpc2.0/util"
)

func (r *RpcClient) callServer(fnName string, args []interface{}, timeOut time.Duration) (result interface{}, err error) {
	jsonRequest := model.NewJsonRpcRequest(fnName, args)

	body, _ := json.Marshal(jsonRequest)
	if r.Config.EnableLogging {
		log.Println("Sending request:" + string(body))
	}
	reader := bytes.NewReader(body)

	req, _ := http.NewRequest("POST", r.Config.GetAddr(), reader)
	req.Header["Content-Type"] = []string{"application/json"}

	client := http.DefaultClient
	client.Timeout = timeOut
	response, err := client.Do(req)
	if err != nil {
		log.Println("Call method", fnName, "resulting in Error.", err.Error())
		return nil, err
	} else {
		strResponseBody, _ := ioutil.ReadAll(response.Body)
		if r.Config.EnableLogging {
			log.Println("Receiving response:" + string(strResponseBody))
		}

		jsonResponse := model.JsonRpcResponse{}
		json.Unmarshal(strResponseBody, &jsonResponse)
		return jsonResponse.Result, err
	}
}

func Register(r *RpcClient) {
	// Client use TCP transport.
	t := reflect.TypeOf(r.Service)
	var v reflect.Value
	if t.Kind() == reflect.Ptr {
		v = reflect.ValueOf(r.Service).Elem()
		t = t.Elem()
	} else {
		v = reflect.ValueOf(r.Service)
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		fn := field.Type

		newFunc := reflect.MakeFunc(fn, func(in []reflect.Value) (result []reflect.Value) {
			param := make([]interface{}, len(in))
			for j := 0; j < len(in); j++ {
				param[j] = in[j].Interface()
			}
			fnName := field.Name

			method := reflect.ValueOf(r.Service).Elem().FieldByName(fnName)
			methodType, _ := reflect.TypeOf(r.Service).Elem().FieldByName(fnName)

			tag := methodType.Tag
			customName := tag.Get(util.MetaTag)
			if customName != "" {
				fnName = strings.Split(customName, ",")[0]
			}

			processedData, err := r.callServer(fnName, param, r.Config.Timeout)

			methodOut := method.Type().Out(0)
			var processedValue = util.GetVal(methodOut, nil)
			if err == nil {
				processedValue = util.GetVal(methodOut, processedData)
			}
			result = append(result, processedValue)
			return result
		})

		f := v.FieldByName(field.Name)
		if !f.CanSet() {
			panic("Client service is need to be a pointer to a struct.")
		}
		f.Set(newFunc)
	}
}
