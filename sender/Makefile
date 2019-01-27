IMG=actionloop/golang-v1.11:ws4
MSG?=hello

.PHONY: url send

url: sender.done
	wsk action get mesg/sender --url

send:
	curl "$(shell wsk action get mesg/sender --url | tail -1)?message=$(MSG)"

sender.done: sender.zip
	wsk action update mesg/sender $< \
	--docker $(IMG) --web true \
	-P ../cred.json -p topic queue && \
	touch sender.done


sender.zip: sender_src.zip
	docker run -i $(IMG) -compile main <$< >$@

sender_src.zip: | src/klient/vendor
	cd src ; zip -r ../$@ sender.go klient

src/klient/vendor:
	export GOPATH="$(shell pwd)" ; cd src/klient ;  dep ensure

clean:
	-wsk delete mesg/sender
	-rm *.zip *.done
	-rm -Rvf src/ksend/vendor

#image.done:
#	docker login
#	docker build . -t $(IMG)
#	docker push $(IMG) && touch image.done

