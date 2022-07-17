//nolint:tagliatelle
package network

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
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

	PlayItMinecraftJavaTunnelType string = "minecraft-java"
)

var (
	ErrUnexpectedStatusCode = errors.New("unexpected status code")
	ErrUnauthorized         = errors.New("unauthorized access, token probably must be refreshed")
)

func Login(email, password string) (SignInResponse, error) {
	var resBody SignInResponse

	b := new(bytes.Buffer)

	err := json.NewEncoder(b).Encode(SignInRequest{Type: PlayItTypeSignIn, Email: email, Password: password})
	if err != nil {
		return resBody, err
	}

	res, err := http.Post("https://api.playit.cloud/login", "application/json", b)
	if err != nil {
		return resBody, err
	}

	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		return resBody, err
	}

	return resBody, res.Body.Close()
}

func RefreshSession(token string) (SignInResponse, error) {
	var (
		resBody SignInResponse
		request = RefreshSessionRequest{
			Type:              PlayItTypeRefreshSession,
			ExpiredSessionKey: token,
		}
	)

	return resBody, postWithToken("https://api.playit.cloud/login", token, request, &resBody)
}

func GetTunnel(token string) (TunnelResponse, error) {
	var (
		resBody TunnelResponse

		request = CommonPlayItDTO{
			Type: PlayItTypeGetTunnelNetwork,
		}
	)

	return resBody, postWithToken("https://api.playit.cloud/tunnel", token, request, &resBody)
}

func ListPortMappings(token string) (PortMappingsResponse, error) {
	var (
		resBody PortMappingsResponse

		request = CommonPlayItDTO{
			Type: PlayItTypeListPortMappings,
		}
	)

	return resBody, postWithToken("https://api.playit.cloud/account", token, request, &resBody)
}

func ListPortLeases(token string) (PortLeasesResponse, error) {
	var (
		resBody PortLeasesResponse

		request = CommonPlayItDTO{
			Type: PlayItTypeListPortLeases,
		}
	)

	return resBody, postWithToken("https://api.playit.cloud/account", token, request, &resBody)
}

func postWithToken(url, token string, payload, resBody any) error {
	if reflect.ValueOf(resBody).Kind() != reflect.Pointer {
		panic("response body must be a reference")
	}

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
	req.Header.Add("Authorization", "bearer "+token)
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusUnauthorized {
		return ErrUnauthorized
	} else if res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("%w: %s", ErrUnexpectedStatusCode, res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(resBody)
	if err != nil {
		return err
	}

	return res.Body.Close()
}
