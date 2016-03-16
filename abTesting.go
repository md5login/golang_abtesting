package ABTesting

import (
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
	"math/rand"
	"time"
)

type Rule struct{
	Environment string		`json:"env"`
	Exposure	int			`json:"exposure"`
	Id			string		`json:"id"`
}

var rules []*Rule

// Sets cookie with wanted rule id as value on the request
func SetRule( res *http.ResponseWriter , rule Rule ){

	cookie := http.Cookie{
		Value:	rule.Id ,
		Name: "abRule" ,
		Secure: false ,
		HttpOnly: false ,
		Path: "/" ,
	}

	if getRandom( 0 , 100 ) <= rule.Exposure {
		http.SetCookie( *res , &cookie )
	}

}

// Checks whether the request has the wanted rule
func HasRule( req *http.Request , ruleId string ) bool{

	cookies := req.Cookies()

	for _ , v := range cookies {
		if v.Name == "abRule" && v.Value == ruleId {
			return true
		}
	}

	return false
}

func getRandom( min int , max int ) int{
	rand.Seed( time.Now().Unix() )
	return rand.Intn( max - min ) + min
}

// returns all the appended rules
func GetExistingRules() []*Rule{
	return rules
}

func GetRuleById( id string ) Rule{

	rule := Rule{}

	for _ , v := range rules {
		if v.Id == id {
			return *v
		}
	}

	return rule

}

func GetRuleByEnvironment( env string ) Rule{

	rule := Rule{}

	for _ , v := range rules {
		if v.Environment == env {
			return *v
		}
	}

	return rule

}

// Use this in runtime to append rules manually one-by-one
func AppendRule( rule Rule ) {

	rules = append( rules , &rule )

}

// Use this to append a bulk of rules
// json : [ {rule} , {rule} , {rule} , ... ]
func AppendRulesFromJSON( filePath string ) error{

	var file []byte

	if _ , err := os.Stat( filePath ); os.IsNotExist( err ){
		return err
	}

	file , err := ioutil.ReadFile( filePath )

	if err != nil {
		return err
	}

	json.Unmarshal( file , &rules )

	return err

}
