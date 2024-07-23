# ROOMZ User Guide

### Introduction

> The ROOMZ app integrates ROOMZ sensors with the Eliona platform to provide real-time data about workspace occupancy.

## Overview

This guide provides instructions for installing, configuring, and using the ROOMZ app to collect data from ROOMZ sensors.

## Installation

Install the ROOMZ app from the Eliona App Store.

## Configuration

Configure the ROOMZ app through the Eliona settings interface.

### Registering the App in ROOMZ Portal

The ROOMZ app uses an experimental [Webhook API](https://github.com/roomz-io/openapi), which is not generally available yet. Contact ROOMZ support to activate data sending. The webhook endpoint for your Eliona instance is: `https://{your-eliona-instance}/apps-public/roomz/webhook`.

### Configuring the ROOMZ App

Configure the app in Eliona by navigating to `Apps > ROOMZ > Settings`, which opens the app's [Generic Frontend](https://doc.eliona.io/collection/v/eliona-english/manuals/settings/apps). Use the POST method to set up the configuration. Each configuration requires the following parameters:

| Attribute         | Description                                                     |
|-------------------|-----------------------------------------------------------------|
| `enable`          | Flag to enable or disable this configuration.                   |
| `refreshInterval` | Interval in seconds for data synchronization.                   |
| `requestTimeout`  | API query timeout in seconds.                                   |
| `projectIDs`      | List of Eliona project IDs for data collection.                 |

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

Once configured, the app listens for data updates from ROOMZ. Each time there is a presence status change (someone enters or leaves a space), the app receives a message from ROOMZ and writes the data to Eliona.

The app automatically creates all assets for the spaces being monitored - If the app receives data from a sensor for the first time, it creates a new asset under the "ROOMZ root" asset. The asset is named by the space ID (the only information provided by the ROOMZ API) but can be renamed and relocated by the user.

The user who created or last updated the app's configuration will be notified of newly created assets.
