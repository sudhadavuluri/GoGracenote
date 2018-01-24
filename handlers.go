// handlers.go

package main

import (
  "net/http"
  "fmt"
  "strings"
  "io/ioutil"
  "github.com/gin-gonic/gin"
)

func registerUser(c *gin.Context){
	url := "https://cXXXXXXX.web.cddbp.net/webapi/xml/1.0/"

	payload := strings.NewReader("<QUERIES>\r\n  <QUERY CMD=\"REGISTER\">\r\n    <CLIENT>XXXXXXX-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX</CLIENT>\r\n  </QUERY>\r\n</QUERIES>")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		ErrDefaultHandler(err)
    }

	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "48f95100-3048-5db6-32e1-f91f75d7ad3f")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ErrDefaultHandler(err)
    }

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ErrDefaultHandler(err)
    }

	fmt.Println(res)
	fmt.Println(string(body))

}

func postGraceMovie(c *gin.Context){	
url := "https://cXXXXXXX.web.cddbp.net/webapi/xml/1.0/"
	payload := strings.NewReader("<QUERIES>\r\n  <LANG>eng</LANG>\r\n  <AUTH>\r\n    <CLIENT>XXXXXXX-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX</CLIENT>\r\n    <USER>XXXXXXXXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX</USER>\r\n  </AUTH>\r\n  <QUERY CMD=\"ALBUM_SEARCH\">\r\n    <TEXT TYPE=\"ARTIST\">flying lotus</TEXT>\r\n    <TEXT TYPE=\"ALBUM_TITLE\">until the quiet comes</TEXT>\r\n    <TEXT TYPE=\"TRACK_TITLE\">all in</TEXT>\r\n  </QUERY>\r\n</QUERIES>")
	
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		ErrDefaultHandler(err)
    }
	
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "c6e46e0a-6eca-e80c-4938-36c96dfa70dd")
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ErrDefaultHandler(err)
    }
	
	defer res.Body.Close()
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ErrDefaultHandler(err)
    }
	
	fmt.Println(res)
	fmt.Println(string(body))
}

func postTrkSrchAttr(c *gin.Context){

	url := "https://cXXXXXXX.web.cddbp.net/webapi/xml/1.0/"

	payload := strings.NewReader("<QUERIES>\r\n  <AUTH>\r\n\t<CLIENT>XXXXXXX-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX</CLIENT>\r\n    <USER>XXXXXXXXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX</USER>\r\n  </AUTH>\r\n  <QUERY CMD=\"ALBUM_SEARCH\">\r\n    <MODE>SINGLE_BEST_COVER</MODE>\r\n    <TEXT TYPE=\"ARTIST\">flying lotus</TEXT>\r\n    <TEXT TYPE=\"ALBUM_TITLE\">until the quiet comes</TEXT>\r\n    <TEXT TYPE=\"TRACK_TITLE\">all in</TEXT>\r\n    <OPTION>\r\n      <PARAMETER>SELECT_EXTENDED</PARAMETER>\r\n      <VALUE>COVER,REVIEW,ARTIST_BIOGRAPHY,ARTIST_IMAGE,ARTIST_OET,MOOD,TEMPO</VALUE>\r\n    </OPTION>\r\n    <OPTION>\r\n      <PARAMETER>SELECT_DETAIL</PARAMETER>\r\n      <VALUE>GENRE:3LEVEL,MOOD:2LEVEL,TEMPO:3LEVEL,ARTIST_ORIGIN:4LEVEL,ARTIST_ERA:2LEVEL,ARTIST_TYPE:2LEVEL</VALUE>\r\n    </OPTION>\r\n  </QUERY>\r\n</QUERIES>")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		ErrDefaultHandler(err)
    }
	
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "6b77514f-d7cc-815c-f069-768338778aaf")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ErrDefaultHandler(err)
    }

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ErrDefaultHandler(err)
    }
	
	fmt.Println(res)
	fmt.Println(string(body))

}


func postChangeLang(c *gin.Context){
	url := "https://cXXXXXXX.web.cddbp.net/webapi/xml/1.0/"

	payload := strings.NewReader("<QUERIES>\r\n  <AUTH>\r\n    <CLIENT>XXXXXXX-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX</CLIENT>\r\n    <USER>XXXXXXXXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX</USER>\r\n  </AUTH>\r\n  <LANG>ger</LANG>\r\n  <QUERY CMD=\"ALBUM_SEARCH\">\r\n    <MODE>SINGLE_BEST_COVER</MODE>\r\n    <TEXT TYPE=\"ARTIST\">flying lotus</TEXT>\r\n    <TEXT TYPE=\"ALBUM_TITLE\">until the quiet comes</TEXT>\r\n    <TEXT TYPE=\"TRACK_TITLE\">all in</TEXT>\r\n    <OPTION>\r\n      <PARAMETER>SELECT_EXTENDED</PARAMETER>\r\n      <VALUE>COVER,REVIEW,ARTIST_BIOGRAPHY,ARTIST_IMAGE,ARTIST_OET,MOOD,TEMPO</VALUE>\r\n    </OPTION>\r\n    <OPTION>\r\n      <PARAMETER>SELECT_DETAIL</PARAMETER>\r\n      <VALUE>GENRE:3LEVEL,MOOD:2LEVEL,TEMPO:3LEVEL,ARTIST_ORIGIN:4LEVEL,ARTIST_ERA:2LEVEL,ARTIST_TYPE:2LEVEL</VALUE>\r\n    </OPTION>\r\n  </QUERY>\r\n</QUERIES>")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		ErrDefaultHandler(err)
    }
	
	req.Header.Add("Content-Type", "text/plain")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "7a0b28e8-bed3-fbe2-5b45-fd34d1bebd6d")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ErrDefaultHandler(err)
    }

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ErrDefaultHandler(err)
    }
	
	fmt.Println(res)
	fmt.Println(string(body))

}


