package tokenstimator

import (
	//"fmt" // Uncomment for testing purposes to display various stuff on the server console end. 
	"net/http"
	"html/template"
	"strconv"
)

// This is a struct for the variables passed into the results html template. Apparently the various variables have to start in a capital letter to be accessible by other packages (in this case html/template). Who knew?. Thus the rather bland variable names after debugging. 
type unitCost struct {
	UnitsReq, UnitsCost, FirstBand, SecondBand, ThirdBand, Fuel, Forex, Wrma, Inflation, Fixed, Vat, Erc, Rep, Total, FirstB, SecondB, ThirdB float64
}

//This function takes in the units requested and returns the total cost along with the various costs and their associated costs. Fun fact, using switch without an arg enables you to evaluate the case arg as an expression. Yay truth tables!
func unitsCost(unitsReq float64) (unitsCost, firstBandCost, secondBandCost, thirdBandCost, firstBand, secondBand, thirdBand float64) {
		
	switch {
	case unitsReq <= 50.0 : 
		firstBand := unitsReq
		firstBandCost := firstBand * 2.50
		unitsCost := firstBandCost
		secondBand := 0.0 
		thirdBand := 0.0
		return unitsCost, firstBandCost, secondBandCost, thirdBandCost, firstBand, secondBand, thirdBand	
	
	case unitsReq > 50 && unitsReq <=1500 :
		firstBand := 50.0
		firstBandCost := firstBand * 2.50
		secondBand := unitsReq - 50
		secondBandCost := secondBand * 11.62
		unitsCost := firstBandCost + secondBandCost
		thirdBand := 0.0
		return unitsCost, firstBandCost,secondBandCost, thirdBandCost, firstBand, secondBand, thirdBand
	
	case unitsReq > 1500 : 
		firstBand := 50.0
		firstBandCost := firstBand * 2.50
		secondBand := 1500.0 - firstBand
		secondBandCost := secondBand * 11.62
		thirdBand := unitsReq - 1500
		thirdBandCost := thirdBand * 19.57
		unitsCost := firstBandCost + secondBandCost + thirdBandCost
		return unitsCost, firstBandCost,secondBandCost, thirdBandCost, firstBand, secondBand, thirdBand
	
	default:
		return 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0
	}
	
}


//This function calculates the various charges and levies.
func charges(unitsReq float64, unitsCharge float64) (float64, float64, float64, float64, float64, float64, float64, float64) {
	const fcc float64 = 5.19
	fuelCharge := fcc * unitsReq
	
	const ferfa float64 = 0.34
	forexCharge := ferfa * unitsReq
	
	const wrma float64 = 0.06
	wrmaCharge := wrma * unitsReq
	
	const inflation float64 = 0
	inflationCharge := inflation * unitsReq
	
	const erc float64 = 0.03
	ercCharge := erc * unitsReq
	
	const rep float64 = 0.05
	repCharge := rep * unitsCharge
	
	var fixedCharge float64 = 120.00
	
	// Used this segment on a command line confirmation of first purchase. May be useful in the future.
	/*
	if askForConfirmation() == true {
		fixedCharge = 120.00
	} else {
		fixedCharge = 0.00
	}
	*/
	
	const vat float64 = 0.16
	vatableTotal := fuelCharge + forexCharge + wrmaCharge + inflationCharge + fixedCharge + unitsCharge
	vatCharge := vat * vatableTotal
	
	return fuelCharge, forexCharge, wrmaCharge, inflationCharge, fixedCharge, vatCharge, ercCharge, repCharge
}

// Handles the routing. App engine uses func init() instead of a func main(). If used away from an app engine env, rename this to func main() plus the package name may change(not sure about this). HTML provided by bootstrap. If away from app engine you may need to write a router for static files e.g css, javascript.
func init() {
	http.HandleFunc("/", frontPage)
	http.HandleFunc("/faq", faqPage)
	http.HandleFunc("/reciept", recieptPage)
}

// RTFM
func frontPage(w http.ResponseWriter, r *http.Request) {
if r.Method == "GET" {
    t, _ := template.ParseFiles("./templates/index.html")
    t.Execute(w, nil)
	} 
}

//RTFM
func faqPage(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("./templates/faq.html")
    t.Execute(w, nil)

}

func recieptPage(w http.ResponseWriter, r *http.Request) {
	//prevents direct access to the reciept page
	if r.Method == "GET" {
    t, _ := template.ParseFiles("./templates/index.html")
    t.Execute(w, nil)
	} else {
	x := template.New("reciept")
	x, _ = template.ParseFiles("./templates/reciept.html")

	
	unitsReqs := r.FormValue("unitsReq")
// Nested if else statement, was avoiding javascript for user input verification. Also wanted to see err in action, feel free to fix this intelligently.
	unitsReq, err := strconv.ParseFloat(unitsReqs, 64)
		if err != nil {
			t, _ := template.ParseFiles("./templates/index.html")
			t.Execute(w, nil)
			
		} else {
	unitsCost(unitsReq)
	unitsCharge, firstBandCost, secondBandCost, thirdBandCost, firstBand, secondBand, thirdBand := unitsCost(unitsReq)
	
	charges(unitsReq, unitsCharge)
	fuelCharge, forexCharge, wrmaCharge, inflationCharge, fixedCharge, vatCharge, ercCharge, repCharge := charges(unitsReq, unitsCharge)
	
	totalCost := fuelCharge + forexCharge + wrmaCharge + inflationCharge + fixedCharge + vatCharge + ercCharge + repCharge + unitsCharge
	resultsCost := unitCost{UnitsReq: unitsReq, UnitsCost: unitsCharge ,FirstBand: firstBandCost ,SecondBand: secondBandCost ,ThirdBand: thirdBandCost ,Fuel: fuelCharge,Forex: forexCharge,Wrma: wrmaCharge,Inflation: inflationCharge,Erc: ercCharge, Rep: repCharge,Fixed: fixedCharge,Vat: vatCharge, Total: totalCost, FirstB: firstBand, SecondB: secondBand, ThirdB: thirdBand} 
	
	x.Execute(w, resultsCost)
	}
	
	}
	
	
}
