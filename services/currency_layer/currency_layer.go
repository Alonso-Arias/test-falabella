package currencylayer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type CurrencyLayerService struct {
}

var apiKey = os.Getenv("API_KEY")
var urlBase = os.Getenv("URL_BASE")

func Convert(currencyIn string, currencyOut string, amount float64) (float64, error) {

	url := fmt.Sprintf("%s/convert?access_key=%s&from=%s&to=%s&amount=%f", urlBase, apiKey, currencyIn, currencyOut, amount)

	fmt.Println("URLLLALALAL: ", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error al crear la solicitud:", err)
		return 0, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al realizar la solicitud:", err)
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return 0, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error al decodificar JSON:", err)
		return 0, err
	}

	fmt.Println("Resultado de la conversión:", result)

	// Acceder al resultado de la conversión
	if val, ok := result["result"].(float64); ok {
		fmt.Println("Resultado de la conversión:", val)
		return val, nil
	} else {
		fmt.Println("Error: no se pudo obtener el resultado de la conversión")
	}
	return 0, nil
}

func GetCurrencies() ([]string, error) {
	url := fmt.Sprintf("https://api.currencylayer.com/list?access_key=%s", apiKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error al crear la solicitud: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error al realizar la solicitud: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar JSON: %v", err)
	}

	currenciesMap, ok := result["currencies"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("error: no se pudo obtener la lista de monedas")
	}

	var currencies []string
	for code := range currenciesMap {
		currencies = append(currencies, code)
	}

	return currencies, nil
}
