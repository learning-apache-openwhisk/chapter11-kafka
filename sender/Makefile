#IMG=openwhisk/actionloop-golang-v1.11
IMG=actionloop/golang-v1.11:librdkafka
MSG?=hello

.PHONY: url send

url: sender.done
	wsk action get mesg/sender --url

send:
	curl "$(shell wsk action get mesg/sender --url | tail -1)?message=$(MSG)"

sender.done: sender.zip
	wsk action update mesg/sender $< \
	--docker $(IMG) --web true \
	-P ../cred.json -p topic queue -p partition 0 && \
	touch sender.done

sender.zip: sender_src.zip
	docker run -i $(IMG) -compile main <$< >$@

sender_src.zip: src/main/sender.go src/main/send.go | src/main/vendor
	cd src ; zip -r  ../$@ main -x main/main.go -x '*_test'

src/main/vendor:
	export GOPATH="$(shell pwd)" ; cd src/main ;  dep ensure

clean:
	-wsk action delete mesg/sender
	-rm *.zip *.done
	-rm -Rvf src/ksend/vendor

#image.done:
#	docker login
#	docker build . -t $(IMG)
#	docker push $(IMG) && touch image.done

