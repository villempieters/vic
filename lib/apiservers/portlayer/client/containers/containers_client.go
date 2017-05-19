package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new containers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for containers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Commit Commit and close a container handle
*/
func (a *Client) Commit(params *CommitParams) (*CommitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCommitParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Commit",
		Method:             "PUT",
		PathPattern:        "/containers/{handle}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CommitReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CommitOK), nil

}

/*
ContainerRemove Remove a container from existence
*/
func (a *Client) ContainerRemove(params *ContainerRemoveParams) (*ContainerRemoveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerRemoveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerRemove",
		Method:             "DELETE",
		PathPattern:        "/containers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ContainerRemoveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerRemoveOK), nil

}

/*
ContainerRename Rename a container to the new name
*/
func (a *Client) ContainerRename(params *ContainerRenameParams) (*ContainerRenameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerRenameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerRename",
		Method:             "PATCH",
		PathPattern:        "/containers/{handle}/name",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ContainerRenameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerRenameOK), nil

}

/*
ContainerSignal signals a running container

Sends a signal to a container by id
*/
func (a *Client) ContainerSignal(params *ContainerSignalParams) (*ContainerSignalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerSignalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerSignal",
		Method:             "POST",
		PathPattern:        "/containers/{id}/signal",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ContainerSignalReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerSignalOK), nil

}

/*
ContainerWait Wait for the container to stop
*/
func (a *Client) ContainerWait(params *ContainerWaitParams) (*ContainerWaitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerWaitParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerWait",
		Method:             "GET",
		PathPattern:        "/containers/{id}/wait",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ContainerWaitReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerWaitOK), nil

}

/*
Create initiates a container create operation

Initiates a container create operation
*/
func (a *Client) Create(params *CreateParams) (*CreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Create",
		Method:             "POST",
		PathPattern:        "/containers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateOK), nil

}

/*
Get Get a container handle
*/
func (a *Client) Get(params *GetParams) (*GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Get",
		Method:             "GET",
		PathPattern:        "/containers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOK), nil

}

/*
GetContainerInfo Gets information about a container by id
*/
func (a *Client) GetContainerInfo(params *GetContainerInfoParams) (*GetContainerInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContainerInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetContainerInfo",
		Method:             "GET",
		PathPattern:        "/containers/info/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetContainerInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetContainerInfoOK), nil

}

/*
GetContainerList Gets a list of all containers
*/
func (a *Client) GetContainerList(params *GetContainerListParams) (*GetContainerListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContainerListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetContainerList",
		Method:             "GET",
		PathPattern:        "/containers/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetContainerListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetContainerListOK), nil

}

/*
GetContainerLogs gets the container logs

Gets the container logs by id
*/
func (a *Client) GetContainerLogs(params *GetContainerLogsParams, writer io.Writer) (*GetContainerLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContainerLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetContainerLogs",
		Method:             "GET",
		PathPattern:        "/containers/{id}/logs",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetContainerLogsReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetContainerLogsOK), nil

}

/*
GetContainerStats gets the container stats

Gets the container stats by id
*/
func (a *Client) GetContainerStats(params *GetContainerStatsParams, writer io.Writer) (*GetContainerStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContainerStatsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetContainerStats",
		Method:             "GET",
		PathPattern:        "/containers/{id}/stats",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetContainerStatsReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetContainerStatsOK), nil

}

/*
GetState Get the current state of the a container
*/
func (a *Client) GetState(params *GetStateParams) (*GetStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetState",
		Method:             "GET",
		PathPattern:        "/containers/{handle}/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStateOK), nil

}

/*
StateChange Changes the state of a container
*/
func (a *Client) StateChange(params *StateChangeParams) (*StateChangeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStateChangeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StateChange",
		Method:             "PUT",
		PathPattern:        "/containers/{handle}/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StateChangeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StateChangeOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
