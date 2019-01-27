.PHONY: deploy send poll
MSG?=hello

deploy: trigger.done
send:
	wsk action invoke demoq/messageHubProduce \
	-P ../cred.json -p topic queue -p value "$(MSG)"

poll:
	wsk activation poll mesg/logme

bind.done:
	wsk package bind /whisk.system/messaging demoq -P ../cred.json
	wsk package update mesg
	touch bind.done

logme.done: logme.go bind.done
	wsk action update mesg/logme logme.go
	touch logme.done

trigger.done: bind.done logme.done
	wsk trigger create messages-trigger \
	-f demoq/messageHubFeed \
	-p topic queue -p isJSONData true
	wsk rule update messages-trigger-logme messages-trigger mesg/logme
	wsk rule enable messages-trigger-logme





clean:
	-wsk rule delete messages-queue-echo
	-wsk trigger delete messages-trigger
	-wsk package delete demoq
	-rm *.done
