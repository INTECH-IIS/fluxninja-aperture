package discovery

import (
	"context"

	"github.com/fluxninja/aperture/pkg/config"
	"github.com/fluxninja/aperture/pkg/etcd/election"
	"github.com/fluxninja/aperture/pkg/k8s"
	"github.com/fluxninja/aperture/pkg/log"
	"github.com/fluxninja/aperture/pkg/notifiers"
	autoscalek8sconfig "github.com/fluxninja/aperture/pkg/policies/autoscale/kubernetes/config"
	"github.com/fluxninja/aperture/pkg/status"
	"go.uber.org/fx"
)

var (
	// FxTagBase is the tag's base used to identify the Kubernetes Control Points Tracker.
	FxTagBase = "kubernetes_control_points"
	// FxTag is the tag used to identify the Kubernetes Control Points Tracker.
	FxTag = config.NameTag(FxTagBase)
)

// Module returns the fx options for Kubernetes Control Point Discovery.
func Module() fx.Option {
	return fx.Options(
		notifiers.TrackersConstructor{Name: FxTagBase}.Annotate(),
		fx.Provide(provideAutoScaleControlPoints),
	)
}

// FxIn is the input for the ProvideKuberetesControlPointsCache function.
type FxIn struct {
	fx.In
	Unmarshaller     config.Unmarshaller
	Lifecycle        fx.Lifecycle
	StatusRegistry   status.Registry
	KubernetesClient k8s.K8sClient      `optional:"true"`
	Trackers         notifiers.Trackers `name:"kubernetes_control_points"`
	Election         *election.Election
	Config           autoscalek8sconfig.AutoScaleKubernetesConfig
}

// provideAutoScaleControlPoints provides Kubernetes AutoScaler and starts Kubernetes control point discovery if enabled.
func provideAutoScaleControlPoints(in FxIn) (AutoScaleControlPoints, error) {
	controlPointCache := newAutoScaleControlPoints(in.Trackers, in.KubernetesClient)

	if !in.Config.Enabled {
		log.Info().Msg("Skipping Kubernetes Control Point Discovery since AutoScale is disabled")
		return controlPointCache, nil
	}
	if in.KubernetesClient == nil {
		log.Error().Msg("Kubernetes client is not available, skipping Kubernetes Control Point Discovery")
		return controlPointCache, nil
	}
	cpd, err := newControlPointDiscovery(in.Election, in.KubernetesClient, controlPointCache)
	if err != nil {
		log.Info().Err(err).Msg("Failed to create Kubernetes Control Point Discovery")
		return nil, err
	}

	in.Lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			controlPointCache.start()
			cpd.start()
			return nil
		},
		OnStop: func(_ context.Context) error {
			cpd.stop()
			controlPointCache.stop()
			return nil
		},
	})

	return controlPointCache, nil
}