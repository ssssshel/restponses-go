package methods

import "github.com/ssssshel/restponses-go/internal/interfaces"

type BaseResponseOpt func(*interfaces.BaseResponse)

type Status2xxResponseOpt func(*interfaces.Status2xx_Response)
