VERSION 0.7

all-build:
    WAIT
        BUILD ./service/go-svc1+build
        BUILD ./libs/addr+build
        BUILD ./web/svelte-app+build
        BUILD ./service/py_poetry_svc1+build
    END
    #Example remote build
    #BUILD  github.com/earthly/earthly/examples/go+build

all-unit-test:
    WAIT
        BUILD ./service/go-svc1+unit-test
        BUILD ./libs/addr+unit-test
        BUILD ./web/svelte-app+unit-test
        BUILD ./service/py_poetry_svc1+unit-test
    END



all-docker:
    WAIT
        BUILD ./service/go-svc1+docker
        BUILD ./service/py_poetry_svc1+docker
    END
    # WAIT
    #     DO +SCAN_WITH_TRIVY --image_name="nikmmd/go_svc1:latest"
    #     DO +SCAN_WITH_TRIVY --image_name="nikmmd/py_poetry_svc1:latest"
    # END


SCAN_WITH_TRIVY:
    COMMAND
    LOCALLY
    ARG image_name
    RUN docker run -v /tmp/trivy-cache:/root/.cache/ -v /var/run/docker.sock:/var/run/docker.sock aquasec/trivy:0.44.1 image --exit-code 1 --severity CRITICAL ${image_name}


# We bootstrap the npm workspace deps
# We track all js monorepo integrations here
# Subpackages exntend this
npm-workspace-deps:
    FROM node:20-alpine3.17
    WORKDIR /monorepo
    COPY package.json package-lock.json tsconfig.json jest.config.js ./
    COPY libs ./libs
    COPY web ./web
    RUN npm ci
    SAVE ARTIFACT package-lock.json AS LOCAL package-lock.json
    SAVE ARTIFACT package.json AS LOCAL package.json
    SAVE ARTIFACT node_modules