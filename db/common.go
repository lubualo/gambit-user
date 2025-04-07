package db

import (
	"os"

	"github.com/lubualo/gambit-user/models"
	"github.com/lubualo/gambit-user/secretm"
)

var SecretModel models.SecretRDSJson
var err error

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	/** escriba en la db usando el secreto y escribiendo la info del user q se registró*/
	return err
}
