![Geo Smart Logo](http://supanadit.com/wp-content/uploads/2019/11/Geo-Smart-Logo.png)

# GEO Smart System
This is Tile38 Implementation for Golang, and also this software has a purpose to be real time tracking system 
simulation such as Uber, Gojek, Grab, etc. The main feature of this software is that it must **lightweight**, 
**less memory usage**, and **fast**, and for the live map it will integration with [Geo Smart Map](https://github.com/supanadit/geosmartmap) and [Geo Smart App](https://github.com/supanadit/geosmartapp)
![Workflow](http://supanadit.com/wp-content/uploads/2019/11/geosmart-work.png)

[![Go Report Card](https://goreportcard.com/badge/github.com/supanadit/geo-smart-system)](https://goreportcard.com/report/github.com/supanadit/geo-smart-system)

## Requirements
- [Tile38 Server](https://tile38.com/)
- [Golang](https://golang.org/)

## Todo
- Change to SSE from Socket IO (OK)
- Connect With Tile38 (OK)
- Get Data From Tile38 by Command SCAN (OK)
- Receive New Point using SSE (OK)
- Send Realtime Point using POST Method (OK)
- Set HOOK by GeoFencing Trigger ( OK )
- Support Nearby Trigger Feature ( OK )
- Support Enter Area Trigger Feature ( OK )
- Support Exit Area Trigger Feature ( OK )
- [Documentation](https://github.com/supanadit/geosmartdocumentation) ( In Progress )

## Notes

This project will always maintained, but the problem is that it will slowly developed because the only contributor just my self, this project can be customize for any model of tracking system, and it's open for any contributor who really want to help this project

## License
Copyright 2019 Supan Adit Pratama

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
