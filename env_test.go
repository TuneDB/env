package env

import (
	"testing"
	"time"
	"fmt"
	"runtime"
	"path/filepath"
	"reflect"
	"os"
)

func TestParseDuration(t *testing.T) {
	var cfg = struct {
		Duration time.Duration `env:"key=DURATION default=5s"`
	}{}

	if err := Process(&cfg); err != nil {
		t.Fatal(err)
	}

	if cfg.Duration != time.Second*5 {
		t.Fatalf("%v != %f ", cfg.Duration, time.Second*5)
	}
}

// Test Slice parsing with YAML value
func TestYamlList(t *testing.T) {
	os.Setenv("LIST", "[a,b,c]")

	var cfg struct {
		List        []string `env:"key=LIST decode=yaml default=[]"`
	}

	err := Process(&cfg);

	ok(t, err)
	assert(t, cfg.List != nil, "List is nil")
	assert(t, len(cfg.List) == 3, "List is not length 3")
	assert(t, cfg.List[0] == "a", "List[0] is not a")
	assert(t, cfg.List[1] == "b", "List[1] is not b")
	assert(t, cfg.List[2] == "c", "List[2] is not c")
}

// Test Map parsing with YAML value
func TestYamlMap(t *testing.T) {
	os.Setenv("MAP", "{a: A, b: B}")

	var cfg struct {
		Map         map[string]string `env:"key=MAP decode=yaml default={}"`
	}

	err := Process(&cfg);

	ok(t, err)
	assert(t, cfg.Map != nil, "Map is nil")
	assert(t, len(cfg.Map) == 2, "Map is not length 2")
	assert(t, cfg.Map["a"] == "A", "Map[a] is not A")
	assert(t, cfg.Map["b"] == "B", "Map[b] is not B")
}

// Test MapSlice parsing with YAML value
func TestYamlMapList(t *testing.T) {
	os.Setenv("MAP_LIST", "{a: [A1, A2], b: [B1, B2]}")

	var cfg struct {
		MapList     map[string][]string `env:"key=MAP_LIST decode=yaml default={}"`
	}

	err := Process(&cfg);

	ok(t, err)
	assert(t, cfg.MapList != nil, "MapList is nil")
	assert(t, len(cfg.MapList) == 2, "MapList is not length 2")
	assert(t, cfg.MapList["a"][0] == "A1", "MapList[a][0] is not A1")
	assert(t, cfg.MapList["a"][1] == "A2", "MapList[a][1] is not A2")
	assert(t, cfg.MapList["b"][0] == "B1", "MapList[b][0] is not B1")
	assert(t, cfg.MapList["b"][1] == "B2", "MapList[b][1] is not B2")
}

// Test Slice parsing with JSON value
func TestJsonList(t *testing.T) {
	os.Setenv("LIST", `["a","b","c"]`)

	var cfg struct {
		List        []string `env:"key=LIST decode=yaml default=[]"`
	}

	err := Process(&cfg);

	ok(t, err)
	assert(t, cfg.List != nil, "List is nil")
	assert(t, len(cfg.List) == 3, "List is not length 3")
	assert(t, cfg.List[0] == "a", "List[0] is not a")
	assert(t, cfg.List[1] == "b", "List[1] is not b")
	assert(t, cfg.List[2] == "c", "List[2] is not c")
}

// Test Map parsing with JSON value
func TestJsonMap(t *testing.T) {
	os.Setenv("MAP", `{"a": "A", "b": "B"}`)

	var cfg struct {
		Map         map[string]string `env:"key=MAP decode=yaml default={}"`
	}

	err := Process(&cfg);

	ok(t, err)
	assert(t, cfg.Map != nil, "Map is nil")
	assert(t, len(cfg.Map) == 2, "Map is not length 2")
	assert(t, cfg.Map["a"] == "A", "Map[a] is not A")
	assert(t, cfg.Map["b"] == "B", "Map[b] is not B")
}
// Test MapSlice parsing with JSON
func TestJsonMapList(t *testing.T) {
	os.Setenv("MAP_LIST", `{"a": ["A1", "A2"], "b": ["B1", "B2"]}`)

	var cfg struct {
		MapList     map[string][]string `env:"key=MAP_LIST decode=yaml default={}"`
	}

	err := Process(&cfg);

	ok(t, err)
	assert(t, cfg.MapList != nil, "MapList is nil")
	assert(t, len(cfg.MapList) == 2, "MapList is not length 2")
	assert(t, cfg.MapList["a"][0] == "A1", "MapList[a][0] is not A1")
	assert(t, cfg.MapList["a"][1] == "A2", "MapList[a][1] is not A2")
	assert(t, cfg.MapList["b"][0] == "B1", "MapList[b][0] is not B1")
	assert(t, cfg.MapList["b"][1] == "B2", "MapList[b][1] is not B2")
}

// Test Slice parsing with no decode specified (default)
// Should accept YAML
func TestDefaultList(t *testing.T) {
	os.Setenv("LIST", "[a,b,c]")

	var cfg struct {
		List        []string `env:"key=LIST default=[]"`
	}

	err := Process(&cfg);

	ok(t, err)
	assert(t, cfg.List != nil, "List is nil")
	assert(t, len(cfg.List) == 3, "List is not length 3")
	assert(t, cfg.List[0] == "a", "List[0] is not a")
	assert(t, cfg.List[1] == "b", "List[1] is not b")
	assert(t, cfg.List[2] == "c", "List[2] is not c")
}

// Test Map parsing with no decode specified (default)
// Should accept YAML
func TestDefaultMap(t *testing.T) {
	os.Setenv("MAP", "{a: A, b: B}")

	var cfg struct {
		Map         map[string]string `env:"key=MAP default={}"`
	}

	err := Process(&cfg);

	ok(t, err)
	assert(t, cfg.Map != nil, "Map is nil")
	assert(t, len(cfg.Map) == 2, "Map is not length 2")
	assert(t, cfg.Map["a"] == "A", "Map[a] is not A")
	assert(t, cfg.Map["b"] == "B", "Map[b] is not B")
}

// Test MapSlice parsing with no decode specified (default)
// Should accept YAML
func TestDefaultMapList(t *testing.T) {
	os.Setenv("MAP_LIST", "{a: [A1, A2], b: [B1, B2]}")

	var cfg struct {
		MapList     map[string][]string `env:"key=MAP_LIST default={}"`
	}

	err := Process(&cfg);

	ok(t, err)
	assert(t, cfg.MapList != nil, "MapList is nil")
	assert(t, len(cfg.MapList) == 2, "MapList is not length 2")
	assert(t, cfg.MapList["a"][0] == "A1", "MapList[a][0] is not A1")
	assert(t, cfg.MapList["a"][1] == "A2", "MapList[a][1] is not A2")
	assert(t, cfg.MapList["b"][0] == "B1", "MapList[b][0] is not B1")
	assert(t, cfg.MapList["b"][1] == "B2", "MapList[b][1] is not B2")
}

// Add some simple funcs to allow quick testing
// copied from: https://github.com/benbjohnson/testing
// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
