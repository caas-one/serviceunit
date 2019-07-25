package service

import (
	"fmt"
	"strings"
	"time"

	mlog "github.com/maxwell92/log"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

var log = mlog.Log

// Controller will List/Watch the service events
type Controller struct {
	indexer  cache.Indexer
	queue    workqueue.RateLimitingInterface
	informer cache.Controller
}

// QueueItem defines the item cached by indexer
type QueueItem struct {
	Key       string
	Type      cache.DeltaType
	Object    interface{}
	OldObject interface{}
}

// NewController return a new Controller instance
func NewController(
	queue workqueue.RateLimitingInterface,
	indexer cache.Indexer,
	informer cache.Controller) *Controller {
	return &Controller{
		informer: informer,
		indexer:  indexer,
		queue:    queue,
	}
}

// Run func will start the controller loop
func (c *Controller) Run(threadiness int, stopCh chan struct{}) {
	defer func() { recover() }()
	defer runtime.HandleCrash()
	defer c.queue.ShutDown()

	go c.informer.Run(stopCh)

	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<-stopCh
}

func (c *Controller) runWorker() {
	for c.processNextItem() {
	}
}

func (c *Controller) handleErr(err error, key interface{}) {
	if err == nil {
		c.queue.Forget(key)
		return
	}

	if c.queue.NumRequeues(key) < 5 {
		c.queue.AddRateLimited(key)
		return
	}

	c.queue.Forget(key)
	log.Infof("Dropping pods out of the queue: key=%v, err=%s", key, err)

	runtime.HandleError(err)
}

func (c *Controller) process(item QueueItem) error {

	obj, exists, err := c.indexer.GetByKey(item.Key)
	if err != nil {
		log.Errorf("Fetching object with key %s from store failed with %v", item.Key, err)
		return err
	}

	if exists {
		if cache.Sync == item.Type {
			log.Infof("Service Sync: namespace=%s, name=%s", obj.(*v1.Pod).Namespace, obj.(*v1.Pod).Name)
		}
	} else {
		log.Errorf("Service not exist")
	}
	return nil
}

func (c *Controller) processNextItem() bool {
	item, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(item)

	// ToDo: process item
	c.process(item.(QueueItem))
	return true
}

func processAdd(svc *v1.Service) {
	service := transfer(svc)
	// log.Infof("Wanted to add a service: namespace=%s, name=%s, nodeport=%s",
	// 	service.Namespace, service.Name, service.Nodeport)
	if !strings.EqualFold("0", service.Nodeport) {
		Instance().Add(svc.Namespace, service)
	}
}

func processUpdate(oldSvc *v1.Service, newSvc *v1.Service) {
	service := transfer(newSvc)
	log.Infof("Wanted to update a service: namespace=%s, name=%s, nodeport=%s",
		service.Namespace, service.Name, service.Nodeport)
	if !strings.EqualFold("0", service.Nodeport) {
		Instance().Update(newSvc.Namespace, service)
	}
}

func processDelete(svc *v1.Service) {
	Instance().Del(svc.Namespace, svc.Name)
}

// Informer is the key component to implements the list/watch
type Informer struct {
	Indexer    cache.Indexer
	Informer   cache.Controller
	Controller *Controller
	Ch         chan struct{}
}

// NewInformer creates a new informer instance
func NewInformer(lw cache.ListerWatcher, ch chan struct{}) *Informer {
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	indexer, informer := cache.NewIndexerInformer(lw, &v1.Service{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(QueueItem{key, cache.Added, obj, nil})
				processAdd(obj.(*v1.Service))
				ch <- struct{}{}
			}
		},
		UpdateFunc: func(oldObj interface{}, newObj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			if err == nil {
				queue.Add(QueueItem{key, cache.Updated, newObj, oldObj})
				processUpdate(oldObj.(*v1.Service), newObj.(*v1.Service))
				ch <- struct{}{}
			}
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(QueueItem{key, cache.Deleted, obj, nil})
				processDelete(obj.(*v1.Service))
				ch <- struct{}{}
			}
		},
	}, cache.Indexers{})

	return &Informer{
		Informer:   informer,
		Indexer:    indexer,
		Controller: NewController(queue, indexer, informer),
	}
}

// Run is a controll-loop for Informer
func (i *Informer) Run() {
	defer func() { recover() }()
	log.Infof("Service informer start......")

	stopCh := make(chan struct{})
	defer close(stopCh)

	i.Controller.Run(1, stopCh)
}
