Para ejecutar proyecto abrir el cmd en la carpeta y escribir:
go run 
	   envio.go "Key de autorizacion","Path del archivo a firmar","Nombre de archivo que aparecera en el email","correo 1","correo 2"...

	   status.go "Key de autorizacion","Id contrato"
	   
	   statusAll.go "Key de autorizacion"
	   

envio.go: Simplemente envia un pdf para firmar.
status.go: Devuelve unicamente el estado del contrato si este fue especificado. 
Los diferentes estados pueden ser:
	*'OUT_FOR_SIGNATURE'
	*'OUT_FOR_DELIVERY'
	*'OUT_FOR_ACCEPTANCE'
	*'OUT_FOR_FORM_FILLING'
	*'OUT_FOR_APPROVAL'
	*'AUTHORING'
	*'CANCELLED'
	*'SIGNED'
	*'APPROVED'
	*'DELIVERED'
	*'ACCEPTED'
	*'FORM_FILLED'
	*'EXPIRED'
	*'ARCHIVED'
	*'PREFILL'
	*'WIDGET_WAITING_FOR_VERIFICATION'
	*'DRAFT'
	*'DOCUMENTS_NOT_YET_PROCESSED'
	*'WAITING_FOR_FAXIN'
	*'WAITING_FOR_VERIFICATION'
	*'WAITING_FOR_NOTARIZATION'