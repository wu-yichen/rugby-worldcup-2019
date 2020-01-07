# wcwhen

A Golang-based CLI for retrieving information of Rugby world cup 2019
![wcwhen](https://keyassets.timeincuk.net/inspirewp/live/wp-content/uploads/sites/7/2019/09/Wallchart-630x425.jpg)

## Usage

```
wcwhen --help
```

Supported commands:

```
COMMANDS:
     team               team -name <teamName>
     group              group -name <groupName>
```

## Docker

- build image

```
     docker build -t {image:tag} -f ci/image/Dockerfile .
```

- run container

```
     docker run --rm -it {image:tag} arg1 arg2...
```

## pipeline

### this application uses concourse pipeline

- to start concourse pipeline service

  > `./setup-pipeline.sh`

- to set up pipeline for wcwhen
  > `fly -t {target} set-pipeline -c ci/pipeline.yml -p wcwhen`
