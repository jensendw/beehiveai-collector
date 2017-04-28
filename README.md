#BeeHiveAI-Collector
Sends local system level data to the beehive AI custom integration API

## Configuration

```shell
BEEHIVEAI_CHECK_INTERVAL=300
BEEHIVEAI_TOKEN=xxxxx
BEEHIVEAI_MARATHON_URL=https://marathon
BEEHIVEAI_JENKINS_URL=https://jenkins
BEEHIVEAI_JENKINS_USERNAME=someusername
BEEHIVEAI_JENKINS_PASSWRD=somepassword
```

BEEHIVEAI_CHECK_INTERVAL
* How often information from the various services is gathered

BEEHIVEAI_TOKEN
* The token used for connecting to the BeeHiveAI API

BEEHIVEAI_MARATHON_URL (Optional)
* The url for the marathon server to collect the data from

BEEHIVEAI_JENKINS_URL (Optional)
* The url for the jenkins servers to collect the data from

BEEHIVEAI_JENKINS_USERNAME (Optional)
* Username to use when connecting to jenkins

BEEHIVEAI_JENKINS_PASSWORD (Optional)
* Passowrd to use when connecting to jenkins

The various check will activate based on if you specify URL's for the services (Marathon, Jenkins etc...)

If no jenkisn username or password is specified it is assumed that anonymous access is enabled to the jenkins servers

## Running

```shell
docker-compose up
```

## License
[Apache 2](http://www.apache.org/licenses/LICENSE-2.0)

## Contributing

1. Fork it ( https://github.com/jensendw/beehiveai-go )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
