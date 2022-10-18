package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
	"flag"
	"bytes"
	

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
	

	
		if len(parametros)==1 {
			/*Consulta de todos los contratos*/
    		agreement, err:= client.AgreementService.GetAgreementAll(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(bytes.NewBuffer(agreement).String())

		} else {
			/*Consulta contrato especifico*/
			AgreementID = parametros[1]
			agreement, err:= client.AgreementService.GetAgreement(context.Background(), AgreementID)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(agreement.Status)
			time.Sleep(10 * time.Second)
		}

}
