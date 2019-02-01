deploy:
	$(MAKE) -C sender
	$(MAKE) -C receiver
	$(MAKE) -C chat

clean:
	$(MAKE) clean -C sender
	$(MAKE) clean -C receiver
	$(MAKE) clean -C chat
