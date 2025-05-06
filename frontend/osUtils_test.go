package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDir(t *testing.T) {
	assert.True(t, isDir("./fixtures"), "should return true for director path")
	assert.False(t, isDir("./invalid_xd"), "should return false for invalid directory")
	assert.False(t, isDir("./fixtures/markdown_test.md"), "should return false for valid files")
}

func TestValidPath(t *testing.T) {
	assert.True(t, validPath("./fixtures/markdown_test.md"), "should return true for correct path")
	assert.True(t, validPath("./fixtures"), "should return true for directories")
	assert.False(t, validPath("./fixtures/invalid_xd.txt"), "should return false for invalid files")
}

func TestGetFiles(t *testing.T) {
	validFiles := getFiles("./fixtures")
	assert.NotNil(t, validFiles, "should return valid pointer to array of files")
	assert.Equal(t, len(*validFiles), 4)
}

func TestGetFileNames(t *testing.T) {

}
