test:
	@echo 'Makefile ongoing'

run:
	docker-compose up --build project
stop:
	docker-compose stop