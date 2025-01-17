package config

import (
	"io/ioutil"
	"testing"
)

func TestUserConfig(t *testing.T) {
	userCfg := &userGitConfig{
		dashboards:   make(map[string][]*UserConfigFile),
		monitors:     make(map[string][]*UserConfigFile),
		screenboards: make(map[string][]*UserConfigFile),
		downtimes:    make(map[string][]*UserConfigFile),

		pullMasterFn: func() error { return nil },
		readDirFn:    ioutil.ReadDir,
		readFileFn:   ioutil.ReadFile,

		basePath: "./fixtures/configs",
	}

	err := userCfg.Reload()
	if err != nil {
		t.Fatal(err)
	}

	if len(userCfg.dashboards) != 6 {
		t.Fatalf("expect 6 dashboards. Got %d", len(userCfg.dashboards))
	}

	for _, expectedID := range []string{"1", "2", "955878", "917832", "10", "20"} {
		if _, ok := userCfg.dashboards[expectedID]; !ok {
			t.Fatalf("expect dashboard id %s", expectedID)
		}
	}

	if len(userCfg.monitors) != 6 {
		t.Fatalf("expect 6 monitors. Got %d", len(userCfg.monitors))
	}

	for _, expectedID := range []string{"3", "4", "6065878", "4891392", "30", "40"} {
		if _, ok := userCfg.monitors[expectedID]; !ok {
			t.Fatalf("expect monitor id %s", expectedID)
		}
	}

	if len(userCfg.screenboards) != 2 {
		t.Fatalf("expect 2 screenboards. Got %d", len(userCfg.screenboards))
	}

	for _, expectedID := range []string{"42", "43"} {
		if _, ok := userCfg.screenboards[expectedID]; !ok {
			t.Fatalf("expect screenboard id %s", expectedID)
		}
	}

	if len(userCfg.downtimes) != 2 {
		t.Fatalf("expect 2 downtimes. Got %d", len(userCfg.downtimes))
	}

	for _, expectedID := range []string{"55", "66"} {
		if _, ok := userCfg.downtimes[expectedID]; !ok {
			t.Fatalf("expect downtime id %s", expectedID)
		}
	}

	// test reload, it should clear the 6 dashboards and load just 2
	userCfg.basePath = "./fixtures/configs/a/1/"
	err = userCfg.Reload()
	if err != nil {
		t.Fatal(err)
	}

	if len(userCfg.dashboards) != 2 {
		t.Fatalf("expect 2 dashboards. Got %d", len(userCfg.dashboards))
	}

	for _, expectedID := range []string{"1", "2"} {
		if _, ok := userCfg.dashboards[expectedID]; !ok {
			t.Fatalf("expect dashboard id %s", expectedID)
		}
	}

}
