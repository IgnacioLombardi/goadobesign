package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"flag"
	//"bytes"
	

	"github.com/IgnacioLombardi/goadobesign/adobesign"
)



type Parametros []string

func (p *Parametros) String() string {
	return fmt.Sprintln(*p)
}

func (p *Parametros) Set(s string) error {
    *p = strings.Split(s, ",")
    return nil
}


func main() {
	var parametros Parametros
	flag.Var(&parametros, "p", "parametros")
	flag.Parse()

	/* KeyAutorizacion --> "3AAABLblqZhCaURh6vYfwDez_B2c4Yv71KXV5-uxpjnh5mvpWlmkxempMvK_mPyPux0Hafc61YsHlfjXH8VZ1HOComunmCWQS"*/
	KeyAutorizacion := parametros[0]
	client := adobesign.NewClient(KeyAutorizacion, "na4", "")

			/*Link de descarga para contrato especifico*/
			AgreementID := parametros[1]

			archivo, err:= client.AgreementService.GetCombinedDocument(context.Background(), AgreementID)
			if err != nil {
				log.Fatal(err)
			}
		
			array:= string(archivo[:])
			borrado1 := strings.Replace(array, "\"","",-1)
			borrado2 := strings.Replace(borrado1, "{url:","",-1)
			borrado3 := strings.Replace(borrado2, "}","",-1)
			fmt.Println(borrado3)
}