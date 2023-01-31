/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"sync"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	gkeclusterapis "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/container/v1beta1"
	gitopsv1alpha1 "github.com/GoogleContainerTools/kpt/rollouts/api/v1alpha1"
	"github.com/GoogleContainerTools/kpt/rollouts/pkg/clusterstore"
	"github.com/GoogleContainerTools/kpt/rollouts/pkg/packagediscovery"
	cloudresourcemanagerv1 "google.golang.org/api/cloudresourcemanager/v1"
	gkehubv1 "google.golang.org/api/gkehub/v1"
	"google.golang.org/api/option"
)

type Options struct {
}

type RemoteClient struct {
	restConfig *rest.Config
}

func (r *RemoteClient) DynamicClient() (dynamic.Interface, error) {
	dynamicClient, err := dynamic.NewForConfig(r.restConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new dynamic client: %w", err)
	}
	return dynamicClient, nil
}

type MemberEndpoint struct {
	GkeCluster *GkeCluster `json:"gkeCluster"`
}

type GkeCluster struct {
}

func (o *Options) InitDefaults() {
}

func (o *Options) BindFlags(prefix string, flags *flag.FlagSet) {
}

var (
	namespaceGVR = schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "namespaces",
	}
)

// RolloutReconciler reconciles a Rollout object
type RolloutReconciler struct {
	client.Client

	store *clusterstore.ClusterStore

	Scheme *runtime.Scheme

	mutex                 sync.Mutex
	packageDiscoveryCache map[types.NamespacedName]*packagediscovery.PackageDiscovery
}

//+kubebuilder:rbac:groups=gitops.kpt.dev,resources=rollouts,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=gitops.kpt.dev,resources=rollouts/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=gitops.kpt.dev,resources=rollouts/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Rollout object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.1/pkg/reconcile
//
// Dumb reconciliation of Rollout API includes the following:
// Fetch the READY kcc clusters.
// For each kcc cluster, fetch RootSync objects in each of the KCC clusters.
func (r *RolloutReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	logger.Info("reconciling", "key", req.NamespacedName)

	projectId := "christopher-exp-anthos-dev"

	accessToken, err := google.DefaultTokenSource(ctx, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("unable to get default access-token from gcloud: %w", err)
	}

	token, err := accessToken.Token()
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("get access token failed: %w", err)
	}

	thisAccessToken := token.AccessToken
	crmClient, err := cloudresourcemanagerv1.NewService(ctx, option.WithTokenSource(accessToken))
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to create new cloudresourcemanager client: %w", err)
	}

	project, err := crmClient.Projects.Get(projectId).Context(ctx).Do()
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("error querying project %q: %w", projectId, err)
	}

	projectNum := project.ProjectNumber
	logger.Info("reconciling", "project number %d", project.ProjectNumber)

	memberships, err := listMemberships(ctx, projectId, accessToken) //, thisAccessToken)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("memberships failed: %w", err)
	}

	fmt.Println(memberships)

	for i := range memberships.Resources {
		membership := memberships.Resources[i]

		restConfig, err := buildMembershipRESTConfig(projectNum, membership, thisAccessToken)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("memberships failed: %w", err)
		}

		remoteClient := &RemoteClient{
			restConfig: restConfig,
		}

		dynamicClient, err := remoteClient.DynamicClient()

		fmt.Println("BEFORE", restConfig.Host)

		// SUNIL - Correct way of adding a timeout
		// ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
		// defer cancel()

		lst, err := dynamicClient.Resource(namespaceGVR).List(context.Background(), v1.ListOptions{})

		if err != nil {
			fmt.Printf("%s\n", err.Error())
			continue
		}
		fmt.Println(lst.Items)
		fmt.Println("AFTER", restConfig.Host)
		fmt.Println()
	}

	fmt.Println("DONE")

	return ctrl.Result{}, nil
}

func buildMembershipRESTConfig(projectNumber int64, membership *gkehubv1.Membership, accessToken string) (*rest.Config, error) {
	restConfig := &rest.Config{}

	membershipName := membership.Name[strings.LastIndex(membership.Name, "/")+1:]

	isGKE := membership.Endpoint.GkeCluster != nil

	membershipUrl := "memberships"

	if isGKE {
		membershipUrl = "gkeMemberships"
	}

	host := fmt.Sprintf("https://connectgateway.googleapis.com/v1/projects/%d/locations/global/%s/%s", projectNumber, membershipUrl, membershipName)
	restConfig.Host = host
	restConfig.BearerToken = accessToken

	return restConfig, nil
}

func listMemberships(ctx context.Context, projectId string, token oauth2.TokenSource) (*gkehubv1.ListMembershipsResponse, error) {
	hubClient, err := gkehubv1.NewService(ctx, option.WithTokenSource(token))
	if err != nil {
		return nil, err
	}

	parent := fmt.Sprintf("projects/%s/locations/-", projectId)
	resp, err := hubClient.Projects.Locations.Memberships.List(parent).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RolloutReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if err := gkeclusterapis.AddToScheme(mgr.GetScheme()); err != nil {
		return err
	}
	if err := gitopsv1alpha1.AddToScheme(mgr.GetScheme()); err != nil {
		return err
	}
	r.Client = mgr.GetClient()

	r.packageDiscoveryCache = make(map[types.NamespacedName]*packagediscovery.PackageDiscovery)

	// setup the clusterstore
	r.store = &clusterstore.ClusterStore{
		Config: mgr.GetConfig(),
		Client: r.Client,
	}
	if err := r.store.Init(); err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&gitopsv1alpha1.Rollout{}).
		Owns(&gitopsv1alpha1.RemoteRootSync{}).
		Complete(r)
}
