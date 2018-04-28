package model

type JsonRpcRequest struct {
	JsonRpc string `json:"jsonrpc"`
	Method string `json:"method"`
	Params []interface{} `json:"params"`
	Id string `json:"id"`
}

type JsonRpcResponse struct {
	JsonRpc string `json:"jsonrpc"`
	Error map[string]string`json:"error,omitempty"`
	Result interface{} `json:"result,omitempty"`
	Id string `json:"id"`
}

func NewJsonRpcRequest(method string, param []interface{}) *JsonRpcRequest {
	return &JsonRpcRequest{
		JsonRpc:"2.0",
		Method:method,
		Params: param,
		Id:"",
	}
}

func NewJsonRpcResponseWithResult(result interface{}) *JsonRpcResponse{
	return &JsonRpcResponse{
		JsonRpc:"2.0",
		Result: result,
		Id:"",
	}
}

func NewJsonRpcResponseWithError(err map[string]string) *JsonRpcResponse{
	return &JsonRpcResponse{
		JsonRpc:"2.0",
		Error: err,
		Id:"",
	}
}