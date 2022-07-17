//nolint:tagliatelle
package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type CommonPlayItDTO struct {
	Type string `json:"type"`
}

type SignInRequest struct {
	Type     string `json:"type"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshSessionRequest struct {
	Type              string `json:"type"`
	ExpiredSessionKey string `json:"expired_session_key"`
}

type SignInResponse struct {
	Type          string `json:"type"`
	ID            string `json:"id"`
	SessionKey    string `json:"session_key"`
	IsGuest       bool   `json:"is_guest"`
	EmailVerified bool   `json:"email_verified"`
}

type TunnelResponse struct {
	Type         string   `json:"type"`
	Agents       []string `json:"agents"`
	PortReleases []string `json:"port_releases"`
}

type (
	PortMappingsResponse struct {
		Type     string `json:"type"`
		Mappings []string
	}

	Mapping struct {
		ID                  string `json:"id"`
		Name                string `json:"name"`
		Proto               string `json:"proto"`
		TunnelIP            string `json:"tunnel_ip"`
		TunnelFromPort      string `json:"tunnel_from_port"`
		TunnelToPort        string `json:"tunnel_to_port"`
		BindIP              string `json:"bind_ip"`
		LocalIP             string `json:"local_ip"`
		LocalPort           string `json:"local_port"`
		TunnelType          string `json:"tunnel_type"`
		GeneratedDomainName string `json:"generated_domain_name"`
		CustomDomain        string `json:"custom_domain"`
	}
)

const (
	PlayItTypeGetTunnelNetwork string = "get-tunnel-network"
	PlayItTypeListPortMappings string = "list-port-mappings"
	PlayItTypePortMappings     string = "port-mappings"
	PlayItTypeRefreshSession   string = "refresh-session"
	PlayItTypeTunnelNetwork    string = "tunnel-network"
	PlayItTypeSignedIn         string = "signed-in"
	PlayItTypeSignIn           string = "sign-in"
	PlayItTypeError            string = "error"
)

const PlayItMinecraftJavaTunnelType string = "minecraft-java"

func SignInPlayItAPI(email, password string) (SignInResponse, error) {
	var response SignInResponse

	b := new(bytes.Buffer)

	err := json.NewEncoder(b).Encode(SignInRequest{
		Type:     PlayItTypeSignIn,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return response, err
	}

	res, err := http.Post("https://api.playit.cloud/login", "application/json", b)
	if err != nil {
		return response, err
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, res.Body.Close()
}

func RefreshSessionPlayItAPI(token string) (SignInResponse, error) {
	var response SignInResponse

	b := new(bytes.Buffer)

	err := json.NewEncoder(b).Encode(RefreshSessionRequest{
		Type:              PlayItTypeRefreshSession,
		ExpiredSessionKey: token,
	})
	if err != nil {
		return response, err
	}

	res, err := http.Post("https://api.playit.cloud/login", "application/json", b)
	if err != nil {
		return response, err
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, res.Body.Close()
}

func GetTunnelPlayItAPI(token string) error {
	b := new(bytes.Buffer)

	err := json.NewEncoder(b).Encode(CommonPlayItDTO{
		Type: PlayItTypeGetTunnelNetwork,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.playit.cloud/tunnel", b)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return res.Body.Close()
}

func ListPortMappings() error {
	return nil
}
