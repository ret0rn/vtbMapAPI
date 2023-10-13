prod:
	export ENVIRONMENT=prod && \
	docker compose up --build $(DOKERFLAGS)

develop:
	export ENVIRONMENT=develop && \
	docker compose up --build $(DOKERFLAGS)

stop:
	docker compose down

.DEFAULT_GOAL := develop