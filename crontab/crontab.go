package crontab

import (
	"fmt"
	"log"
	"sync"
	"errors"
	cron "github.com/robfig/cron/v3"
)

// Crontab crontab manager
type Crontab struct {
	inner *cron.Cron
	ids   map[string]cron.EntryID
	mutex sync.Mutex
}

// NewCrontab new crontab
func NewCrontab() *Crontab {
	return &Crontab{
		inner: cron.New(cron.WithSeconds()),
		ids:   make(map[string]cron.EntryID),
	}
}

//// NewCrontab new crontab
//func  (c *Crontab) New() *Crontab {
//	return &Crontab{
//		inner: cron.New(cron.WithSeconds()),
//		ids:   make(map[string]cron.EntryID),
//	}
//}


// IDs ...
func (c *Crontab) IDs() []string {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	validIDs := make([]string, 0, len(c.ids))
	invalidIDs := make([]string, 0)
	for sid, eid := range c.ids {
		if e := c.inner.Entry(eid); e.ID != eid {
			invalidIDs = append(invalidIDs, sid)
			continue
		}
		validIDs = append(validIDs, sid)
	}
	for _, id := range invalidIDs {
		delete(c.ids, id)
	}
	return validIDs
}

// Start start the crontab engine
func (c *Crontab) Start() {
	c.inner.Start()
}

// Stop stop the crontab engine
func (c *Crontab) Stop() {
	c.inner.Stop()
}

// DelByID remove one crontab task
func (c *Crontab) DelByID(id string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	fmt.Println("id:", id)
	eid, ok := c.ids[id]
	if !ok {
		return
	}
	c.inner.Remove(eid)
	delete(c.ids, id)
}

// AddByID add one crontab task
// id is unique
// spec is the crontab expression
func (c *Crontab) AddByID(id string, spec string, cmd cron.Job) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.ids[id]; ok {
		return errors.New("crontab id exists")
	}
	eid, err := c.inner.AddJob(spec, cmd)
	if err != nil {
		return err
	}
	log.Println("AddByID eid:", eid)

	c.ids[id] = eid
	return nil
}

// AddByFunc add function as crontab task
func (c *Crontab) AddByFunc(id string, spec string, f func()) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.ids[id]; ok {
		return errors.New("crontab id exists")
	}

	eid, err := c.inner.AddFunc(spec, f)
	if err != nil {
		return err
	}

	c.ids[id] = eid
	log.Println("AddByFunc eid:", c.ids)
	return nil
}

// IsExists check the crontab task whether existed with job id
func (c *Crontab) IsExists(jid string) bool {
	_, exist := c.ids[jid]
	return exist
}


//secondStr := "*/30 * * * * ?"
//crontab := crontab.NewCrontab()
// 实现接口的方式添加定时任务
//task := &testTask{}

//type testTask struct{}

//func (t *testTask) Run() {
//	log.Println("hello world")
//}

//if err := crontab.AddByID("111-222", "*/5 * * * * ?", task); err != nil {
//	fmt.Printf("error to add crontab task:%s", err)
//	ch <- 111
//}

//// 添加函数作为定时任务
//taskFunc := func() {
//	log.Println("hello add")
//}
//
//if err := crontab.AddByFunc("333-222", "*/12 * * * * ?", taskFunc); err != nil {
//	log.Println("error to add crontab task:", err)
//	ch <- "333-222"
//}

//crontab.AddByFunc("36015d73-d192-4f2e-a69d-423e8076482e", "*/12 * * * * ?", func() {
//	log.Println("error to add crontab task")
//	if bl {
//		crontab.DelByID("36015d73-d192-4f2e-a69d-423e8076482e")
//	}
//})

//crontab.Start()
//fmt.Println(crontab.IDs())
//select {}