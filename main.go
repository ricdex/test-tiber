package main

import (
        "crypto/tls"
        "fmt"
        "unsafe"

        "my-go-project/sgx"

        "github.com/intel/trustauthority-client/go-connector"
)

func main() {
        // Configuración básica para el cliente de Trust Authority
        cfg := connector.Config{
                BaseUrl: "https://portal.trustauthority.intel.com",
                ApiUrl:  "https://api.trustauthority.intel.com",
                TlsCfg:  &tls.Config{},
                ApiKey:  "---", // Aqui clave API
        }

        // Crear el cliente
        client, err := connector.New(&cfg)
        if err != nil {
                fmt.Printf("Error al crear el cliente: %v\n", err)
                return
        }

        // Imprimir que el cliente se creó
        fmt.Println("Cliente de Trust Authority creado correctamente:", client)

        //obtener el nounce
        reqNounce := connector.GetNonceArgs{}
        respNounce, err := client.GetNonce(reqNounce)
        if err != nil {
                fmt.Printf("Something bad happened: %s\n\n", err)
                return
        }

        fmt.Println("Get nounce:", respNounce.Nonce)

        //obtener la evidencia
        //obtener la evidencia
        var reportData [64]byte // Datos específicos del entorno SGX
        copy(reportData[:], "example report data") // Reemplaza con datos reales


        var someUint64 uint64 = 12345                                      // Reemplaza con el valor adecuado
        var someByteArray []byte = reportData[:]                    // Reemplaza con el valor adecuado
        var somePointer unsafe.Pointer = unsafe.Pointer(&reportData[0])  // Reemplaza con el valor adecuado


            fmt.Printf("someUint64: %d\n", someUint64)
            fmt.Printf("someByteArray: %v\n", someByteArray)
            fmt.Printf("somePointer: %v\n", somePointer)

        adapter, err := sgx.NewEvidenceAdapter(someUint64, someByteArray, somePointer)
        if err != nil {
                return
        }

        fmt.Println("Get adapter:", adapter)

        req := connector.AttestArgs{
                Adapter:   adapter,
                //PolicyIds: policyIds,
                //RequestId: "",
        }
        respToken, err := client.Attest(req)
        if err != nil {
            return
        }
                                                                                                                                                                                        1,12          Top

		fmt.Println("Get token:", respToken)
	/*
			evidence, err := adapter.CollectEvidence(respNounce.Nonce)
			if err != nil {
				return
			}

			fmt.Println("Get evidence:", evidence)
			// Crear el cliente usando sgxAdapter
		    /*adapter := &sgx.SgxAdapter{}
		    if err != nil {
		        fmt.Printf("Error al crear el cliente: %v\n", err)
		        return
		    }

		    evidence, err := adapter.CollectEvidence(respNounce.Nonce)
		    if err != nil {
		        return
		    }

			//hacer la atestacion
			/*reqAttest := connector.GetTokenArgs{
				Nonce:    respNounce.Nonce,
				Evidence: evidence.Evidence,
			}
			respAttest, err := client.GetToken(reqAttest)
			if err != nil {
				fmt.Printf("Something bad happened: %s\n\n", err)
				return
			}

			fmt.Println("Get attest:", respAttest)
	*/
}
