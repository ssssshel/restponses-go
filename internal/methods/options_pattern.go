package methods

import "github.com/ssssshel/restponses-go/internal/interfaces"

type BaseResponseOpt func(*interfaces.GeneralInput)

type Response2xxOpt func(*interfaces.GenericSuccessfullResponse)

type Response3xxOpt func(*interfaces.GenericRedirectionResponse)

type ResponseErrorOpt func()
