package gateway

import (
	"github.com/zgwit/iot-master/v4/connect"
	"github.com/zgwit/iot-master/v4/log"
	"time"
)

type Keeper struct {
	Tunnel connect.Tunnel
}

func (k *Keeper) Keep() {
	go func() {
		for {
			//10秒自动重连
			time.Sleep(time.Second * 10)

			if k.Tunnel.Running() {
				continue
			}
			if k.Tunnel.Closed() {
				break
			}

			err := k.Tunnel.Open()
			if err != nil {
				log.Error(err)
			}
		}
	}()
}

func Keep(tunnel connect.Tunnel) *Keeper {
	keeper := &Keeper{Tunnel: tunnel}
	keeper.Keep()
	return keeper
}

//
//func (k *Keeper) Retry() {
//	//重连
//	if k.RetryMaximum == 0 || k.retried < k.RetryMaximum {
//		k.retried++
//
//		timeout := k.RetryTimeout
//		if timeout == 0 {
//			timeout = 10 //默认10秒
//		}
//
//		k.retryTimer = time.AfterFunc(time.Second*time.Duration(timeout), func() {
//			k.retryTimer = nil
//			err := k.Start()
//			if err != nil {
//				log.Error(err)
//			}
//		})
//	}
//}
