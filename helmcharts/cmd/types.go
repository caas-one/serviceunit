package cmd

import (
	mlog "github.com/maxwell92/log"
)

var log = mlog.Log

// build cmd
var profile string
var root string

// install cmd
var charts string
var installAll bool
var installSu string
var installApp string

// setup cmd
var outputProfile string
var srcNamespace string
var dstNamespace string

// update cmd
var updateAll bool
var updateSu string
var updateApp string

// clean cmd
var output string

//delete cmd with --purge
var deleteAll bool
var deleteSu string
var deleteApp string
