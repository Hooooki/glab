/* MIT License

Copyright (c) 2023 Fragan Gourvil

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE */

package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitlab-environment/pkg/config"
	"gitlab-environment/pkg/entity"
	"gitlab-environment/pkg/rest"
	"net/url"
	"strings"
)

const Endpoint = "/environments"

func ListEnvironment(c *config.Config) (*[]entity.Environment, error) {

	var environments []entity.Environment

	r := &rest.Request{
		Endpoint: Endpoint,
		Method:   rest.GET,
		Config:   *c,
	}

	resp := r.Send()

	if resp.StatusCode != 200 {
		return nil, errors.New(string(resp.Body))
	}

	err := json.Unmarshal(resp.Body, &environments)
	return &environments, err
}

func GetEnvironment(c *config.Config, id int) (*entity.Environment, error) {

	var environment entity.Environment

	r := &rest.Request{
		Endpoint: fmt.Sprintf("%s/%d", Endpoint, id),
		Method:   rest.GET,
		Config:   *c,
	}

	resp := r.Send()

	if resp.StatusCode != 200 {
		return nil, errors.New(string(resp.Body))
	}

	err := json.Unmarshal(resp.Body, &environment)
	return &environment, err
}

func AddEnvironment(c *config.Config, e entity.Environment) (*entity.Environment, error) {

	d := url.Values{
		"name": {e.Name},
	}

	if "" != e.ExternalUrl {
		d.Add("external_url", e.ExternalUrl)
	}

	r := rest.Request{
		Endpoint: Endpoint,
		Method:   rest.POST,
		Config:   *c,
		Body:     strings.NewReader(d.Encode()),
	}

	resp := r.Send()

	if resp.StatusCode != 201 {
		return nil, errors.New(string(resp.Body))
	}

	err := json.Unmarshal(resp.Body, &e)
	return &e, err
}

func UpdateEnvironment(c *config.Config, e entity.Environment) (*entity.Environment, error) {

	d := url.Values{}

	if "" == e.ExternalUrl && "" == e.Tier {
		return nil, errors.New("nothing to update")
	}

	if "" != e.ExternalUrl {
		d.Add("external_url", e.ExternalUrl)
	}

	if "" != e.Tier {
		d.Add("tier", e.Tier)
	}

	r := rest.Request{
		Endpoint: fmt.Sprintf("%s/%d", Endpoint, e.Id),
		Method:   rest.PUT,
		Config:   *c,
		Body:     strings.NewReader(d.Encode()),
	}

	resp := r.Send()

	if resp.StatusCode != 200 {
		return nil, errors.New(string(resp.Body))
	}

	err := json.Unmarshal(resp.Body, &e)
	return &e, err

}

func StopEnvironment(c *config.Config, id int) (*entity.Environment, error) {

	var environment entity.Environment

	r := rest.Request{
		Endpoint: fmt.Sprintf("%s/%d/stop", Endpoint, id),
		Method:   rest.POST,
		Config:   *c,
	}

	resp := r.Send()

	if resp.StatusCode != 200 {
		return nil, errors.New(string(resp.Body))
	}

	err := json.Unmarshal(resp.Body, &environment)
	return &environment, err
}

func DeleteEnvironment(c *config.Config, id int) error {

	r := &rest.Request{
		Endpoint: fmt.Sprintf("%s/%d", Endpoint, id),
		Method:   rest.DELETE,
		Config:   *c,
	}

	resp := r.Send()

	if resp.StatusCode != 204 {
		return errors.New(string(resp.Body))
	}

	return nil
}
