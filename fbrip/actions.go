package fbrip

import (
	"os"
	"fmt"
	"net/url"
)

type ActionConfig struct{
	GetBasicInfo bool
	React react
	Post post
	Comment comment
	Scrap scrap
}

// ACTIONS
func(u *UserRip) getBasicInfo(){
	// Making GET request
	URL_struct,_ := url.Parse("https://mbasic.facebook.com/profile.php?v=info")
	response := u.GET(URL_struct)
	// Searching for user basic info -> {Name:,Birthday:,Gender:}
	bi := searchBasicInfo(response.Body)
	//Setting basic info for user
	u.Info.setInfo(bi)
}

func(u *UserRip) makeReaction(Url *url.URL, reaction string){
	//Fixing Url & Making GET request in the publication link
	Url = fixUrl(Url)
	response := u.GET(Url)
	//Searching for Reaction Url (it contains specific Query Parameters) 
	tempUrl := searchReactionPickerUrl(response.Body)
	//Making GET request for the reaction selection link
	response = u.GET(tempUrl)
	//Searching for `ufi/reaction` (it contains specific Query Parameters) 
	tempUrl = searchUfiReactionUrl(response.Body,reaction)
	//Doing reaction
	u.GET(tempUrl)
}


//scrap Urls
func (u *UserRip) scrap(Urls []*url.URL){
	for _,Url := range Urls{
		path := fmt.Sprintf("./scraps/%s-%s",u.Parameters["email"],Url.Path)
		f, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		response := u.GET()
		_, err2 := f.WriteString(bodyToString(response.Body))
		if err2 != nil {
			fmt.Println(err2)
		}
	}
}
