package squadcast

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

type Service struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	APIKey      string `json:"api_key"`
	Slug        string `json:"slug"`
	// More fields ...
}

type ServicesResponse struct {
	Services []*Service `json:"data"`
}

type ServiceResponse struct {
	Service *Service `json:"data"`
}

type CreateServiceRequest struct {
	Name               string `json:"name"`
	EscalationPolicyId string `json:"escalation_policy_id"`
	Description        string `json:"description"`
	EmailPrefix        string `json:"email_prefix"`
}
// https://apidocs.squadcast.com/#abb07c8a-d547-46eb-88f1-19378314ec4e
func (c *Client) GetAllServices(ctx context.Context) ([]*Service, error) {
	params := &requestParams{
		method:  "GET",
		subPath: "/services",
	}

	var servicesResponse ServicesResponse
	if err := c.doAPIRequest(ctx, params, &servicesResponse); err != nil {
		return nil, err
	}

	return servicesResponse.Services, nil
}

// https://apidocs.squadcast.com/#abb07c8a-d547-46eb-88f1-19378314ec4e
func (c *Client) GetServiceByName(ctx context.Context, name string) (*Service, error) {
	params := &requestParams{
		method:  "GET",
		subPath: "/services",
		queries: map[string]string{"name": name},
	}

	var serviceResponse ServiceResponse
	if err := c.doAPIRequest(ctx, params, &serviceResponse); err != nil {
		return nil, err
	}

	return serviceResponse.Service, nil
}

// https://apidocs.squadcast.com/#b9722ea8-f97d-4017-b5b0-80986d1ae654
func (c *Client) GetServiceByID(ctx context.Context, id string) (*Service, error) {
	params := &requestParams{
		method:  "GET",
		subPath: fmt.Sprintf("/services/%s", id),
	}

	var serviceResponse ServiceResponse
	if err := c.doAPIRequest(ctx, params, &serviceResponse); err != nil {
		return nil, err
	}

	return serviceResponse.Service, nil
}

func (c *Client) CreateService(ctx context.Context, name, escalationPolicyId, description, emailPrefix string) (*Service, error) {

	request := CreateServiceRequest{
		Name:               name,
		EscalationPolicyId: escalationPolicyId,
		Description:        description,
		EmailPrefix:        emailPrefix,
	}

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(request)

	params := &requestParams{
		method:  "POST",
		subPath: "/services/",
		queries: map[string]string{},
		body:    buffer,
	}

	var serviceResponse ServiceResponse
	if err := c.doAPIRequest(ctx, params, &serviceResponse); err != nil {
		return nil, err
	}
	return serviceResponse.Service, nil
}
