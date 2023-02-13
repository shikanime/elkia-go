package apiserver

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"hash/fnv"

	fleetv1alpha1pb "github.com/infinity-blackhole/elkia/pkg/api/fleet/v1alpha1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type FleetServiceConfig struct {
	KubernetesClientSet *kubernetes.Clientset
	SessionStore        *SessionStore
	IdentityProvider    *IdentityProvider
}

func NewFleetService(config FleetServiceConfig) *FleetService {
	return &FleetService{
		kubernetesClientSet: config.KubernetesClientSet,
		sessionStore:        config.SessionStore,
		identityProvider:    config.IdentityProvider,
	}
}

type FleetService struct {
	fleetv1alpha1pb.UnimplementedFleetServiceServer
	kubernetesClientSet *kubernetes.Clientset
	sessionStore        *SessionStore
	identityProvider    *IdentityProvider
}

func (s *FleetService) ListWorlds(
	ctx context.Context,
	in *fleetv1alpha1pb.ListWorldRequest,
) (*fleetv1alpha1pb.ListWorldResponse, error) {
	nss, err := s.kubernetesClientSet.
		CoreV1().
		Namespaces().
		List(context.Background(), metav1.ListOptions{
			LabelSelector: "fleet.elkia.io/world=true",
		})
	if err != nil {
		return nil, err
	}
	worlds := make([]*fleetv1alpha1pb.World, len(nss.Items))
	for i, ns := range nss.Items {
		worlds[i] = &fleetv1alpha1pb.World{
			Id:   ns.Name,
			Name: ns.Labels["fleet.elkia.io/world-name"],
		}
	}
	return &fleetv1alpha1pb.ListWorldResponse{
		Worlds: worlds,
	}, nil
}

func (s *FleetService) ListGateways(
	ctx context.Context,
	in *fleetv1alpha1pb.ListGatewayRequest,
) (*fleetv1alpha1pb.ListGatewayResponse, error) {
	svcs, err := s.kubernetesClientSet.
		CoreV1().
		Services(in.WorldId).
		List(context.Background(), metav1.ListOptions{
			LabelSelector: "fleet.elkia.io/gateway=true",
			FieldSelector: "spec.type=LoadBalancer",
		})
	if err != nil {
		return nil, err
	}
	gateways := make([]*fleetv1alpha1pb.Gateway, len(svcs.Items))
	for i, svc := range svcs.Items {
		addr, err := s.getGatewayAddrFromService(&svc)
		if err != nil {
			return nil, err
		}
		gateways[i] = &fleetv1alpha1pb.Gateway{
			Id:      svc.Name,
			Address: addr,
		}
	}
	return &fleetv1alpha1pb.ListGatewayResponse{
		Gateways: gateways,
	}, nil
}

func (s *FleetService) getGatewayAddrFromService(
	svc *corev1.Service,
) (string, error) {
	ip, err := s.getGatewayIpFromService(svc)
	if err != nil {
		return "", err
	}
	port, err := s.getGatewayPortFromService(svc)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(ip, port), nil
}

func (s *FleetService) getGatewayIpFromService(
	svc *corev1.Service,
) (string, error) {
	if svc.Spec.LoadBalancerIP == "" {
		return "", fmt.Errorf(
			"service %s/%s has no LoadBalancerIP",
			svc.Namespace,
			svc.Name,
		)
	}
	return svc.Spec.LoadBalancerIP, nil
}

func (s *FleetService) getGatewayPortFromService(
	svc *corev1.Service,
) (int32, error) {
	if len(svc.Spec.Ports) == 0 {
		return 0, fmt.Errorf(
			"service %s/%s has no ports",
			svc.Namespace,
			svc.Name,
		)
	}
	for _, port := range svc.Spec.Ports {
		if port.Name == "nostale" {
			return port.Port, nil
		}
	}
	return 0, fmt.Errorf(
		"service %s/%s has no port named 'nostale'",
		svc.Namespace,
		svc.Name,
	)
}

func (s *FleetService) CreateHandoff(
	ctx context.Context,
	in *fleetv1alpha1pb.CreateHandoffRequest,
) (*fleetv1alpha1pb.CreateHandoffResponse, error) {
	session, token, err := s.identityProvider.
		PerformLoginFlowWithPasswordMethod(
			ctx,
			in.Identifier,
			in.Token,
		)
	if err != nil {
		return nil, err
	}
	h := fnv.New32a()
	if err := gob.
		NewEncoder(h).
		Encode(session.Id); err != nil {
		panic(err)
	}
	key := h.Sum32()
	if err := s.sessionStore.SetHandoffSession(
		ctx,
		key,
		&fleetv1alpha1pb.Handoff{
			Identifier: in.Identifier,
			Token:      token,
		},
	); err != nil {
		return nil, err
	}
	return &fleetv1alpha1pb.CreateHandoffResponse{
		Key: key,
	}, nil
}

func (s *FleetService) PerformHandoff(
	ctx context.Context,
	in *fleetv1alpha1pb.PerformHandoffRequest,
) (*emptypb.Empty, error) {
	handoff, err := s.sessionStore.GetHandoffSession(ctx, in.Key)
	if err != nil {
		return nil, err
	}
	if err := s.sessionStore.DeleteHandoffSession(ctx, in.Key); err != nil {
		return nil, err
	}
	keySession, err := s.identityProvider.GetSession(ctx, handoff.Token)
	if err != nil {
		return nil, err
	}
	if err := s.identityProvider.Logout(ctx, handoff.Token); err != nil {
		return nil, err
	}
	credentialsSession, _, err := s.identityProvider.
		PerformLoginFlowWithPasswordMethod(
			ctx,
			handoff.Identifier,
			in.Token,
		)
	if err != nil {
		return nil, err
	}
	if keySession.Identity.Id != credentialsSession.Identity.Id {
		return nil, errors.New("invalid credentials")
	}
	return &emptypb.Empty{}, nil
}
