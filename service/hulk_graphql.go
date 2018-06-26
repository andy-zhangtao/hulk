package service

import (
	"github.com/graphql-go/graphql"
	"github.com/andy-zhangtao/hulk/model"
	"encoding/json"
	"errors"
	"fmt"
)

var HulkType = graphql.NewObject(graphql.ObjectConfig{
	Name: "hulk",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if h, ok := p.Source.(model.Hulk); ok {
					return h.Name, nil
				}
				return nil, nil
			},
		},
		"version": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if h, ok := p.Source.(model.Hulk); ok {
					return h.Version, nil
				}
				return nil, nil
			},
		},
		"time": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if h, ok := p.Source.(model.Hulk); ok {
					return h.Time, nil
				}
				return nil, nil
			},
		},
		"configure": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if h, ok := p.Source.(model.Hulk); ok {
					data, err := json.Marshal(h.Configure)
					if err != nil {
						return nil, err
					}
					return string(data), nil
				}
				return nil, nil
			},
		},
	},
})

var QueryAllHulk = &graphql.Field{
	Type:        graphql.NewList(HulkType),
	Description: "query all the hulk configure",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"version": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["name"].(string)
		version, _ := p.Args["version"].(string)

		if (name != "" && version != "") {
			return querySpecifyVersionHulk(name, version)
		} else if (name != "") {
			return querySpecifyHulk(name)
		} else {
			return queryALLHulk()
		}
	},
}

var QuerySpecifyHulk = &graphql.Field{
	Type:        graphql.NewList(HulkType),
	Description: "query specify hulk configure",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"version": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["name"].(string)
		version, _ := p.Args["version"].(string)

		if version == "" {
			return querySpecifyHulk(name)
		}

		return querySpecifyVersionHulk(name, version)
	},
}

var SaveHulk = &graphql.Field{
	Type:        HulkType,
	Description: "add a new Hulk configure",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"version": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"configure": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["name"].(string)
		version, _ := p.Args["version"].(string)
		configure, _ := p.Args["configure"].(string)

		c := make(map[string]interface{})

		err := json.Unmarshal([]byte(configure), &c)
		if err != nil {
			return nil, err
		}

		h := model.Hulk{
			Name:      name,
			Version:   version,
			Configure: c,
		}
		return h, newHulk(h)
	},
}

var UpdateHulk = &graphql.Field{
	Type:        HulkType,
	Description: "add a new Hulk configure",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"version": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"configure": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["name"].(string)
		version, _ := p.Args["version"].(string)
		configure, _ := p.Args["configure"].(string)

		c := make(map[string]interface{})

		err := json.Unmarshal([]byte(configure), &c)
		if err != nil {
			return nil, err
		}

		h := model.Hulk{
			Name:      name,
			Version:   version,
			Configure: c,
		}
		return h, updateHulk(h)
	},
}

var DeleteHulk = &graphql.Field{
	Type:        HulkType,
	Description: "Delete Specify Version Hulk",
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

		return model.Hulk{
			Name:    name,
			Version: version,
		}, deleteHulk(name, version)
	},
}

var CopyHulk = &graphql.Field{
	Type:        HulkType,
	Description: "Copy Specify Version Hulk",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"version": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"newname": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"newversion": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["name"].(string)
		version, _ := p.Args["version"].(string)
		newname, _ := p.Args["newname"].(string)
		newversion, _ := p.Args["newversion"].(string)

		h, err := querySpecifyVersionHulk(name, version)
		if err != nil {
			return nil, err
		}

		if len(h) == 0 {
			return nil, errors.New(fmt.Sprintf("Can Not Find This Hulk Name:[%s] Version:[%s]", name, version))
		}

		_h := model.Hulk{
			Name:      newname,
			Version:   newversion,
			Configure: h[0].Configure,
		}

		return _h, newHulk(_h)
	},
}
