package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
	"regexp"
	"flag"

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

	fmt.Println(parametros)
	/* KeyAutorizacion --> "3AAABLblqZhCaURh6vYfwDez_B2c4Yv71KXV5-uxpjnh5mvpWlmkxempMvK_mPyPux0Hafc61YsHlfjXH8VZ1HOComunmCWQS"*/
	KeyAutorizacion := parametros[0]
	Path := parametros[1]
	NombreArchivo := parametros[2]
	EmailDestino := parametros[3]
	
	Archivo := strings.Split(Path, "/")
	fmt.Println(Archivo)
	client := adobesign.NewClient(KeyAutorizacion, "na4", "")

	file, err := os.Open(Path)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	document, err := client.TransientDocumentService.UploadTransientDocument(context.Background(), data, Archivo[(len(Archivo))-1])
	if err != nil {
		log.Fatal(err)
	}
	
	var agreement *adobesign.CreateAgreementResponse
	var agreement2 []uint8

	/*ENVIO PARA UN SOLO FIRMANTE*/
	
	if ((len(parametros))==4) {
		

		agreement, err = client.AgreementService.CreateAgreement(context.Background(), adobesign.Agreement{
		FileInfos: []adobesign.FileInfo{{TransientDocumentId: document.TransientDocumentId}},
		Name:      NombreArchivo,
		ParticipantSetsInfo: []adobesign.ParticipantSetInfo{
		{
			MemberInfos: []adobesign.MemberInfo{{Email: EmailDestino}},
			Order:       1,                               
			Role:        adobesign.ParticipantRole.Signer,
		},
		},
		SignatureType: adobesign.SignatureType.Esign,
		State:         adobesign.AgreementState.InProcess,
		})
		if err != nil {
		log.Fatal(err)
		}

		fmt.Println(agreement.Id)
		time.Sleep(10 * time.Second)
		agreement2, err = client.AgreementService.GetAgreementUrl(context.Background(), agreement.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
	/////////*ENVIO PARA DOS FIRMANTES*/////////
	if ((len(parametros))==5) {
		agreement, err := client.AgreementService.CreateAgreement(context.Background(), adobesign.Agreement{
		FileInfos: []adobesign.FileInfo{{TransientDocumentId: document.TransientDocumentId}},
		Name:      NombreArchivo,
		ParticipantSetsInfo: []adobesign.ParticipantSetInfo{
		{
			MemberInfos: []adobesign.MemberInfo{{Email: EmailDestino}},
			Order:       1,                               
			Role:        adobesign.ParticipantRole.Signer,
		},
		{
			MemberInfos: []adobesign.MemberInfo{{Email: parametros[4]}},
			Order:       2,                               
			Role:        adobesign.ParticipantRole.Signer,
		},
		},
		SignatureType: adobesign.SignatureType.Esign,
		State:         adobesign.AgreementState.InProcess,
		})
		if err != nil {
		log.Fatal(err)
		}
		fmt.Println(agreement.Id)
		time.Sleep(10 * time.Second)
		agreement2, err = client.AgreementService.GetAgreementUrl(context.Background(), agreement.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
	/////////*ENVIO PARA TRES FIRMANTES*/////////
	if ((len(parametros))==6) {
		agreement, err := client.AgreementService.CreateAgreement(context.Background(), adobesign.Agreement{
		FileInfos: []adobesign.FileInfo{{TransientDocumentId: document.TransientDocumentId}},
		Name:      NombreArchivo,
		ParticipantSetsInfo: []adobesign.ParticipantSetInfo{
		{
			MemberInfos: []adobesign.MemberInfo{{Email: EmailDestino}},
			Order:       1,                               
			Role:        adobesign.ParticipantRole.Signer,
		},
		{
			MemberInfos: []adobesign.MemberInfo{{Email: parametros[4]}},
			Order:       2,                               
			Role:        adobesign.ParticipantRole.Signer,
		},
		{
			MemberInfos: []adobesign.MemberInfo{{Email: parametros[5]}},
			Order:       3,                               
			Role:        adobesign.ParticipantRole.Signer,
		},
		},
		SignatureType: adobesign.SignatureType.Esign,
		State:         adobesign.AgreementState.InProcess,
		})
		if err != nil {
		log.Fatal(err)
		}
		fmt.Println(agreement.Id)
		time.Sleep(10 * time.Second)
		agreement2, err = client.AgreementService.GetAgreementUrl(context.Background(), agreement.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
		/////////*ENVIO PARA CUATRO FIRMANTES*/////////
	if ((len(parametros))==5) {
		agreement, err := client.AgreementService.CreateAgreement(context.Background(), adobesign.Agreement{
		FileInfos: []adobesign.FileInfo{{TransientDocumentId: document.TransientDocumentId}},
		Name:      NombreArchivo,
		ParticipantSetsInfo: []adobesign.ParticipantSetInfo{
		{
			MemberInfos: []adobesign.MemberInfo{{Email: EmailDestino}},
			Order:       1,                               
			Role:        adobesign.ParticipantRole.Signer,
		},
		{
			MemberInfos: []adobesign.MemberInfo{{Email: parametros[4]}},
			Order:       2,                               
			Role:        adobesign.ParticipantRole.Signer,
		},
		{
			MemberInfos: []adobesign.MemberInfo{{Email: parametros[5]}},
			Order:       3,                               
			Role:        adobesign.ParticipantRole.Signer,
		},
		{
			MemberInfos: []adobesign.MemberInfo{{Email: parametros[6]}},
			Order:       4,                               
			Role:        adobesign.ParticipantRole.Signer,
		},
		},
		SignatureType: adobesign.SignatureType.Esign,
		State:         adobesign.AgreementState.InProcess,
		})
		if err != nil {
		log.Fatal(err)
		}
		fmt.Println(agreement.Id)
		time.Sleep(10 * time.Second)
		agreement2, err = client.AgreementService.GetAgreementUrl(context.Background(), agreement.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
		info := string(agreement2[:])
		r, _  := regexp.Compile(`https?://([-\w\.]+)+(:\d+)?(/([\w/_\.]*(\?\S+)?)?)?`)
		link := strings.Split((r.FindString(info)), `"`)
		fmt.Println(r.FindString(link[0]))
	



	/*webhook, err := client.WebhookService.CreateWebhook(context.Background(), adobesign.CreateWebhookRequest{
		Name:                      "testing webhook",
		Scope:                     adobesign.Scope.Resource,
		State:                     "ACTIVE",
		WebhookSubscriptionEvents: []string{adobesign.WebhookSubscriptionEvent.AgreementAll},
		ResourceType:              adobesign.Resource.Agreement,
		ResourceId:                agreement.Id,
		WebhookUrlInfo:            adobesign.WebhookUrlInfo{Url: ""},
		WebhookConditionalParams: adobesign.WebhookConditionalParams{WebhookAgreementEvents: adobesign.WebhookAgreementEvents{
			IncludeDetailedInfo:     true,
			IncludeDocumentsInfo:    true,
			IncludeParticipantsInfo: true,
			IncludeSignedDocuments:  true,
		}},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(webhook)*/
}
