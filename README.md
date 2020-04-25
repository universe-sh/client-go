# Go API client for openapi

Universe.sh houston API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *http://api-houston-$.endpoints.$.cloud.goog*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountApi* | [**ListAccount**](docs/AccountApi.md#listaccount) | **Get** /v1/account | 
*EventsApi* | [**AddEvents**](docs/EventsApi.md#addevents) | **Post** /v1/events | 
*EventsApi* | [**ListEvents**](docs/EventsApi.md#listevents) | **Get** /v1/events | 
*EventsApi* | [**ReadEvent**](docs/EventsApi.md#readevent) | **Get** /v1/events/{event} | 
*SatellitesApi* | [**AddPools**](docs/SatellitesApi.md#addpools) | **Post** /v1/satellites/{cloud}/{satellite}/pools | 
*SatellitesApi* | [**AddSatellites**](docs/SatellitesApi.md#addsatellites) | **Post** /v1/satellites/{cloud} | 
*SatellitesApi* | [**CreateMetrics**](docs/SatellitesApi.md#createmetrics) | **Post** /v1/satellites/{cloud}/{satellite}/{pool}/metrics | 
*SatellitesApi* | [**DeletePool**](docs/SatellitesApi.md#deletepool) | **Delete** /v1/satellites/{cloud}/{satellite}/{pool} | 
*SatellitesApi* | [**DeleteSatellite**](docs/SatellitesApi.md#deletesatellite) | **Delete** /v1/satellites/{cloud}/{satellite} | 
*SatellitesApi* | [**ListMetrics**](docs/SatellitesApi.md#listmetrics) | **Get** /v1/satellites/{cloud}/{satellite}/{pool}/metrics | 
*SatellitesApi* | [**ListPools**](docs/SatellitesApi.md#listpools) | **Get** /v1/satellites/{cloud}/{satellite}/pools | 
*SatellitesApi* | [**ListSatellites**](docs/SatellitesApi.md#listsatellites) | **Get** /v1/satellites/{cloud} | 
*SatellitesApi* | [**ReadPool**](docs/SatellitesApi.md#readpool) | **Get** /v1/satellites/{cloud}/{satellite}/{pool} | 
*SatellitesApi* | [**ReadSatellite**](docs/SatellitesApi.md#readsatellite) | **Get** /v1/satellites/{cloud}/{satellite} | 
*SatellitesApi* | [**UpdatePool**](docs/SatellitesApi.md#updatepool) | **Put** /v1/satellites/{cloud}/{satellite}/{pool} | 
*SatellitesApi* | [**UpdateSatellite**](docs/SatellitesApi.md#updatesatellite) | **Put** /v1/satellites/{cloud}/{satellite} | 
*TeamsApi* | [**AddTeams**](docs/TeamsApi.md#addteams) | **Post** /v1/teams | 
*TeamsApi* | [**DeleteInvitations**](docs/TeamsApi.md#deleteinvitations) | **Delete** /v1/teams/invitations/{invitation} | 
*TeamsApi* | [**DeleteTeam**](docs/TeamsApi.md#deleteteam) | **Delete** /v1/teams/{team} | 
*TeamsApi* | [**ListInvitations**](docs/TeamsApi.md#listinvitations) | **Get** /v1/teams/{team}/invitations | 
*TeamsApi* | [**ListInvoices**](docs/TeamsApi.md#listinvoices) | **Get** /v1/teams/{team}/invoices | 
*TeamsApi* | [**ListTeams**](docs/TeamsApi.md#listteams) | **Get** /v1/teams | 
*TeamsApi* | [**ReadInvitations**](docs/TeamsApi.md#readinvitations) | **Get** /v1/teams/invitations/{invitation} | 
*TeamsApi* | [**ReadInvitationsAccept**](docs/TeamsApi.md#readinvitationsaccept) | **Get** /v1/teams/invitations/{invitation}/accept | 
*TeamsApi* | [**ReadInvoice**](docs/TeamsApi.md#readinvoice) | **Get** /v1/teams/{team}/invoices/{invoice} | 
*TeamsApi* | [**ReadPreferences**](docs/TeamsApi.md#readpreferences) | **Get** /v1/teams/{team}/preferences | 
*TeamsApi* | [**ReadTeam**](docs/TeamsApi.md#readteam) | **Get** /v1/teams/{team} | 
*TeamsApi* | [**UpdateInvitations**](docs/TeamsApi.md#updateinvitations) | **Put** /v1/teams/invitations/{invitation} | 
*TeamsApi* | [**UpdatePreferences**](docs/TeamsApi.md#updatepreferences) | **Put** /v1/teams/{team}/preferences | 
*TeamsApi* | [**UpdateTeam**](docs/TeamsApi.md#updateteam) | **Put** /v1/teams/{team} | 


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountDefaultTeam](docs/AccountDefaultTeam.md)
 - [Event](docs/Event.md)
 - [Generic](docs/Generic.md)
 - [Invitation](docs/Invitation.md)
 - [InvitationInvitedBy](docs/InvitationInvitedBy.md)
 - [InvitationTeam](docs/InvitationTeam.md)
 - [InvitationUser](docs/InvitationUser.md)
 - [Invoice](docs/Invoice.md)
 - [Metric](docs/Metric.md)
 - [MetricData](docs/MetricData.md)
 - [Pool](docs/Pool.md)
 - [PoolAutoscaling](docs/PoolAutoscaling.md)
 - [PoolLaunchSpecification](docs/PoolLaunchSpecification.md)
 - [PoolLaunchSpecificationConfig](docs/PoolLaunchSpecificationConfig.md)
 - [PoolLaunchSpecificationRootDisk](docs/PoolLaunchSpecificationRootDisk.md)
 - [Preferences](docs/Preferences.md)
 - [Satellite](docs/Satellite.md)
 - [Team](docs/Team.md)
 - [TeamUsers](docs/TeamUsers.md)


## Documentation For Authorization



## okta_jwt


- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: 
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```



## Author

oss@universe.sh

