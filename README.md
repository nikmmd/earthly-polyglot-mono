# All
```
#Build dists
earthly +all-build 

#Build docker
earthly +all-docker

#Unit test
earthly +all-unit

```


# Features

| Lang      | +build | +docker | +lint| +unit-test| +integration-test | +release| +dev| +docker-scan |
| ----------- | ----------- | ---- | ---| ---| --- | -- |--| -- |
| Go | :white_check_mark:  | :white_check_mark: | |  :white_check_mark:|
| JS| :white_check_mark:   | | | :white_check_mark: |
| Typescript| :white_check_mark: | | | :white_check_mark: |
| Python|:white_check_mark:|:white_check_mark:||:white_check_mark:|||:white_check_mark:|


# Stack
|+dev-docker|+dev-kuberentes|
|---|---|
