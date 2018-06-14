package main
 
import (
  "net/http"
  "encoding/json"
  "errors"
  "net/url"
  "time"
  "strings"
  "log"
  "gopkg.in/mgo.v2"
   // "gopkg.in/mgo.v2/bson"  
)

//HttpData struct is defined to hold the prepared http request data to push to MongoDB
type HttpData struct{
	Time string
	URL string
	HTTPMethod string
	RemoteAddr string
	Header map[string][]string
	LengthContent int64
	Protocol string
}
type Message struct{
	StatusCode string
	Message string
}

//Function handler is the handler function for http calls
func handler(res http.ResponseWriter, r *http.Request) {
	httpd :=new(HttpData)
	//get current time
	t:=time.Now().Format(time.RFC1123)
	httpd.Time = t
	//get the complete URL including scheme, host, path, and query
	pathString:=strings.Join([]string {"http://",r.Host},"")
	u,err := url.Parse(pathString)
	if err != nil {
		log.Fatal(err)
	}
	//get the rawquery
	u.RawQuery = r.URL.RawQuery
	u.Path = r.URL.Path
	s:=u.String()
	//fmt.Printf("Reconstructed Path:%s",s)
	httpd.URL = s
   //get HTTP method
	httpd.HTTPMethod = r.Method
   //Get Remote Client Address
	httpd.RemoteAddr = r.RemoteAddr
	//Get All headers
	httpd.Header = r.Header
	// Length of the request body
	httpd.LengthContent = r.ContentLength
	//Get Protocol	
	httpd.Protocol = r.Proto
	//Call function to write to database
	//Database is Cloud MongoDB by mlab, You can Create an Account:  https://mlab.com/login
	status,errdb := calldb(httpd)
	log.Print("Status from DB operation:",status,". ",errdb)
	m:=new(Message)
	if status=="success"{
		m.StatusCode=status
		m.Message="Inserted HTTP data successfully"
	}else{m.StatusCode=status
		  m.Message="Failed to insert HTTP data to database"}	
	j , err := json.Marshal(m)
		if err!=nil{
			log.Print("Failed to convert response message to JSON. Error: ",err)
		}
		res.Write(j)
	
}

//Function calldb connects to mongodb and inserts the data passed in as parameter. Returns status and error if any
func calldb(h *HttpData) (string,error){
	mongoDialInfo:= & mgo.DialInfo {
		Addrs: [] string {
		 "ds259620.mlab.com:59620"},
		Database: "samples",
		Username: "fnanda", 
		Password: "welcome1",
		Timeout: 60 * time.Second,
	   }
	session, err:= mgo.DialWithInfo(mongoDialInfo)
	if err != nil {
		return "error",errors.New(strings.Join([]string{"Failed to Connect to database.Error Message:",err.Error()},""))
		}		
	defer session.Close()
 	session.SetMode(mgo.Monotonic, true)
 	c := session.DB("samples").C("httprequestdata")
 	err = c.Insert(h)
 	if err != nil {
		return "error",errors.New(strings.Join([]string{"Failed to Insert data to database.Error Message:",err.Error()},""))
	}
	return "success",nil
}

//Function handlerFav is defined to suppress favicon call from Browser
func handlerFav(res http.ResponseWriter,req *http.Request){}
 
func main(){
	var Port = "8000"
	http.HandleFunc("/favicon.ico", handlerFav) 
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe(strings.Join([]string{":",Port},""),nil))	
}