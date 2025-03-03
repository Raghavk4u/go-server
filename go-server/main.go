package main
import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter , r *http.Request){
	if err := r.ParseForm(); err != nil {
          fmt.Fprintf(w , "ParseForm() err : %v",err)
		  return		
	}
	fmt.Fprintf(w , "Post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	
	fmt.Fprintf(w , "Name = %s\n Address = %s\n", name , address)
}

func helloHandler(w http.ResponseWriter , r *http.Request){

	if r.URL.Path != "./hello"{
		http.Error(w, "404 not found",http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w, "method not allowed",http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w,"/hello")
}

func main(){

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form",formHandler)
	http.HandleFunc("./hello",helloHandler)

	fmt.Println("Starting server at 8080 \n")
	if err := http.ListenAndServe(":8080",nil); err != nil{
        log.Fatal(err)
	}

}