package auth

import (
	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
	"log"
	"os"
)

type Auth0Client struct {
	client *management.Management
}

var instance *Auth0Client

func GetAuth0Client() *Auth0Client {
	if instance == nil {
		domain := os.Getenv("AUTH0_DOMAIN")
		id := os.Getenv("AUTH0_CLIENT_ID")
		secret := os.Getenv("AUTH0_CLIENT_SECRET")
		m, err := management.New(domain, management.WithClientCredentials(id, secret))
		if err != nil {
			log.Fatalf("Failed to get auth0 client: %+v", err)
		}
		instance = &Auth0Client{
			m,
		}
		return instance
	} else {
		return instance
	}
}

type Permission struct {
	Name       string
	Identifier string
}

type Role struct {
	Id          string
	Name        string
	Permissions []*Permission
}

func (c *Auth0Client) GetRoles() ([]*Role, error) {
	list, err := c.client.Role.List()
	if err != nil {
		return nil, err
	}

	result := make([]*Role, 0, len(list.Roles))

	for _, v := range list.Roles {
		p, err := c.getPermissions(auth0.StringValue(v.ID))
		if err != nil {
			return nil, err
		}
		result = append(result, &Role{
			auth0.StringValue(v.ID), auth0.StringValue(v.Name), p,
		})
	}
	return result, nil
}

func (c *Auth0Client) getPermissions(roleId string) ([]*Permission, error) {
	list, err := c.client.Role.Permissions(roleId)
	if err != nil {
		return nil, err
	}

	result := make([]*Permission, 0, len(list.Permissions))
	for _, v := range list.Permissions {
		result = append(result, &Permission{
			auth0.StringValue(v.Name), auth0.StringValue(v.ResourceServerIdentifier),
		})
	}
	return result, nil
}
