package resolvers

import (
	"github.com/graphql-go/graphql"
	"os"
)

var getSpotifyApiKey = &graphql.Field{
	Type: graphql.String,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return os.Getenv("SPOTIFY_API_CLIENT_ID"), nil
	},
}

var getSpotifyApiSecret = &graphql.Field{
	Type: graphql.String,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return os.Getenv("SPOTIFY_API_CLIENT_SECRET"), nil
	},
}

var getNasaApiKey = &graphql.Field{
	Type: graphql.String,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return os.Getenv("NASA_API_KEY"), nil
	},
}
