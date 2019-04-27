package main

import (
	"math"
	"encoding/json"
	"fmt"
	"net/http"
)

type engineData struct {
	Gamma float64 `json:"gamma"`
	Tc float64 `json:"tc"`
	M float64 `json:"m"`
	Pc float64 `json:"pc"`
	Pe float64 `json:"pe"`
	Epsilon float64 `json:"epsilon"`
	Isp float64 `json:"isp"`
	Fth float64 `json:"fth"`
	Theta float64 `json:"theta"`
	Beta float64 `json:"beta"`
	G float64 `json:"g"`
	Sigmac float64 `json:"sigmac"`
}

type engineResponseData struct {
	Ve float64 `json:"ve"`
	Md float64 `json:"md"`
	At float64 `json:"at"`
	Rt float64 `json:"rt"`
	Ldn float64 `json:"ldn"`
	Ac float64 `json:"ac"`
	Rc float64 `json:"rc"`
	Lcn float64 `json:"lcn"`
	Lc float64 `json:"lc"`
	Twall float64 `json:"twall"`
}

func radian(deg float64) (float64) {
	rad := deg*(3.14/180)
	return rad
}

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}

func process(edata engineData) (engineResponseData) {
	pi := 3.14
	Re := 0.56
	Ae := 0.92

	var erdata engineResponseData
	t := edata.G*edata.Isp
	erdata.Ve = toFixed(t, 2)
	erdata.Md = toFixed(edata.Fth/t, 2)

	t = (pi*Re*Re)/edata.Epsilon
	erdata.At = toFixed(t, 3)
	erdata.Rt = toFixed(math.Sqrt(t/pi), 3)

	erdata.Ldn = toFixed(math.Sqrt(Ae/pi)*(1/math.Tan(radian(edata.Theta))), 2)

	t = 3*t
	erdata.Ac = toFixed(t, 3)

	t = math.Sqrt(t/pi)
	erdata.Rc = toFixed(t, 3)

	erdata.Lcn = toFixed(math.Sqrt(t/pi)*(1/math.Tan(radian(edata.Beta))), 3)

	erdata.Lc = toFixed((t*1)/(pi*t*t), 2)

	erdata.Twall = toFixed((edata.Pc*t)/(2*edata.Sigmac), 3)

	return erdata
}

func calc(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var edata engineData
	var erdata engineResponseData
	
	decoder.Decode(&edata)

	// if err != nil {
	// 	panic(err)
	// }

	erdata = process(edata)

	fmt.Println(erdata)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(erdata); err != nil {
        panic(err)
    }
}

func main() {
	http.HandleFunc("/", calc)
	http.ListenAndServe(":8090", nil)
}