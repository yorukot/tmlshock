## About tmlshock
Terminal ttl clock, including customizable clock and stopwatch

## Features

![feature](https://github.com/MHNightCat/tmlshock/blob/main/img/feature.png)

## Contents
* [Install](#install)
* [Usage](#Usage)
* [Setting](#setting)
  * [Color](#color)
    * [Colon color](#colon-color)
  * [Disable second](#disable-second)
  * [Enable date](#enable-date)
  * [Date formate](#date-formate)
  * [Disable hour](#disable-hour)
  * [12 hours format](#12-hours-format)
## Install

see [release](https://github.com/MHNightCat/tmlshock/releases)

download tmlshock file and move it to /usr/local/bin

```bash
sudo mv ./tmlshock /usr/local/bin
```

## Usage

### clock
```bash
tmlshock clock
```
### stopwatch
```bash
tmlshock stopwatch
```

### timer
You can set the timer time using two types

classic:
```bash
tmlshock timer -hr 1 -m 1 -s 1
```

lazy:
```bash
tmlshock timer -t 1:20:01
```

## Setting

## **Color**

You can use color codes (down below) or use color names

![color](https://github.com/MHNightCat/tmlshock/blob/main/img/color.png)

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
```bash
tmlshock clock -c red
```
![red-clock](https://github.com/MHNightCat/tmlshock/blob/main/img/red-clock.png)

### **Colon-color**

To use a custom colon color just enter `-cc color`

**Example**
```bash
tmlshock clock -cc color
```

![custom-colon-color](https://github.com/MHNightCat/tmlshock/blob/main/img/custom-colon-color.png)

## **Disable-second**

To disable the second just enter `-s false`

**Example**
```bash
tmlshock clock -s false
```

![no-sec-clock](https://github.com/MHNightCat/tmlshock/blob/main/img/no-sec-clock.png)

## **Enable-date**

To enable the date just enter `-d true`

**Example**
```bash
tmlshock clock -d true
```

![date-clock](https://github.com/MHNightCat/tmlshock/blob/main/img/date-clock.png)

## **Date-formate**

To use a custom date format just enter `-df 2006/02/01`(YYYY/MM/DD)

**Example**
```bash
tmlshock clock -d true -df 2006/01/02 
```
(YYYY/DD/MM)

![date-format-clock](https://github.com/MHNightCat/tmlshock/blob/main/img/date-format-clock.png)

```bash
tmlshock clock -d true -df 02/01/2006
```
(MM/DD/YYYY)

![date-format-2-clock](https://github.com/MHNightCat/tmlshock/blob/main/img/date-format-2-clock.png)

## **Disable-hour**

* This option only for stopwatch and timer

To use a custom date format just enter `-dh true`(YYYY/MM/DD)

**Example**
```bash
tmlshock stopwatch -dh true
```

![disable-hour-stopwatch](https://github.com/MHNightCat/tmlshock/blob/main/img/disable-hour-stopwatch.png)

## **12-hours-format**

* This option only for clock

To use a custom date format just enter `-hf 12`(YYYY/MM/DD)

**Example**
```bash
tmlshock clock -hf 12
```
