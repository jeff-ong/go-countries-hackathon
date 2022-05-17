package main


import (
    "html/template"
    "net/http"
    "io/ioutil"
    "log"
    "github.com/gorilla/mux"
)


type Country struct {
    Name string
    Continents []string
}


type CountriesData struct {
    PageTitle string
    Data string
}

func getCountryData(url string) string {
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    return string(body)
}

func main() {
    r := mux.NewRouter();
   
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        template.Must(template.ParseFiles("home.html")).Execute(w, CountriesData{
            PageTitle: "Countries of the world \n Hackathon",
            Data: getCountryData("https://restcountries.com/v3.1/all"),
        })
    })

    r.HandleFunc("/detail/{country}", func(w http.ResponseWriter, r *http.Request) {
        template.Must(template.ParseFiles("detail.html")).Execute(w, CountriesData{
            PageTitle: "Countries of the world | Hackathon",
            Data: getCountryData("https://restcountries.com/v3.1/name/" + mux.Vars(r)["country"]),
        }  )
    })

    http.ListenAndServe(":8080", r)
}

