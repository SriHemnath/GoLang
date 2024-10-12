package main

import (
	"fmt"
	"reflect"
)

var _ Response = (*SampleResponse)(nil)

type SampleResponse struct {
	Result *Result
}

type Response interface {
	BusinessLogicError() error
}

func (r SampleResponse) BusinessLogicError() error {
	//manually return nil for reference struct. If not a nil pointer to a struct will be returned instead of nil and it will throw panic when accessed
	//if r.Result == nil {
	//	return nil
	//}
	return r.Result
}

var _ error = (*Result)(nil)

type Result struct {
	ErrorCode    int
	ErrorMessage string
}

func (r *Result) Error() string {
	return fmt.Sprintf("%d %s", r.ErrorCode, r.ErrorMessage)
}

func main() {
	HandleResponse(SampleResponse{
		Result: nil,
	})

	HandleResponse(SampleResponse{
		Result: &Result{
			ErrorCode:    1,
			ErrorMessage: "sample business logic error",
		},
	})
}

func HandleResponse(response interface{}) {
	r, ok := response.(Response)
	if !ok {
		fmt.Println("response is not valid")
	}

	err := r.BusinessLogicError()
	fmt.Println(reflect.TypeOf(err))
	fmt.Println(err)
	if err != nil {
		HandleError(err)

		return
	}

	fmt.Println("response is successful")
}

func HandleError(err error) {
	fmt.Printf("response has an error, %s\n", err.Error())
}
