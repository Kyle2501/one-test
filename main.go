//  _webapp labs foundation
// . . Ejected . media

package main

// .

import (
		"os"
		"log"
		
		"text/template"
		"net/http"
)


type htmlPageData struct {
    pageTitle string
    pagePath string
    pageList []pageNav
    
}


type pageNav struct {
    pageTitle string
    pageLink string
}


// . indexHandler
func indexHandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/" {
    	http.NotFound(w, r)
    	return
    }
    
// ,

  pageTitle := "Ejected . media"
  pagePath := r.URL.Path
    
  pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  
  pageFilePath := template.Must(template.ParseFiles("page_layout.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  indexHandler



//  .  html url routes as well as terminal cli logs

func main() {

    appName := "Ejected . media"
    
    
    http.HandleFunc("/", indexHandler)



// -- -
        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
                log.Printf("Loading _webapp with default port")
        }
        log.Printf("_webapp is active and Listening on port %s", port)
        
        log.Printf("// -- - %s", appName)
        log.Printf("_webapp now loaded and running at http://localhost:%s", port)
        
// -- - 
        if err := http.ListenAndServe(":"+port, nil); err != nil {
                log.Fatal("Error Starting the HTTP Server :", err)
                return
        }


}


