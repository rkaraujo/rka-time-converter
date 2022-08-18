# Time Converter

This was written in Go Lang. It converts time between timezones.

## Options

´´´
-h Shows help
-tz Timezone to convert to (default America/New_York). List of timezones: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
-time Time to convert in HH:mm format, ex.: 22:25 (default current time)
´´´

Example:

´´´
go run .\rka-time-converter.go -tz America/Los_Angeles -time 22:25
22:25 is 18:25 in America/Los_Angeles
´´´
