package methods

import interfaces "github.com/ssssshel/restponses-go/internal/interfaces"

type Response2xxOpt func(*interfaces.GenericSuccessfullResponse)

type Response3xxOpt func(*interfaces.GenericRedirectionResponse)

type Response4xxOpt func(*interfaces.GenericErrorResponse)
