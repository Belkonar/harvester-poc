package core

import (
	"context"
	"time"

	etcd "go.etcd.io/etcd/client/v3"
)

func WriteKV(key string, val string) error {
	cli, err := etcd.New(etcd.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		return err
	}

	_, err = cli.KV.Put(context.Background(), key, val)

	if err != nil {
		return err
	}

	defer cli.Close()

	return nil
}
