#!/bin/bash

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

source ${DIR}/test_helper.sh

start_container

poll_log 'JBoss AS 7.2.0.Final "Janus" started in' 'JBoss AS 7.2.0.Final "Janus" started (with errors) in' || _exit 1 "Server not started"

_log "Server started"

grep_log 'Deployed "camunda-example-invoice-7.6.0' || _exit 2 "Process application not deployed"

_log "Process application deployed"

test_login admin || _exit 3 "Unable to login to admin"
test_login cockpit || _exit 4 "Unable to login to cockpit"
test_login tasklist || _exit 5 "Unable to login to tasklist"

_log "Login successfull"

test_encoding || _exit 6 "Wrong encoding detected"

_exit 0 "Test successfull"
