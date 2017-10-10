# golang_abtesting
A utility to manage rules for a connection

Written by Michael Sazonov

## Install
    $ go get github.com/md5login/abTesting

## Usage

Can be used anywhere with *http.Request, http.ResponseWriter*

```go

package main

import (
  "github.com/md5login/abTesting"
  "net/http"
  "strings"
)

func init(){
  err := ABTesting.AppendRulesFromJSON( "conf/abtesting.json" )
  // handle err
}

func handler(w http.ResponseWriter, req *http.Request) {

	if ABTesting.HasRule( req , "testRule" ) {
	
	  http.Redirect( 302 , "/success.html" )
	  
	} else if strings.HasSuffix( req.RequestURI , "?trule=secret_word" ) {
	
	  ABTesting.SetRule( &w , ABTesting.GetRuleById( "testRule" ) )
	  http.Redirect( 302 , "/success.html" )
	  
	} else {
	  
  	  w.Header().Set("Content-Type", "text/plain")
  	  w.Write([]byte("No secret here.\n"))
	  
	}
}

func main(){
  
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":12345", nil))
  
}
```

## Append rules in runtime
```go
rule := ABTesting.Rule{
		Id: "testRule" ,
		Exposure: 20 ,
	}
	
ABTesting.AppendRule( rule )
```

## Rules config
```json
[
  {
    "exposure" : 100 ,
    "id" : "testRule"
  } ,
  {
    "exposure" : 50 ,
    "id" : "anotherTest"
  } ,
]
