# Grafana Backup Tool

This tool is based on [ysde/grafana-backup-tool](https://github.com/ysde/grafana-backup-tool) version 1.4.2.

## The aim of this tool

- Backup all of the Grafana components
    - By local storage and AWS S3
- Restore all of the Grafana components
    - By local storage or AWS S3
- Restore a specific dashboard's JSON.
    - By only local storage

## Support components

- Folder
- Folder Permissions
- Library Elements (doesn't work with Grafana 8.0.0 but 8.4.3)
- Dashboard (contains Alert)
- Datasource
- Alert Channel
- Alert Rules (Supported in version 9.4.0 of grafana and up.)
- Teams
- Team Members (Needs Basic Authentication (username and password, see [grafana doc](https://grafana.com/docs/grafana/latest/http_api/org/#admin-organizations-api))
    - You need to set `Admin's account and password` in `grafanaSettings.json`, or set the base64 encoded `admin account and password` in ENV `GRAFANA_BASIC_AUTH`. E.g `export GRAFANA_BASIC_AUTH=YWRtaW46YWRtaW4=`
    - Or Sets this ENV of the Grafana server `GF_USERS_ALLOW_ORG_CREATE=true`. see [grafana doc](https://grafana.com/docs/grafana/latest/http_api/org/#create-organization)
- Organization (Needs Basic Authentication (username and password, see [grafana doc](https://grafana.com/docs/grafana/latest/http_api/org/#admin-organizations-api))
    - You need to set `Admin's account and password` in `grafanaSettings.json`, or set the base64 encoded `admin account and password` in ENV `GRAFANA_BASIC_AUTH`. E.g `export GRAFANA_BASIC_AUTH=YWRtaW46YWRtaW4=`
    - Or Sets this ENV of the Grafana server `GF_USERS_ALLOW_ORG_CREATE=true`. see [grafana doc](https://grafana.com/docs/grafana/latest/http_api/org/#create-organization)
- User (Needs Basic Authentication (username and password, see [grafana doc](https://grafana.com/docs/grafana/latest/http_api/org/#admin-organizations-api))
    - You need to set `Admin's account and password` in `grafanaSettings.json`, or set the base64 encoded `admin account and password` in ENV `GRAFANA_BASIC_AUTH`. E.g `export GRAFANA_BASIC_AUTH=YWRtaW46YWRtaW4=`
    - Grafana's api doesn't provide user's password when backing up, so the `default_user_password` in `grafanaSettings.json`, or in ENV `DEFAULT_USER_PASSWORD`, E.g `export DEFAULT_USER_PASSWORD=supersecret` will be used when restoring.
- Snapshots
- Dashboard Versions (only backup, no restore)
- Annotations

**NOTE** The only supported `orgId` right now is `1`, the default organization will be backed up only!

## Quick Guide

### Build & Start

1. **Build**

```bash
make build
```

2. **Fill variables in `docker-compose.yml` file**

```bash
cp docker-compose.yml.example docker-compose.yml
vim docker-compose.yml
```

3. **Start**

```bash
docker compose up -d
```

### How to Use

There’s 3 features: **backup all**, **restore all**, **restore dashboard’s JSON**.

1. **Backup all of the Grafana components**
    - Backup file would be store in **local storage and aws s3**

```bash
curl localhost:9111/backup
```

2. **Restore all of the Grafana components**
    - Try restore all of Grafana components from **local storage or AWS S3**

```bash
curl -G localhost:9111/restore/all \
	--data-urlencode archive="202501220533.tar.gz"
```

3. **Restore a specific dashboard's JSON**
    - Try restore a specific dashboard’s JSON from **local storage**
    - Response of below command is dashboard’s JSON

```bash
curl -G http://localhost:9111/restore/dashboard \
	--data-urlencode archive="202501220533.tar.gz" \
	--data-urlencode title="Cosmos Validator"
```
