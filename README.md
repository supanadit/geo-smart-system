![Logo](https://i.ibb.co/80c1hWg/Logo.png)

# GEO Smart System

[![Build Status](https://travis-ci.com/supanadit/geo-smart-system.svg?branch=master)](https://travis-ci.com/supanadit/geo-smart-system)
[![Go Report Card](https://goreportcard.com/badge/github.com/supanadit/geo-smart-system)](https://goreportcard.com/report/github.com/supanadit/geo-smart-system)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/c184e682b4944bb2b68a665f47869cd3)](https://www.codacy.com/manual/supanadit/geo-smart-system?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=supanadit/geo-smart-system&amp;utm_campaign=Badge_Grade)

This is Tile38 Implementation for Golang, and also this software has a purpose to be real time tracking system 
simulation such as Uber, Gojek, Grab, etc. The main feature of this software is the real time tracking system also must be   **lightweight**, 
**less memory usage**, and **fast**, for the live map it has integrate with [Geo Smart Map](https://github.com/supanadit/geosmartmap) and [Geo Smart App](https://github.com/supanadit/geosmartapp)

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

This project will always be maintained, but the problem is, it will slowly maintained because my self is the only contributor, i want this project customizable for any model of tracking system, and i still finding contributor to help this project grow up and keep alive also straight forward following the mission.
