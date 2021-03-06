//Heavily inspired/derived from discordgo https://github.com/bwmarrin/discordgo

package gobeam

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// ErrJSONUnmarshal is returned for JSON Unmarshall errors.
var ErrJSONUnmarshal = errors.New("json unmarshal")

// Request makes a (GET/POST/...) Requests to Beam REST API with JSON data.
// All the other Beam REST Calls in this file use this function.
func (s *Session) Request(method, urlStr string, data interface{}) (response []byte, err error) {

	if s.Debug {
		fmt.Println("API REQUEST  PAYLOAD :: [" + fmt.Sprintf("%+v", data) + "]")
	}

	var body []byte
	if data != nil {
		body, err = json.Marshal(data)
		if err != nil {
			return
		}
	}

	return s.request(method, urlStr, "application/json", body)
}

// request makes a (GET/POST/...) Requests to Beam REST API.
func (s *Session) request(method, urlStr, contentType string, b []byte) (response []byte, err error) {

	if s.Debug {
		fmt.Printf("API REQUEST %8s :: %s\n", method, urlStr)
	}

	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(b))
	if err != nil {
		return
	}

	if s.UseCookies {
		for _, c := range s.Cookies {
			if c.Name != "sid" {
				continue
			}
			req.Header.Add("Cookie", fmt.Sprintf("%s=%s", c.Name, c.Value))
		}
	} else {
		req.Header.Set("X-CSRF-Token", s.CsrfToken)
	}

	// Not used on initial login..
	if s.Token != "" {
		req.Header.Set("authorization", s.Token)
	}

	req.Header.Set("Content-Type", contentType)
	// TODO: Make a configurable static variable.
	req.Header.Set("User-Agent", fmt.Sprintf("GoBeamBot (https://github.com/xackery/gobeam, v%s)", VERSION))

	if s.Debug {
		for k, v := range req.Header {
			fmt.Printf("API REQUEST   HEADER :: [%s] = %+v\n", k, v)
		}
	}

	client := &http.Client{Timeout: (20 * time.Second)}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println("error closing resp body")
		}
	}()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if s.UseCookies {
		s.Cookies = resp.Cookies()
	} else {
		s.CsrfToken = resp.Header.Get("X-CSRF-Token")
	}

	if s.Debug {

		fmt.Printf("API RESPONSE  STATUS :: %s\n", resp.Status)
		for k, v := range resp.Header {
			fmt.Printf("API RESPONSE  HEADER :: [%s] = %+v\n", k, v)
		}
		fmt.Printf("API RESPONSE    BODY :: [%s]\n", response)
	}

	// See http://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html
	switch resp.StatusCode {

	case 200: // OK
	case 204: // No Content

		// TODO check for 401 response, invalidate token if we get one.

	case 429: // TOO MANY REQUESTS - Rate limiting
		//err = fmt.Errorf("Request unmarshal rate limit error : %+v", err)
	default: // Error condition
		err = fmt.Errorf("HTTP %s, %s", resp.Status, response)
	}

	return
}

func unmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return ErrJSONUnmarshal
	}

	return nil
}

// ------------------------------------------------------------------------------------------------
// Functions specific to Beam Sessions
// ------------------------------------------------------------------------------------------------

// Login asks the Beam server for an authentication token.
func (s *Session) Login() (err error) {

	if s.config == nil {
		s.config = &GoBeamConfig{}
		err = s.config.Load("gobeam.json")
		if err != nil {
			err = fmt.Errorf("Failed to load gobeam.json config: %s", err.Error())
			return
		}
	}

	data := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{s.config.Username, s.config.Password}

	response, err := s.Request("POST", LOGIN, data)
	if err != nil {
		return
	}

	temp := &LoginPayload{}

	err = unmarshal(response, &temp)
	if err != nil {
		return
	}

	s.LoginPayload = temp
	return
}

// ------------------------------------------------------------------------------------------------
// Functions specific to Beam Websockets
// ------------------------------------------------------------------------------------------------

// Gateway returns the a websocket Gateway address
func (s *Session) Gateway() (gateway string, err error) {

	response, err := s.Request("GET", fmt.Sprintf("%s%d", GATEWAY, s.LoginPayload.Channel.ID), nil)
	if err != nil {
		return
	}

	temp := struct {
		Endpoints []string `json:"endpoints,omitempty"`
		Authkey   string   `json:"authkey"`
	}{}

	err = unmarshal(response, &temp)
	if err != nil {
		return
	}
	if len(temp.Endpoints) > 0 {
		gateway = temp.Endpoints[0]
	}
	//if len(temp.Authkey) < 1 {
	//	err = fmt.Errorf("No authkey provided")
	//	return
	//}
	s.authKey = temp.Authkey

	return
}

// Gateway returns the a websocket Gateway address
func (s *Session) TetrisGateway() (gateway string, err error) {
	if s.LoginPayload == nil {
		err = fmt.Errorf("No login payload set")
		return
	}

	response, err := s.Request("GET", fmt.Sprintf("%s%d/robot", TETRIS, s.LoginPayload.Channel.ID), nil)
	if err != nil {
		return
	}

	temp := struct {
		Address string `json:"address,omitempty"`
		Key     string `json:"key,omitempty"`
	}{}

	err = unmarshal(response, &temp)
	if err != nil {
		return
	}
	if len(temp.Key) < 1 {
		err = fmt.Errorf("No authkey provided")
		return
	}
	gateway = temp.Address
	s.authKey = temp.Key

	return
}

//Return channel details
func (s *Session) Channels(id uint32) (channel *Channel, err error) {
	if s.LoginPayload == nil {
		err = fmt.Errorf("Not logged in")
		return
	}

	if id == 0 {
		id = s.LoginPayload.Channel.ID
	}

	response, err := s.Request("GET", fmt.Sprintf("%s%d", CHANNELS, id), nil)
	if err != nil {
		return
	}

	err = unmarshal(response, &channel)
	if err != nil {
		return
	}
	return
}

//Return tetris details
func (s *Session) TetrisGamesPublic(id uint32) (versions *VersionsResponse, err error) {
	if id == 0 && s.LoginPayload != nil {
		id = s.LoginPayload.Channel.ID
	}
	if id < 1 {
		err = fmt.Errorf("Version ID not found")
		return
	}

	response, err := s.Request("GET", fmt.Sprintf("%s%s%d", TETRIS_GAMES, "versions/", id), nil)
	if err != nil {
		return
	}

	err = unmarshal(response, &versions)
	if err != nil {
		return
	}
	return
}

func (s *Session) TetrisGamesOwned() (versions *VersionsResponse, err error) {
	if s.LoginPayload == nil {
		err = fmt.Errorf("Not logged in")
		return
	}

	response, err := s.Request("GET", fmt.Sprintf("%s%s", TETRIS_GAMES, "owned"), nil)
	if err != nil {
		return
	}

	err = unmarshal(response, &versions)
	if err != nil {
		return
	}
	return
}
