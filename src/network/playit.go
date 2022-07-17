//nolint:tagliatelle
package network

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

type ErrResponse struct {
	Type    string `json:"type"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type (
	PortMappingsResponse struct {
		Type     string    `json:"type"`
		Mappings []Mapping `json:"mappings"`
	}

	Mapping struct {
		ID                  string `json:"id"`
		Name                string `json:"name"`
		Proto               string `json:"proto"`
		TunnelIP            string `json:"tunnel_ip"`
		TunnelFromPort      int    `json:"tunnel_from_port"`
		TunnelToPort        int    `json:"tunnel_to_port"`
		BindIP              string `json:"bind_ip"`
		LocalIP             string `json:"local_ip"`
		LocalPort           int    `json:"local_port"`
		TunnelType          string `json:"tunnel_type"`
		GeneratedDomainName string `json:"generated_domain_name"`
		CustomDomain        string `json:"custom_domain"`
	}
)

type (
	PortLeasesResponse struct {
		Type   string  `json:"type"`
		Leases []Lease `json:"leases"`
	}

	Lease struct {
		ID                 string `json:"id"`
		IPAddress          string `json:"ip_address"`
		StripeSubID        string `json:"stripe_sub_id"`
		IsRandomAllocation bool   `json:"is_random_allocation"`
		Ports              Ports  `json:"ports"`
	}

	Ports struct {
		Proto    string `json:"proto"`
		FromPort int    `json:"from_port"`
		ToPort   int    `json:"to_port"`
	}
)

const (
	PlayItTypeGetTunnelNetwork string = "get-tunnel-network"
	PlayItTypeListPortMappings string = "list-port-mappings"
	PlayItTypeListPortLeases   string = "list-port-leases"
	PlayItTypeRefreshSession   string = "refresh-session"
	PlayItTypeTunnelNetwork    string = "tunnel-network"
	PlayItTypePortMappings     string = "port-mappings"
	PlayItTYpePortLeases       string = "port-leases"
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

func GetTunnelPlayItAPI(token string) (TunnelResponse, error) {
	var (
		response TunnelResponse

		request = CommonPlayItDTO{
			Type: PlayItTypeGetTunnelNetwork,
		}
	)

	return response, postWithToken("https://api.playit.cloud/tunnel", token, request, &response)
}

func ListPortMappings(token string) (PortMappingsResponse, error) {
	var (
		response PortMappingsResponse

		request = CommonPlayItDTO{
			Type: PlayItTypeListPortMappings,
		}
	)

	return response, postWithToken("https://api.playit.cloud/account", token, request, &response)
}

func ListPortLeases(token string) (PortLeasesResponse, error) {
	var (
		response PortLeasesResponse

		request = CommonPlayItDTO{
			Type: PlayItTypeListPortLeases,
		}
	)

	return response, postWithToken("https://api.playit.cloud/account", token, request, &response)
}

func postWithToken(url, token string, payload, v any) error {
	b := new(bytes.Buffer)

	err := json.NewEncoder(b).Encode(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		return err
	}

	return res.Body.Close()
}
