.DEFAULT_GOAL = format

ifdef PATHS
go-fmt:
	$(AT) if command -v goimports >/dev/null; then \
		goimports -w $(PATHS); \
	else \
		gofmt -s -w $(PATHS); \
	fi
else
go-fmt: ;
endif
.PHONY: go-fmt

format: go-fmt
.PHONY: format
