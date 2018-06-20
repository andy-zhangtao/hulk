package service

import (
	"github.com/graphql-go/graphql"
	"github.com/andy-zhangtao/hulk/model"
	"github.com/andy-zhangtao/hulk/env"
)

var RegisterType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Register",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if r, ok := p.Source.(model.Register); ok {
					return r.Name, nil
				}
				return nil, nil
			},
		},
		"resume": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if r, ok := p.Source.(model.Register); ok {
					return r.Resume, nil
				}
				return nil, nil
			},
		},
		"version": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if r, ok := p.Source.(model.Register); ok {
					return r.Version, nil
				}
				return nil, nil
			},
		},
		"time": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if r, ok := p.Source.(model.Register); ok {
					return r.Time, nil
				}
				return nil, nil
			},
		},
		"ip": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if r, ok := p.Source.(model.Register); ok {
					return r.IP, nil
				}
				return nil, nil
			},
		},
	},
})

var NewServiceRegister = &graphql.Field{
	Type:        RegisterType,
	Description: "Register A New Service",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"resume": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"version": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["name"].(string)
		resume, _ := p.Args["resume"].(string)
		version, _ := p.Args["version"].(string)

		r := model.Register{
			Name:    name,
			Version: version,
			Resume:  resume,
			IP:      p.Context.Value(env.RIP("RemoteIP")).(string),
		}

		return r, newRegister(r)
	},
}

var UpdateServiceRegister = &graphql.Field{
	Type:        RegisterType,
	Description: "Update A Specify Register",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"resume": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"version": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["name"].(string)
		resume, _ := p.Args["resume"].(string)
		version, _ := p.Args["version"].(string)

		r := model.Register{
			Name:    name,
			Version: version,
			Resume:  resume,
		}

		return r, updateRegister(r)
	},
}

var DeleteServiceRegister = &graphql.Field{
	Type:        RegisterType,
	Description: "Delete A Specify Register",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"version": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["name"].(string)
		version, _ := p.Args["version"].(string)

		r := model.Register{
			Name:    name,
			Version: version,
		}

		return r, deleteRegister(name, version)
	},
}

var QueryServiceRegister = &graphql.Field{
	Type:        graphql.NewList(RegisterType),
	Description: "Query All Register Service",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return queryAllRegister()
	},
}
