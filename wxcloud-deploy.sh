
wxcloud run:deploy \
    --targetDir=. \
    --containerPort=8100 \
    --envId=prod-2goqrtua754c85a7 \
    --serviceName=golang-kk00-001 \
    --remark=cli-upload \
    --dockerfile=DockerfileTest \
    --noConfirm