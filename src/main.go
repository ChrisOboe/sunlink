package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetOwnedGamesResponse struct {
	Response struct {
		GameCount int `json:"game_count"`
		Games     []struct {
			Appid                    int    `json:"appid"`
			Name                     string `json:"name"`
			PlaytimeForever          int    `json:"playtime_forever"`
			ImgIconURL               string `json:"img_icon_url"`
			PlaytimeWindowsForever   int    `json:"playtime_windows_forever"`
			PlaytimeMacForever       int    `json:"playtime_mac_forever"`
			PlaytimeLinuxForever     int    `json:"playtime_linux_forever"`
			HasCommunityVisibleStats bool   `json:"has_community_visible_stats,omitempty"`
		} `json:"games"`
	} `json:"response"`
}

type App struct {
	Name     string   `json:"name"`
	Output   string   `json:"output"`
	Detached []string `json:"detached"`
}

type SunshineGamelist struct {
	Env struct {
		Path string `json:"PATH"`
	} `json:"env"`
	Apps []App `json:"apps"`
}

func doSteam(key string, steamid string) (SunshineGamelist, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.steampowered.com/IPlayerService/GetOwnedGames/v1/", nil)
	if err != nil {
		return SunshineGamelist{}, err
	}

	q := req.URL.Query()
	q.Add("key", key)
	q.Add("steamid", steamid)
	q.Add("include_appinfo", "true")
	q.Add("include_played_free_games", "true")
	q.Add("include_free_sub", "true")

	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return SunshineGamelist{}, err
	}

	defer resp.Body.Close()
	rawgames, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return SunshineGamelist{}, err
	}
	var games GetOwnedGamesResponse
	err = json.Unmarshal(rawgames, &games)

	var out SunshineGamelist
	for _, game := range games.Response.Games {

		runCmd := fmt.Sprintf("steam -applaunch %d", game.Appid)
		a := App{
			Name:     game.Name,
			Detached: []string{runCmd},
		}

		out.Apps = append(out.Apps, a)
	}

	return out, nil
}

func main() {
	gamelist, err := doSteam("60137AE59F3F69FCD2D2D7829C0D9C9F", "76561198119107992")
	if err != nil {
		fmt.Println(err)
		return
	}

	j, err:=json.MarshalIndent(gamelist, "","  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(j))
}
