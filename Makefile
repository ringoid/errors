all: clean stage-deploy

build:
	@echo '--- Building timeout function ---'
	GOOS=linux go build lambda-timeout/timeout.go
	@echo '--- Building invalidtoken function ---'
	GOOS=linux go build lambda-invalidtoken/invalidtoken.go
	@echo '--- Building nonok function ---'
	GOOS=linux go build lambda-nonok/nonok.go
	@echo '--- Building old function ---'
	GOOS=linux go build lambda-toooldversion/old.go
	@echo '--- Building internalerror function ---'
	GOOS=linux go build lambda-internal-error/internalerror.go

zip_lambda: build
	@echo '--- Zip timeoutfunction ---'
	zip timeout.zip ./timeout
	@echo '--- Zip invalidtoken ---'
	zip invalidtoken.zip ./invalidtoken
	@echo '--- Zip nonok ---'
	zip nonok.zip ./nonok
	@echo '--- Zip old ---'
	zip old.zip ./old
	@echo '--- Zip internalerror ---'
	zip internalerror.zip ./internalerror

stage-deploy: zip_lambda
	@echo '--- Build lambda stage ---'
	@echo 'Package template'
	sam package --template-file errors-template.yaml --s3-bucket ringoid-cloudformation-template --output-template-file errors-template-packaged.yaml
	@echo 'Deploy errors-image-stack'
	sam deploy --template-file errors-template-packaged.yaml --s3-bucket ringoid-cloudformation-template --stack-name errors-image-stack --capabilities CAPABILITY_IAM --parameter-overrides Env=stage --no-fail-on-empty-changeset

clean:
	@echo '--- Delete old artifacts ---'
	rm -rf timeout
	rm -rf timeout.zip
	rm -rf invalidtoken
	rm -rf invalidtoken.zip
	rm -rf nonok
	rm -rf nonok.zip
	rm -rf old.zip
	rm -rf old
	rm -rf internalerror
	rm -rf internalerror.zip