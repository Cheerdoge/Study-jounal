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

// NewServiceRegistry 初始化 etcd 客户端
func NewServiceRegistry(endpoints []string) (*ServiceRegistry, error) {
	// 1. 配置并创建 etcd 客户端
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,       // etcd 服务器地址列表，如 ["127.0.0.1:2379"]
		DialTimeout: 5 * time.Second, // 连接超时时间
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to etcd: %v", err)
	}

	// 2. 实例化并返回自定义的 ServiceRegistry
	return &ServiceRegistry{
		client: client,
	}, nil
}

// Register 注册服务并维持心跳
func (r *ServiceRegistry) Register(serviceName, serviceAddr string, ttl int64) error {
	r.key = fmt.Sprintf("/services/%s/%s", serviceName, serviceAddr)
	r.value = `{"Addr":"` + serviceAddr + `"}`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. 向 etcd 申请一个租约 (Lease)
	leaseResp, err := r.client.Grant(ctx, ttl)
	if err != nil {
		return fmt.Errorf("failed to grant lease: %v", err)
	}
	r.leaseID = leaseResp.ID

	// 2. 将服务的 Key-Value 注册到 etcd 中，并绑定刚刚申请的租约
	_, err = r.client.Put(ctx, r.key, r.value, clientv3.WithLease(r.leaseID))
	if err != nil {
		return fmt.Errorf("failed to put key-value to etcd: %v", err)
	}

	// 3. 设置自动续租 (KeepAlive)
	// etcd 会在后台自动发送心跳来刷新这个租约
	r.keepAliveChan, err = r.client.KeepAlive(context.Background(), r.leaseID)
	if err != nil {
		return fmt.Errorf("failed to keep alive: %v", err)
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

	log.Printf("Successfully registered service '%s' with address '%s' to etcd", serviceName, serviceAddr)

	return nil
}

// Deregister 注销服务
func (r *ServiceRegistry) Deregister() error {
	if r.client == nil || r.leaseID == 0 {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 撤销租约，etcd 会自动删除与该租约绑定的所有 Key
	_, err := r.client.Revoke(ctx, r.leaseID)
	if err != nil {
		return fmt.Errorf("failed to revoke lease for %s: %v", r.key, err)
	}

	log.Printf("Successfully deregistered service '%s' from etcd", r.key)
	return r.client.Close()
}
