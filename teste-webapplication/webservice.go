package  main

import (
	"html/template"
	"log"
	"net/http"
)

type RadioButton  struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

type PageVariables struct {
	PageTitle        string
	PageRadioButtons []RadioButton
	Answer           string
  }

func main() {
	http.HandleFunc("/", DisplayRadioButtons)
	http.HandleFunc("/selected", UserSelected)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func DisplayRadioButtons(w http.ResponseWriter, r *http.Request){
	Title := "Qual você prefere?"

	MyRadioButtons := []RadioButton{
		RadioButton{"animalselect", "gatos", false, false, "Gatos"},
		RadioButton{"animalselect", "cachorros", false, false, "Cachorros"},
	}

	MyPageVariables := PageVariables{
		PageTitle: Title,
		PageRadioButtons : MyRadioButtons,
		}
	
	   t, err := template.ParseFiles("home.html") //parse the html file homepage.html
	   if err != nil { // if there is an error
		 log.Print("template parsing error: ", err) // log it
	   }
	
	   err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	   if err != nil { // if there is an error
		 log.Print("template executing error: ", err) //log it
	}
}

func UserSelected(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	// r.Form is now either
	// map[animalselect:[cats]] OR
	// map[animalselect:[dogs]]
   // so get the animal which has been selected
	youranimal := r.Form.Get("animalselect")
  
	Title := "Você prefere o animal"
	MyPageVariables := PageVariables{
	  PageTitle: Title,
	  Answer : youranimal,
	  }
  
   // generate page by passing page variables into template
	  t, err := template.ParseFiles("home.html") //parse the html file homepage.html
	  if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	  }
  
	  err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	  if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	  }
  }
