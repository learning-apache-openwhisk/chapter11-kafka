OW_USER?=openwhisk
OW_RUNTIME?=$(OW_USER)/actionloop
OW_COMPILER?=$(OW_USER)/actionloop-golang-v1.11
WSK?=wsk
MAIN=main
PACKAGE=mesg
SRCS=src/$(MAIN).go src/app/assets.go src/app/box.go
VENDORS=src/app/vendor
NAME=chat
BINZIP=$(MAIN)-bin.zip
SRCZIP=$(MAIN)-src.zip

deploy: package.done $(BINZIP)
	$(WSK) action update $(PACKAGE)/$(NAME) $(BINZIP) --main $(MAIN) --kind go:1.11 --web true
	$(WSK) action get $(PACKAGE)/$(NAME) --url

src/%/vendor:
	cd $(@D) ; GOPATH=$$PWD/../.. dep ensure

$(BINZIP): $(SRCS) $(VENDORS)
	cd src ; packr ; zip - -r * | docker run -i $(OW_COMPILER) -compile $(MAIN) >../$(BINZIP)

clean:
	-wsk action delete $(PACKAGE)/$(NAME)
	-rm src/app/*-packr.go
	-rm  $(BINZIP) package.done

package.done:
	$(WSK) package update $(PACKAGE)
	touch package.done

.PHONY: deploy devel test clean
