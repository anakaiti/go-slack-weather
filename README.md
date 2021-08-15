# go-slack-weather
[![go-slack-weather](https://github.com/fadhilthomas/go-slack-weather/actions/workflows/go-slack-weather.yml/badge.svg?branch=main)](https://github.com/fadhilthomas/go-slack-weather/actions/workflows/go-slack-weather.yml)
[![license](https://img.shields.io/badge/license-MIT-_red.svg)](https://opensource.org/licenses/MIT)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/fadhilthomas/go-slack-weather/issues)

Send current weather updates from OpenWeatherMap API to your Slack profile status using GitHub Action

![slack](https://user-images.githubusercontent.com/29804796/129493231-2de98bac-09ac-4686-a97d-487a344dd6a1.png)

---

## Resources

- [Setup](#setup)
- [Help & Bugs](#help--bugs)
- [License](#license)

---

## Setup
* Fork this repository.
* Set these environment variable in GitHub repository secrets.

| **Variable** | **Value** |
|--|--|
| WEATHER_API | Get OpenWeatherMap API [here](https://home.openweathermap.org/users/sign_up). |
| CITY | Find your city id [here](http://bulk.openweathermap.org/sample/). |
| TIMEZONE | Find your timezone city [here](https://www.iana.org/time-zones). |
| SLACK_TOKEN | Your Slack User OAuth Token. |


## Help & Bugs

If you are still confused or found a bug, please [open the issue](https://github.com/fadhilthomas/go-slack-weather/issues). All bug reports are appreciated, some features have not been tested yet due to lack of free time.

## License

**go-slack-weather** released under MIT. See `LICENSE` for more details.