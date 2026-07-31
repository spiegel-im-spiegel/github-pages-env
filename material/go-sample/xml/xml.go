package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

var xmldata = `
<?xml version="1.0" encoding="utf-8"?>
<rdf:RDF
  xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
  xmlns:foaf="http://xmlns.com/foaf/0.1/"
  xmlns:dc="http://purl.org/dc/terms/"
  xmlns:cc="http://web.resource.org/cc/"
>
  <foaf:Document rdf:about="https://text.baldanders.info/">
    <dc:title lang="ja">text.Baldanders.info</dc:title>
	<dc:creator>Spiegel</dc:creator>
	<dc:date>2018-03-13T11:57:52+09:00</dc:date>
	<cc:license rdf:resource="http://creativecommons.org/licenses/by-sa/4.0/"/>
  </foaf:Document>
</rdf:RDF>
`

//Time is wrapper class of time.Time
type Time struct {
	time.Time
}

//UnmarshalXML returns result of Unmarshal for xml.Unmarshal()
func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	tm, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return err
	}
	*t = Time{tm}
	return nil
}

type RDF struct {
	XMLName  xml.Name `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# RDF"`
	Document struct {
		URI   string `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# about,attr"`
		Title struct {
			Language string `xml:"lang,attr,omitempty"`
			Name     string `xml:",chardata"`
		} `xml:"http://purl.org/dc/terms/ title"`
		Creator string `xml:"http://purl.org/dc/terms/ creator"`
		Date    Time   `xml:"http://purl.org/dc/terms/ date"`
		License struct {
			URI string `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# resource,attr"`
		} `xml:"http://web.resource.org/cc/ license"`
	} `xml:"http://xmlns.com/foaf/0.1/ Document"`
}

/*
type RDF struct {
	Title string `xml:"Document>title"`
}
*/

func main() {
	rdf := &RDF{}
	if err := xml.Unmarshal([]byte(xmldata), rdf); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(rdf.Document.Date.Format(time.RFC3339))
}
