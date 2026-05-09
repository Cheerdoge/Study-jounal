package registry

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type ServiceRegistry struct {
	client        *clientv3.Client
	leaseID       clientv3.LeaseID
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	key           string
	value         string
}

func NewServiceRegistry(endpoints []string) (*ServiceRegistry, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return &ServiceRegistry{client: cli}, nil
}

func (r *ServiceRegistry) Register(serviceName, serviceAddr string, ttl int64) error {
	r.key = "/services/" + serviceName + "/" + serviceAddr
	r.value = `{"Addr":"` + serviceAddr + `"}`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	leaseResp, err := r.client.Grant(ctx, ttl)
	if err != nil {
		return err
	}
	r.leaseID = leaseResp.ID

	_, err = r.client.Put(ctx, r.key, r.value, clientv3.WithLease(r.leaseID))
	if err != nil {
		return err
	}

	r.keepAliveChan, err = r.client.KeepAlive(context.Background(), r.leaseID)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case keepAliveResp := <-r.keepAliveChan:
				if keepAliveResp == nil {
					return
				}
			}
		}
	}()

	return nil
}

func (r *ServiceRegistry) Deregister() error {
	if r.client == nil || r.leaseID == 0 {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.client.Revoke(ctx, r.leaseID)
	if err != nil {
		return fmt.Errorf("failed to revoke lease for %s: %v", r.key, err)
	}

	log.Printf("Successfully deregistered service '%s' from etcd", r.key)
	return r.client.Close()
}
