package main

import (
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"github.com/dragtor/gopherism/htmllink/pkg"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	domain *string
	output *string
)

var (
	INVALID_DOMAIN   = errors.New("Invalid Domain")
	DIFFERENT_DOMAIN = errors.New("Different Domain")
)

func validateFlag() {
	if strings.TrimSpace(*domain) == "" {
		panic("domain name not specified")
	}
}

func init() {
	domain = flag.String("d", "", "Specify domain name")
	output = flag.String("o", "./output.xml", "Output XML file path")
}

func IsPathInDomain(givenDomain, urlStr string) (string, error) {
	givenURL, err := url.Parse(givenDomain)
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", INVALID_DOMAIN
	}
	if u.Host != "" {
		// if subdomain is different
		splittedGivenURL := strings.Split(givenURL.Host, ".")
		splittedURL := strings.Split(u.Host, ".")

		domainname, extension := splittedGivenURL[len(splittedGivenURL)-2], splittedGivenURL[len(splittedGivenURL)-1]
		domainnameAtr, extensionAtr := splittedURL[len(splittedURL)-2], splittedURL[len(splittedURL)-1]
		if domainname != domainnameAtr || extension != extensionAtr {
			return "", DIFFERENT_DOMAIN
		}
	}
	if u.Host == "" {
		u.Scheme = givenURL.Scheme
		u.Host = givenURL.Host
	}
	newURL := u.String()
	return newURL, nil
}

type SiteMap struct {
	Pages map[string]PageRoutes
}

type PageRoutes struct {
	SiteRoute []string
}

func FetchHTMLData(url string) ([]byte, error) {
	resp, err := http.Get(*domain)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
		return nil, err
	}
	return body, nil

}

func GetSiteLocationsForURL(userdomain, url string) ([]string, error) {
	var validSiteLocationInPage []string
	pageData, err := FetchHTMLData(url)
	if err != nil {
		return nil, err
	}
	links, err := pkg.GetLinks(pageData)
	if err != nil {
		return nil, err
	}
	for _, l := range links {
		path, err := IsPathInDomain(userdomain, l.Href)
		if err != nil {
			continue
		}
		validSiteLocationInPage = append(validSiteLocationInPage, path)
	}
	return validSiteLocationInPage, nil
}

var locationVisitMap = map[string]bool{}

func GenerateSiteMapFromLocations(originDomain string, urlList []string) ([]string, error) {
	newUrlList := make([]string, len(urlList))
	copy(newUrlList, urlList)
	unvisitedLocationCount := 0
	for _, path := range urlList {
		if isVisited, present := locationVisitMap[path]; present {
			if isVisited {
				continue
			}
		}

		unvisitedLocationCount++
		locationList, err := GetSiteLocationsForURL(originDomain, path)
		if err != nil {
			fmt.Printf("%v", err)
			panic(err)
		}
		//mark path as visited
		locationVisitMap[path] = true
		//append all location in that page to urlList for next iteration
		for _, p := range locationList {
			if _, present := locationVisitMap[p]; !present {
				locationVisitMap[p] = false
				newUrlList = append(newUrlList, p)
			}
		}
	}
	if unvisitedLocationCount == 0 {
		return urlList, nil
	}
	return GenerateSiteMapFromLocations(originDomain, newUrlList)
}

type XML struct {
	XMLName xml.Name `xml:"xml"`
	Urlset  UrlSet   `xml:"urlset"`
}

type UrlSet struct {
	Xmlns string `xml:"xmlns,attr"`
	Url   []URL  `xml:"url"`
}

type URL struct {
	Loc string `xml:"loc"`
}

func WriteToYaml(filepath string, sitemap []string) error {
	var urlList []URL
	for _, path := range sitemap {
		urlData := URL{Loc: path}
		urlList = append(urlList, urlData)
	}

	samplestruct := XML{
		Urlset: UrlSet{
			Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
			Url:   urlList,
		},
	}
	xmldata, err := xml.MarshalIndent(samplestruct, "  ", "    ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filepath, xmldata, 777)
	return err
}

func main() {
	flag.Parse()
	validateFlag()
	initialPathList := []string{*domain}
	sitemapList, err := GenerateSiteMapFromLocations(*domain, initialPathList)
	if err != nil {
		panic(err)
	}
	err = WriteToYaml(*output, sitemapList)
	if err != nil {
		panic(err)
	}

}
