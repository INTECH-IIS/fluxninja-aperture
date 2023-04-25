package delete

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes/scheme"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"

	languagev1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/policy/language/v1"
	"github.com/fluxninja/aperture/cmd/aperturectl/cmd/utils"
	"github.com/fluxninja/aperture/operator/api"
	policyv1alpha1 "github.com/fluxninja/aperture/operator/api/policy/v1alpha1"
	"github.com/fluxninja/aperture/pkg/log"
)

// DeletePolicyCmd is the command to apply a policy to the cluster.
var DeletePolicyCmd = &cobra.Command{
	Use:           "policy",
	Short:         "Delete Aperture Policy from the cluster",
	Long:          `Use this command to delete the Aperture Policy from the cluster.`,
	SilenceErrors: true,
	Example:       `aperturectl delete policy --policy=static-rate-limiting`,
	RunE: func(_ *cobra.Command, _ []string) error {
		return deletePolicy()
	},
}

// deletePolicy deletes the policy from the cluster.
func deletePolicy() error {
	deployment, err := utils.GetControllerDeployment(kubeRestConfig, controllerNs)
	if err != nil {
		return err
	}

	err = api.AddToScheme(scheme.Scheme)
	if err != nil {
		return fmt.Errorf("failed to connect to Kubernetes: %w", err)
	}

	kubeClient, err := k8sclient.New(kubeRestConfig, k8sclient.Options{
		Scheme: scheme.Scheme,
	})
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes client: %w", err)
	}

	policyCR := &policyv1alpha1.Policy{}
	policyCR.Name = policyName
	policyCR.Namespace = deployment.GetNamespace()
	err = kubeClient.Delete(context.Background(), policyCR)
	if err != nil && !apierrors.IsNotFound(err) {
		if strings.Contains(err.Error(), "no matches for kind") {
			err = deletePolicyUsingAPI()
		}

		if err != nil {
			return fmt.Errorf("failed to delete policy in Kubernetes: %w", err)
		}
	}

	log.Info().Str("policy", policyName).Str("namespace", deployment.GetNamespace()).Msg("Deleted Policy successfully")
	return nil
}

// deletePolicyUsingAPI deletes the policy using the API.
func deletePolicyUsingAPI() error {
	policyRequest := languagev1.DeletePolicyRequest{
		Name: policyName,
	}
	_, err := client.DeletePolicy(context.Background(), &policyRequest)
	if err != nil {
		log.Warn().Err(err).Str("policy", policyName).Msg("failed to delete Policy")
	}

	return nil
}