## About Tshock
Terminal ttl clock, including customizable clock and stopwatch

## Features

[![feature](https://github.com/MHNightCat/tshock/blob/main/img/feature.png)]

## Contents
* [Install](#install)
* [Usage](#Usage)
* [Setting](#setting)
  * [Color](#color)
  * [Disable second](#disable-second)
  * [Enable date](#disable-date)
  * [Date formate](#date-formate)

## Install

## Usage
run clock:
```sh
tshock clock
```
run stopwatch:
```sh
tshock stopwatch
```

## Setting

### Color

You can use color codes (down below) or use color names

![color](https://github.com/MHNightCat/tshock/blob/main/img/color.png)

```
black
blue
cyan
dark-gray
green
light-green
light-blue
light-cyan
light-gray
light-magenta
light-red
light-yellow
magenta
light-red
light-yellow
magenta
red
white
yellow
```

Example:
```sh
tshock clock -c red
```
![red-clock](https://github.com/MHNightCat/tshock/blob/main/img/red-clock.png)

### Disable-secnod

To disable the second just enter `-s false`

Example
```sh
tshock clock -s false
```

![no-sec-clock](https://github.com/MHNightCat/tshock/blob/main/img/no-sec-clock.png)

### Enable-date

To enable the date just enter `-d true`

Example
```sh
tshock clock -d true
```

![date-clock](https://github.com/MHNightCat/tshock/blob/main/img/date-clock.png)

### Date-formate

To use a custom date format just enter `-df 2006/02/01`(YYYY/MM/DD)

Example
```sh
tshock clock -d true -df 2006/01/02 
```
(YYYY/DD/MM)

![date-format-clock](https://github.com/MHNightCat/tshock/blob/main/img/date-format-clock.png)

```sh
tshock clock -d true -df 02/01/2006
```
(MM/DD/YYYY)

![date-format-2-clock](https://github.com/MHNightCat/tshock/blob/main/img/date-format-2-clock.png)



