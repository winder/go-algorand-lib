UNAME := $(shell uname)
ifneq (,$(findstring MINGW,$(UNAME)))
#Gopath is not saved across sessions, probably existing Windows env vars, override them
export GOPATH := $(HOME)/go
GOPATH1 := $(GOPATH)
export PATH := $(PATH):$(GOPATH)/bin
else
export GOPATH := $(shell go env GOPATH)
GOPATH1 := $(firstword $(subst :, ,$(GOPATH)))
endif

MSGP_GENERATE := ./basics ./crypto

msgp: $(patsubst %,%/msgp_gen.go,$(MSGP_GENERATE))

%/msgp_gen.go: ALWAYS
		@set +e; \
		printf "msgp: $(@D)..."; \
		$(GOPATH1)/bin/msgp -file ./$(@D) -o $@ -warnmask github.com/algorand/go-algorand > ./$@.out 2>&1; \
		if [ "$$?" != "0" ]; then \
			printf "failed:\n$(GOPATH1)/bin/msgp -file ./$(@D) -o $@ -warnmask github.com/algorand/go-algorand\n"; \
			cat ./$@.out; \
			rm ./$@.out; \
			exit 1; \
		else \
			echo " done."; \
		fi; \
		rm -f ./$@.out
ALWAYS:
