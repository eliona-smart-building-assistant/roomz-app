# ROOMZ User Guide

### Introduction

> The ROOMZ app provides Eliona with data from ROOMZ sensors.

## Overview

This guide provides instructions on configuring, installing, and using the ROOMZ app to gather data from ROOMZ sensors.

## Installation

Install the ROOMZ app via the Eliona App Store.

## Configuration

The ROOMZ app requires configuration through Eliona’s settings interface.

### Registering the app in ROOMZ Portal

This app uses an experimental [Webhook API](https://github.com/roomz-io/openapi), which is not yet generally available, therefore ROOMZ support must be contacted to activate sending data. The webhook endpoint is located at `https://{your-eliona-instance}/apps-public/roomz/webhook`.

### Configure the ROOMZ app 

Configurations can be created in Eliona under `Apps > ROOMZ > Settings` which opens the app's [Generic Frontend](https://doc.eliona.io/collection/v/eliona-english/manuals/settings/apps). Here you can use the appropriate endpoint with the POST method. Each configuration requires the following data:

| Attribute         | Description                                                                     |
|-------------------|---------------------------------------------------------------------------------|
| `enable`          | Flag to enable or disable this configuration.                                   |
| `refreshInterval` | Interval in seconds for data synchronization.                                   |
| `requestTimeout`  | API query timeout in seconds.                                                   |
| `projectIDs`      | List of Eliona project IDs for data collection.                                 |

Example configuration JSON:

```json
{
  "enable": true,
  "refreshInterval": 60,
  "requestTimeout": 120,
  "projectIDs": [
    "10"
  ]
}
```

## Continuous Asset Creation

Once configured, the app listens for data updates from ROOMZ. With every presence status change (someone walks into or out of the space), the app receives a message from ROOMZ and writes the data to Eliona.

If the app has received a data from the sensor for the first time, then the app creates a new asset located under "ROOMZ root" asset. The asset is named by the space ID (that is the only information that ROOMZ API provides), but can be renamed and moved by the user.

This means that the app creates all assets automatically for all the spaces that are used.

The user who created or last updated the app's configuration will get notified of newly created assets.
