package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_watcher_watch(t *testing.T) {

	if err := initLogger(); err != nil {
		return
	}

	wd, _ := os.Getwd()
	directory := filepath.Join(wd, "testdata")
	file := filepath.Join(directory, "testdata.txt")

	err := os.MkdirAll(directory, os.ModePerm)
	assert.NoError(t, err)

	t.Run("Creating file event notified", func(t *testing.T) {
		var testWatcher = &grafanaWatcherMock{0, &grafanaAttributes{path: directory}}
		err = start(testWatcher)
		assert.NoError(t, err)

		_, err = os.Create(file)
		assert.NoError(t, err)
		time.Sleep(500 * time.Millisecond)
		assert.Equal(t, 1, testWatcher.eventCount)

		err = os.Remove(file)
		assert.NoError(t, err)
		time.Sleep(500 * time.Millisecond)
		assert.Equal(t, 1, testWatcher.eventCount)
		err = testWatcher.stop()
		assert.NoError(t, err)
	})

	t.Cleanup(func() {
		err := os.RemoveAll(directory)
		assert.NoError(t, err)
	})

}

type grafanaWatcherMock struct {
	eventCount int
	attr       *grafanaAttributes
}

func (g *grafanaWatcherMock) stop() error {
	return nil
}
func (g *grafanaWatcherMock) startGrafana() error {
	return nil
}
func (g *grafanaWatcherMock) attributes() *grafanaAttributes {
	return g.attr
}

func (g *grafanaWatcherMock) killProcess() error {
	g.eventCount++
	return nil
}
