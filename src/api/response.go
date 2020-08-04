package api

type ResponseBody struct {
	Status int
	Result interface{}
}

func setResponceBody(result interface{}, errorCode int) ResponseBody {
	res := ResponseBody{errorCode, result}
	return res
}