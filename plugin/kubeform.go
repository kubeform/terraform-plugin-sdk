package plugin

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/internal/tfplugin5"
	"github.com/pkg/errors"

	"github.com/hashicorp/terraform-plugin-sdk/internal/helper/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

type KubeformServer struct {
	provider *plugin.GRPCProviderServer
}

func NewKubeformServer(p terraform.ResourceProvider) *KubeformServer {
	return &KubeformServer{
		provider: plugin.NewGRPCProviderServerShim(p),
	}
}

func (k *KubeformServer) PrepareProviderConfig(config []byte) ([]byte, error) {
	req := tfplugin5.PrepareProviderConfig_Request{
		Config: &tfplugin5.DynamicValue{
			Msgpack: config,
		},
	}
	resp, err := k.provider.PrepareProviderConfig(context.TODO(), &req)
	if err != nil {
		return nil, err
	}

	return resp.PreparedConfig.Msgpack, diagToError(resp.Diagnostics)
}

func (k *KubeformServer) Configure(config []byte) error {
	req := tfplugin5.Configure_Request{
		Config: &tfplugin5.DynamicValue{
			Msgpack: config,
		},
	}
	resp, err := k.provider.Configure(context.TODO(), &req)
	if err != nil {
		return err
	}

	return diagToError(resp.Diagnostics)
}

func (k *KubeformServer) ValidateResourceTypeConfig(typeName string, config []byte) error {
	internalConfig := tfplugin5.DynamicValue{
		Msgpack: config,
	}
	req := tfplugin5.ValidateResourceTypeConfig_Request{
		TypeName: typeName,
		Config:   &internalConfig,
	}
	resp, err := k.provider.ValidateResourceTypeConfig(context.TODO(), &req)
	if err != nil {
		return err
	}

	return diagToError(resp.Diagnostics)
}

func (k *KubeformServer) PlanResourceChange(typeName string, priorState, proposedNewState []byte) ([]byte, []byte, bool, error) {
	internalPriorState := tfplugin5.DynamicValue{
		Msgpack: priorState,
	}
	internalProposedState := tfplugin5.DynamicValue{
		Msgpack: proposedNewState,
	}
	req := tfplugin5.PlanResourceChange_Request{
		TypeName:         typeName,
		PriorState:       &internalPriorState,
		ProposedNewState: &internalProposedState,
	}
	resp, err := k.provider.PlanResourceChange(context.TODO(), &req)
	if err != nil {
		return nil, nil, false, err
	}

	return resp.PlannedState.Msgpack, resp.PlannedPrivate, len(resp.RequiresReplace) > 0, diagToError(resp.Diagnostics)
}

func (k *KubeformServer) ApplyResourceChange(typeName string, priorState, plannedState, config, plannedPrivate []byte) ([]byte, error) {
	req := tfplugin5.ApplyResourceChange_Request{
		TypeName: typeName,
		PriorState: &tfplugin5.DynamicValue{
			Msgpack: priorState,
		},
		PlannedState: &tfplugin5.DynamicValue{
			Msgpack: plannedState,
		},
		Config: &tfplugin5.DynamicValue{
			Msgpack: config,
		},
		PlannedPrivate: plannedPrivate,
	}
	resp, err := k.provider.ApplyResourceChange(context.TODO(), &req)
	if err != nil {
		return nil, err
	}

	return resp.NewState.Msgpack, diagToError(resp.Diagnostics)
}

func diagToError(d []*tfplugin5.Diagnostic) error {
	var err error
	var flag bool
	for idx, key := range d {
		if key.Severity.String() == "WARNING" || key.Summary == "Invalid or unknown key" {
			continue
		}
		if !flag {
			err = errors.New("")
			flag = true
		}
		err = errors.Wrapf(err, "%s %d: %s", "Error", idx, key.Summary)
	}
	return err
}
