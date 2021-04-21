package entity

import "github.com/Galdoba/TR_Dynasty/pkg/entity/asset"

const (
	EntityTypeTraveller = 1
)

type Entity struct {
	Name        string
	Description string
	Usage       string
	Type        int
	Call        map[string]asset.Asset
}
