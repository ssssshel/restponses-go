package restponses

import (
	"fmt"

	internal_interfaces "github.com/ssssshel/restponses-go/internal/interfaces"
)

func Tests() {
	d := Response100Continue(Get)
	fmt.Println(d)
	res := &internal_interfaces.GenericSuccessResponse{
		BaseResponse: internal_interfaces.BaseResponse{
			HttpStatus: 200,
		},
	}
}
