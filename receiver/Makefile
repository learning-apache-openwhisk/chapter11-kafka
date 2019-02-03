#IMG=openwhisk/actionloop-golang-v1.11
IMG=actionloop/golang-v1.11:librdkafka

.PHONY: url

url: receiver.done 
	wsk action get mesg/receiver --url

receiver.done: receiver.zip
	wsk action update mesg/receiver $< \
	--docker $(IMG) --web true \
	-P ../cred.json -p secret s3cr3t \
	-p topic queue -p partition 0 && \
	touch receiver.done

receiver.zip: receiver_src.zip
	docker run -i $(IMG) -compile main <$< >$@

receiver_src.zip: src/main/receiver.go src/main/receive.go | src/main/vendor
	cd src ; zip -r ../$@ main -x main/main.go

src/main/vendor:
	export GOPATH="$(shell pwd)" ; cd src/main ;  dep ensure

clean:
	-wsk action delete mesg/receiver
	-rm -Rvf src/main/vendor
	-rm *.zip *.done
